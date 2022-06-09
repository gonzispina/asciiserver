package asciidrawer

type FigureType int

const (
	// TypeRectangle FigureType
	TypeRectangle FigureType = iota
)

type FigureVisitor interface {
	visitRectangle(r *Rectangle)
}

// Figure contract
type Figure interface {
	accept(d FigureVisitor)
	Type() FigureType
}

// Rectangle representation
type Rectangle struct {
	Vertex  Vertex `bson:"vertex"`
	Height  int    `bson:"height"`
	Width   int    `bson:"width"`
	Outline string `bson:"outline"`
	Fill    string `bson:"fill"`
}

func (r *Rectangle) accept(f FigureVisitor) {
	f.visitRectangle(r)
}

func (r *Rectangle) Type() FigureType {
	return TypeRectangle
}
