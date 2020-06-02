package user

import (
	"go-rest-api/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// API for user.
type API struct {
	Users  *Service
	Logger *logrus.Logger
	Db     *db.Conn
}

// New creates new user api.
func New(log *logrus.Logger, db *db.Conn) *API {
	service := NewService()

	return &API{
		Users:  service,
		Logger: log,
		Db:     db,
	}
}

// Router creates routes for user.
func (api *API) Router(router *gin.RouterGroup) {
	router.GET("/ping", api.get)
}

// Result gives api result.
type Result struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (api *API) get(c *gin.Context) {

	var result Result
	api.Db.Table("traits").Select("name,id").Scan(&result)

	api.Logger.Info(result)

	c.JSON(200, gin.H{
		"data":    result,
		"message": "hello there pong",
	})
}
