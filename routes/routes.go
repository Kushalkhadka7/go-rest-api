package routes

import (
	"go-rest-api/db"
	"go-rest-api/logger"
	middleware "go-rest-api/middlewares"
	"go-rest-api/user"

	"github.com/gin-gonic/gin"
)

// Router to accepts incoming requests.
type Router struct {
	*gin.Engine
}

// New creates new router.
func New(db *db.Conn) *Router {
	logger := logger.NewLogger()

	r := gin.New()

	r.Use(middleware.RequestLogger())

	r.Use(gin.Recovery())

	userAPI := user.New(logger, db)

	routesGroup := r.Group("/")

	userAPI.Router(routesGroup)

	return &Router{
		r,
	}
}
