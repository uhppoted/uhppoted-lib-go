package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	lib "github.com/uhppoted/uhppoted-lib-go/uhppoted"
)

func listen(u lib.Uhppoted, args []string) error {
	events := make(chan lib.ListenerEvent)
	errors := make(chan error)
	interrupt := make(chan os.Signal, 1)

	defer close(events)
	defer close(errors)
	defer close(interrupt)

	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for evt := range events {
			if bytes, err := json.MarshalIndent(evt, "   ", "   "); err == nil {
				fmt.Printf("event\n")
				fmt.Printf("   %v\n", string(bytes))
				fmt.Println()
			}
		}
	}()

	go func() {
		for err := range errors {
			fmt.Printf("*** ERROR %v\n", err)
		}
	}()

	if err := lib.Listen(u, events, errors, interrupt); err != nil {
		return err
	}

	return nil
}
