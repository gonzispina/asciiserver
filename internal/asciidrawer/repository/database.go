package repository

import (
	"github.com/gonzispina/asciiserver/internal/asciidrawer"
	"github.com/gonzispina/gokit/context"
	"github.com/gonzispina/gokit/logs"
	"github.com/gonzispina/gokit/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DefaultDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{CollectionName: "canvases"}
}

type DatabaseConfig struct {
	CollectionName string
}

// NewCanvasStorage constructor
func NewCanvasStorage(c *DatabaseConfig, m *mongo.Mongo, l logs.Logger) *CanvasStorage {
	// fail fast
	if c == nil {
		panic("config must be initialized")
	}
	if m == nil {
		panic("mongo client must be initialized")
	}
	if l == nil {
		panic("logger must be initialized")
	}

	return &CanvasStorage{
		config: c,
		mongo:  m,
		log:    l,
	}
}

// CanvasStorage repository
type CanvasStorage struct {
	config *DatabaseConfig
	mongo  *mongo.Mongo
	log    logs.Logger
}

// SaveSerialization saves a new serialization into the database
func (s *CanvasStorage) SaveSerialization(ctx context.Context, serialization string) (string, error) {
	hexID := primitive.NewObjectID()

	insert := bson.D{
		{Key: "_id", Value: hexID},
		{Key: "str", Value: serialization},
	}

	_, err := s.mongo.Collection(s.config.CollectionName).InsertOne(ctx, insert)
	if err != nil {
		s.log.Error(ctx, "Couldn't save serialization", logs.Error(err))
		return "", err
	}

	return hexID.Hex(), nil
}

func (s *CanvasStorage) GetSerialization(ctx context.Context, id string) (*asciidrawer.Serialization, error) {
	hexID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		s.log.Error(ctx, "Couldn't generate hex id", logs.Error(err))
		return nil, err
	}

	filter := bson.M{"_id": hexID}
	res := s.mongo.Collection(s.config.CollectionName).FindOne(ctx, filter)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, asciidrawer.ErrSerializationDoesNotExists
		}
		s.log.Error(ctx, "Couldn't fetch the person from the db", logs.Error(err))
		return nil, err
	}

	sr := asciidrawer.Serialization{}
	if err := res.Decode(&sr); err != nil {
		s.log.Error(ctx, "Couldn't decode person", logs.Error(err))
		return nil, err
	}

	return &sr, nil
}
