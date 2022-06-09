package factory

import (
	"github.com/gonzispina/asciiserver/cmd/app/web"
	"github.com/gonzispina/gokit/logs"
)

// NewHandlersFactory returns an instance of a Handlers factory
func NewHandlersFactory(cases *Cases, logger logs.Logger) *Handlers {
	if cases == nil {
		panic("use case factory must be initialized")
	}
	if logger == nil {
		panic("logger must be initialized")
	}
	return &Handlers{
		cases:  cases,
		logger: logger,
	}
}

// Handlers factory
type Handlers struct {
	cases  *Cases
	logger logs.Logger

	// i18n
	canvasHandler *web.CanvasHandler
}

// CanvasHandler controller
func (f *Handlers) CanvasHandler() *web.CanvasHandler {
	if f.canvasHandler == nil {
		f.canvasHandler = web.NewCanvasHandler(
			f.cases.ASCIIDrawer(),
			f.logger,
		)
	}
	return f.canvasHandler
}
