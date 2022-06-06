package asciidrawer

import "github.com/gonzispina/gokit/context"

type Drawer interface {
	Draw(ctx context.Context, sID string) (*Canvas, error)
	drawRectangle(canvas *Canvas, r *Rectangle) error
}

type drawer struct {
	serializer Serializer
}

func (d *drawer) Draw(ctx context.Context, sID string) (*Canvas, error) {
	s, err := d.serializer.GetSerialization(ctx, sID)
	if err != nil {
		return nil, err
	}

	canvas := newCanvas(s.CanvasSize)
	for _, f := range s.Figures {
		err := f.Draw(d, canvas)
		if err != nil {
			return nil, err
		}
	}

	return canvas, nil
}

func (d *drawer) drawRectangle(c *Canvas, r *Rectangle) error {
	panic("implement me")
}
