# Librarian API

Only users with role librarian can call this API.
Request must contains proper Token in Authorization Header.

- [Create Book](#create-book)
- [Get Book](#get-book)
- [Get All Books](#get-all-books)
- [Update Book](#update-book)
- [Delete Book](#delete-book)

## Create Book

Creates new book.

Method: POST

URI: /api/librarian/book

Request:
* BookName (string)
* AuthorName (string)
* Publisher (string)
* Description (string?)

Response:
* BookID (string)
* BookName (string)
* AuthorName (string)
* Publisher (string)
* Description (string?)
* Status (int)
* CreatedAt (timestamp)
* UpdatedAt (timestamp)
* BorrowerID (string?)

## Get Book

Returns single book.

Method: GET

URI: /api/librarian/book/{book-id}

Response:
* BookID (string)
* BookName (string)
* AuthorName (string)
* Publisher (string)
* Description (string?)

## Get All Books

Returns all books that fit search criteria.
All parameters are oprional.

Method: GET

URI: /api/librarian/book/all/?searchTerm={searchTerm}&offset={offset}&limit={limit}

Response:
* Data (BookData)
* Meta (MetaData)

BookData is array of:
* BookID (string)
* BookName (string)
* AuthorName (string)
* Publisher (string)
* Status (int)
* Borrower (string?)

MetaData:
* SearchTerm (string?)
* RowOffset (int?)
* RowLimit (int)

## Update Book

Updates book.

Method: PUT

URI: /api/librarian/book

Request:
* BookID (string)
* BookName (string)
* AuthorName (string)
* Publisher (string)
* Description (string?)

Response:
* UpdatedAt (timestamp)

## Delete Book

Deletes single book.

Method: DELETE

URI: /api/librarian/book/{book-id}
