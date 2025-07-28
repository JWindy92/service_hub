package database

import (
	"fmt"
	"log"

	"github.com/JWindy92/service_hub/service_hub/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var DB *gorm.DB

type SQLiteImpl struct {
	DB *gorm.DB
}

func (impl *SQLiteImpl) ConnectDB() *gorm.DB {
	dbPath := "../../perfectpint.db"
	log.Printf("Connecting to SQLite DB %s\n", dbPath)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	impl.DB = db
	return db
}

type PostgresImpl struct {
	DB *gorm.DB
}

func (impl *PostgresImpl) ConnectDB() *gorm.DB {
	// Example: replace with real values or use environment variables
	host := "localhost"
	port := 5555
	user := "postgres"
	password := "dbpass"
	dbname := "postgres"

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	log.Printf("Connecting to Postgres DB: %s\n", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to Postgres: %v", err))
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		panic(fmt.Sprintf("failed to migrate schema: %v", err))
	}

	impl.DB = db
	return db
}
