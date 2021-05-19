package main

import (
	"reflect"
	"testing"

	"github.com/mniak/goauthserver/domain"
	"github.com/mniak/goauthserver/mongo_providers"
	"github.com/stretchr/testify/assert"
)

func TestMongoProviders(t *testing.T) {
	method := mongo_providers.NewKeyProvider
	vv := reflect.ValueOf(method).Type()

	assert.IsType(t, reflect.TypeOf((domain.KeyProvider)(nil)), vv)
}
