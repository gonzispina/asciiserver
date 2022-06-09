package app

import (
	"github.com/gonzispina/asciiserver/cmd/app/web"
	"github.com/gonzispina/asciiserver/cmd/factory"
	"github.com/gonzispina/gokit/context"
	"github.com/gonzispina/gokit/logs"
	"github.com/gonzispina/gokit/mongo"
	"net/http"
	"time"
)

// New application
func New(logger logs.Logger) *App {
	return &App{log: logger}
}

// App ...
type App struct {
	done chan struct{}
	log  logs.Logger

	mongo        *mongo.Mongo
	cases        *factory.Cases
	repositories *factory.Repositories

	mux http.Handler
}

// Cases layer instances
func (a *App) Cases() *factory.Cases {
	return a.cases
}

// Repositories layer instances
func (a *App) Repositories() *factory.Repositories {
	return a.repositories
}

// Mux of the application
func (a *App) Mux() http.Handler {
	return a.mux
}

// Init the application
func (a *App) Init() {
	a.mongo = mongo.NewMongoClient(mongo.DefaultConnectionString(), "ascii", nil)

	a.repositories = factory.NewRepositoriesFactory(a.mongo, a.log)
	a.cases = factory.NewCasesFactory(a.repositories)
	handlers := factory.NewHandlersFactory(a.cases, a.log)

	installCanvases(context.Background(), a.repositories.CanvasStorage(), a.log)

	a.mux = web.NewRouter(
		handlers.CanvasHandler(),
		a.log,
	)

	ctx := context.Background()
	stopServer := InitializeServer(ctx, "8080", a.mux, a.log)

	a.done = make(chan struct{}, 1)
	go func() {
		<-a.done

		a.log.Info(ctx, "Stopping server...")
		stopServer()

		ctx, _ = context.WithTimeout(context.Background(), time.Second*10)
		_ = a.mongo.Close(ctx)
	}()
}

// Stop the application gracefully
func (a *App) Stop() {
	close(a.done)
}
