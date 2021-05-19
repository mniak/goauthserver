package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mniak/goauthserver/domain"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Host    string
	Port    int
	BaseURL string

	DevMode bool

	Router *gin.Engine
	Logger *zap.Logger

	KeyProvider domain.KeyProvider
}

func NewServerConfig() (*ServerConfig, error) {
	var port int
	var err error
	var logger *zap.Logger

	devMode := getEnvBool("DEV_MODE", false)

	if devMode {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return nil, err
	}

	host := os.Getenv("HOST")
	if host == "" {
		logger.Warn("invalid host, using default",
			zap.String("host", host),
			zap.String("default_host", DefaultHost))
		host = DefaultHost
	}

	portStr := os.Getenv("PORT")
	if port, err = strconv.Atoi(portStr); err != nil {
		logger.Warn("invalid port, using default",
			zap.String("port", portStr),
			zap.Int("default_port", DefaultPort))
		port = DefaultPort
	}

	var baseURL string
	if devMode {
		baseURL = fmt.Sprintf("http://%s:%d", host, port)
	} else {
		if port == 443 {
			baseURL = fmt.Sprintf("https://%s", host)
		} else {
			baseURL = fmt.Sprintf("https://%s:%d", host, port)
		}
	}
	baseURL = getEnvString("EXTERNAL_URL", baseURL)

	logger.Info("server configuration loaded",
		zap.String("Host", host),
		zap.Int("Port", port),
		zap.String("BaseURL", baseURL),
		zap.Bool("DevMode", devMode),
	)
	return &ServerConfig{
		Host:    host,
		Port:    port,
		BaseURL: baseURL,

		DevMode: devMode,

		Router: gin.Default(),
		Logger: logger,
	}, nil
}

func (s *ServerConfig) RunServer() error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), s.Router)
	if err != nil {
		s.Logger.Fatal("could not start server", zap.Any("error", err))
	}
	return err
}

func (c *ServerConfig) Close() {
}
