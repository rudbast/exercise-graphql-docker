package database

import "database/sql"

var globalDB *sql.DB

func Init(driver, connString string) error {
	db, err := sql.Open(driver, connString)
	if err != nil {
		return err
	}

	globalDB = db
	return globalDB.Ping()
}

func Get() *sql.DB {
	return globalDB
}
