package database

import (
	"database/sql"
	"golang-chapter-37/implem-gin-shipping/util"

	_ "github.com/lib/pq"

	"fmt"
	"time"
)

func InitDB(config util.Configuration) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s",
		config.DBConfig.DBUsername, config.DBConfig.DBPassword, config.DBConfig.DBName, config.DBConfig.DBHost)
	db, err := sql.Open("postgres", connStr)

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db, err
}
