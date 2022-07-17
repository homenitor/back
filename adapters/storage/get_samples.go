package storage

import (
	"context"
	"time"

	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoGetSample struct {
	Probe string  `bson:"probe"`
	Value float64 `bson:"value"`
}

type MongoGetSamples struct {
	MeasuredAt time.Time        `bson:"_id"`
	Samples    []MongoGetSample `bson:"samples"`
	Average    float64          `bson:"average"`
}

func (r *MongoDBRepository) GetSamples(category values.SampleCategory, query libraries.GetSamplesQuery) ([]*entities.GetSamplesView, error) {
	pipeline := r.getSamplesPipeline(category, query)

	cursor, err := r.samples().Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var results []*entities.GetSamplesView
	for cursor.Next(context.TODO()) {
		var doc MongoGetSamples
		err := cursor.Decode(&doc)
		if err != nil {
			return nil, err
		}
		values := make(map[string]float64, 0)
		for _, sample := range doc.Samples {
			values[sample.Probe] = sample.Value
		}
		results = append(results, &entities.GetSamplesView{
			MeasuredAt: doc.MeasuredAt,
			Values:     values,
			Average:    doc.Average,
		})
	}

	return results, nil
}

func (r *MongoDBRepository) getSamplesPipeline(category values.SampleCategory, query libraries.GetSamplesQuery) primitive.A {
	return bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"category", category},
					{"measured_at",
						bson.D{
							{"$gte", query.From},
							{"$lte", query.To},
						},
					},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$measured_at"},
					{"samples",
						bson.D{
							{"$push",
								bson.D{
									{"probe", "$probe"},
									{"value", "$value"},
								},
							},
						},
					},
					{"average", bson.D{{"$avg", "$value"}}},
				},
			},
		},
		bson.D{{"$sort", bson.D{{"_id", 1}}}},
	}
}
