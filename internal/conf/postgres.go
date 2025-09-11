package conf

import (
	"errors"
	"os"
)

func GetPostgresDSN() string {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		panic(errors.New("POSTGRES_DSN is not set"))
	}
	if !validatePostgresDSN(dsn) {
		panic(errors.New("invalid POSTGRES_DSN format"))
	}
	return dsn
}

func validatePostgresDSN(dsn string) bool {
	// Basic validation for DSN format
	// Check if DSN is empty
	if dsn == "" {
		return false
	}

	// Check if DSN contains required components
	// Expected format: postgres://username:password@host:port/dbname
	if len(dsn) < 11 || dsn[:11] != "postgres://" {
		return false
	}

	// Check for @ symbol separating credentials from host
	atIndex := -1
	for i := 11; i < len(dsn); i++ {
		if dsn[i] == '@' {
			atIndex = i
			break
		}
	}
	if atIndex == -1 {
		return false
	}

	// Check for : in host:port section
	colonFound := false
	for i := atIndex + 1; i < len(dsn); i++ {
		if dsn[i] == ':' {
			colonFound = true
			break
		}
	}
	if !colonFound {
		return false
	}

	// Check for / after host:port
	slashFound := false
	for i := atIndex + 1; i < len(dsn); i++ {
		if dsn[i] == '/' {
			slashFound = true
			break
		}
	}

	return slashFound
}
