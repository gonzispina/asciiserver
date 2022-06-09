package asciidrawer

import (
	"github.com/gonzispina/gokit/context"
)

type CanvasStorage interface {
	CreateSerialization(ctx context.Context, canvasSize int, figures []Figure) (*Serialization, error)
	GetSerialization(ctx context.Context, id string) (*Serialization, error)
}

// Drawer use case
type Drawer interface {
	CreateDrawing(ctx context.Context, canvasSize int, figures []Figure) (*Serialization, error)
	Draw(ctx context.Context, sID string) (*Canvas, error)
}

// NewDrawer constructor
func NewDrawer(s CanvasStorage) Drawer {
	if s == nil {
		panic("storage must be initialized")
	}
	return &drawer{storage: s}
}

type drawer struct {
	storage CanvasStorage
}

func (d *drawer) CreateDrawing(ctx context.Context, canvasSize int, figures []Figure) (*Serialization, error) {
	return d.storage.CreateSerialization(ctx, canvasSize, figures)
}

func (d *drawer) Draw(ctx context.Context, sID string) (*Canvas, error) {
	s, err := d.storage.GetSerialization(ctx, sID)
	if err != nil {
		return nil, err
	}

	canvas := newCanvas(s.CanvasSize)
	for _, f := range s.Figures {
		f.Accept(canvas)
	}

	return canvas, nil
}
