package storage

import (
	"context"

	"github.com/homenitor/back/core/entities"
)

func (r *MongoDBRepository) ListProbes() ([]*entities.ProbeListingView, error) {
	cursor, err := r.probes().Find(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return r.decodeProbeListingViews(cursor)
}
