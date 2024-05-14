package main

import (
	"log"

	"github.com/karokojnr/rocketbin/internal/db"
	"github.com/karokojnr/rocketbin/internal/rocket"
)

func Run() error {
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	_ = rocket.New(rocketStore)
	return nil
}

func main() {

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
