package asciidrawer

import "github.com/gonzispina/gokit/errors"

func newValidator(canvasSize int) *validator {
	return &validator{canvasSize: canvasSize}
}

type validator struct {
	canvasSize int
	err        error
}

func (v *validator) addErr(err errors.Error) {
	if v.err == nil {
		v.err = err
		return
	}
	v.err = err.Wrap(v.err)
}

func (v *validator) Err() error {
	return v.err
}

func (v *validator) visitRectangle(r *Rectangle) {
	if r.Outline == "" && r.Fill == "" {
		v.addErr(RecMustHaveFillOrOutlineErr)
	}

	if r.Height < 1 || r.Width < 1 {
		v.addErr(NoDimRecErr)
	}

	if r.Vertex.Column < 0 ||
		r.Vertex.Column+r.Width > v.canvasSize ||
		r.Vertex.Row < 0 ||
		r.Vertex.Row+r.Height > v.canvasSize {
		v.addErr(RecOutOfSquare)
	}
}
