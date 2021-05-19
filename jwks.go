package main

import "github.com/gin-gonic/gin"

func (s *ServerConfig) JWKS() gin.HandlerFunc {
	return func(c *gin.Context) {
		keys, err := s.KeyProvider.GetKeys()
		if err != nil {
			c.Error(err)
		}

		keysEntry := make([]gin.H, len(keys))
		for idx, _ := range keysEntry {
			keysEntry[idx] = gin.H{}
		}

		c.JSON(200, gin.H{
			"keys": []gin.H{},
		})
	}
}
