package core

import (
	"building-restful-web-services-with-go/chapter2/library/data"
	"building-restful-web-services-with-go/chapter2/library/util"
	"building-restful-web-services-with-go/chapter2/library/values"
	"context"
	"encoding/json"
	"io"
	"strings"
)

var (
	// Login will return user's auth token
	Login = login

	// AuthorizeUser returns user's role if user token exists
	// if token not exists, empty string will be returned
	AuthorizeUser = authorizeUser
)

func login(ctx context.Context, requestBody io.Reader) (response interface{}, err error) {
	type loginRequest struct {
		Username string
		Password string
	}

	request := &loginRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		cause := "Failed to decode JSON"
		err = util.NewError(cause, util.ErrorCodeInvalidJSONBody, util.ErrBadRequest, err)
		return
	}

	request.Username = strings.TrimSpace(request.Username)
	request.Password = strings.TrimSpace(request.Password)
	if request.Username == "" || request.Password == "" {
		cause := "Username or password are empty"
		err = util.NewError(cause, util.ErrorCodeValidation, util.ErrBadRequest, err)
		return
	}

	token, err := data.LoginUser(ctx, request.Username, request.Password)
	if err != nil {
		cause := "Failed to login user"
		err = util.NewError(cause, util.ErrorCodeInternal, util.ErrInternal, err)
		return
	}

	if token == "" {
		cause := "Invalid username or password"
		err = util.NewError(cause, util.ErrorCodeInvalidCredentials, util.ErrNotAuthenticated, err)
		return
	}

	type loginResponse struct {
		Token string
	}

	response = &loginResponse{
		Token: token,
	}

	return
}

func authorizeUser(ctx context.Context, token string) (response int, err error) {
	token = strings.TrimSpace(token)
	if token == "" {
		cause := "Invalid value for token parameter"
		err = util.NewError(cause, util.ErrorCodeValidation, util.ErrBadRequest, err)
		return
	}

	userRole, err := data.AuthorizeUser(ctx, token)
	if err != nil {
		cause := "Failed to authorize user"
		err = util.NewError(cause, util.ErrorCodeInternal, util.ErrInternal, err)
		return
	}

	if userRole == values.UserRoleUnknown {
		cause := "User not found"
		err = util.NewError(cause, util.ErrorCodeEntityNotFound, util.ErrResourceNotFound, err)
		return
	}

	response = int(userRole)

	return
}
