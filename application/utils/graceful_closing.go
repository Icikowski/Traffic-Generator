package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func CreateShutdownHook(hook func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		hook()
		os.Exit(0)
	}()
}
