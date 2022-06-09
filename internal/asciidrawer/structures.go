package asciidrawer

// Serialization representation
type Serialization struct {
	ID           string   `bson:"_id"`
	CanvasHeight int      `bson:"canvasHeight"`
	CanvasWidth  int      `bson:"canvasWidth"`
	Figures      []Figure `bson:"figures"`
}

// Vertex representation
type Vertex struct {
	Row    int `bson:"row"`
	Column int `bson:"column"`
}
