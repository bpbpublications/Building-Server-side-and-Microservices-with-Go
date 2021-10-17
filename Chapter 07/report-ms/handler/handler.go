package handler

import (
	"context"
	"gomodules/errormodule"
	"io"
	"net/http"
	"net/url"
	"report-ms/core"
	"report-ms/server/dbserver"
	"report-ms/server/grpcserver"
	"report-ms/values"
	"strings"

	"cloud.google.com/go/civil"
)

var (
	// Handle handle all http requests
	Handle = handle
)

// Request contains information about request
type Request struct {
	Authorization string
	Body          io.Reader
	URL           *url.URL
	Method        string
}

func handle(ctx context.Context, request *Request) (interface{}, error) {
	if request.Method != http.MethodGet {
		return nil, errormodule.ErrInvalidAPICall
	}

	if !strings.HasPrefix(request.URL.Path, "/api/report") {
		return nil, errormodule.ErrInvalidAPICall
	}

	uri := request.URL.Path[11:]

	ctx = dbserver.PrepareDbRunner(ctx)

	userRole, err := grpcserver.AuthorizeUser(ctx, request.Authorization)
	if err != nil {
		return nil, errormodule.ErrNotAuthenticated
	}

	if userRole != values.UserRoleLibrarian {
		return nil, errormodule.ErrNotAuthenticated
	}

	switch {
	case strings.HasPrefix(uri, "/daily"):
		startDate, endDate, err := getParams(request.URL)
		if err != nil {
			return nil, errormodule.ErrInvalidAPICall
		}

		return core.GetDailyReport(ctx, startDate, endDate)
	case strings.HasPrefix(uri, "/books"):
		return core.GetBooksReport(ctx)
	default:
		return nil, errormodule.ErrInvalidAPICall
	}
}

func getParams(uri *url.URL) (startDate, endDate civil.Date, err error) {
	params := uri.Query()
	param, ok := params["startDate"]
	if ok {
		startDate, err = civil.ParseDate(param[0])
		if err != nil {
			return
		}
	} else {
		err = errormodule.ErrBadRequest
		return
	}

	param, ok = params["endDate"]
	if ok {
		endDate, err = civil.ParseDate(param[0])
		if err != nil {
			return
		}
	} else {
		err = errormodule.ErrBadRequest
		return
	}

	return
}
