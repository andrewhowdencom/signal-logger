package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handle(sig os.Signal) {
	switch sig {
	case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
		fmt.Printf("Received %s. Exiting\n", sig)
		os.Exit(0)
	default:
		fmt.Printf("Received %s\n", sig)
	}
}

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	defer close(done)

	signal.Notify(sigs)

	for sig := range sigs {
		handle(sig)
	}
}
