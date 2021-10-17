package core

import (
	"book-ms/data"
	"book-ms/server/grpcserver"
	"book-ms/server/rabbitmq"
	"book-ms/values"
	"context"
	"encoding/json"
	"gomodules/databasemodule"
	"gomodules/errormodule"
	"io"
	"log"
	"strings"
	"time"
)

var (
	// CreateBook creates new Book
	CreateBook = createBook

	// GetBook returns book
	GetBook = getBook

	// GetAllBooks returns list of books
	GetAllBooks = getAllBooks

	// UpdateBook updates book
	UpdateBook = updateBook

	// DeleteBook deletes book
	DeleteBook = deleteBook

	// BorrowOrReturnBook borrows book if is available
	// or returns book if is borrowed
	BorrowOrReturnBook = borrowOrReturnBook
)

func createBook(ctx context.Context, requestBody io.Reader) (response interface{}, err error) {
	type createBookRequest struct {
		BookName    string
		AuthorName  string
		Publisher   string
		Description string
	}

	request := &createBookRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		cause := "Failed to decode JSON"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInvalidJSONBody, errormodule.ErrBadRequest, err)
		return
	}

	request.BookName = strings.TrimSpace(request.BookName)
	if request.BookName == "" {
		cause := "Trying to create book with empty name"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	request.AuthorName = strings.TrimSpace(request.AuthorName)
	if request.AuthorName == "" {
		cause := "Trying to create book with empty author name"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	request.Publisher = strings.TrimSpace(request.Publisher)
	if request.Publisher == "" {
		cause := "Trying to create book with empty publisher"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	response, err = data.CreateBook(ctx, request.BookName, request.AuthorName, request.Publisher, databasemodule.NewNullableString(request.Description))
	if err != nil {
		cause := "Failed to create book"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	return
}

func getBook(ctx context.Context, bookID string) (response interface{}, err error) {
	if bookID == "" {
		cause := "Invalid value for bookID parameter"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	book, err := data.GetBook(ctx, bookID)
	if err != nil {
		cause := "Failed to get book"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	if book == nil {
		cause := "Book not found"
		err = errormodule.NewError(cause, errormodule.ErrorCodeEntityNotFound, errormodule.ErrResourceNotFound, err)
		return
	}

	response = book

	return
}

func getAllBooks(ctx context.Context, searchTerm string, rowOffset, rowLimit, userRole int) (response interface{}, err error) {
	if rowOffset < 0 {
		cause := "Invalid value for row offser parameter"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	if rowLimit < 0 || rowLimit > values.MaxRowLimit {
		cause := "Invalid value for row limit parameter"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	if rowLimit == 0 {
		rowLimit = values.MaxRowLimit
	}

	var books interface{}

	if userRole == values.UserRoleMember {
		books, err = data.GetAllBooksForMember(ctx, searchTerm, rowOffset, rowLimit)
	} else {
		books, err = getAllBooksLibrarian(ctx, searchTerm, rowOffset, rowLimit)
	}

	if err != nil {
		cause := "Failed to get all books"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	type metaData struct {
		SearchTerm string `json:",omitempty"`
		RowOffset  int    `json:",omitempty"`
		RowLimit   int
	}

	meta := &metaData{
		SearchTerm: searchTerm,
		RowOffset:  rowOffset,
		RowLimit:   rowLimit,
	}

	type getAllResponse struct {
		Data interface{} `json:"data"`
		Meta interface{} `json:"meta"`
	}

	response = &getAllResponse{
		Data: books,
		Meta: meta,
	}

	return
}

func updateBook(ctx context.Context, requestBody io.Reader) (response interface{}, err error) {
	type updateBookRequest struct {
		BookID      string
		BookName    string
		AuthorName  string
		Publisher   string
		Description string
	}

	request := &updateBookRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		cause := "Failed to decode JSON"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInvalidJSONBody, errormodule.ErrBadRequest, err)
		return
	}

	request.BookID = strings.TrimSpace(request.BookID)
	if request.BookID == "" {
		cause := "Invalid value for bookID parameter"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	request.BookName = strings.TrimSpace(request.BookName)
	if request.BookName == "" {
		cause := "Invalid value for book name"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	request.AuthorName = strings.TrimSpace(request.AuthorName)
	if request.AuthorName == "" {
		cause := "Invalid value for author name"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	request.Publisher = strings.TrimSpace(request.Publisher)
	if request.Publisher == "" {
		cause := "Invalid value for publisher"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	updatedAt, err := data.UpdateBook(ctx, request.BookID, request.BookName, request.AuthorName, request.Publisher, databasemodule.NewNullableString(request.Description))
	if err != nil {
		cause := "Failed to update book"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	if updatedAt == time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC) {
		cause := "Book not found"
		err = errormodule.NewError(cause, errormodule.ErrorCodeEntityNotFound, errormodule.ErrResourceNotFound, err)
		return
	}

	type updateBookResponse struct {
		UpdatedAt time.Time
	}

	response = &updateBookResponse{
		UpdatedAt: updatedAt,
	}

	return
}

func deleteBook(ctx context.Context, bookID string) (err error) {
	bookID = strings.TrimSpace(bookID)
	if bookID == "" {
		cause := "Invalid value for bookID parameter"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	rowsAffected, err := data.DeleteBook(ctx, bookID)
	if err != nil {
		cause := "Failed to delete book"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	if rowsAffected == 0 {
		cause := "Book not found"
		err = errormodule.NewError(cause, errormodule.ErrorCodeEntityNotFound, errormodule.ErrResourceNotFound, err)
		return
	}

	return
}

func borrowOrReturnBook(ctx context.Context, token string, requestBody io.Reader) (err error) {
	type borrowOrReturnBookRequest struct {
		BookID string
	}

	request := &borrowOrReturnBookRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		cause := "Failed to decode JSON"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInvalidJSONBody, errormodule.ErrBadRequest, err)
		return
	}

	request.BookID = strings.TrimSpace(request.BookID)
	if request.BookID == "" {
		cause := "Invalid value for bookID parameter"
		err = errormodule.NewError(cause, errormodule.ErrorCodeValidation, errormodule.ErrBadRequest, err)
		return
	}

	status, err := data.GetBookStatus(ctx, request.BookID)
	if err != nil {
		cause := "Failed to get book status"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	if status == values.BookStatusUnknown {
		cause := "Book not found"
		err = errormodule.NewError(cause, errormodule.ErrorCodeEntityNotFound, errormodule.ErrResourceNotFound, err)
		return
	}

	userUID := ""
	newStatus := values.BookStatusAvailable
	if status == values.BookStatusAvailable {
		userUID, err = grpcserver.GetUserID(ctx, token)
		if err != nil {
			return
		}

		newStatus = values.BookStatusBorrowed

		err = rabbitmq.SendMessage(request.BookID)
		if err != nil {
			// Just log error, no rollback
			log.Fatalf("Failed to send message: %v\n", err)
			err = nil
		}

	}

	err = data.ChangeBookStatus(ctx, request.BookID, newStatus, databasemodule.NewNullableString(userUID))
	if err != nil {
		cause := "Failed to change book status"
		err = errormodule.NewError(cause, errormodule.ErrorCodeInternal, errormodule.ErrInternal, err)
		return
	}

	return
}

func getAllBooksLibrarian(ctx context.Context, searchTerm string, rowOffset, rowLimit int) (books []*data.BookInfoLibrarian, err error) {
	books, err = data.GetAllBooksForLibrerian(ctx, searchTerm, rowOffset, rowLimit)
	if err != nil {
		return
	}

	fullName := ""
	for _, book := range books {
		if book.BorrowerID != "" {
			fullName, err = grpcserver.GetUserFullName(ctx, book.BorrowerID)
			if err != nil {
				return
			}

			book.Borrower = fullName
		}
	}

	return
}
