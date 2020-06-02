package db

import (
	"fmt"
	"go-rest-api/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Conn for database.
type Conn struct {
	*gorm.DB
}

// New database connection.
func New(config *config.Config) (*Conn, error) {

	dbURL := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Db, config.Database.Password)

	db, err := gorm.Open(config.Database.Client, dbURL)
	if err != nil {
		return nil, err
	}

	return &Conn{db}, nil
}
