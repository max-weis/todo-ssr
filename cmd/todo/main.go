package main

import (
	"log/slog"
	"sync"
)

func main() {
	app, err := InitApp()
	if err != nil {
		panic(err)
	}

	app.Log.Info("starting server", slog.String("port", app.cfg.Port))

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := app.Start(app.cfg.HttpPort()); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
