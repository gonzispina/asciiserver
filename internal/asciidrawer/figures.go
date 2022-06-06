package asciidrawer

// Figure contract
type figure interface {
	Draw(d Drawer, c *Canvas) error
	Serialize(s Serializer) (string, error)
}

// Rectangle representation
type Rectangle struct {
	vertex  vertex
	height  uint8
	width   uint8
	Outline byte
	Fill    byte
}

func (r *Rectangle) Draw(d Drawer, c *Canvas) error {
	return d.drawRectangle(c, r)
}

func (r *Rectangle) Serialize(s Serializer) string {
	return s.serializeRectangle(r)
}
