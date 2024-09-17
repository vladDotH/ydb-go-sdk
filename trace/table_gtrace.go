// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// tableComposeOptions is a holder of options
type tableComposeOptions struct {
	panicCallback func(e interface{})
}

// TableOption specified Table compose option
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
type TableComposeOption func(o *tableComposeOptions)

// WithTablePanicCallback specified behavior on panic
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func WithTablePanicCallback(cb func(e interface{})) TableComposeOption {
	return func(o *tableComposeOptions) {
		o.panicCallback = cb
	}
}

// Compose returns a new Table which has functional fields composed both from t and x.
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func (t *Table) Compose(x *Table, opts ...TableComposeOption) *Table {
	var ret Table
	options := tableComposeOptions{}
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	{
		h1 := t.OnInit
		h2 := x.OnInit
		ret.OnInit = func(t TableInitStartInfo) func(TableInitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableInitDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableInitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnClose
		h2 := x.OnClose
		ret.OnClose = func(t TableCloseStartInfo) func(TableCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableCloseDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnDo
		h2 := x.OnDo
		ret.OnDo = func(t TableDoStartInfo) func(TableDoDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableDoDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableDoDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnDoTx
		h2 := x.OnDoTx
		ret.OnDoTx = func(t TableDoTxStartInfo) func(TableDoTxDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableDoTxDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableDoTxDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnCreateSession
		h2 := x.OnCreateSession
		ret.OnCreateSession = func(t TableCreateSessionStartInfo) func(TableCreateSessionDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableCreateSessionDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableCreateSessionDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionNew
		h2 := x.OnSessionNew
		ret.OnSessionNew = func(t TableSessionNewStartInfo) func(TableSessionNewDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionNewDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionNewDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionDelete
		h2 := x.OnSessionDelete
		ret.OnSessionDelete = func(t TableSessionDeleteStartInfo) func(TableSessionDeleteDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionDeleteDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionDeleteDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionKeepAlive
		h2 := x.OnSessionKeepAlive
		ret.OnSessionKeepAlive = func(t TableKeepAliveStartInfo) func(TableKeepAliveDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableKeepAliveDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableKeepAliveDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionBulkUpsert
		h2 := x.OnSessionBulkUpsert
		ret.OnSessionBulkUpsert = func(t TableBulkUpsertStartInfo) func(TableBulkUpsertDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableBulkUpsertDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableBulkUpsertDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionQueryPrepare
		h2 := x.OnSessionQueryPrepare
		ret.OnSessionQueryPrepare = func(t TablePrepareDataQueryStartInfo) func(TablePrepareDataQueryDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TablePrepareDataQueryDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TablePrepareDataQueryDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionQueryExecute
		h2 := x.OnSessionQueryExecute
		ret.OnSessionQueryExecute = func(t TableExecuteDataQueryStartInfo) func(TableExecuteDataQueryDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableExecuteDataQueryDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableExecuteDataQueryDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionQueryExplain
		h2 := x.OnSessionQueryExplain
		ret.OnSessionQueryExplain = func(t TableExplainQueryStartInfo) func(TableExplainQueryDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableExplainQueryDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableExplainQueryDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionQueryStreamExecute
		h2 := x.OnSessionQueryStreamExecute
		ret.OnSessionQueryStreamExecute = func(t TableSessionQueryStreamExecuteStartInfo) func(TableSessionQueryStreamExecuteDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionQueryStreamExecuteDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionQueryStreamExecuteDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnSessionQueryStreamRead
		h2 := x.OnSessionQueryStreamRead
		ret.OnSessionQueryStreamRead = func(t TableSessionQueryStreamReadStartInfo) func(TableSessionQueryStreamReadDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableSessionQueryStreamReadDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableSessionQueryStreamReadDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnTxBegin
		h2 := x.OnTxBegin
		ret.OnTxBegin = func(t TableTxBeginStartInfo) func(TableTxBeginDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableTxBeginDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableTxBeginDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnTxExecute
		h2 := x.OnTxExecute
		ret.OnTxExecute = func(t TableTransactionExecuteStartInfo) func(TableTransactionExecuteDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableTransactionExecuteDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableTransactionExecuteDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnTxExecuteStatement
		h2 := x.OnTxExecuteStatement
		ret.OnTxExecuteStatement = func(t TableTransactionExecuteStatementStartInfo) func(TableTransactionExecuteStatementDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableTransactionExecuteStatementDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableTransactionExecuteStatementDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnTxCommit
		h2 := x.OnTxCommit
		ret.OnTxCommit = func(t TableTxCommitStartInfo) func(TableTxCommitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableTxCommitDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableTxCommitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnTxRollback
		h2 := x.OnTxRollback
		ret.OnTxRollback = func(t TableTxRollbackStartInfo) func(TableTxRollbackDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TableTxRollbackDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TableTxRollbackDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnPoolStateChange
		h2 := x.OnPoolStateChange
		ret.OnPoolStateChange = func(t TablePoolStateChangeInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(t)
			}
			if h2 != nil {
				h2(t)
			}
		}
	}
	{
		h1 := t.OnPoolSessionAdd
		h2 := x.OnPoolSessionAdd
		ret.OnPoolSessionAdd = func(info TablePoolSessionAddInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(info)
			}
			if h2 != nil {
				h2(info)
			}
		}
	}
	{
		h1 := t.OnPoolSessionRemove
		h2 := x.OnPoolSessionRemove
		ret.OnPoolSessionRemove = func(info TablePoolSessionRemoveInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			if h1 != nil {
				h1(info)
			}
			if h2 != nil {
				h2(info)
			}
		}
	}
	{
		h1 := t.OnPoolPut
		h2 := x.OnPoolPut
		ret.OnPoolPut = func(t TablePoolPutStartInfo) func(TablePoolPutDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TablePoolPutDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TablePoolPutDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnPoolGet
		h2 := x.OnPoolGet
		ret.OnPoolGet = func(t TablePoolGetStartInfo) func(TablePoolGetDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TablePoolGetDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TablePoolGetDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	{
		h1 := t.OnPoolWait
		h2 := x.OnPoolWait
		ret.OnPoolWait = func(t TablePoolWaitStartInfo) func(TablePoolWaitDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(TablePoolWaitDoneInfo)
			if h1 != nil {
				r = h1(t)
			}
			if h2 != nil {
				r1 = h2(t)
			}
			return func(t TablePoolWaitDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(t)
				}
				if r1 != nil {
					r1(t)
				}
			}
		}
	}
	return &ret
}
func (t *Table) onInit(t1 TableInitStartInfo) func(TableInitDoneInfo) {
	fn := t.OnInit
	if fn == nil {
		return func(TableInitDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableInitDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onClose(t1 TableCloseStartInfo) func(TableCloseDoneInfo) {
	fn := t.OnClose
	if fn == nil {
		return func(TableCloseDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onDo(t1 TableDoStartInfo) func(TableDoDoneInfo) {
	fn := t.OnDo
	if fn == nil {
		return func(TableDoDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableDoDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onDoTx(t1 TableDoTxStartInfo) func(TableDoTxDoneInfo) {
	fn := t.OnDoTx
	if fn == nil {
		return func(TableDoTxDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableDoTxDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onCreateSession(t1 TableCreateSessionStartInfo) func(TableCreateSessionDoneInfo) {
	fn := t.OnCreateSession
	if fn == nil {
		return func(TableCreateSessionDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableCreateSessionDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionNew(t1 TableSessionNewStartInfo) func(TableSessionNewDoneInfo) {
	fn := t.OnSessionNew
	if fn == nil {
		return func(TableSessionNewDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionNewDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionDelete(t1 TableSessionDeleteStartInfo) func(TableSessionDeleteDoneInfo) {
	fn := t.OnSessionDelete
	if fn == nil {
		return func(TableSessionDeleteDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionDeleteDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionKeepAlive(t1 TableKeepAliveStartInfo) func(TableKeepAliveDoneInfo) {
	fn := t.OnSessionKeepAlive
	if fn == nil {
		return func(TableKeepAliveDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableKeepAliveDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionBulkUpsert(t1 TableBulkUpsertStartInfo) func(TableBulkUpsertDoneInfo) {
	fn := t.OnSessionBulkUpsert
	if fn == nil {
		return func(TableBulkUpsertDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableBulkUpsertDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionQueryPrepare(t1 TablePrepareDataQueryStartInfo) func(TablePrepareDataQueryDoneInfo) {
	fn := t.OnSessionQueryPrepare
	if fn == nil {
		return func(TablePrepareDataQueryDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TablePrepareDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionQueryExecute(t1 TableExecuteDataQueryStartInfo) func(TableExecuteDataQueryDoneInfo) {
	fn := t.OnSessionQueryExecute
	if fn == nil {
		return func(TableExecuteDataQueryDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableExecuteDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionQueryExplain(t1 TableExplainQueryStartInfo) func(TableExplainQueryDoneInfo) {
	fn := t.OnSessionQueryExplain
	if fn == nil {
		return func(TableExplainQueryDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableExplainQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionQueryStreamExecute(t1 TableSessionQueryStreamExecuteStartInfo) func(TableSessionQueryStreamExecuteDoneInfo) {
	fn := t.OnSessionQueryStreamExecute
	if fn == nil {
		return func(TableSessionQueryStreamExecuteDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionQueryStreamExecuteDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onSessionQueryStreamRead(t1 TableSessionQueryStreamReadStartInfo) func(TableSessionQueryStreamReadDoneInfo) {
	fn := t.OnSessionQueryStreamRead
	if fn == nil {
		return func(TableSessionQueryStreamReadDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableSessionQueryStreamReadDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onTxBegin(t1 TableTxBeginStartInfo) func(TableTxBeginDoneInfo) {
	fn := t.OnTxBegin
	if fn == nil {
		return func(TableTxBeginDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableTxBeginDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onTxExecute(t1 TableTransactionExecuteStartInfo) func(TableTransactionExecuteDoneInfo) {
	fn := t.OnTxExecute
	if fn == nil {
		return func(TableTransactionExecuteDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableTransactionExecuteDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onTxExecuteStatement(t1 TableTransactionExecuteStatementStartInfo) func(TableTransactionExecuteStatementDoneInfo) {
	fn := t.OnTxExecuteStatement
	if fn == nil {
		return func(TableTransactionExecuteStatementDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableTransactionExecuteStatementDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onTxCommit(t1 TableTxCommitStartInfo) func(TableTxCommitDoneInfo) {
	fn := t.OnTxCommit
	if fn == nil {
		return func(TableTxCommitDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableTxCommitDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onTxRollback(t1 TableTxRollbackStartInfo) func(TableTxRollbackDoneInfo) {
	fn := t.OnTxRollback
	if fn == nil {
		return func(TableTxRollbackDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TableTxRollbackDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onPoolStateChange(t1 TablePoolStateChangeInfo) {
	fn := t.OnPoolStateChange
	if fn == nil {
		return
	}
	fn(t1)
}
func (t *Table) onPoolSessionAdd(info TablePoolSessionAddInfo) {
	fn := t.OnPoolSessionAdd
	if fn == nil {
		return
	}
	fn(info)
}
func (t *Table) onPoolSessionRemove(info TablePoolSessionRemoveInfo) {
	fn := t.OnPoolSessionRemove
	if fn == nil {
		return
	}
	fn(info)
}
func (t *Table) onPoolPut(t1 TablePoolPutStartInfo) func(TablePoolPutDoneInfo) {
	fn := t.OnPoolPut
	if fn == nil {
		return func(TablePoolPutDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TablePoolPutDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onPoolGet(t1 TablePoolGetStartInfo) func(TablePoolGetDoneInfo) {
	fn := t.OnPoolGet
	if fn == nil {
		return func(TablePoolGetDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TablePoolGetDoneInfo) {
			return
		}
	}
	return res
}
func (t *Table) onPoolWait(t1 TablePoolWaitStartInfo) func(TablePoolWaitDoneInfo) {
	fn := t.OnPoolWait
	if fn == nil {
		return func(TablePoolWaitDoneInfo) {
			return
		}
	}
	res := fn(t1)
	if res == nil {
		return func(TablePoolWaitDoneInfo) {
			return
		}
	}
	return res
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnInit(t *Table, c *context.Context, call call) func(limit int) {
	var p TableInitStartInfo
	p.Context = c
	p.Call = call
	res := t.onInit(p)
	return func(limit int) {
		var p TableInitDoneInfo
		p.Limit = limit
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnClose(t *Table, c *context.Context, call call) func(error) {
	var p TableCloseStartInfo
	p.Context = c
	p.Call = call
	res := t.onClose(p)
	return func(e error) {
		var p TableCloseDoneInfo
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnDo(t *Table, c *context.Context, call call, label string, idempotent bool, nestedCall bool) func(attempts int, _ error) {
	var p TableDoStartInfo
	p.Context = c
	p.Call = call
	p.Label = label
	p.Idempotent = idempotent
	p.NestedCall = nestedCall
	res := t.onDo(p)
	return func(attempts int, e error) {
		var p TableDoDoneInfo
		p.Attempts = attempts
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnDoTx(t *Table, c *context.Context, call call, label string, idempotent bool, nestedCall bool) func(attempts int, _ error) {
	var p TableDoTxStartInfo
	p.Context = c
	p.Call = call
	p.Label = label
	p.Idempotent = idempotent
	p.NestedCall = nestedCall
	res := t.onDoTx(p)
	return func(attempts int, e error) {
		var p TableDoTxDoneInfo
		p.Attempts = attempts
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnCreateSession(t *Table, c *context.Context, call call) func(session sessionInfo, attempts int, _ error) {
	var p TableCreateSessionStartInfo
	p.Context = c
	p.Call = call
	res := t.onCreateSession(p)
	return func(session sessionInfo, attempts int, e error) {
		var p TableCreateSessionDoneInfo
		p.Session = session
		p.Attempts = attempts
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnSessionNew(t *Table, c *context.Context, call call) func(session sessionInfo, _ error) {
	var p TableSessionNewStartInfo
	p.Context = c
	p.Call = call
	res := t.onSessionNew(p)
	return func(session sessionInfo, e error) {
		var p TableSessionNewDoneInfo
		p.Session = session
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnSessionDelete(t *Table, c *context.Context, call call, session sessionInfo) func(error) {
	var p TableSessionDeleteStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	res := t.onSessionDelete(p)
	return func(e error) {
		var p TableSessionDeleteDoneInfo
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnSessionKeepAlive(t *Table, c *context.Context, call call, session sessionInfo) func(error) {
	var p TableKeepAliveStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	res := t.onSessionKeepAlive(p)
	return func(e error) {
		var p TableKeepAliveDoneInfo
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnSessionBulkUpsert(t *Table, c *context.Context, call call, session sessionInfo) func(error) {
	var p TableBulkUpsertStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	res := t.onSessionBulkUpsert(p)
	return func(e error) {
		var p TableBulkUpsertDoneInfo
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnSessionQueryPrepare(t *Table, c *context.Context, call call, session sessionInfo, query string) func(result tableDataQuery, _ error) {
	var p TablePrepareDataQueryStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	p.Query = query
	res := t.onSessionQueryPrepare(p)
	return func(result tableDataQuery, e error) {
		var p TablePrepareDataQueryDoneInfo
		p.Result = result
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnSessionQueryExecute(t *Table, c *context.Context, call call, session sessionInfo, query tableDataQuery, parameters tableQueryParameters, keepInCache bool) func(tx txInfo, prepared bool, result tableResult, _ error) {
	var p TableExecuteDataQueryStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	p.Query = query
	p.Parameters = parameters
	p.KeepInCache = keepInCache
	res := t.onSessionQueryExecute(p)
	return func(tx txInfo, prepared bool, result tableResult, e error) {
		var p TableExecuteDataQueryDoneInfo
		p.Tx = tx
		p.Prepared = prepared
		p.Result = result
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnSessionQueryExplain(t *Table, c *context.Context, call call, session sessionInfo, query string) func(aST string, plan string, _ error) {
	var p TableExplainQueryStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	p.Query = query
	res := t.onSessionQueryExplain(p)
	return func(aST string, plan string, e error) {
		var p TableExplainQueryDoneInfo
		p.AST = aST
		p.Plan = plan
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnSessionQueryStreamExecute(t *Table, c *context.Context, call call, session sessionInfo, query tableDataQuery, parameters tableQueryParameters) func(error) {
	var p TableSessionQueryStreamExecuteStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	p.Query = query
	p.Parameters = parameters
	res := t.onSessionQueryStreamExecute(p)
	return func(e error) {
		var p TableSessionQueryStreamExecuteDoneInfo
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnSessionQueryStreamRead(t *Table, c *context.Context, call call, session sessionInfo) func(error) {
	var p TableSessionQueryStreamReadStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	res := t.onSessionQueryStreamRead(p)
	return func(e error) {
		var p TableSessionQueryStreamReadDoneInfo
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnTxBegin(t *Table, c *context.Context, call call, session sessionInfo) func(tx txInfo, _ error) {
	var p TableTxBeginStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	res := t.onTxBegin(p)
	return func(tx txInfo, e error) {
		var p TableTxBeginDoneInfo
		p.Tx = tx
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnTxExecute(t *Table, c *context.Context, call call, session sessionInfo, tx txInfo, query tableDataQuery, parameters tableQueryParameters) func(result tableResult, _ error) {
	var p TableTransactionExecuteStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	p.Tx = tx
	p.Query = query
	p.Parameters = parameters
	res := t.onTxExecute(p)
	return func(result tableResult, e error) {
		var p TableTransactionExecuteDoneInfo
		p.Result = result
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnTxExecuteStatement(t *Table, c *context.Context, call call, session sessionInfo, tx txInfo, statementQuery tableDataQuery, parameters tableQueryParameters) func(result tableResult, _ error) {
	var p TableTransactionExecuteStatementStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	p.Tx = tx
	p.StatementQuery = statementQuery
	p.Parameters = parameters
	res := t.onTxExecuteStatement(p)
	return func(result tableResult, e error) {
		var p TableTransactionExecuteStatementDoneInfo
		p.Result = result
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnTxCommit(t *Table, c *context.Context, call call, session sessionInfo, tx txInfo) func(error) {
	var p TableTxCommitStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	p.Tx = tx
	res := t.onTxCommit(p)
	return func(e error) {
		var p TableTxCommitDoneInfo
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnTxRollback(t *Table, c *context.Context, call call, session sessionInfo, tx txInfo) func(error) {
	var p TableTxRollbackStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	p.Tx = tx
	res := t.onTxRollback(p)
	return func(e error) {
		var p TableTxRollbackDoneInfo
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnPoolStateChange(t *Table, size int, event string) {
	var p TablePoolStateChangeInfo
	p.Size = size
	p.Event = event
	t.onPoolStateChange(p)
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnPoolSessionAdd(t *Table, session sessionInfo) {
	var p TablePoolSessionAddInfo
	p.Session = session
	t.onPoolSessionAdd(p)
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnPoolSessionRemove(t *Table, session sessionInfo) {
	var p TablePoolSessionRemoveInfo
	p.Session = session
	t.onPoolSessionRemove(p)
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnPoolPut(t *Table, c *context.Context, call call, session sessionInfo) func(error) {
	var p TablePoolPutStartInfo
	p.Context = c
	p.Call = call
	p.Session = session
	res := t.onPoolPut(p)
	return func(e error) {
		var p TablePoolPutDoneInfo
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnPoolGet(t *Table, c *context.Context, call call) func(session sessionInfo, attempts int, _ error) {
	var p TablePoolGetStartInfo
	p.Context = c
	p.Call = call
	res := t.onPoolGet(p)
	return func(session sessionInfo, attempts int, e error) {
		var p TablePoolGetDoneInfo
		p.Session = session
		p.Attempts = attempts
		p.Error = e
		res(p)
	}
}
// Internals: https://github.com/ydb-platform/ydb-go-sdk/blob/master/VERSIONING.md#internals
func TableOnPoolWait(t *Table, c *context.Context, call call) func(session sessionInfo, _ error) {
	var p TablePoolWaitStartInfo
	p.Context = c
	p.Call = call
	res := t.onPoolWait(p)
	return func(session sessionInfo, e error) {
		var p TablePoolWaitDoneInfo
		p.Session = session
		p.Error = e
		res(p)
	}
}
