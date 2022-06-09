package app

import (
	"github.com/gonzispina/asciiserver/internal/asciidrawer"
	"github.com/gonzispina/asciiserver/internal/asciidrawer/repository"
	"github.com/gonzispina/gokit/context"
	"github.com/gonzispina/gokit/logs"
)

func installDrawings(ctx context.Context, storage *repository.CanvasStorage, logger logs.Logger) {
	logger.Info(ctx, "Installing canvases...")

	sr := &asciidrawer.Serialization{
		ID:           "62a22d7b8a95a01c6aedfb0f",
		CanvasWidth:  21,
		CanvasHeight: 8,
		Figures: []asciidrawer.Figure{
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
		},
	}

	if err := storage.SaveSerialization(ctx, sr); err != nil {
		panic(err)
	}
}
