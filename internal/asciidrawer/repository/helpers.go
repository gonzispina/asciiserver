package repository

import (
	"github.com/gonzispina/asciiserver/internal/asciidrawer"
	"go.mongodb.org/mongo-driver/bson"
)

type figure struct {
	Type   asciidrawer.FigureType `bson:"type"`
	Figure interface{}            `bson:"data"`
}

func mapFigures(fs []asciidrawer.Figure) bson.A {
	res := make(bson.A, len(fs))
	for i, f := range fs {
		res[i] = figure{
			Type:   f.Type(),
			Figure: f,
		}
	}
	return res
}
