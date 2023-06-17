package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dgraph-io/dgo/v230"
	"github.com/jack-watts/go-dgraph-starter/app/services/api/config"
)

type Application struct {
	DG *dgo.Dgraph
}

func main() {
	var cfg config.Config
	cfg.Initialize()

	// Set CMD Flags
	flag.StringVar(&cfg.Web.Addr, "web-host", "localhost:8081", "set the hostname for the service. default: localhost:8081")
	flag.StringVar(&cfg.DgURL, "dgraph-host", "localhost:9080", "set the DGraph host url. default: localhost:9080")

	if err := run(&cfg); err != nil {
		log.Fatalf("Unable to start service: %s\n", err)
	}
}

func run(cfg *config.Config) error {

	// =============================================================================
	// inititate service startup

	log.Print("starting service")
	defer log.Print("shutdown complete")

	// =============================================================================
	// get a new DGraph client connection
	dg, cancel := getDgraphClient(cfg)
	defer cancel()

	app := Application{
		DG: dg,
	}

	// =============================================================================
	// start the API service
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	api := http.Server{
		Addr:         cfg.Web.Addr,
		Handler:      app.routes(),
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Starting api V1 router, host: %s\n", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// =============================================================================
	// set up graceful shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Printf("graceful shutdown started:signal:%s\n", sig)
		defer log.Printf("graceful shutdown complete:signal:%s\n", sig)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil

}
