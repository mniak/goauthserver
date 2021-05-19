package mongo_providers

import (
	"context"
	"crypto/x509"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/pkcs12"
)

type mongoKeyProvider struct {
	collection *mongo.Collection
}

type dbKey struct {
	PrivateKey     []byte `bson:"private.key"`
	PublicCrt      []byte `bson:"public.crt"`
	CertificatePfx []byte `bson:"certificate.pfx"`
	Password       string `bson:"password"`
}

func NewKeyProvider(collection *mongo.Collection) *mongoKeyProvider {
	return &mongoKeyProvider{
		collection: collection,
	}
}

func (p *mongoKeyProvider) Keys(ctx context.Context) ([]*x509.Certificate, error) {
	cur, err := p.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	keys := make([]*x509.Certificate, 0)
	for cur.Next(ctx) {
		var k dbKey
		err = cur.Decode(&k)
		if err != nil {
			return nil, err
		}

		_, xcert, err := pkcs12.Decode(k.CertificatePfx, k.Password)
		if err != nil {
			return nil, err
		}
		keys = append(keys, xcert)
	}
	return keys, nil
}
