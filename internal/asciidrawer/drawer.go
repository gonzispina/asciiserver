package asciidrawer

import (
	"github.com/gonzispina/gokit/context"
)

type CanvasStorage interface {
	SaveSerialization(ctx context.Context, serialization string) (string, error)
	GetSerialization(ctx context.Context, id string) (*Serialization, error)
}

// Drawer use case
type Drawer interface {
	Draw(ctx context.Context, sID string) (*Canvas, error)
	drawRectangle(canvas *Canvas, r *Rectangle)
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

func (d *drawer) CreateDrawing(ctx context.Context, serialization string) (*Serialization, error) {
	// First we validate the string
	s, err := newSerializer(serialization).deserialize()
	if err != nil {
		return nil, err
	}

	id, err := d.storage.SaveSerialization(ctx, serialization)
	if err != nil {
		return nil, err
	}

	s.ID = id
	return s, nil
}

func (d *drawer) Draw(ctx context.Context, sID string) (*Canvas, error) {
	s, err := d.storage.GetSerialization(ctx, sID)
	if err != nil {
		return nil, err
	}

	canvas := newCanvas(s.CanvasSize)
	for _, f := range s.Figures {
		f.Draw(d, canvas)
	}

	return canvas, nil
}

func (d *drawer) drawRectangle(c *Canvas, r *Rectangle) {
	getChar := func(fr, lr, fc, lc bool) byte {
		if r.Outline == "" {
			return r.Fill[0]
		}

		if fr || lr || fc || lc {
			return r.Outline[0]
		}

		if r.Fill == "" {
			return " "[0]
		}

		return r.Fill[0]
	}

	for i := r.vertex.row; i < r.Height; i++ {
		fr := i == r.vertex.row
		lr := i == r.Height-1
		for j := r.vertex.column; j < r.Width; j++ {
			fc := j == r.vertex.column
			lc := j == r.Width-1

			c.Rows[j][i] = getChar(fr, lr, fc, lc)
		}
	}

	return
}
