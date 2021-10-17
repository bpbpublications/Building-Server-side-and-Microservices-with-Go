package data

import (
	"book-ms/values"
	"context"
	"gomodules/databasemodule"
	"time"
)

func executeQueryWithTimeResponse(ctx context.Context, query string, params ...interface{}) (result time.Time, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(databasemodule.Runner)

	rows, err := dbRunner.Query(ctx, query, params...)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := databasemodule.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		result = rr.ReadByIdxTime(0)
	}

	err = rr.Error()

	return
}

func executeQueryWithRowsAffected(ctx context.Context, query string, params ...interface{}) (result int64, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(databasemodule.Runner)

	res, err := dbRunner.Exec(ctx, query, params...)
	if err != nil {
		return
	}

	result, err = res.RowsAffected()

	return
}

func executeQueryWithInt64Response(ctx context.Context, query string, params ...interface{}) (result int64, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(databasemodule.Runner)

	rows, err := dbRunner.Query(ctx, query, params...)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := databasemodule.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		result = rr.ReadByIdxInt64(0)
	}

	err = rr.Error()

	return
}
