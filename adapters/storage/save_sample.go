package storage

import (
	"context"

	"github.com/homenitor/back/core/entities"
)

func (r *MongoDBRepository) SaveSample(probeID string, sample *entities.Sample) error {
	coll := r.collection("samples")
	doc := MongoSample{
		Probe:      probeID,
		MeasuredAt: sample.MeasuredAt(),
		Value:      sample.Value(),
		Category:   sample.Category(),
	}

	result, err := coll.InsertOne(context.TODO(), doc)

	r.logging.Debugf("document inserted: %v\n", result.InsertedID)
	return err
}
