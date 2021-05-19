package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mniak/goauthserver/mongo_providers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	config.Router.GET("/.well-known/openid-configuration", config.Discovery())
	config.Router.GET("/auth/keys", config.JWKS())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONNECTIONSTRING")))
	config.KeyProvider = mongo_providers.NewKeyProvider(client.Database(os.Getenv("MONGO_DATABASE")).Collection(os.Getenv("MONGO_KEYS_COLLECTION")))

	err = config.RunServer()
	if err != nil {
		os.Exit(-1)
	}
}
