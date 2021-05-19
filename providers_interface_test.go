package main

import (
	"testing"

	"github.com/mniak/goauthserver/domain"
	"github.com/mniak/goauthserver/mongo_providers"
)

func TestMongoProviders(t *testing.T) {
	var _ domain.KeyProvider = mongo_providers.NewKeyProvider()
}
