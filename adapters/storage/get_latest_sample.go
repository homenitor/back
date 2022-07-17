package storage

import (
	"context"

	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *MongoDBRepository) GetLatestSample(probeID string, category values.SampleCategory) (*entities.Sample, error) {
	opts := options.FindOne().SetSort(bson.D{{"measured_at", -1}})
	filter := bson.D{
		{"probe", probeID},
		{"category", category},
	}

	var result MongoSample
	err := r.samples().FindOne(context.TODO(), filter, opts).Decode(&result)
	if err != nil {
		return nil, err
	}

	return entities.NewSample(result.Category, result.MeasuredAt, result.Value), nil
}
