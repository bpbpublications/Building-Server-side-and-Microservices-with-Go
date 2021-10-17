package dbserver

import (
	"building-restful-web-services-with-go/chapter2/library/config"
	"building-restful-web-services-with-go/chapter2/library/values"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var (
	// InitializeDb initializes database
	InitializeDb = initializeDb

	// PrepareDbRunner creates db runner for master database
	// and puts it in context
	PrepareDbRunner = prepareDbRunner
)

var dbHandler *sql.DB

func initializeDb() (err error) {
	connectionString := config.GetDatabaseConnectionString()
	maxIdleConnections := config.GetDatabaseMaxIdleConnections()
	maxOpenConnections := config.GetDatabaseMaxOpenConnections()
	connectionMaxLifetime := config.GetDatabaseConnectionMaxLifetime()

	dbHandler, err = initDbHandle("master", "postgres", connectionString,
		maxIdleConnections, maxOpenConnections, connectionMaxLifetime)
	if err != nil {
		return
	}

	return
}

func prepareDbRunner(ctx context.Context) context.Context {
	return context.WithValue(ctx, values.ContextKeyDbRunner, createRunner(dbHandler))
}

func initDbHandle(name, dbType, connectionString string, maxIdleConnections, maxOpenConnections int, connectionMaxLifetime time.Duration) (dbHandler *sql.DB, err error) {
	if dbType == "" {
		return nil, errors.New("Database type is empty")
	}

	if connectionString == "" {
		return nil, errors.New("Connection string is empty")
	}

	dbHandler, err = sql.Open(dbType, connectionString)
	if err != nil {
		return nil, err
	}

	// Initialize connection pool
	dbHandler.SetMaxIdleConns(maxIdleConnections)
	dbHandler.SetMaxOpenConns(maxOpenConnections)
	dbHandler.SetConnMaxLifetime(connectionMaxLifetime)

	err = validateDB(dbHandler)
	if err != nil {
		dbHandler.Close()
	}

	return
}

func validateDB(dbHandler *sql.DB) (err error) {
	err = dbHandler.Ping()
	if err != nil {
		return
	}

	timeZone, err := readDatabaseTimeZone(context.Background(), dbHandler)
	if err != nil {
		return
	}

	if timeZone != "UTC" {
		err = fmt.Errorf("Database 'timezone' must be set to 'UTC'. Currently it's '%v'", timeZone)
		return
	}

	return
}

func readDatabaseTimeZone(ctx context.Context, dbHandler *sql.DB) (timeZone string, err error) {
	rowsTimeZone, err := dbHandler.QueryContext(ctx, "show timezone")
	if err != nil {
		return
	}

	defer rowsTimeZone.Close()

	if !rowsTimeZone.Next() {
		err = fmt.Errorf("No time zone")
		return
	}

	err = rowsTimeZone.Scan(&timeZone)

	return
}

func createRunner(db *sql.DB) Runner {
	run := new(dbRunner)
	run.db = db
	return run
}
