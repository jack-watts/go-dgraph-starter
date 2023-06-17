package main

import (
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Unable to start sevrice: %s\n", err)
	}
}

func run() error {

	// =============================================================================
	// inititate service startup

	log.Print("starting service")
	defer log.Print("shutdown complete")

	api := http.Server{}

	if err := api.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
