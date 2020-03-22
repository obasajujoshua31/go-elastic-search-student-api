package services

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-elastic-search-student-api/config"
	"time"
)

type DB struct {
	*sql.DB
	GoDB *gorm.DB
}

const dialect = "postgres"

func ConnectToDatabase(config config.Config) (*DB, error) {
	db, err := gorm.Open(dialect, constructConnString(config))

	if err != nil {
		return nil, err
	}

	return &DB{db.DB(), db}, nil

}

func constructConnString(config config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword, "disable")
}

func (d *DB) SetPool() {
	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(100)
	d.SetConnMaxLifetime(time.Hour)
}
