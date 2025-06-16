package main

import (
	"context"
	"log"

	"gophermarket/internal/app"
	"gophermarket/internal/app/config"
	"gophermarket/internal/libs/graceful"
	"gophermarket/internal/libs/logger"
)

func main() {
	l, err := logger.New()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		_ = l.Sync()
	}()

	cfg, err := config.New()
	if err != nil {
		l.Fatal(err.Error())
		return
	}

	a := app.New(l, cfg)
	gr := graceful.New(
		graceful.NewProcess(a.HTTPAdapter),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = gr.Start(ctx)
	if err != nil {
		l.Fatal(err.Error())
		panic(err)
	}
}
