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

// CreateSerialization saves a new serialization into the database
func (s *CanvasStorage) CreateSerialization(ctx context.Context, canvasHeight, canvasWidth int, figures []asciidrawer.Figure) (*asciidrawer.Serialization, error) {
	hexID := primitive.NewObjectID()

	insert := bson.D{
		{Key: "_id", Value: hexID},
		{Key: "canvasHeight", Value: canvasHeight},
		{Key: "canvasWidth", Value: canvasWidth},
		{Key: "figures", Value: mapFigures(figures)},
	}

	_, err := s.mongo.Collection(s.config.CollectionName).InsertOne(ctx, insert)
	if err != nil {
		s.log.Error(ctx, "Couldn't save serialization", logs.Error(err))
		return nil, err
	}

	return &asciidrawer.Serialization{
		ID:           hexID.Hex(),
		CanvasHeight: canvasHeight,
		CanvasWidth:  canvasWidth,
		Figures:      figures,
	}, nil
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
		s.log.Error(ctx, "Couldn't fetch the serialization from the db", logs.Error(err))
		return nil, err
	}

	b, err := res.DecodeBytes()
	if err != nil {
		s.log.Error(ctx, "Couldn't decode serialization bytes", logs.Error(err))
		return nil, err
	}

	sr := asciidrawer.Serialization{
		ID:           b.Lookup("_id").ObjectID().Hex(),
		CanvasHeight: int(b.Lookup("canvasHeight").AsInt64()),
		CanvasWidth:  int(b.Lookup("canvasWidth").AsInt64()),
		Figures:      []asciidrawer.Figure{},
	}

	elems, err := b.Lookup("figures").Array().Elements()
	if err != nil {
		s.log.Error(ctx, "Couldn't look up figures", logs.Error(err))
		return nil, err
	}

	for _, elem := range elems {
		var f asciidrawer.Figure

		document := elem.Value().Document()
		t := document.Lookup("type").AsInt64()

		switch asciidrawer.FigureType(t) {
		case asciidrawer.TypeRectangle:
			f = &asciidrawer.Rectangle{}
			break
		default:
			s.log.Warn(ctx, "Couldn't match database type")
			continue
		}

		err := bson.Unmarshal(document.Lookup("data").Value, f)
		if err != nil {
			s.log.Error(ctx, "Couldn't unmarshal rectangle", logs.Error(err))
			return nil, err
		}
		sr.Figures = append(sr.Figures, f)
	}

	return &sr, nil
}
