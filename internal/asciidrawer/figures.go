package asciidrawer

// Figure contract
type figure interface {
	Draw(d Drawer, c *Canvas)
	Serialize(s *serializer) string
}

// Rectangle representation
type Rectangle struct {
	vertex  vertex
	Height  int
	Width   int
	Outline string
	Fill    string
}

func (r *Rectangle) Draw(d Drawer, c *Canvas) {
	d.drawRectangle(c, r)
}

func (r *Rectangle) Serialize(s *serializer) string {
	return s.serializeRectangle(r)
}
