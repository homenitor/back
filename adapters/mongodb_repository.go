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

func (r *MongoDBRepository) ListProbes() ([]*entities.ProbeListingView, error) {
	coll := r.client.Database("homenitor").Collection("probes")
	cursor, err := coll.Find(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
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

func (r *MongoDBRepository) GetProbe(id string) (*entities.Probe, error) {
	coll := r.client.Database("homenitor").Collection("probes")
	filter := FindByIDFilter{ID: id}

	var result MongoProbe
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.ErrProbeNotFound
		}

		return nil, err
	}

	probe := entities.NewProbeWithID(result.ID)
	return probe, nil
}

func (r *MongoDBRepository) SaveProbe(probe *entities.Probe) error {
	coll := r.client.Database("homenitor").Collection("probes")
	doc := MongoProbe{
		ID: probe.ID(),
	}

	result, err := coll.InsertOne(context.TODO(), doc)

	r.logging.Debugf("document inserted: %v\n", result.InsertedID)
	return err
}

func (r *MongoDBRepository) SaveSample(probeID string, sample *entities.Sample) error {
	coll := r.client.Database("homenitor").Collection("samples")
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
	coll := r.client.Database("homenitor").Collection("samples")
	opts := options.FindOne().SetSort(bson.D{{"measured_at", -1}})
	filter := bson.D{
		{"probe", probeID},
		{"category", category},
	}

	var result MongoSample
	err := coll.FindOne(context.TODO(), filter, opts).Decode(&result)
	if err != nil {
		return nil, err
	}

	return entities.NewSample(result.Category, result.MeasuredAt, result.Value)
}

func (r *MongoDBRepository) Disconnect() error {
	return r.client.Disconnect(context.TODO())
}
