package databasemodule

import "database/sql"

var (
	// NewNullableString return a null String if the parameter is empty,
	// a valid String otherwise
	NewNullableString = newNullableString

	// GetNullStringValue returns empty string if null.String is null,
	// a valid String othervise
	GetNullStringValue = getNullStringValue
)

// NullString represents nullable string.
// Intended for use in fetching data from SQL queries
type NullString struct {
	sql.NullString
}

func newNullableString(x string) NullString {
	if x == "" {
		return NullString{}
	}

	return NullString{sql.NullString{String: x, Valid: true}}
}

func getNullStringValue(x NullString) string {
	if x.Valid == false {
		return ""
	}

	return x.String
}
