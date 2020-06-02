package server

import (
	"go-rest-api/config"
	"go-rest-api/routes"
	"log"
	"net/http"
)

// Conn struct to construct new http server.
type Conn struct {
	*http.Server
}

// appServer, all server instance should have start method.
type starter interface {
	Start()
}

// New creates a new http server.
func New(router *routes.Router, config *config.Config) *Conn {
	srv := &http.Server{
		Addr:              config.Server.Port,
		ReadHeaderTimeout: config.Server.ReadTimeOut,
		WriteTimeout:      config.Server.WriteTimeOut,
		Handler:           router,
	}

	return &Conn{
		srv,
	}
}

// Start starts the server
func (srv *Conn) Start() error {
	log.Printf("Starting server on port %s...", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
