package main

import (
	"context"
	"log"

	"golang.design/x/hotkey/mainthread"

	"macos_keymapper/config"
)

func main() {
	ctx := context.Background()
	mainthread.Init(func() {
		parser := config.NewParser()
		parser.ParseConfigFile()
		parser.RegisterHotkeys()

		suspendThread(ctx)
	})
}

func suspendThread(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("Context is stopped")
		}
	}
}
