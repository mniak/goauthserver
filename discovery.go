package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *ServerConfig) Discovery() gin.HandlerFunc {
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
