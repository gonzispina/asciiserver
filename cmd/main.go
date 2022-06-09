package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gonzispina/asciiserver/cmd/app"
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
	application := app.New()
	application.Init()

	// Gracefully shutdown
	defer application.Stop(time.Second * 30)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, SIGINT)
	signal.Notify(sigchan, SIGTERM)
	signal.Notify(sigchan, SIGTSTP)

	<-sigchan
}
