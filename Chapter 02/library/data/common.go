package data

import (
	"building-restful-web-services-with-go/chapter2/library/server/dbserver"
	"building-restful-web-services-with-go/chapter2/library/values"
	"context"
	"time"
)

func executeQueryWithStringResponse(ctx context.Context, query string, params ...interface{}) (result string, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(dbserver.Runner)

	rows, err := dbRunner.Query(ctx, query, params...)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := dbserver.GetRowReader(rows)
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
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(dbserver.Runner)

	rows, err := dbRunner.Query(ctx, query, params...)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := dbserver.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		result = rr.ReadByIdxInt64(0)
	}

	err = rr.Error()

	return
}

func executeQueryWithRowsAffected(ctx context.Context, query string, params ...interface{}) (result int64, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(dbserver.Runner)

	res, err := dbRunner.Exec(ctx, query, params...)
	if err != nil {
		return
	}

	result, err = res.RowsAffected()

	return
}

func executeQueryWithTimeResponse(ctx context.Context, query string, params ...interface{}) (result time.Time, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(dbserver.Runner)

	rows, err := dbRunner.Query(ctx, query, params...)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := dbserver.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		result = rr.ReadByIdxTime(0)
	}

	err = rr.Error()

	return
}
