package storage

import (
	"context"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *MongoDBRepository) decodeProbe(mongo_result *mongo.SingleResult) (*entities.Probe, error) {
	var result MongoProbe
	err := mongo_result.Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.ErrProbeNotFound
		}

		return nil, err
	}

	return entities.NewProbeWithID(result.ID), nil
}

func (r *MongoDBRepository) decodeProbeListingViews(cursor *mongo.Cursor) ([]*entities.ProbeListingView, error) {
	defer cursor.Close(context.TODO())

	var results []*entities.ProbeListingView
	for cursor.Next(context.TODO()) {
		var doc MongoProbe
		err := cursor.Decode(&doc)
		if err != nil {
			return nil, err
		}
		results = append(results, &entities.ProbeListingView{
			ID: doc.ID,
		})
	}

	return results, nil
}
