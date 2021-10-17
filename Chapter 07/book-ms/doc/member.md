# Member API

Only users with role member can call this API.
Request must contains proper Token in Authorization Header.

- [Get Book](#get-book)
- [Get All Books](#get-all-books)
- [Borrow or Return Book](#borrow-or-return-book)

## Get Book

Returns single book.

Method: GET

URI: /api/member/book/{book-id}

Response:
* BookID (string)
* BookName (string)
* AuthorName (string)
* Publisher (string)
* Description (string?)

## Get All Books

Returns all available books that fit search criteria.
All parameters are oprional.

Method: GET

URI: /api/member/book/all/?searchTerm={searchTerm}&offset={offset}&limit={limit}

Response:
* Data (BookData)
* Meta (MetaData)

BookData is array of:
* BookID (string)
* BookName (string)
* AuthorName (string)
* Publisher (string)

MetaData:
* SearchTerm (string?)
* RowOffset (int?)
* RowLimit (int)

## Borrow or Return Book

Change book's status (depends on current status)

Method: PATCH

URI: /api/member/book

Request:
* BookID
