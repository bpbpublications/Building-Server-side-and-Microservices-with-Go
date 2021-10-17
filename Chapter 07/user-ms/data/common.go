package data

import (
	"context"
	"gomodules/databasemodule"
	"user-ms/values"
)

func executeQueryWithStringResponse(ctx context.Context, query string, params ...interface{}) (result string, err error) {
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
		result = rr.ReadByIdxString(0)
	}

	err = rr.Error()

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
