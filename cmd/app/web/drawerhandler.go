package web

import (
	"errors"
	"github.com/gonzispina/asciiserver/internal/asciidrawer"
	"github.com/gonzispina/gokit/logs"
	"github.com/gonzispina/gokit/rest"
)

// NewCanvasHandler constructor
func NewCanvasHandler(
	drawer asciidrawer.Drawer,
	logger logs.Logger,
) *CanvasHandler {
	if drawer == nil {
		panic("drawer must be initialized")
	}
	if logger == nil {
		panic("must be initialized")
	}
	return &CanvasHandler{
		drawer: drawer,
		log:    logger,
	}
}

// CanvasHandler holds the gateway the drawer API
type CanvasHandler struct {
	drawer asciidrawer.Drawer
	log    logs.Logger
}

// GetCanvas entry point
func (h *CanvasHandler) GetCanvas(r *rest.Request) *rest.Response {
	ctx := r.Context()
	canvasID := r.URLParam("id")
	res, err := h.drawer.Draw(ctx, canvasID)
	if err != nil {
		if errors.Is(err, asciidrawer.ErrSerializationDoesNotExists) {
			return rest.NotFound(err)
		}
		h.log.Error(ctx, "Couldn't get drawing", logs.Error(err))
		return rest.InternalServerError()
	}

	return rest.OK(res)
}
