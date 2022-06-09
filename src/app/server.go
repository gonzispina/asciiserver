package app

import (
	"github.com/gonzispina/gokit/context"
	"github.com/gonzispina/gokit/logs"
	"net/http"
	"strconv"
	"time"
)

// InitializeServer runs ListenAndServe on the http.Server
func InitializeServer(ctx context.Context, port string, mux http.Handler, logger logs.Logger) context.CancelFunc {
	s := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic("Could not initialize server " + err.Error())
		}
	}()

	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		<-ctx.Done()

		logger.Info(ctx, "Shutting down server...")
		logger.Info(ctx, "Shutdown time out: "+strconv.Itoa(30)+" seconds")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		s.SetKeepAlivesEnabled(false)
		if err := s.Shutdown(ctx); err != nil {
			logger.Error(ctx, "Could not shut down server gracefully", logs.Error(err))
		}

		logger.Info(ctx, "Server stopped")
	}(ctx)

	return cancel
}

