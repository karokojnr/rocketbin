package main

import (
	"log"

	"github.com/karokojnr/rocketbin/internal/db"
	"github.com/karokojnr/rocketbin/internal/rocket"
	"github.com/karokojnr/rocketbin/internal/transport/grpc"
)

func Run() error {
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("error running db migrations")
		return err
	}

	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
