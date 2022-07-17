package storage

import (
	"context"

	"github.com/homenitor/back/core/entities"
)

func (r *MongoDBRepository) GetProbe(id string) (*entities.Probe, error) {
	filter := FindByIDFilter{ID: id}

	mongo_result := r.probes().FindOne(context.TODO(), filter)

	return r.decodeProbe(mongo_result)
}
