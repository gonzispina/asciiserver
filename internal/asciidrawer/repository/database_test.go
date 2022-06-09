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
			canvasSize := 10
			figures := []asciidrawer.Figure{
				&asciidrawer.Rectangle{
					Vertex: asciidrawer.Vertex{
						Row:    2,
						Column: 4,
					},
					Height:  1,
					Width:   2,
					Outline: "X",
					Fill:    "A",
				},
				&asciidrawer.Rectangle{
					Vertex: asciidrawer.Vertex{
						Row:    5,
						Column: 6,
					},
					Height:  1,
					Width:   2,
					Outline: "A",
					Fill:    "A",
				},
			}

			s, err := storage.SaveSerialization(ctx, canvasSize, figures)
			require.Nil(t, err)

			res, err := storage.GetSerialization(ctx, s.ID)
			require.Nil(t, err)

			assert.Equal(t, s.ID, res.ID)
			assert.Equal(t, canvasSize, res.CanvasSize)
			assert.Equal(t, figures, res.Figures)
		})
	})
}
