package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gonzispina/asciiserver/cmd/app"
	"github.com/gonzispina/gokit/context"
	"github.com/gonzispina/gokit/logs"
)

var (
	// SIGTERM os.Signal
	SIGTERM os.Signal = syscall.SIGTERM
	// SIGTSTP os.Signal
	SIGTSTP os.Signal = syscall.SIGTSTP
	// SIGINT os.Signal
	SIGINT os.Signal = syscall.SIGINT
)

func main() {
	logger := logs.InitDefault()
	application := app.New(logger)
	application.Init()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, SIGINT)
	signal.Notify(sigchan, SIGTERM)
	signal.Notify(sigchan, SIGTSTP)

	<-sigchan

	logger.Info(context.Background(), "Shutting down application")
	application.Stop(time.Second * 30)
}
