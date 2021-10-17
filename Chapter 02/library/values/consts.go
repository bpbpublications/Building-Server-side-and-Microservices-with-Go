package values

// Values for user role
const (
	UserRoleUnknown   = 0
	UserRoleMember    = 1
	UserRoleLibrarian = 2
)

// Values for book status
const (
	BookStatusUnknown   = 0
	BookStatusAvailable = 1
	BookStatusBorrowed  = 2
)

// MaxRowLimit is max row limit for get all API
const MaxRowLimit = 1000
