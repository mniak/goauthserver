package domain

import (
	"context"
	"crypto/x509"
)

type KeyProvider interface {
	Keys(ctx context.Context) ([]*x509.Certificate, error)
}
