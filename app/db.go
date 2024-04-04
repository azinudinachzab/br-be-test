package app

import (
	"database/sql"
	"net/url"
)

func dbConnection(dsn string) (*sql.DB, error) {
	conn, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem&TimeZone=Asia/Jakarta"
	db, err := sql.Open("postgres", conn.String()) //user:password@/dbname
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
