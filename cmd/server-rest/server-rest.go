package main

import (
	"flag"
	"log"

	"github.com/fredrikaverpil/go-api-std/internal/rest"
	"github.com/fredrikaverpil/go-api-std/internal/services/user"
	"github.com/fredrikaverpil/go-api-std/internal/stores"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	listenAddr := flag.String("listenaddr", ":8080", "server listen address")
	flag.Parse()

	// store := stores.NewDummyStore()
	store, err := stores.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	userService := user.NewService(store)

	server := rest.NewServer(*listenAddr, *userService)
	log.Printf("Server running on http://localhost%s", *listenAddr)

	log.Fatal(server.Run())
}
