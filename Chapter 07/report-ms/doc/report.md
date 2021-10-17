# Report API

Only users with role librarian can call this API.
Request must contains proper Token in Authorization Header.

- [Daily Report](#daily-report)
- [Book Report](#book-report)

## Daily Report

Returns how many books are borrowed for each day in data range.

Method: GET

URI: /api/report/daily?startDate={startDate}&endDate={endDate}

Response:
* Data (DailyReportData)
* Meta (MetaData)

DailyReportData is array of:
* ReportDate (timestamp)
* BookCount (int)

MetaData:
* StartDate (timestamp)
* EndDate (timestamp)

## Book Report

Returns how many time each book has been borrowed.

Method: GET

URI: /api/report/book

Response:
* Data (BookReportData)

BookReportData is array of:
* BookID (string)
* BookCount (int)
