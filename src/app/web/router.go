package web

import (
	"github.com/go-chi/chi"
	"github.com/gonzispina/gokit/logs"
	"github.com/gonzispina/gokit/rest"
)

// NewRouter add routes to web multiplexer
func NewRouter(
	canvas *CanvasHandler,
	logger logs.Logger,
) *chi.Mux {
	mux := chi.NewRouter()

	upgradeMiddleware := rest.UpgradeMiddleware(logger)
	mux.Group(func(r chi.Router) {
		r.Get("/canvas/{id}", upgradeMiddleware(canvas.GetCanvas))
	})

	return mux
}
