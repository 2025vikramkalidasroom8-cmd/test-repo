package main

import (
	"fmt"
	"log"

	watchdog "github.com/Mritunjay2005/watchDog"
)

func main() {
	watcher := watchdog.New()

	if err := watcher.Watch("./**/*.go"); err != nil {
		log.Fatal(err)
	}

	watcher.OnCreate(func(event watchdog.CreateEvent) {
		fmt.Println("Created:", event.Path)
	})

	watcher.OnWrite(func(event watchdog.WriteEvent) {
		fmt.Println("Modified:", event.Path)
	})

	watcher.OnRemove(func(event watchdog.RemoveEvent) {
		fmt.Println("Deleted:", event.Path)
	})

	fmt.Println("Watching Go files...")

	if err := watcher.Start(); err != nil {
		log.Fatal(err)
	}
}
