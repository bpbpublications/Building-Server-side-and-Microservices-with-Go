package core

import (
	"context"
	"encoding/json"
	"gomodules/errormodule"
	"io"
	"strings"
	"user-ms/data"
	"user-ms/server/dbserver"
)

var (
	// Login will return user's auth token
	Login = login

	// AuthorizeUser returns user's role if user token exists
	// if token not exists, empty string will be returned
	AuthorizeUser = authorizeUser

	// GetUserID returns userID for provided token
	GetUserID = getUserID

	// GetUserFullName returns full name of user
	GetUserFullName = getUserFullName
)

const UserRoleUnknown = 0

func login(ctx context.Context, requestBody io.Reader) (response interface{}, err error) {
	type loginRequest struct {
		Username string
		Password string
	}

	request := &loginRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		cause := "Failed to decode JSON"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInvalidJSONBody, errormodule.ErrBadRequest, err)
		return
	}

	request.Username = strings.TrimSpace(request.Username)
	request.Password = strings.TrimSpace(request.Password)
	if request.Username == "" || request.Password == "" {
		cause := "Username or password are empty"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	token, err := data.LoginUser(ctx, request.Username, request.Password)
	if err != nil {
		cause := "Failed to login user"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	if token == "" {
		cause := "Invalid username or password"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInvalidCredentials, errormodule.ErrNotAuthenticated, err)
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

func authorizeUser(ctx context.Context, token string) (response int64, err error) {
	ctx = dbserver.PrepareDbRunner(ctx)

	token = strings.TrimSpace(token)
	if token == "" {
		cause := "Invalid value for token parameter"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	userRole, err := data.AuthorizeUser(ctx, token)
	if err != nil {
		cause := "Failed to authorize user"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	if userRole == UserRoleUnknown {
		cause := "User not found"
		err = errormodule.NewError(cause, errormodule.ErrorCodeEntityNotFound, errormodule.ErrResourceNotFound, err)
		return
	}

	response = userRole

	return
}

func getUserID(ctx context.Context, token string) (response string, err error) {
	ctx = dbserver.PrepareDbRunner(ctx)

	token = strings.TrimSpace(token)
	if token == "" {
		cause := "Invalid value for token parameter"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	userID, err := data.GetUserID(ctx, token)
	if err != nil {
		cause := "Failed to get user"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	if userID == "" {
		cause := "User not found"
		err = errormodule.NewError(cause, errormodule.ErrorCodeEntityNotFound, errormodule.ErrResourceNotFound, err)
		return
	}

	response = userID

	return
}

func getUserFullName(ctx context.Context, userID string) (response string, err error) {
	ctx = dbserver.PrepareDbRunner(ctx)

	userID = strings.TrimSpace(userID)
	if userID == "" {
		cause := "Invalid value for userID"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	fullName, err := data.GetUserFullName(ctx, userID)
	if err != nil {
		cause := "Failed to get user's full name"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	if fullName == "" {
		cause := "User not found"
		err = errormodule.NewError(cause, errormodule.ErrorCodeEntityNotFound, errormodule.ErrResourceNotFound, err)
		return
	}

	response = fullName

	return
}
