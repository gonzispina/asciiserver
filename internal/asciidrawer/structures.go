package asciidrawer

// Serialization representation
type Serialization struct {
	ID         string   `bson:"_id"`
	CanvasSize int      `bson:"canvasSize"`
	Figures    []Figure `bson:"figures"`
}

// Vertex representation
type Vertex struct {
	Row    int `bson:"row"`
	Column int `bson:"column"`
}
