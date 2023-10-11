package main

import (
	"flag"
	"log"

	"github.com/fredrikaverpil/go-api-std/internal/rest"
	"github.com/fredrikaverpil/go-api-std/internal/services"
	"github.com/fredrikaverpil/go-api-std/internal/stores"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "server listen address")
	flag.Parse()

	// store := stores.NewDummyStore()
	store, err := stores.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	userService := services.NewService(store)

	server := rest.NewServer(*listenAddr, *userService)
	println("Server running on port", *listenAddr)

	log.Fatal(server.Run())
}
