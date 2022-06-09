package asciidrawer_test

import (
	"github.com/gonzispina/asciiserver/internal/asciidrawer"
	"github.com/gonzispina/asciiserver/internal/asciidrawer/repository"
	"github.com/gonzispina/gokit/context"
	"github.com/gonzispina/gokit/errors"
	"github.com/gonzispina/gokit/logs"
	"github.com/gonzispina/gokit/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestDrawer_CreateDrawing(t *testing.T) {
	ctx := context.Background()
	m := mongo.NewMongoClient(mongo.DefaultConnectionString(), "test", nil)
	storage := repository.NewCanvasStorage(repository.DefaultDatabaseConfig(), m, logs.InitTest())
	drawer := asciidrawer.NewDrawer(storage)

	t.Run("Given we want to create a canvas", func(t *testing.T) {
		canvasWidth := 24
		canvasHeight := 10
		t.Run("When we create a rectangle", func(t *testing.T) {
			t.Run("If we create a rectangle without fill and outline it returns an error", func(t *testing.T) {
				r := &asciidrawer.Rectangle{
					Vertex: asciidrawer.Vertex{
						Row:    1,
						Column: 1,
					},
					Height:  1,
					Width:   1,
					Outline: "",
					Fill:    "",
				}

				_, err := drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, []asciidrawer.Figure{r})
				assert.True(t, errors.Is(err, asciidrawer.RecMustHaveFillOrOutlineErr))
			})

			t.Run("If we create a rectangle in a bad position it returns an error", func(t *testing.T) {
				r := &asciidrawer.Rectangle{
					Vertex: asciidrawer.Vertex{
						Row:    0,
						Column: -1,
					},
					Height:  1,
					Width:   1,
					Outline: "",
					Fill:    "",
				}

				_, err := drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, []asciidrawer.Figure{r})
				assert.True(t, errors.Is(err, asciidrawer.RecOutOfSquare))

				r.Vertex.Column = canvasWidth
				_, err = drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, []asciidrawer.Figure{r})
				assert.True(t, errors.Is(err, asciidrawer.RecOutOfSquare))

				r.Vertex.Column = 0
				r.Vertex.Row = -1
				_, err = drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, []asciidrawer.Figure{r})
				assert.True(t, errors.Is(err, asciidrawer.RecOutOfSquare))

				r.Vertex.Row = canvasHeight
				_, err = drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, []asciidrawer.Figure{r})
				assert.True(t, errors.Is(err, asciidrawer.RecOutOfSquare))
			})

			t.Run("If we create a rectangle with no height or width returns an error", func(t *testing.T) {
				r := &asciidrawer.Rectangle{
					Vertex: asciidrawer.Vertex{
						Row:    0,
						Column: 0,
					},
					Height:  1,
					Width:   0,
					Outline: "A",
					Fill:    "A",
				}

				_, err := drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, []asciidrawer.Figure{r})
				assert.True(t, errors.Is(err, asciidrawer.NoDimRecErr))

				r.Height = 0
				_, err = drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, []asciidrawer.Figure{r})
				assert.True(t, errors.Is(err, asciidrawer.NoDimRecErr))
			})

			t.Run("If we create a valid rectangle it returns the serialization", func(t *testing.T) {
				r := &asciidrawer.Rectangle{
					Vertex: asciidrawer.Vertex{
						Row:    0,
						Column: 0,
					},
					Height:  1,
					Width:   1,
					Outline: "A",
					Fill:    "A",
				}

				s, err := drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, []asciidrawer.Figure{r})
				assert.Nil(t, err)
				assert.True(t, s.ID != "")
				assert.Equal(t, canvasWidth, s.CanvasWidth)
				assert.Equal(t, canvasHeight, s.CanvasHeight)
			})
		})
	})
}

func TestDrawer_Draw(t *testing.T) {
	ctx := context.Background()
	m := mongo.NewMongoClient(mongo.DefaultConnectionString(), "test", nil)
	storage := repository.NewCanvasStorage(repository.DefaultDatabaseConfig(), m, logs.InitTest())
	drawer := asciidrawer.NewDrawer(storage)

	t.Run("Given we want to draw a canvas", func(t *testing.T) {
		t.Run("When the canvas doesn't exist it returns an error", func(t *testing.T) {
			_, err := drawer.Draw(ctx, primitive.NewObjectID().Hex())
			assert.Equal(t, asciidrawer.ErrSerializationDoesNotExists, err)
		})

		t.Run("When we draw squares", func(t *testing.T) {
			t.Run("Ir doesn't fill squares without fill", func(t *testing.T) {
				canvasWidth := 3
				canvasHeight := 4
				f := []asciidrawer.Figure{
					&asciidrawer.Rectangle{
						Vertex: asciidrawer.Vertex{
							Row:    0,
							Column: 0,
						},
						Height:  3,
						Width:   3,
						Outline: "X",
						Fill:    "",
					},
				}

				expected :=
					"XXX\n" +
						"X X\n" +
						"XXX\n" +
						"   "

				s, err := drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, f)
				require.Nil(t, err)

				got, err := drawer.Draw(ctx, s.ID)
				require.Nil(t, err)

				assert.Equal(t, expected, got)
			})

			t.Run("It outlines with the fill when the outline char is not specified", func(t *testing.T) {
				canvasWidth := 5
				canvasHeight := 4
				f := []asciidrawer.Figure{
					&asciidrawer.Rectangle{
						Vertex: asciidrawer.Vertex{
							Row:    0,
							Column: 0,
						},
						Height:  4,
						Width:   4,
						Outline: "",
						Fill:    "O",
					},
				}

				expected :=
					"OOOO \n" +
						"OOOO \n" +
						"OOOO \n" +
						"OOOO "

				s, err := drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, f)
				require.Nil(t, err)

				got, err := drawer.Draw(ctx, s.ID)
				require.Nil(t, err)

				assert.Equal(t, expected, got)
			})

			t.Run("It sets the fill and outline properly ", func(t *testing.T) {
				canvasWidth := 24
				canvasHeight := 10
				f := []asciidrawer.Figure{
					&asciidrawer.Rectangle{
						Vertex: asciidrawer.Vertex{
							Row:    2,
							Column: 3,
						},
						Height:  3,
						Width:   5,
						Outline: "@",
						Fill:    "X",
					},
					&asciidrawer.Rectangle{
						Vertex: asciidrawer.Vertex{
							Row:    3,
							Column: 10,
						},
						Height:  6,
						Width:   14,
						Outline: "X",
						Fill:    "O",
					},
				}

				expected :=
					"                        \n" +
						"                        \n" +
						"   @@@@@                \n" +
						"   @XXX@  XXXXXXXXXXXXXX\n" +
						"   @@@@@  XOOOOOOOOOOOOX\n" +
						"          XOOOOOOOOOOOOX\n" +
						"          XOOOOOOOOOOOOX\n" +
						"          XOOOOOOOOOOOOX\n" +
						"          XXXXXXXXXXXXXX\n" +
						"                        "

				s, err := drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, f)
				require.Nil(t, err)

				got, err := drawer.Draw(ctx, s.ID)
				require.Nil(t, err)

				assert.Equal(t, expected, got)
			})

			t.Run("It overlaps the squares properly ", func(t *testing.T) {
				canvasWidth := 21
				canvasHeight := 8
				f := []asciidrawer.Figure{
					&asciidrawer.Rectangle{
						Vertex: asciidrawer.Vertex{
							Row:    0,
							Column: 14,
						},
						Width:   7,
						Height:  6,
						Outline: "",
						Fill:    ".",
					},
					&asciidrawer.Rectangle{
						Vertex: asciidrawer.Vertex{
							Row:    3,
							Column: 0,
						},
						Width:   8,
						Height:  4,
						Outline: "O",
						Fill:    "",
					},
					&asciidrawer.Rectangle{
						Vertex: asciidrawer.Vertex{
							Row:    5,
							Column: 5,
						},
						Width:   5,
						Height:  3,
						Outline: "X",
						Fill:    "X",
					},
				}

				expected :=
					"              .......\n" +
						"              .......\n" +
						"              .......\n" +
						"OOOOOOOO      .......\n" +
						"O      O      .......\n" +
						"O    XXXXX    .......\n" +
						"OOOOOXXXXX           \n" +
						"     XXXXX           "

				s, err := drawer.CreateDrawing(ctx, canvasHeight, canvasWidth, f)
				require.Nil(t, err)

				got, err := drawer.Draw(ctx, s.ID)
				require.Nil(t, err)

				assert.Equal(t, expected, got)
			})
		})
	})
}
