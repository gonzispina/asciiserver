package asciidrawer

// newCanvas constructor
func newCanvas(height, width int) *canvas {
	rows := make([][]rune, height)
	for i := 0; i < height; i++ {
		rows[i] = make([]rune, width)
	}
	return &canvas{Height: height, Width: width, Rows: rows}
}

// canvas representation. Is a square matrix where every position represents a character.
type canvas struct {
	Height int
	Width  int
	Rows   [][]rune
}

func (c *canvas) String() string {
	str := ""
	for i, row := range c.Rows {
		for j, col := range row {
			if string(col) == "\x00" {
				c.Rows[i][j] = ' '
			}

			str += string(c.Rows[i][j])
		}

		if i != c.Height-1 {
			str += "\n"
			continue
		}
	}
	return str
}

func (c *canvas) visitRectangle(r *Rectangle) {
	getChar := func(fr, lr, fc, lc bool) rune {
		if r.Outline == "" {
			return rune(r.Fill[0])
		}

		if fr || lr || fc || lc {
			return rune(r.Outline[0])
		}

		if r.Fill == "" {
			return ' '
		}

		return rune(r.Fill[0])
	}

	for i := r.Vertex.Row; i < r.Vertex.Row+r.Height; i++ {
		fr := i == r.Vertex.Row
		lr := i == r.Vertex.Row+r.Height-1
		for j := r.Vertex.Column; j < r.Vertex.Column+r.Width; j++ {
			fc := j == r.Vertex.Column
			lc := j == r.Vertex.Column+r.Width-1
			char := getChar(fr, lr, fc, lc)
			c.Rows[i][j] = char
		}
	}

	return
}
