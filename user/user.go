package user

import (
	"go-rest-api/db"
	"go-rest-api/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// API for user.
type API struct {
	Users  *ServiceStore
	Logger *logrus.Logger
}

// New creates new user api.
func New(log *logrus.Logger, db *db.Conn) *API {

	userModal := store.NewUser(db)
	service := NewService(userModal, log)

	return &API{
		Users:  service,
		Logger: log,
	}
}

// Router creates routes for user.
func (api *API) Router(router *gin.RouterGroup) {
	router.GET("/ping", api.Get)
}

// Get user information.
func (api *API) Get(c *gin.Context) {
	data, err := api.Users.Get()

	if err != nil {
		c.JSON(err.HTTPStatusCode, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}
