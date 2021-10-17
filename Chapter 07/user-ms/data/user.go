package data

import "context"

var (
	// LoginUser find user with provided username and password
	// and returns user's token
	LoginUser = loginUser

	// AuthorizeUser returns user's role if user token exists
	// if token not exists, empty string will be returned
	AuthorizeUser = authorizeUser

	// GetUserID returns userID for provided token
	GetUserID = getUserID

	// GetUserFullName returns full name of user
	GetUserFullName = getUserFullName
)

func loginUser(ctx context.Context, username, password string) (response string, err error) {
	query := `
		select token
		from library_user
		where username = $1 and user_password = crypt($2, user_password)`

	return executeQueryWithStringResponse(ctx, query, username, password)
}

func authorizeUser(ctx context.Context, token string) (response int64, err error) {
	query := `
		select user_role
		from library_user
		where token = $1`

	return executeQueryWithInt64Response(ctx, query, token)
}

func getUserID(ctx context.Context, token string) (response string, err error) {
	query := `
		select user_id
		from library_user
		where token = $1`

	return executeQueryWithStringResponse(ctx, query, token)
}

func getUserFullName(ctx context.Context, userID string) (response string, err error) {
	query := `
		select full_name
		from library_user
		where user_id = $1`

	return executeQueryWithStringResponse(ctx, query, userID)
}
