package main

import (
	"os"
	"os/signal"

	_ "github.com/lets-go-go/helper/process/rerun"
	"github.com/lets-go-go/logger"
)

func main() {
	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	logger.Infoln("hello world")

	<-stopChan // wait for SIGINT

	logger.Infoln("gracefully stopped stopped")
}
