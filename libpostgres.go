package libpostgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	_ "go.elastic.co/apm/module/apmsql/pq"

	"go.elastic.co/apm/module/apmsql"
)

// Options is a struct
type Options struct {
	Host           string
	Port           int
	Username       string
	Password       string
	DBName         string
	ConnectTimeout int
	SSLCert        string
	SSLKey         string
	SSLRootCert    string
	SSLMode        string
}

var sslModes []string = []string{"disable", "allow", "prefer", "require", "verify-ca", "verify-full"}

func isValidSSLMode(sslMode string) bool {
	for _, v := range sslModes {
		if sslMode == v {
			return true
		}
	}

	return false
}

// Connect is ..
func Connect(options Options) (*sql.DB, error) {
	sslMode := "sslmode=disabled"

	if options.SSLMode != "" && options.SSLMode != "disabled" {
		if !isValidSSLMode(options.SSLMode) {
			return nil, errors.New("invalid ssl mode")
		}

		sslMode = fmt.Sprintf("sslmode=%s&sslrootcert=%s&sslcert=%s&sslkey=%s",
			options.SSLMode,
			options.SSLRootCert,
			options.SSLCert,
			options.SSLKey)
	}

	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s",
		options.Username,
		options.Password,
		options.Host,
		options.Port,
		options.DBName,
		sslMode)

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ConnectWithAMP is ...
func ConnectWithAMP(options Options) (*sql.DB, error) {
	sslMode := "sslmode=disabled"

	if options.SSLMode != "" && options.SSLMode != "disabled" {
		if !isValidSSLMode(options.SSLMode) {
			return nil, errors.New("invalid ssl mode")
		}

		sslMode = fmt.Sprintf("sslmode=%s&sslrootcert=%s&sslcert=%s&sslkey=%s",
			options.SSLMode,
			options.SSLRootCert,
			options.SSLCert,
			options.SSLKey)
	}

	dbConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s",
		options.Username,
		options.Password,
		options.Host,
		options.Port,
		options.DBName,
		sslMode)

	db, err := apmsql.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
