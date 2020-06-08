package db

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/logger"
	modal "go-rest-api/modal"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Conn for database.
type Conn struct {
	*gorm.DB
}

// Database functions.
type Database interface {
	Migrate()
	Close()
}

// New database connection.
func New(config *config.Config) (*Conn, error) {

	dbURL := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Db, config.Database.Password)

	db, err := gorm.Open(config.Database.Client, dbURL)
	if err != nil {
		logger.Logger.Info(err)
		return nil, err
	}

	return &Conn{db}, nil
}

// Migrate migrate table to db.
func (db *Conn) Migrate() {
	db.AutoMigrate(&modal.User{})

	db.Model(&modal.User{}).AddIndex("idx_user_name", "name")
}

// Close closes the databse connection.
func (db *Conn) Close() {
	db.Close()
}
