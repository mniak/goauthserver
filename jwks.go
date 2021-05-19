package main

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"math/big"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme"
)

func (s *ServerConfig) JWKS() gin.HandlerFunc {
	return func(c *gin.Context) {
		keys, err := s.KeyProvider.Keys(c)
		if err != nil {
			c.Error(err)
		}

		jwKeys := make([]gin.H, len(keys))
		for idx, k := range keys {
			jwk, err := jwk(k)
			if err != nil {
				c.Error(err)
			}
			jwKeys[idx] = jwk
		}

		c.JSON(200, gin.H{
			"keys": jwKeys,
		})
	}
}

func jwk(cert *x509.Certificate) (map[string]interface{}, error) {
	thumbprint, err := acme.JWKThumbprint(cert.PublicKey)
	if err != nil {
		return nil, err
	}
	switch pub := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		n := pub.N
		e := big.NewInt(int64(pub.E))

		return map[string]interface{}{
			"kty": "RSA",
			"e":   base64.RawURLEncoding.EncodeToString(e.Bytes()),
			"n":   base64.RawURLEncoding.EncodeToString(n.Bytes()),
			"kid": thumbprint,
		}, nil
	case *ecdsa.PublicKey:
		p := pub.Curve.Params()
		n := p.BitSize / 8
		if p.BitSize%8 != 0 {
			n++
		}
		x := pub.X.Bytes()
		if n > len(x) {
			x = append(make([]byte, n-len(x)), x...)
		}
		y := pub.Y.Bytes()
		if n > len(y) {
			y = append(make([]byte, n-len(y)), y...)
		}
		return map[string]interface{}{
			"crv": p.Name,
			"kty": "EC",
			"x":   base64.RawURLEncoding.EncodeToString(x),
			"y":   base64.RawURLEncoding.EncodeToString(y),
			"kid": thumbprint,
		}, nil
	default:
		return nil, acme.ErrUnsupportedKey
	}
}
