package data

import (
	"building-restful-web-services-with-go/chapter2/library/server/dbserver"
	"building-restful-web-services-with-go/chapter2/library/util"
	"building-restful-web-services-with-go/chapter2/library/values"
	"context"
	"time"
)

var (
	// CreateBook creates new Book
	CreateBook = createBook

	// GetBook returns book
	GetBook = getBook

	// GetAllBooksForMember returns list of books for memeber
	GetAllBooksForMember = getAllBooksForMember

	// GetAllBooksForLibrerian returns list of books for librerian
	GetAllBooksForLibrerian = getAllBooksForLibrerian

	// UpdateBook updates book
	UpdateBook = updateBook

	// DeleteBook deletes book
	DeleteBook = deleteBook

	// GetBookStatus returns book status
	GetBookStatus = getBookStatus

	// ChangeBookStatus changes book status
	ChangeBookStatus = changeBookStatus
)

func createBook(ctx context.Context, bookName, authorName, publisher string, description util.NullString) (response *BookEntity, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(dbserver.Runner)

	query := `
		insert into book(book_name, author_name, publisher, book_description)
		values ($1, $2, $3, $4)
		returning book_id, created_at`

	rows, err := dbRunner.Query(ctx, query, bookName, authorName, publisher, description)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := dbserver.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		response = &BookEntity{}
		response.BookID = rr.ReadByIdxString(0)
		response.BookName = bookName
		response.AuthorName = authorName
		response.Publisher = publisher
		response.Description = util.GetNullStringValue(description)
		response.Status = values.BookStatusAvailable
		response.CreatedAt = rr.ReadByIdxTime(1)
		response.UpdatedAt = rr.ReadByIdxTime(1)
		response.BorrowerID = ""
	}

	err = rr.Error()

	return
}

func getBook(ctx context.Context, bookID string) (response *BookDetails, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(dbserver.Runner)

	query := `
		select
			book_id as "BookID",
			book_name as "BookName",
			author_name as "AuthorName",
			publisher as "Publisher",
			book_description as "Description"
		from book
		where book_id = $1`

	rows, err := dbRunner.Query(ctx, query, bookID)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := dbserver.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		response = &BookDetails{}
		rr.ReadAllToStruct(response)
	}

	err = rr.Error()

	return
}

func getAllBooksForMember(ctx context.Context, searchTerm string, rowOffset, rowLimit int) (response []*BookInfoMember, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(dbserver.Runner)

	query := `
		select
			book_id as "BookID",
			book_name as "BookName",
			author_name as "AuthorName",
			publisher as "Publisher"
		from book
		where book_name like '%%' || $1 || '%%' and book_status = $2
		offset $3
		limit $4`

	rows, err := dbRunner.Query(ctx, query, searchTerm, values.BookStatusAvailable, rowOffset, rowLimit)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := dbserver.GetRowReader(rows)
	if err != nil {
		return
	}

	response = make([]*BookInfoMember, 0)
	for rr.ScanNext() {
		book := &BookInfoMember{}
		rr.ReadAllToStruct(book)
		response = append(response, book)
	}

	err = rr.Error()

	return
}

func getAllBooksForLibrerian(ctx context.Context, searchTerm string, rowOffset, rowLimit int) (response []*BookInfoLibrarian, err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(dbserver.Runner)

	query := `
		select
			b.book_id as "BookID",
			b.book_name as "BookName",
			b.author_name as "AuthorName",
			b.publisher as "Publisher",
			b.book_status as "Status",
			u.full_name as "Borrower"
		from book b
		left join library_user u on u.user_id = b.borrower_id
		where b.book_name like '%%' || $1 || '%%'
		offset $2
		limit $3`

	rows, err := dbRunner.Query(ctx, query, searchTerm, rowOffset, rowLimit)
	if err != nil {
		return
	}

	defer rows.Close()

	rr, err := dbserver.GetRowReader(rows)
	if err != nil {
		return
	}

	response = make([]*BookInfoLibrarian, 0)
	for rr.ScanNext() {
		book := &BookInfoLibrarian{}
		rr.ReadAllToStruct(book)
		response = append(response, book)
	}

	err = rr.Error()

	return
}

func updateBook(ctx context.Context, bookID, bookName, authorName, publisher string, description util.NullString) (response time.Time, err error) {
	query := `
		update book
		set
			book_name = $1,
			author_name = $2,
			publisher = $3,
			book_description = $4
		where book_id = $5
		returning updated_at`

	return executeQueryWithTimeResponse(ctx, query, bookName, authorName, publisher, description, bookID)
}

func deleteBook(ctx context.Context, bookID string) (response int64, err error) {
	query := `delete from book where book_id = $1`

	return executeQueryWithRowsAffected(ctx, query, bookID)
}

func getBookStatus(ctx context.Context, bookID string) (response int64, err error) {
	query := `select book_status from book where book_id  = $1`

	return executeQueryWithInt64Response(ctx, query, bookID)
}

func changeBookStatus(ctx context.Context, bookID string, status int, userID util.NullString) (err error) {
	dbRunner := ctx.Value(values.ContextKeyDbRunner).(dbserver.Runner)

	query := `
		update book
		set
			book_status = $1,
			borrower_id = $2
		where book_id = $3`

	_, err = dbRunner.Exec(ctx, query, status, userID, bookID)

	return
}
