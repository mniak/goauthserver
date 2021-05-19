package mongo_providers

import (
	"github.com/mniak/goauthserver/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoKeyProvider struct {
	collection *mongo.Collection
}

func NewKeyProvider(collection *mongo.Collection) *mongoKeyProvider {
	return &mongoKeyProvider{
		collection: collection,
	}
}

func (p *mongoKeyProvider) GetKeys() ([]domain.Key, error) {
	return []domain.Key{}, nil
}
