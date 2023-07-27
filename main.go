package main

import (
	"flag"
	"log"

	"github.com/fredrikaverpil/go-api-std/api"
	"github.com/fredrikaverpil/go-api-std/stores"
)

func main() {
	// Parse command line flags
	listenAddr := flag.String("listenaddr", ":8080", "server listen address")
	flag.Parse()

	// store := stores.DummyStore{}

	store, err := stores.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new server
	server := api.NewServer(*listenAddr, &store)
	println("Server running on port", *listenAddr)

	log.Fatal(server.Run())
}
