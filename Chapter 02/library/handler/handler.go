package handler

import (
	"building-restful-web-services-with-go/chapter2/library/core"
	"building-restful-web-services-with-go/chapter2/library/server/dbserver"
	"building-restful-web-services-with-go/chapter2/library/util"
	"building-restful-web-services-with-go/chapter2/library/values"
	"context"
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
		return nil, util.ErrInvalidAPICall
	}

	uri := request.URL.Path[4:]

	ctx = dbserver.PrepareDbRunner(ctx)

	switch {
	case strings.HasPrefix(uri, "/open"):
		return handleOpen(ctx, uri[5:], request)
	case strings.HasPrefix(uri, "/member"):
		userRole, err := core.AuthorizeUser(ctx, request.Authorization)
		if err != nil {
			return nil, util.ErrNotAuthenticated
		}

		if userRole != values.UserRoleMember {
			return nil, util.ErrNotAuthenticated
		}

		return handleMember(ctx, uri[7:], request)
	case strings.HasPrefix(uri, "/librarian"):
		userRole, err := core.AuthorizeUser(ctx, request.Authorization)
		if err != nil {
			return nil, util.ErrNotAuthenticated
		}

		if userRole != values.UserRoleLibrarian {
			return nil, util.ErrNotAuthenticated
		}

		return handleLibrarian(ctx, uri[10:], request)
	default:
		return nil, util.ErrInvalidAPICall
	}
}

func handleOpen(ctx context.Context, uri string, request *Request) (response interface{}, err error) {
	if uri == "/login" && request.Method == http.MethodPost {
		return core.Login(ctx, request.Body)
	}

	return nil, util.ErrInvalidAPICall
}

func handleMember(ctx context.Context, uri string, request *Request) (response interface{}, err error) {
	if !strings.HasPrefix(uri, "/book") {
		return nil, util.ErrInvalidAPICall
	}

	uri = uri[5:]

	switch request.Method {
	case http.MethodGet:
		if uri == "" {
			return nil, util.ErrInvalidAPICall
		}

		if strings.HasPrefix(uri, "/all") {
			searchTerm, rowOffset, rowLimit, err := getParams(request.URL)
			if err != nil {
				return nil, util.ErrInvalidAPICall
			}

			return core.GetAllBooks(ctx, searchTerm, rowOffset, rowLimit, values.UserRoleMember)
		}

		return core.GetBook(ctx, uri[1:])
	case http.MethodPatch:
		return nil, core.BorrowOrReturnBook(ctx, request.Authorization, request.Body)
	default:
		return nil, util.ErrInvalidAPICall
	}
}

func handleLibrarian(ctx context.Context, uri string, request *Request) (response interface{}, err error) {
	if !strings.HasPrefix(uri, "/book") {
		return nil, util.ErrInvalidAPICall
	}

	uri = uri[5:]

	switch request.Method {
	case http.MethodPost:
		return core.CreateBook(ctx, request.Body)
	case http.MethodGet:
		if uri == "" {
			return nil, util.ErrInvalidAPICall
		}

		if strings.HasPrefix(uri, "/all") {
			searchTerm, rowOffset, rowLimit, err := getParams(request.URL)
			if err != nil {
				return nil, util.ErrInvalidAPICall
			}

			return core.GetAllBooks(ctx, searchTerm, rowOffset, rowLimit, values.UserRoleLibrarian)
		}

		return core.GetBook(ctx, uri[1:])
	case http.MethodPut:
		return core.UpdateBook(ctx, request.Body)
	case http.MethodDelete:
		if uri == "" {
			return nil, util.ErrInvalidAPICall
		}

		return nil, core.DeleteBook(ctx, uri[1:])
	default:
		return nil, util.ErrInvalidAPICall
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
