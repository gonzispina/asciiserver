package asciidrawer

import (
	"fmt"
	"github.com/gonzispina/gokit/errors"
)

// ErrSerializationDoesNotExists ...
var ErrSerializationDoesNotExists = errors.New("serialization does not exist", "SerializationNotFound")

// UnexpectedEOF ...
var UnexpectedEOFErr = errors.New("unexpected end of string", "UnexpectedEndOfString")

// RecMustHaveFillOrOutlineErr ...
var RecMustHaveFillOrOutlineErr = errors.New("rectangles must have at least one of fill or outline", "InvalidRec")

// RecOutOfSquare ...
var RecOutOfSquare = errors.New("rectangle is bigger than the canvas or is not in a correct position", "InvalidRec")

// NoDimRecErr ...
var NoDimRecErr = errors.New("rec must have at least one unit for the height and one unit for the width", "InvalidRec")

// ExpectedTokenErr ...
func ExpectedTokenErr(s string, c int) error {
	return errors.New(fmt.Sprintf("expected %s token on column %v", s, c), "ExpectedToken")
}

// UnexpectedTokenErr ...
func UnexpectedTokenErr(s string, c int) error {
	return errors.New(fmt.Sprintf("unexpected %s token on column %v", s, c), "UnexpectedToken")
}

// ExpectedNumberErr ...
func ExpectedNumberErr(c int) error {
	return errors.New(fmt.Sprintf("expected number on column %v", c), "ExpectedNumber")
}
