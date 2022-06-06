package asciidrawer

import "github.com/gonzispina/gokit/context"

// Serializer use case
type Serializer interface {
	GetSerialization(ctx context.Context, id string) (*Serialization, error)
	CreateSerialization(ctx context.Context, serialization string) (*Serialization, error)
	serializeRectangle(r *Rectangle) string
}
