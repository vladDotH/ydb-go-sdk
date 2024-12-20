package xsql

import (
	"context"
	"database/sql/driver"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/stack"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql/iface"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql/table/conn/badconn"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

type txWrapper struct {
	conn *conn
	tx   iface.Tx
	ctx  context.Context //nolint:containedctx
}

func (tx *txWrapper) ID() string {
	return tx.tx.ID()
}

var (
	_ driver.Tx             = &txWrapper{}
	_ driver.ExecerContext  = &txWrapper{}
	_ driver.QueryerContext = &txWrapper{}
)

func (tx *txWrapper) Commit() (finalErr error) {
	var (
		ctx    = tx.ctx
		onDone = trace.DatabaseSQLOnTxCommit(tx.conn.connector.Trace(), &ctx,
			stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql.(*txWrapper).Commit"),
			tx,
		)
	)
	defer func() {
		onDone(finalErr)
	}()

	if err := tx.tx.Commit(tx.ctx); err != nil {
		return xerrors.WithStackTrace(err)
	}

	return nil
}

func (tx *txWrapper) Rollback() (finalErr error) {
	var (
		ctx    = tx.ctx
		onDone = trace.DatabaseSQLOnTxRollback(tx.conn.connector.Trace(), &ctx,
			stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql.(*txWrapper).Rollback"),
			tx,
		)
	)
	defer func() {
		onDone(finalErr)
	}()

	err := tx.tx.Rollback(tx.ctx)
	if err != nil {
		return xerrors.WithStackTrace(err)
	}

	return err
}

func (tx *txWrapper) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (
	_ driver.Rows, finalErr error,
) {
	onDone := trace.DatabaseSQLOnTxQuery(tx.conn.connector.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql.(*txWrapper).QueryContext"),
		tx.ctx, tx, query,
	)
	defer func() {
		onDone(finalErr)
	}()

	query, params, err := tx.conn.normalize(query, args...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	if isExplain(ctx) {
		ast, plan, err := tx.conn.cc.Explain(ctx, query, params)
		if err != nil {
			return nil, xerrors.WithStackTrace(err)
		}

		return rowByAstPlan(ast, plan), nil
	}

	rows, err := tx.tx.Query(ctx, query, params)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return rows, nil
}

func (tx *txWrapper) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (
	_ driver.Result, finalErr error,
) {
	onDone := trace.DatabaseSQLOnTxExec(tx.conn.connector.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql.(*txWrapper).ExecContext"),
		tx.ctx, tx, query,
	)
	defer func() {
		onDone(finalErr)
	}()

	query, params, err := tx.conn.normalize(query, args...)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	result, err := tx.tx.Exec(ctx, query, params)
	if err != nil {
		return nil, xerrors.WithStackTrace(err)
	}

	return result, nil
}

func (tx *txWrapper) PrepareContext(ctx context.Context, sql string) (_ driver.Stmt, finalErr error) {
	onDone := trace.DatabaseSQLOnTxPrepare(tx.conn.connector.Trace(), &ctx,
		stack.FunctionID("github.com/ydb-platform/ydb-go-sdk/v3/internal/xsql.(*txWrapper).PrepareContext"),
		tx.ctx, tx, sql,
	)
	defer func() {
		onDone(finalErr)
	}()
	if !tx.conn.cc.IsValid() {
		return nil, badconn.Map(xerrors.WithStackTrace(errNotReadyConn))
	}

	return &stmt{
		conn:      tx.conn,
		processor: tx.tx,
		ctx:       ctx,
		sql:       sql,
	}, nil
}
