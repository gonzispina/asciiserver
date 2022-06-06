package asciidrawer

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// newSerializer constructor
func newSerializer(serialization string) *serializer {
	return &serializer{str: strings.ToUpper(serialization), i: 0}
}

type serializer struct {
	str string
	i   int
}

func (s *serializer) advance() (rune, error) {
	if s.i >= len(s.str) {
		return ' ', UnexpectedEOFErr
	}
	s.i++
	return rune(s.str[s.i]), nil
}

func (s *serializer) next() (rune, error) {
	if s.i+1 >= len(s.str) {
		return ' ', io.EOF
	}
	return rune(s.str[s.i+1]), nil
}

func (s *serializer) match(r rune) bool {
	return rune(s.str[s.i]) == r
}

func (s *serializer) getNumber() (int, error) {
	c, err := s.advance()
	if err != nil {
		return 0, err
	}

	if !isNumber(c) {
		return 0, ExpectedNumberErr(s.i)
	}

	number := string(c)
	for c, _ = s.advance(); isNumber(c); {
		number += string(c)
	}

	res, _ := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(res), nil
}

func (s *serializer) deserialize() (*Serialization, error) {
	current, _ := s.advance()
	if current != 'C' {
		return nil, ExpectedTokenErr("C", s.i)
	}

	current, _ = s.advance()
	if !isNumber(current) {
		return nil, ExpectedNumberErr(s.i)
	}

	size, _ := strconv.ParseInt(string(current), 10, 64)
	res := &Serialization{
		Str:        s.str,
		CanvasSize: int(size),
	}

	for current, err := s.advance(); err != nil; {
		if !isLetter(current) {
			return nil, ExpectedTokenErr("letter", s.i)
		}

		// If this grows a lot we can implement a double dispatch as we did with the s.str and drawing
		switch current {
		case 'R':
			r, err := s.deserializeRectangle(res)
			if err != nil {
				return nil, err
			}
			res.Figures = append(res.Figures, r)
			break
		default:
			return nil, UnexpectedTokenErr(string(current), s.i)
		}
	}

	return res, nil
}

func (s *serializer) deserializeRectangle(sr *Serialization) (*Rectangle, error) {
	var c rune
	var err error
	var res Rectangle
	for c, err = s.advance(); err != nil; {
		switch c {
		case 'V':
			vx, err := s.getNumber()
			if err != nil {
				return nil, err
			}

			c, err = s.advance()
			if err != nil {
				return nil, err
			}

			if c != ';' {
				return nil, ExpectedTokenErr(";", s.i)
			}

			vy, err := s.getNumber()
			if err != nil {
				return nil, err
			}
			res.vertex = vertex{row: vy, column: vx}
			break
		case 'H':
			h, err := s.getNumber()
			if err != nil {
				return nil, err
			}
			res.Height = h
			break
		case 'W':
			w, err := s.getNumber()
			if err != nil {
				return nil, err
			}
			res.Width = w
			break
		case 'O':
			c, err = s.advance()
			if err != nil {
				return nil, err
			}
			res.Outline = string(c)
			break
		case 'F':
			c, err = s.advance()
			if err != nil {
				return nil, err
			}
			res.Fill = string(c)
			break
		default:
			// Let's validate the rectangle
			if res.Outline == "" && res.Fill == "" {
				return nil, RecMustHaveFillOrOutlineErr
			}

			if res.vertex.column < 1 || res.Width >= sr.CanvasSize ||
				res.vertex.row < 0 || res.Height >= sr.CanvasSize {
				return nil, RecOutOfSquare
			}

			if res.Height < 1 || res.Width < 1 {
				return nil, NoDimRecErr
			}
		}
	}

	if err != nil {
		return nil, err
	}

	// Unreachable
	return nil, nil
}

func (s *serializer) serializeRectangle(r *Rectangle) string {
	str := fmt.Sprintf("RV%v;%vH%vW%v")
	if r.Outline != "" {
		str = fmt.Sprintf("%sO%s", str, r.Outline)
	}
	if r.Fill != "" {
		str = fmt.Sprintf("%sF%s", str, r.Fill)
	}
	return str
}
