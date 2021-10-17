package handler

import (
	"book-ms/core"
	"book-ms/server/dbserver"
	"book-ms/server/grpcserver"
	"book-ms/values"
	"context"
	"gomodules/errormodule"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
	if !strings.HasPrefix(request.URL.Path, "/api") {
		return nil, errormodule.ErrInvalidAPICall
	}

	uri := request.URL.Path[4:]

	ctx = dbserver.PrepareDbRunner(ctx)

	userRole, err := grpcserver.AuthorizeUser(ctx, request.Authorization)
	if err != nil {
		return nil, errormodule.ErrNotAuthenticated
	}

	switch {
	case strings.HasPrefix(uri, "/member"):
		if userRole != values.UserRoleMember {
			return nil, errormodule.ErrNotAuthenticated
		}

		return handleMember(ctx, uri[7:], request)
	case strings.HasPrefix(uri, "/librarian"):
		if userRole != values.UserRoleLibrarian {
			return nil, errormodule.ErrNotAuthenticated
		}

		return handleLibrarian(ctx, uri[10:], request)
	default:
		return nil, errormodule.ErrInvalidAPICall
	}
}

func handleMember(ctx context.Context, uri string, request *Request) (response interface{}, err error) {
	if !strings.HasPrefix(uri, "/book") {
		return nil, errormodule.ErrInvalidAPICall
	}

	uri = uri[5:]

	switch request.Method {
	case http.MethodGet:
		if uri == "" {
			return nil, errormodule.ErrInvalidAPICall
		}

		if strings.HasPrefix(uri, "/all") {
			searchTerm, rowOffset, rowLimit, err := getParams(request.URL)
			if err != nil {
				return nil, errormodule.ErrInvalidAPICall
			}

			return core.GetAllBooks(ctx, searchTerm, rowOffset, rowLimit, values.UserRoleMember)
		}

		return core.GetBook(ctx, uri[1:])
	case http.MethodPatch:
		return nil, core.BorrowOrReturnBook(ctx, request.Authorization, request.Body)
	default:
		return nil, errormodule.ErrInvalidAPICall
	}
}

func handleLibrarian(ctx context.Context, uri string, request *Request) (response interface{}, err error) {
	if !strings.HasPrefix(uri, "/book") {
		return nil, errormodule.ErrInvalidAPICall
	}

	uri = uri[5:]

	switch request.Method {
	case http.MethodPost:
		return core.CreateBook(ctx, request.Body)
	case http.MethodGet:
		if uri == "" {
			return nil, errormodule.ErrInvalidAPICall
		}

		if strings.HasPrefix(uri, "/all") {
			searchTerm, rowOffset, rowLimit, err := getParams(request.URL)
			if err != nil {
				return nil, errormodule.ErrInvalidAPICall
			}

			return core.GetAllBooks(ctx, searchTerm, rowOffset, rowLimit, values.UserRoleLibrarian)
		}

		return core.GetBook(ctx, uri[1:])
	case http.MethodPut:
		return core.UpdateBook(ctx, request.Body)
	case http.MethodDelete:
		if uri == "" {
			return nil, errormodule.ErrInvalidAPICall
		}

		return nil, core.DeleteBook(ctx, uri[1:])
	default:
		return nil, errormodule.ErrInvalidAPICall
	}
}

func getParams(uri *url.URL) (searchTerm string, rowOffset, rowLimit int, err error) {
	params := uri.Query()

	param, ok := params["searchTerm"]
	if ok {
		searchTerm = param[0]
	}

	param, ok = params["offset"]
	if ok {
		rowOffset, err = strconv.Atoi(param[0])
		if err != nil {
			return
		}
	}

	param, ok = params["limit"]
	if ok {
		rowLimit, err = strconv.Atoi(param[0])
		if err != nil {
			return
		}
	}

	return
}
