package xsql

import (
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/bind"
	querySql "github.com/ydb-platform/ydb-go-sdk/v3/internal/query/conn"
	tableSql "github.com/ydb-platform/ydb-go-sdk/v3/internal/table/conn"
	"github.com/ydb-platform/ydb-go-sdk/v3/retry/budget"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

type (
	Option interface {
		Apply(c *Connector) error
	}
	QueryBindOption interface {
		Option
		bind.Bind
	}
	tableQueryOptionsOption struct {
		tableOps  []tableSql.Option
		queryOpts []querySql.Option
	}
	traceDatabaseSQLOption struct {
		t    *trace.DatabaseSQL
		opts []trace.DatabaseSQLComposeOption
	}
	traceRetryOption struct {
		t    *trace.Retry
		opts []trace.RetryComposeOption
	}
	disableServerBalancerOption struct{}
	onCloseOption               func(*Connector)
	retryBudgetOption           struct {
		budget budget.Budget
	}
	bindOption struct {
		bind.Bind
	}
	queryProcessorOption queryProcessor
)

func (processor queryProcessorOption) Apply(c *Connector) error {
	c.queryProcessor = queryProcessor(processor)

	return nil
}

func (opt bindOption) Apply(c *Connector) error {
	c.tableOpts = append(c.tableOpts, tableSql.WithQueryBindings(opt.Bind))
	c.queryOpts = append(c.queryOpts, querySql.WithQueryBindings(opt.Bind))

	return nil
}

func (opt retryBudgetOption) Apply(c *Connector) error {
	c.tableOpts = append(c.tableOpts, tableSql.WithRetryBudget(opt.budget))
	c.queryOpts = append(c.queryOpts, querySql.WithRetryBudget(opt.budget))

	return nil
}

func (opt traceRetryOption) Apply(c *Connector) error {
	c.traceRetry = c.traceRetry.Compose(opt.t, opt.opts...)
	c.tableOpts = append(c.tableOpts, tableSql.WithTraceRetry(opt.t, opt.opts...))
	c.queryOpts = append(c.queryOpts, querySql.WithTraceRetry(opt.t, opt.opts...))

	return nil
}

func (onClose onCloseOption) Apply(c *Connector) error {
	c.onCLose = append(c.onCLose, onClose)

	return nil
}

func (disableServerBalancerOption) Apply(c *Connector) error {
	c.disableServerBalancer = true

	return nil
}

func (opt traceDatabaseSQLOption) Apply(c *Connector) error {
	c.trace = c.trace.Compose(opt.t, opt.opts...)
	c.tableOpts = append(c.tableOpts, tableSql.WithTrace(opt.t, opt.opts...))
	c.queryOpts = append(c.queryOpts, querySql.WithTrace(opt.t, opt.opts...))

	return nil
}

func (opt tableQueryOptionsOption) Apply(c *Connector) error {
	c.queryOpts = append(c.queryOpts, opt.queryOpts...)
	c.tableOpts = append(c.tableOpts, opt.tableOps...)

	return nil
}

func WithTrace(
	t *trace.DatabaseSQL, //nolint:gocritic
	opts ...trace.DatabaseSQLComposeOption,
) Option {
	return traceDatabaseSQLOption{
		t:    t,
		opts: opts,
	}
}

func WithDisableServerBalancer() Option {
	return disableServerBalancerOption{}
}

func WithOnClose(onClose func(*Connector)) Option {
	return onCloseOption(onClose)
}

func WithTraceRetry(
	t *trace.Retry,
	opts ...trace.RetryComposeOption,
) Option {
	return traceRetryOption{
		t:    t,
		opts: opts,
	}
}

func WithRetryBudget(budget budget.Budget) Option {
	return retryBudgetOption{
		budget: budget,
	}
}

func WithQueryBind(bind bind.Bind) QueryBindOption {
	return bindOption{
		Bind: bind,
	}
}

func WithDefaultQueryMode(mode tableSql.QueryMode) Option {
	return tableQueryOptionsOption{
		tableOps: []tableSql.Option{
			tableSql.WithDefaultQueryMode(mode),
		},
	}
}

func WithFakeTx(modes ...tableSql.QueryMode) Option {
	return tableQueryOptionsOption{
		tableOps: []tableSql.Option{
			tableSql.WithFakeTxModes(modes...),
		},
	}
}

func WithIdleThreshold(idleThreshold time.Duration) Option {
	return tableQueryOptionsOption{
		tableOps: []tableSql.Option{
			tableSql.WithIdleThreshold(idleThreshold),
		},
	}
}

func WithTableOptions(opts ...tableSql.Option) Option {
	return tableQueryOptionsOption{
		tableOps: opts,
	}
}

func WithQueryOptions(opts ...querySql.Option) Option {
	return tableQueryOptionsOption{
		queryOpts: opts,
	}
}

func OverQueryService() Option {
	return queryProcessorOption(QUERY_SERVICE)
}

func OverTableService() Option {
	return queryProcessorOption(TABLE_SERVICE)
}
