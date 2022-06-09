package asciidrawer

import (
	"github.com/gonzispina/gokit/errors"
)

// ErrSerializationDoesNotExists ...
var ErrSerializationDoesNotExists = errors.New("serialization does not exist", "SerializationNotFound")

// RecMustHaveFillOrOutlineErr ...
var RecMustHaveFillOrOutlineErr = errors.New("rectangles must have at least one of fill or outline", "InvalidRec")

// RecOutOfSquare ...
var RecOutOfSquare = errors.New("rectangle is bigger than the canvas or is not in a correct position", "InvalidRec")

// NoDimRecErr ...
var NoDimRecErr = errors.New("rec must have at least one unit for the height and one unit for the width", "InvalidRec")
