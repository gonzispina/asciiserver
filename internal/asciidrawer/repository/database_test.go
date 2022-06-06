package repository

import (
	"github.com/gonzispina/asciiserver/internal/asciidrawer"
	"github.com/gonzispina/gokit/context"
	"github.com/gonzispina/gokit/logs"
	"github.com/gonzispina/gokit/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestCanvasStorage(t *testing.T) {
	m := mongo.NewMongoClient(mongo.DefaultConnectionString(), "test", nil)
	l := logs.InitTest()
	c := DefaultDatabaseConfig()
	ctx := context.Background()

	storage := NewCanvasStorage(c, m, l)
	t.Run("Given we want to get a canvas", func(t *testing.T) {
		t.Run("When the canvas doesn't exists it returns an error", func(t *testing.T) {
			_, err := storage.GetSerialization(ctx, primitive.NewObjectID().Hex())
			assert.Equal(t, asciidrawer.ErrSerializationDoesNotExists, err)
		})

		t.Run("When the canvas exists it returns the proper one", func(t *testing.T) {
			id, err := storage.SaveSerialization(ctx, "test")
			require.Nil(t, err)

			res, err := storage.GetSerialization(ctx, id)
			require.Nil(t, err)

			assert.Equal(t, id, res.ID)
			assert.Equal(t, "test", res.Str)
		})
	})
}
