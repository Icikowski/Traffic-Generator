package utils

import (
	"os"
	"os/signal"
	"syscall"
)

// RegisterGracefulExitHooks adds support for OS-specific interruption signals and pre-exit hooks
func RegisterGracefulExitHooks(hook func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		hook()
		os.Exit(0)
	}()
}
