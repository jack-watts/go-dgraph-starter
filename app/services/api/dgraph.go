package main

import (
	"log"

	"github.com/dgraph-io/dgo/v230"
	"github.com/dgraph-io/dgo/v230/protos/api"
	"github.com/jack-watts/go-dgraph-starter/app/services/api/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CancelFunc func()

func getDgraphClient(cfg *config.Config) (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial(cfg.DgURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}
