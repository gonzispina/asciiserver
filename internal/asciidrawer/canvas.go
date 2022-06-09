package asciidrawer

// newCanvas constructor
func newCanvas(size int) *Canvas {
	rows := make([][]byte, size)
	for i := 0; i < size; i++ {
		rows[i] = make([]uint8, size)
	}
	return &Canvas{Size: size, Rows: rows}
}

// Canvas representation. Is a square matrix where every position represents a character.
type Canvas struct {
	Size int
	Rows [][]byte
}

func (c *Canvas) visitRectangle(r *Rectangle) {
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

	for i := r.Vertex.Row; i < r.Height; i++ {
		fr := i == r.Vertex.Row
		lr := i == r.Height-1
		for j := r.Vertex.Column; j < r.Width; j++ {
			fc := j == r.Vertex.Column
			lc := j == r.Width-1

			c.Rows[j][i] = getChar(fr, lr, fc, lc)
		}
	}

	return
}
