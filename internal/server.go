package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	config *Config
	Route  *gin.RouterGroup
}

func New(config *Config) *Server {

	engine := gin.Default()
	engine.Use(cors.Default())
	route := engine.Group("/api")

	return &Server{
		engine: engine,
		Route:  route,
		config: config,
	}
}

func (s *Server) Run() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.config.Port),
		Handler: s.engine.Handler(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("[API] Server Started on ", s.config.Port)
	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)
	<-quite
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited")
}
