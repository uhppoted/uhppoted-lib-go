package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	lib "github.com/uhppoted/uhppoted-lib-go/src/uhppoted"
)

type listener struct {
	events chan lib.ListenerEvent
	errors chan error
}

func (l listener) OnEvent(evt lib.ListenerEvent) {
	l.events <- evt
}

func (l listener) OnError(err error) {
	l.errors <- err
}

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

	l := listener{
		events: events,
		errors: errors,
	}

	if err := lib.Listen(u, l, interrupt); err != nil {
		return err
	}

	return nil
}
