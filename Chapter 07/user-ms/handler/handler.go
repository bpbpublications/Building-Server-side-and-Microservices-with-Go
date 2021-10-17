package handler

import (
	"context"
	"gomodules/errormodule"
	"io"
	"net/http"
	"net/url"
	"strings"
	"user-ms/core"
	"user-ms/server/dbserver"
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

	switch {
	case strings.HasPrefix(uri, "/open"):
		return handleOpen(ctx, uri[5:], request)
	default:
		return nil, errormodule.ErrInvalidAPICall
	}
}

func handleOpen(ctx context.Context, uri string, request *Request) (response interface{}, err error) {
	if uri == "/login" && request.Method == http.MethodPost {
		return core.Login(ctx, request.Body)
	}

	return nil, errormodule.ErrInvalidAPICall
}
