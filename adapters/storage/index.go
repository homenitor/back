package storage

import (
	"context"

	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/values"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBRepository struct {
	client  *mongo.Client
	logging libraries.Logging
}

type FindByIDFilter struct {
	ID string `bson:"id"`
}

type FindByProbeAndCategoryFilter struct {
	Probe    string                `bson:"probe"`
	Category values.SampleCategory `bson:"category"`
}

func NewMongoDBRepository(logging libraries.Logging) libraries.Repository {
	const uri = "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	logging.Infof("Connected to MongoDB server")

	return &MongoDBRepository{
		logging: logging,
		client:  client,
	}
}
func (r *MongoDBRepository) Disconnect() error {
	return r.client.Disconnect(context.TODO())
}

func (r *MongoDBRepository) collection(name string) *mongo.Collection {
	return r.client.Database("homenitor").Collection(name)
}
