package adapters

import (
	"context"
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
	"go.mongodb.org/mongo-driver/bson"
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

type MongoSample struct {
	Probe      string                `bson:"probe"`
	MeasuredAt time.Time             `bson:"measured_at"`
	Value      float64               `bson:"value"`
	Category   values.SampleCategory `bson:"category"`
}

type MongoProbe struct {
	ID string `bson:"id"`
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
	pipeline := bson.A{
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

func (r *MongoDBRepository) ListProbes() ([]*entities.ProbeListingView, error) {
	cursor, err := r.probes().Find(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return r.decodeProbeListingViews(cursor)
}

func (r *MongoDBRepository) GetProbe(id string) (*entities.Probe, error) {
	filter := FindByIDFilter{ID: id}

	mongo_result := r.probes().FindOne(context.TODO(), filter)

	return r.decodeProbe(mongo_result)
}

func (r *MongoDBRepository) SaveProbe(probe *entities.Probe) error {
	doc := MongoProbe{
		ID: probe.ID(),
	}

	result, err := r.probes().InsertOne(context.TODO(), doc)

	r.logging.Debugf("document inserted: %v\n", result.InsertedID)
	return err
}

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

func (r *MongoDBRepository) Disconnect() error {
	return r.client.Disconnect(context.TODO())
}

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

func (r *MongoDBRepository) samples() *mongo.Collection {
	return r.collection("samples")
}

func (r *MongoDBRepository) probes() *mongo.Collection {
	return r.collection("probes")
}

func (r *MongoDBRepository) collection(name string) *mongo.Collection {
	return r.client.Database("homenitor").Collection(name)
}
