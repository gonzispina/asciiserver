package factory

import (
	"github.com/gonzispina/asciiserver/internal/asciidrawer/repository"
	"github.com/gonzispina/gokit/logs"
	"github.com/gonzispina/gokit/mongo"
)

func NewRepositoriesFactory(
	mongo *mongo.Mongo,
	logger logs.Logger,
) *Repositories {
	if mongo == nil {
		panic("mongo must be initialized")
	}
	if logger == nil {
		panic("logger must be initialized")
	}
	return &Repositories{
		mongo:  mongo,
		logger: logger,
	}
}

// Repositories factory
type Repositories struct {
	mongo  *mongo.Mongo
	logger logs.Logger

	canvasStorage *repository.CanvasStorage
}

/*
 *
 * Canvas
 *
 */

// CanvasStorage repository
func (f *Repositories) CanvasStorage() *repository.CanvasStorage {
	if f.canvasStorage == nil {
		f.canvasStorage = repository.NewCanvasStorage(repository.DefaultDatabaseConfig(), f.mongo, f.logger)
	}
	return f.canvasStorage
}
