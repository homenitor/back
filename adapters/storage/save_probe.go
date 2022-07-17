package storage

import (
	"context"

	"github.com/homenitor/back/core/entities"
)

func (r *MongoDBRepository) SaveProbe(probe *entities.Probe) error {
	doc := MongoProbe{
		ID: probe.ID(),
	}

	result, err := r.probes().InsertOne(context.TODO(), doc)

	r.logging.Debugf("document inserted: %v\n", result.InsertedID)
	return err
}
