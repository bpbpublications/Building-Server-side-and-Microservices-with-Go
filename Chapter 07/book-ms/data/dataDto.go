package data

import (
	"time"
)

// BookEntity is struct used to describe book
// This struct contains all database columns converted to Go types
type BookEntity struct {
	BookID      string
	BookName    string
	AuthorName  string
	Publisher   string
	Description string `json:",omitempty"`
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BorrowerID  string `json:",omitempty"`
}

// BookDetails is struct used to describe book
// This struct is used in get single book API
type BookDetails struct {
	BookID      string
	BookName    string
	AuthorName  string
	Publisher   string
	Description string `json:",omitempty"`
}

// BookInfoLibrarian is struct used to describe book
// This structure is used in librarian get all book API
type BookInfoLibrarian struct {
	BookID     string
	BookName   string
	AuthorName string
	Publisher  string
	Status     int64
	BorrowerID string `json:"-"`
	Borrower   string `json:",omitempty"`
}

// BookInfoMember is struct used to describe book
// This structure is used in member get all book API
type BookInfoMember struct {
	BookID     string
	BookName   string
	AuthorName string
	Publisher  string
}
