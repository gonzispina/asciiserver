package asciidrawer

import (
	"github.com/gonzispina/gokit/context"
)

type CanvasStorage interface {
	CreateSerialization(ctx context.Context, canvasHeight, canvasWidth int, figures []Figure) (*Serialization, error)
	GetSerialization(ctx context.Context, id string) (*Serialization, error)
}

// Drawer use case
type Drawer interface {
	CreateDrawing(ctx context.Context, canvasHeight, canvasWidth int, figures []Figure) (*Serialization, error)
	Draw(ctx context.Context, sID string) (string, error)
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

func (d *drawer) CreateDrawing(ctx context.Context, canvasHeight, canvasWidth int, figures []Figure) (*Serialization, error) {
	v := newValidator(canvasHeight, canvasWidth)
	for _, f := range figures {
		f.accept(v)
	}

	if err := v.Err(); err != nil {
		return nil, err
	}

	return d.storage.CreateSerialization(ctx, canvasHeight, canvasWidth, figures)
}

func (d *drawer) Draw(ctx context.Context, sID string) (string, error) {
	s, err := d.storage.GetSerialization(ctx, sID)
	if err != nil {
		return "", err
	}

	c := newCanvas(s.CanvasHeight, s.CanvasWidth)
	for _, f := range s.Figures {
		f.accept(c)
	}

	return c.String(), nil
}
