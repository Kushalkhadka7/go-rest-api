package db

import (
	"fmt"
	"go-rest-api/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnMock for database.
type ConnMock struct {
	*gorm.DB
}

// NewMock database connection.
func NewMock(config *config.Config) *ConnMock {

	dbURL := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", config.Database.User, config.Database.Password, config.Database.Db)

	db, err := gorm.Open(config.Database.Client, dbURL)
	if err != nil {
		panic(err.Error())
	}

	return &ConnMock{db}
}
