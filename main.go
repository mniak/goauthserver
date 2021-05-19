package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	Host   string
	Port   int
	Router *gin.Engine
	Logger *zap.Logger
}

const (
	DefaultHost = "0.0.0.0"
	DefaultPort = 5000
)

func NewServer() (*Server, error) {
	var port int

	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	host := os.Getenv("HOST")
	if host == "" {
		log.Warn("invalid host, using default",
			zap.String("host", host),
			zap.String("default_host", DefaultHost))
		host = DefaultHost
	}

	portStr := os.Getenv("PORT")
	if port, err = strconv.Atoi(portStr); err != nil {
		log.Warn("invalid port, using default",
			zap.String("port", portStr),
			zap.Int("default_port", DefaultPort))
		port = DefaultPort
	}

	return &Server{
		Host:   host,
		Port:   port,
		Router: gin.Default(),
		Logger: log,
	}, nil
}

func (s *Server) Run() error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), s.Router)
	if err != nil {
		s.Logger.Fatal("could not start server", zap.Any("error", err))
	}
	return err
}

func main() {
	server, err := NewServer()
	if err != nil {
		log.Fatalln(err)
	}

	server.Router.GET("/.well-known/openid-configuration", discoveryEndpoint)

	err = server.Run()
	if err != nil {
		os.Exit(-1)
	}
}

func discoveryEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"issuer": "http://localhost.google.com",
	})
}
