package data

import (
	"context"
	"gomodules/databasemodule"
	"report-ms/values"
	"time"
)

var (
	// CreateReport will insert new row into report table
	CreateReport = createReport

	// GetDailyReport returns how many books are borrowed by day
	GetDailyReport = getDailyReport

	// GetBooksReport returns how many time each book has been borrowed
	GetBooksReport = getBooksReport
)

func createReport(ctx context.Context, bookID string) (err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(databasemodule.Runner)

	query := `insert into report(book_id) values ($1)`

	_, err = dbRunner.Exec(ctx, query, bookID)

	return
}

func getDailyReport(ctx context.Context, startDate, endDate time.Time) (response []*DailyReportData, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(databasemodule.Runner)

	query := `
		select
			report_date::timestamp::date::text as "ReportDate",
			count(*) as "BookCount"
		from report
		where report_date >= $1 and report_date <= $2
		group by report_date::timestamp::date::text`

	rows, err := dbRunner.Query(ctx, query, startDate, endDate)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := databasemodule.GetRowReader(rows)
	if err != nil {
		return
	}

	response = make([]*DailyReportData, 0)
	for rr.ScanNext() {
		report := &DailyReportData{}
		rr.ReadAllToStruct(report)
		response = append(response, report)
	}

	err = rr.Error()

	return
}

func getBooksReport(ctx context.Context) (response []*BookReportData, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(databasemodule.Runner)

	query := `
		select
			book_id as "BookID",
			count(*) as "BookCount"
		from report
		group by book_id`

	rows, err := dbRunner.Query(ctx, query)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := databasemodule.GetRowReader(rows)
	if err != nil {
		return
	}

	response = make([]*BookReportData, 0)
	for rr.ScanNext() {
		report := &BookReportData{}
		rr.ReadAllToStruct(report)
		response = append(response, report)
	}

	err = rr.Error()

	return
}
