package server

import (
	"net/http"
	"zero-kernel-service/internal/infrastructure/config"

	"github.com/gin-gonic/gin"
)

// NewServer -- create new server instance
func NewServer(cfg config.Config) (Server, error) {
	ginHandler := gin.New()
	ginHandler.Use(gin.Recovery())

	server := &http.Server{
		Addr:    cfg.Server.ListenAddress,
		Handler: ginHandler,
	}

	return Server{
		Server: server,
	}, nil
}

// Server -- server structure
type Server struct {
	*http.Server
}
