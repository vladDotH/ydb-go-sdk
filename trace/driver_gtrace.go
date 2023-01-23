// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// driverComposeOptions is a holder of options
type driverComposeOptions struct {
	panicCallback func(e interface{})
}

// DriverOption specified Driver compose option
type DriverComposeOption func(o *driverComposeOptions)

// WithDriverPanicCallback specified behavior on panic
func WithDriverPanicCallback(cb func(e interface{})) DriverComposeOption {
	return func(o *driverComposeOptions) {
		o.panicCallback = cb
	}
}

// Compose returns a new Driver which has functional fields composed both from t and x.
func (t Driver) Compose(x Driver, opts ...DriverComposeOption) (ret Driver) {
	options := driverComposeOptions{}
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	{
		h1 := t.OnInit
		h2 := x.OnInit
		ret.OnInit = func(d DriverInitStartInfo) func(DriverInitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverInitDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverInitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnClose
		h2 := x.OnClose
		ret.OnClose = func(d DriverCloseStartInfo) func(DriverCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverCloseDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnNetRead
		h2 := x.OnNetRead
		ret.OnNetRead = func(d DriverNetReadStartInfo) func(DriverNetReadDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverNetReadDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverNetReadDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnNetWrite
		h2 := x.OnNetWrite
		ret.OnNetWrite = func(d DriverNetWriteStartInfo) func(DriverNetWriteDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverNetWriteDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverNetWriteDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnNetDial
		h2 := x.OnNetDial
		ret.OnNetDial = func(d DriverNetDialStartInfo) func(DriverNetDialDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverNetDialDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverNetDialDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnNetClose
		h2 := x.OnNetClose
		ret.OnNetClose = func(d DriverNetCloseStartInfo) func(DriverNetCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverNetCloseDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverNetCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnResolve
		h2 := x.OnResolve
		ret.OnResolve = func(d DriverResolveStartInfo) func(DriverResolveDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverResolveDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverResolveDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnStateChange
		h2 := x.OnConnStateChange
		ret.OnConnStateChange = func(d DriverConnStateChangeStartInfo) func(DriverConnStateChangeDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverConnStateChangeDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverConnStateChangeDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnInvoke
		h2 := x.OnConnInvoke
		ret.OnConnInvoke = func(d DriverConnInvokeStartInfo) func(DriverConnInvokeDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverConnInvokeDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverConnInvokeDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnNewStream
		h2 := x.OnConnNewStream
		ret.OnConnNewStream = func(d DriverConnNewStreamStartInfo) func(DriverConnNewStreamRecvInfo) func(DriverConnNewStreamDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverConnNewStreamRecvInfo) func(DriverConnNewStreamDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverConnNewStreamRecvInfo) func(DriverConnNewStreamDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(DriverConnNewStreamDoneInfo)
				if r != nil {
					r2 = r(d)
				}
				if r1 != nil {
					r3 = r1(d)
				}
				return func(d DriverConnNewStreamDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(d)
					}
					if r3 != nil {
						r3(d)
					}
				}
			}
		}
	}
	{
		h1 := t.OnConnTake
		h2 := x.OnConnTake
		ret.OnConnTake = func(d DriverConnTakeStartInfo) func(DriverConnTakeDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverConnTakeDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverConnTakeDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnPark
		h2 := x.OnConnPark
		ret.OnConnPark = func(d DriverConnParkStartInfo) func(DriverConnParkDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverConnParkDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverConnParkDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnBan
		h2 := x.OnConnBan
		ret.OnConnBan = func(d DriverConnBanStartInfo) func(DriverConnBanDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverConnBanDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverConnBanDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnAllow
		h2 := x.OnConnAllow
		ret.OnConnAllow = func(d DriverConnAllowStartInfo) func(DriverConnAllowDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverConnAllowDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverConnAllowDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnConnClose
		h2 := x.OnConnClose
		ret.OnConnClose = func(d DriverConnCloseStartInfo) func(DriverConnCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverConnCloseDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverConnCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnRepeaterWakeUp
		h2 := x.OnRepeaterWakeUp
		ret.OnRepeaterWakeUp = func(d DriverRepeaterWakeUpStartInfo) func(DriverRepeaterWakeUpDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverRepeaterWakeUpDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverRepeaterWakeUpDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnBalancerInit
		h2 := x.OnBalancerInit
		ret.OnBalancerInit = func(d DriverBalancerInitStartInfo) func(DriverBalancerInitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverBalancerInitDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverBalancerInitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnBalancerClose
		h2 := x.OnBalancerClose
		ret.OnBalancerClose = func(d DriverBalancerCloseStartInfo) func(DriverBalancerCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverBalancerCloseDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverBalancerCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnBalancerChooseEndpoint
		h2 := x.OnBalancerChooseEndpoint
		ret.OnBalancerChooseEndpoint = func(d DriverBalancerChooseEndpointStartInfo) func(DriverBalancerChooseEndpointDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverBalancerChooseEndpointDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverBalancerChooseEndpointDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnBalancerUpdate
		h2 := x.OnBalancerUpdate
		ret.OnBalancerUpdate = func(d DriverBalancerUpdateStartInfo) func(DriverBalancerUpdateDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverBalancerUpdateDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverBalancerUpdateDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	{
		h1 := t.OnGetCredentials
		h2 := x.OnGetCredentials
		ret.OnGetCredentials = func(d DriverGetCredentialsStartInfo) func(DriverGetCredentialsDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(DriverGetCredentialsDoneInfo)
			if h1 != nil {
				r = h1(d)
			}
			if h2 != nil {
				r1 = h2(d)
			}
			return func(d DriverGetCredentialsDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(d)
				}
				if r1 != nil {
					r1(d)
				}
			}
		}
	}
	return ret
}
func (t Driver) onInit(d DriverInitStartInfo) func(DriverInitDoneInfo) {
	fn := t.OnInit
	if fn == nil {
		return func(DriverInitDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverInitDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClose(d DriverCloseStartInfo) func(DriverCloseDoneInfo) {
	fn := t.OnClose
	if fn == nil {
		return func(DriverCloseDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onNetRead(d DriverNetReadStartInfo) func(DriverNetReadDoneInfo) {
	fn := t.OnNetRead
	if fn == nil {
		return func(DriverNetReadDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverNetReadDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onNetWrite(d DriverNetWriteStartInfo) func(DriverNetWriteDoneInfo) {
	fn := t.OnNetWrite
	if fn == nil {
		return func(DriverNetWriteDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverNetWriteDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onNetDial(d DriverNetDialStartInfo) func(DriverNetDialDoneInfo) {
	fn := t.OnNetDial
	if fn == nil {
		return func(DriverNetDialDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverNetDialDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onNetClose(d DriverNetCloseStartInfo) func(DriverNetCloseDoneInfo) {
	fn := t.OnNetClose
	if fn == nil {
		return func(DriverNetCloseDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverNetCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onResolve(d DriverResolveStartInfo) func(DriverResolveDoneInfo) {
	fn := t.OnResolve
	if fn == nil {
		return func(DriverResolveDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverResolveDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnStateChange(d DriverConnStateChangeStartInfo) func(DriverConnStateChangeDoneInfo) {
	fn := t.OnConnStateChange
	if fn == nil {
		return func(DriverConnStateChangeDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverConnStateChangeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnInvoke(d DriverConnInvokeStartInfo) func(DriverConnInvokeDoneInfo) {
	fn := t.OnConnInvoke
	if fn == nil {
		return func(DriverConnInvokeDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverConnInvokeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnNewStream(d DriverConnNewStreamStartInfo) func(DriverConnNewStreamRecvInfo) func(DriverConnNewStreamDoneInfo) {
	fn := t.OnConnNewStream
	if fn == nil {
		return func(DriverConnNewStreamRecvInfo) func(DriverConnNewStreamDoneInfo) {
			return func(DriverConnNewStreamDoneInfo) {
				return
			}
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverConnNewStreamRecvInfo) func(DriverConnNewStreamDoneInfo) {
			return func(DriverConnNewStreamDoneInfo) {
				return
			}
		}
	}
	return func(d DriverConnNewStreamRecvInfo) func(DriverConnNewStreamDoneInfo) {
		res := res(d)
		if res == nil {
			return func(DriverConnNewStreamDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t Driver) onConnTake(d DriverConnTakeStartInfo) func(DriverConnTakeDoneInfo) {
	fn := t.OnConnTake
	if fn == nil {
		return func(DriverConnTakeDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverConnTakeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnPark(d DriverConnParkStartInfo) func(DriverConnParkDoneInfo) {
	fn := t.OnConnPark
	if fn == nil {
		return func(DriverConnParkDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverConnParkDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnBan(d DriverConnBanStartInfo) func(DriverConnBanDoneInfo) {
	fn := t.OnConnBan
	if fn == nil {
		return func(DriverConnBanDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverConnBanDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnAllow(d DriverConnAllowStartInfo) func(DriverConnAllowDoneInfo) {
	fn := t.OnConnAllow
	if fn == nil {
		return func(DriverConnAllowDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverConnAllowDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnClose(d DriverConnCloseStartInfo) func(DriverConnCloseDoneInfo) {
	fn := t.OnConnClose
	if fn == nil {
		return func(DriverConnCloseDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverConnCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onRepeaterWakeUp(d DriverRepeaterWakeUpStartInfo) func(DriverRepeaterWakeUpDoneInfo) {
	fn := t.OnRepeaterWakeUp
	if fn == nil {
		return func(DriverRepeaterWakeUpDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverRepeaterWakeUpDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onBalancerInit(d DriverBalancerInitStartInfo) func(DriverBalancerInitDoneInfo) {
	fn := t.OnBalancerInit
	if fn == nil {
		return func(DriverBalancerInitDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverBalancerInitDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onBalancerClose(d DriverBalancerCloseStartInfo) func(DriverBalancerCloseDoneInfo) {
	fn := t.OnBalancerClose
	if fn == nil {
		return func(DriverBalancerCloseDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverBalancerCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onBalancerChooseEndpoint(d DriverBalancerChooseEndpointStartInfo) func(DriverBalancerChooseEndpointDoneInfo) {
	fn := t.OnBalancerChooseEndpoint
	if fn == nil {
		return func(DriverBalancerChooseEndpointDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverBalancerChooseEndpointDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onBalancerUpdate(d DriverBalancerUpdateStartInfo) func(DriverBalancerUpdateDoneInfo) {
	fn := t.OnBalancerUpdate
	if fn == nil {
		return func(DriverBalancerUpdateDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverBalancerUpdateDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onGetCredentials(d DriverGetCredentialsStartInfo) func(DriverGetCredentialsDoneInfo) {
	fn := t.OnGetCredentials
	if fn == nil {
		return func(DriverGetCredentialsDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DriverGetCredentialsDoneInfo) {
			return
		}
	}
	return res
}
func DriverOnInit(t Driver, c *context.Context, endpoint string, database string, secure bool) func(error) {
	var p DriverInitStartInfo
	p.Context = c
	p.Endpoint = endpoint
	p.Database = database
	p.Secure = secure
	res := t.onInit(p)
	return func(e error) {
		var p DriverInitDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnClose(t Driver, c *context.Context) func(error) {
	var p DriverCloseStartInfo
	p.Context = c
	res := t.onClose(p)
	return func(e error) {
		var p DriverCloseDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnNetRead(t Driver, address string, buffer int) func(received int, _ error) {
	var p DriverNetReadStartInfo
	p.Address = address
	p.Buffer = buffer
	res := t.onNetRead(p)
	return func(received int, e error) {
		var p DriverNetReadDoneInfo
		p.Received = received
		p.Error = e
		res(p)
	}
}
func DriverOnNetWrite(t Driver, address string, bytes int) func(sent int, _ error) {
	var p DriverNetWriteStartInfo
	p.Address = address
	p.Bytes = bytes
	res := t.onNetWrite(p)
	return func(sent int, e error) {
		var p DriverNetWriteDoneInfo
		p.Sent = sent
		p.Error = e
		res(p)
	}
}
func DriverOnNetDial(t Driver, c *context.Context, address string) func(error) {
	var p DriverNetDialStartInfo
	p.Context = c
	p.Address = address
	res := t.onNetDial(p)
	return func(e error) {
		var p DriverNetDialDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnNetClose(t Driver, address string) func(error) {
	var p DriverNetCloseStartInfo
	p.Address = address
	res := t.onNetClose(p)
	return func(e error) {
		var p DriverNetCloseDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnResolve(t Driver, target string, resolved []string) func(error) {
	var p DriverResolveStartInfo
	p.Target = target
	p.Resolved = resolved
	res := t.onResolve(p)
	return func(e error) {
		var p DriverResolveDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnConnStateChange(t Driver, endpoint EndpointInfo, state ConnState) func(state ConnState) {
	var p DriverConnStateChangeStartInfo
	p.Endpoint = endpoint
	p.State = state
	res := t.onConnStateChange(p)
	return func(state ConnState) {
		var p DriverConnStateChangeDoneInfo
		p.State = state
		res(p)
	}
}
func DriverOnConnInvoke(t Driver, c *context.Context, endpoint EndpointInfo, m Method) func(_ error, issues []Issue, opID string, state ConnState, metadata map[string][]string) {
	var p DriverConnInvokeStartInfo
	p.Context = c
	p.Endpoint = endpoint
	p.Method = m
	res := t.onConnInvoke(p)
	return func(e error, issues []Issue, opID string, state ConnState, metadata map[string][]string) {
		var p DriverConnInvokeDoneInfo
		p.Error = e
		p.Issues = issues
		p.OpID = opID
		p.State = state
		p.Metadata = metadata
		res(p)
	}
}
func DriverOnConnNewStream(t Driver, c *context.Context, endpoint EndpointInfo, m Method) func(error) func(_ error, state ConnState, metadata map[string][]string) {
	var p DriverConnNewStreamStartInfo
	p.Context = c
	p.Endpoint = endpoint
	p.Method = m
	res := t.onConnNewStream(p)
	return func(e error) func(error, ConnState, map[string][]string) {
		var p DriverConnNewStreamRecvInfo
		p.Error = e
		res := res(p)
		return func(e error, state ConnState, metadata map[string][]string) {
			var p DriverConnNewStreamDoneInfo
			p.Error = e
			p.State = state
			p.Metadata = metadata
			res(p)
		}
	}
}
func DriverOnConnTake(t Driver, c *context.Context, endpoint EndpointInfo) func(error) {
	var p DriverConnTakeStartInfo
	p.Context = c
	p.Endpoint = endpoint
	res := t.onConnTake(p)
	return func(e error) {
		var p DriverConnTakeDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnConnPark(t Driver, c *context.Context, endpoint EndpointInfo) func(error) {
	var p DriverConnParkStartInfo
	p.Context = c
	p.Endpoint = endpoint
	res := t.onConnPark(p)
	return func(e error) {
		var p DriverConnParkDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnConnBan(t Driver, c *context.Context, endpoint EndpointInfo, state ConnState, cause error) func(state ConnState) {
	var p DriverConnBanStartInfo
	p.Context = c
	p.Endpoint = endpoint
	p.State = state
	p.Cause = cause
	res := t.onConnBan(p)
	return func(state ConnState) {
		var p DriverConnBanDoneInfo
		p.State = state
		res(p)
	}
}
func DriverOnConnAllow(t Driver, c *context.Context, endpoint EndpointInfo, state ConnState) func(state ConnState) {
	var p DriverConnAllowStartInfo
	p.Context = c
	p.Endpoint = endpoint
	p.State = state
	res := t.onConnAllow(p)
	return func(state ConnState) {
		var p DriverConnAllowDoneInfo
		p.State = state
		res(p)
	}
}
func DriverOnConnClose(t Driver, c *context.Context, endpoint EndpointInfo) func(error) {
	var p DriverConnCloseStartInfo
	p.Context = c
	p.Endpoint = endpoint
	res := t.onConnClose(p)
	return func(e error) {
		var p DriverConnCloseDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnRepeaterWakeUp(t Driver, c *context.Context, name string, event string) func(error) {
	var p DriverRepeaterWakeUpStartInfo
	p.Context = c
	p.Name = name
	p.Event = event
	res := t.onRepeaterWakeUp(p)
	return func(e error) {
		var p DriverRepeaterWakeUpDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnBalancerInit(t Driver, c *context.Context) func(error) {
	var p DriverBalancerInitStartInfo
	p.Context = c
	res := t.onBalancerInit(p)
	return func(e error) {
		var p DriverBalancerInitDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnBalancerClose(t Driver, c *context.Context) func(error) {
	var p DriverBalancerCloseStartInfo
	p.Context = c
	res := t.onBalancerClose(p)
	return func(e error) {
		var p DriverBalancerCloseDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnBalancerChooseEndpoint(t Driver, c *context.Context) func(endpoint EndpointInfo, _ error) {
	var p DriverBalancerChooseEndpointStartInfo
	p.Context = c
	res := t.onBalancerChooseEndpoint(p)
	return func(endpoint EndpointInfo, e error) {
		var p DriverBalancerChooseEndpointDoneInfo
		p.Endpoint = endpoint
		p.Error = e
		res(p)
	}
}
func DriverOnBalancerUpdate(t Driver, c *context.Context, needLocalDC bool) func(endpoints []EndpointInfo, localDC string, _ error) {
	var p DriverBalancerUpdateStartInfo
	p.Context = c
	p.NeedLocalDC = needLocalDC
	res := t.onBalancerUpdate(p)
	return func(endpoints []EndpointInfo, localDC string, e error) {
		var p DriverBalancerUpdateDoneInfo
		p.Endpoints = endpoints
		p.LocalDC = localDC
		p.Error = e
		res(p)
	}
}
func DriverOnGetCredentials(t Driver, c *context.Context) func(token string, _ error) {
	var p DriverGetCredentialsStartInfo
	p.Context = c
	res := t.onGetCredentials(p)
	return func(token string, e error) {
		var p DriverGetCredentialsDoneInfo
		p.Token = token
		p.Error = e
		res(p)
	}
}
