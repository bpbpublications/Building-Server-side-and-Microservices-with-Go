package core

import (
	"context"
	"gomodules/errormodule"
	"report-ms/data"
	"time"

	"cloud.google.com/go/civil"
)

var (
	// GetDailyReport returns how many books are borrowed by day
	GetDailyReport = getDailyReport

	// GetBooksReport returns how many time each book has been borrowed
	GetBooksReport = getBooksReport
)

type getReportResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

func getDailyReport(ctx context.Context, startDate, endDate civil.Date) (response interface{}, err error) {
	today := civil.DateOf(time.Now())

	if startDate.After(today) {
		cause := "Invalid value for start date"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	if endDate.After(today) {
		cause := "Invalid value for end date"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	if startDate.After(endDate) {
		cause := "Start date must be before end date"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	start := startDate.In(time.UTC)
	end := endDate.In(time.UTC)

	report, err := data.GetDailyReport(ctx, start, end)
	if err != nil {
		cause := "Failed to get daily report"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	type metaData struct {
		StartDate civil.Date
		EndDate   civil.Date
	}

	meta := &metaData{
		StartDate: startDate,
		EndDate:   endDate,
	}

	response = &getReportResponse{
		Data: report,
		Meta: meta,
	}

	return
}

func getBooksReport(ctx context.Context) (response interface{}, err error) {
	report, err := data.GetBooksReport(ctx)
	if err != nil {
		cause := "Failed to get books report"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	response = &getReportResponse{
		Data: report,
	}

	return
}
