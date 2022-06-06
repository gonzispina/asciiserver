package asciidrawer

// Serialization representation
type Serialization struct {
	ID         string   `bson:"_id"`
	Str        string   `bson:"str"`
	CanvasSize int      `bson:"-"`
	Figures    []figure `bson:"-"`
}

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

// Vertex representation
type vertex struct {
	row    int
	column int
}
