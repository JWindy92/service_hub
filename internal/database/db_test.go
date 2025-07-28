package database

import "testing"

func TestPostgresConnect(t *testing.T) {
	db := PostgresImpl{}
	db.ConnectDB()
}

func TestSQLiteConnect(t *testing.T) {
	db := SQLiteImpl{}
	db.ConnectDB()
}
