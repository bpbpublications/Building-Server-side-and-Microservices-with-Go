package databasemodule

import (
	"database/sql"
	"errors"
	"reflect"
	"strconv"
	"time"
)

var (
	// GetRowReader returns RowReader interface that can be used to read data from sql.Rows
	GetRowReader = getRowReader
)

type rowReader struct {
	rows      *sql.Rows
	columns   []string
	values    []interface{}
	valuePtrs []interface{}
	lastError error
}

// Errors
var (
	ErrorNullValue   = errors.New("Null value encountered")
	ErrorWrongType   = errors.New("Unable to convert type")
	ErrorUnsupported = errors.New("Unsupported type")
)

// RowReader is used to simplify reading of sql.Rows objects
type RowReader interface {
	ScanNext() bool
	Error() error
	RowReaderFxs
}

// RowReaderFxs contains method to read values from single row in sql.Rows
type RowReaderFxs interface {
	ReadByIdxString(columnIdx int) string
	ReadByIdxInt64(columnIdx int) int64
	ReadByIdxTime(columnIdx int) time.Time
	ReadAllToStruct(p interface{})
}

func getRowReader(rows *sql.Rows) (RowReader, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	n := len(columns)

	rr := new(rowReader)
	rr.rows = rows
	rr.columns = columns
	rr.values = make([]interface{}, n)
	rr.valuePtrs = make([]interface{}, n)
	for i := 0; i < n; i++ {
		rr.valuePtrs[i] = &rr.values[i]
	}

	return rr, nil
}

func (rr *rowReader) Error() error {
	return rr.lastError
}

func (rr *rowReader) ScanNext() (hasMore bool) {
	if hasMore = rr.rows.Next(); hasMore {
		err := rr.rows.Scan(rr.valuePtrs...)
		rr.lastError = err
		if err != nil {
			hasMore = false
		}
	}

	return
}

func (rr *rowReader) ReadByIdxString(columnIdx int) string {
	switch value := rr.values[columnIdx].(type) {
	case string:
		return value
	case []byte:
		return string(value)
	case nil:
		panic(ErrorNullValue)
	default:
		panic(ErrorWrongType)
	}
}

func (rr *rowReader) ReadByIdxInt64(columnIdx int) int64 {
	switch value := rr.values[columnIdx].(type) {
	case int64:
		return value
	case []byte:
		s := string(value)
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(ErrorWrongType)
		}
		return i
	case nil:
		panic(ErrorNullValue)
	default:
		panic(ErrorWrongType)
	}
}

func (rr *rowReader) ReadByIdxTime(columnIdx int) time.Time {
	switch value := rr.values[columnIdx].(type) {
	case time.Time:
		return value
	case []byte:
		time, err := time.Parse(time.RFC3339Nano, string(value))
		if err != nil {
			panic(ErrorWrongType)
		}
		return time
	case string:
		time, err := time.Parse(time.RFC3339Nano, value)
		if err != nil {
			panic(ErrorWrongType)
		}
		return time
	case nil:
		panic(ErrorNullValue)
	default:
		panic(ErrorWrongType)
	}
}

func (rr *rowReader) ReadAllToStruct(p interface{}) {
	var value reflect.Value

	value = reflect.ValueOf(p)
	if value.Kind() != reflect.Ptr {
		return
	}

	value = reflect.Indirect(value)
	if value.Kind() != reflect.Struct {
		return
	}

	for columnIdx, columnName := range rr.columns {
		if rr.values[columnIdx] == nil {
			continue
		}

		column := value.FieldByName(columnName)
		if column == (reflect.Value{}) {
			continue
		}

		switch column.Kind() {
		case reflect.String:
			column.SetString(rr.ReadByIdxString(columnIdx))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			column.SetInt(rr.ReadByIdxInt64(columnIdx))
		default:
			panic(ErrorUnsupported)
		}
	}
}
