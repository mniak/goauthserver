package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Host    string
	Port    int
	BaseURL string

	DevMode bool

	Router *gin.Engine
	Logger *zap.Logger
}

const (
	DefaultHost = "0.0.0.0"
	DefaultPort = 5000
)

func main() {
	config, err := NewServerConfig()
	defer config.Close()

	if err != nil {
		log.Fatalln(err)
	}
	config.Router.GET("/.well-known/openid-configuration", config.discoveryEndpoint())
	config.Router.GET("/auth/keys", config.jwksEndpoint())

	err = config.RunServer()
	if err != nil {
		os.Exit(-1)
	}
}

func (s *ServerConfig) discoveryEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"issuer": s.BaseURL,

			"response_types_supported":              []string{"code"},
			"subject_types_supported":               []string{"public"},
			"id_token_signing_alg_values_supported": []string{"RS256"},

			"authorization_endpoint": fmt.Sprintf("%s/auth/authorization", s.BaseURL),
			"token_endpoint":         fmt.Sprintf("%s/auth/token", s.BaseURL),
			// "userinfo_endpoint":      "bbbbbbbbbbbbb",

			"jwks_uri": fmt.Sprintf("%s/auth/keys", s.BaseURL),

			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
			// "aaaaaaaaaa":             "bbbbbbbbbbbbb",
		})
	}
}

func (s *ServerConfig) jwksEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{})
	}
}
