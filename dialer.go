package ydb

import (
	"context"
	"crypto/tls"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/ydb-platform/ydb-go-genproto/Ydb_Discovery_V1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

var (
	// DefaultKeepaliveInterval contains default duration between grpc keepalive
	DefaultKeepaliveInterval = 5 * time.Minute
	MinKeepaliveInterval     = 1 * time.Minute
	DefaultGRPCMsgSize       = 64 * 1024 * 1024 // 64MB
)

// Dialer contains options of dialing and initialization of particular ydb
// driver.
type Dialer struct {
	// DriverConfig is a driver configuration.
	DriverConfig *DriverConfig

	// TLSConfig specifies the TLS configuration to use for tls client.
	// If TLSConfig is zero then connections are insecure.
	TLSConfig *tls.Config

	// Timeout is the maximum amount of time a dial will wait for a connect to
	// complete.
	// If Timeout is zero then no timeout is used.
	Timeout time.Duration

	// Keepalive is the interval used to check whether inner connections are
	// still valid.
	// Dialer could increase keepalive interval if given value is too small.
	Keepalive time.Duration

	// NetDial is an optional function that may replace default network dialing
	// function such as net.Dial("tcp").
	// Deprecated: Use it for test purposes and special cases only. In most cases should be left empty.
	NetDial func(context.Context, string) (net.Conn, error)
}

// Dial dials given addr and initializes driver instance on success.
func (d *Dialer) Dial(ctx context.Context, addr string) (_ *driver, err error) {
	config := d.DriverConfig.withDefaults()
	grpcKeepalive := d.Keepalive
	if grpcKeepalive == 0 {
		grpcKeepalive = DefaultKeepaliveInterval
	} else if grpcKeepalive < MinKeepaliveInterval {
		grpcKeepalive = MinKeepaliveInterval
	}
	tlsConfig := d.TLSConfig
	if tlsConfig != nil {
		tlsConfig.RootCAs, err = WithYdbCA(tlsConfig.RootCAs)
		if err != nil {
			return nil, err
		}
	}
	return (&dialer{
		netDial:   d.NetDial,
		tlsConfig: tlsConfig,
		keepalive: grpcKeepalive,
		timeout:   d.Timeout,
		config:    config,
		meta: &meta{
			trace:        config.Trace,
			database:     config.Database,
			credentials:  config.Credentials,
			requestsType: config.RequestsType,
		},
	}).dial(ctx, addr)
}

// dialer is an instance holding single Dialer.Dial() configuration parameters.
type dialer struct {
	netDial   func(context.Context, string) (net.Conn, error)
	tlsConfig *tls.Config
	keepalive time.Duration
	timeout   time.Duration
	config    DriverConfig
	meta      *meta
}

func (d *dialer) dial(ctx context.Context, addr string) (_ *driver, err error) {
	endpoint := d.endpointByAddr(addr)
	c := cluster{
		dial:  d.dialHostPort,
		ttl:   d.config.ConnectionTTL,
		trace: d.config.Trace,
		balancer: func() balancer {
			if d.config.DiscoveryInterval == 0 {
				return new(singleConnBalancer)
			}
			if !d.config.PreferLocalEndpoints {
				return d.newBalancer()
			}
			return newMultiBalancer(
				withBalancer(
					d.newBalancer(), func(_ *conn, info connInfo) bool {
						return info.local
					},
				),
				withBalancer(
					d.newBalancer(), func(_ *conn, info connInfo) bool {
						return !info.local
					},
				),
			)
		}(),
	}
	defer func() {
		if err != nil {
			_ = c.Close()
		}
	}()
	c.Insert(ctx, endpoint)
	driver := &driver{
		cluster:              &c,
		meta:                 d.meta,
		trace:                d.config.Trace,
		requestTimeout:       d.config.RequestTimeout,
		streamTimeout:        d.config.StreamTimeout,
		operationTimeout:     d.config.OperationTimeout,
		operationCancelAfter: d.config.OperationCancelAfter,
	}
	// Ensure that endpoint is online.
	_, err = driver.getConn(ctx)
	if err != nil {
		return nil, err
	}
	if d.config.DiscoveryInterval > 0 {
		discoveryClient := &discoveryClient{
			discoveryService: Ydb_Discovery_V1.NewDiscoveryServiceClient(driver),
			database:         d.config.Database,
			ssl:              d.useTLS(),
		}

		curr, err := discoveryClient.Discover(ctx)
		if err != nil {
			return nil, err
		}
		// Endpoints must be sorted to merge
		sortEndpoints(curr)
		wg := newWG()
		wg.Add(len(curr))
		for _, e := range curr {
			go c.Insert(ctx, e, wg)
		}
		if d.config.FastDial {
			wg.WaitFirst()
		} else {
			wg.Wait()
		}
		c.explorer = NewRepeater(
			d.config.DiscoveryInterval,
			func(ctx context.Context) {
				next, err := discoveryClient.Discover(ctx)
				if err != nil {
					return
				}
				// if nothing endpoint - re-discover after one second
				// and use old endpoint list
				if len(next) == 0 {
					go func() {
						time.Sleep(time.Second)
						c.explorer.Force()

					}()
					return
				}
				// NOTE: curr endpoints must be sorted here.
				sortEndpoints(next)

				wg := new(sync.WaitGroup)
				max := len(next) + len(curr)
				wg.Add(max) // set to max possible amount
				actual := 0
				diffEndpoints(curr, next,
					func(i, j int) {
						actual++
						// Endpoints are equal, but we still need to update meta
						// data such that load factor and others.
						go c.Update(ctx, next[j], wg)
					},
					func(i, j int) {
						actual++
						go c.Insert(ctx, next[j], wg)
					},
					func(i, j int) {
						actual++
						go c.Remove(ctx, curr[i], wg)
					},
				)
				wg.Add(actual - max) // adjust
				wg.Wait()
				curr = next
			},
		)
	}
	return driver, nil
}

func (d *dialer) dialHostPort(ctx context.Context, host string, port int) (*grpc.ClientConn, error) {
	rawСtx := ctx
	if d.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, d.timeout)
		defer cancel()
	}
	s := connAddr{addr: host, port: port}.String()
	dialDone := driverTraceOnDial(d.config.Trace, ctx, s)

	cc, err := grpc.DialContext(ctx, s, d.grpcDialOptions()...)

	dialDone(rawСtx, s, err)
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func (d *dialer) grpcDialOptions() (opts []grpc.DialOption) {
	if d.netDial != nil {
		opts = append(opts, grpc.WithContextDialer(d.netDial))
	}
	if d.useTLS() {
		opts = append(opts, grpc.WithTransportCredentials(
			credentials.NewTLS(d.tlsConfig),
		))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts,
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    d.keepalive,
			Timeout: d.timeout,
		}),
	)
	opts = append(opts, grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(DefaultGRPCMsgSize),
		grpc.MaxCallSendMsgSize(DefaultGRPCMsgSize),
	))

	return append(opts, grpc.WithBlock())
}

func (d *dialer) newBalancer() balancer {
	return balancers[d.config.BalancingMethod](d.config.BalancingConfig)
}

func (d *dialer) useTLS() bool {
	return d.tlsConfig != nil
}

func (d *dialer) splitHostPort(addr string) (host string, port int, err error) {
	var prt string
	host, prt, err = net.SplitHostPort(addr)
	if err != nil {
		return
	}
	port, err = strconv.Atoi(prt)
	return
}

func (d *dialer) mustSplitHostPort(addr string) (host string, port int) {
	host, prt, err := net.SplitHostPort(addr)
	if err != nil {
		panic(err)
	}
	port, err = strconv.Atoi(prt)
	if err != nil {
		panic(err)
	}
	return host, port
}

func (d *dialer) endpointByAddr(addr string) (e Endpoint) {
	e.Addr, e.Port = d.mustSplitHostPort(addr)
	return
}

// withContextDialer is an adapter to allow the use of normal go-world net dial
// function as WithDialer option argument for grpc Dial().
func withContextDialer(f func(context.Context, string) (net.Conn, error)) func(string, time.Duration) (net.Conn, error) {
	if f == nil {
		return nil
	}
	return func(addr string, timeout time.Duration) (net.Conn, error) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		return f(ctx, addr)
	}
}
