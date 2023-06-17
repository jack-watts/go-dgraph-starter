package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jack-watts/go-dgraph-starter/app/services/api/config"
)

func main() {
	var cfg config.Config
	cfg.Initialize()

	// Set CMD Flags
	flag.StringVar(&cfg.Web.Addr, "web-host", "localhost:8081", "set the hostname for the service. default: localhost:8081")
	flag.StringVar(&cfg.DgURL, "dgraph-host", "localhost:8080", "set the DGraph host url. default: localhost:8081")

	if err := run(&cfg); err != nil {
		log.Fatalf("Unable to start sevrice: %s\n", err)
	}
}

func run(cfg *config.Config) error {

	// =============================================================================
	// inititate service startup

	log.Print("starting service")
	defer log.Print("shutdown complete")

	api := http.Server{
		Addr:         cfg.Web.Addr,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
	}

	if err := api.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
