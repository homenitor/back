package storage

import (
	"time"

	"github.com/homenitor/back/core/values"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoSample struct {
	Probe      string                `bson:"probe"`
	MeasuredAt time.Time             `bson:"measured_at"`
	Value      float64               `bson:"value"`
	Category   values.SampleCategory `bson:"category"`
}

type MongoProbe struct {
	ID string `bson:"id"`
}

func (r *MongoDBRepository) samples() *mongo.Collection {
	return r.collection("samples")
}

func (r *MongoDBRepository) probes() *mongo.Collection {
	return r.collection("probes")
}
