package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/luquxSentinel/housebid/service"
	"github.com/luquxSentinel/housebid/storage"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// new storage
	storage, err := storage.New()
	if err != nil {
		log.Fatal(err)
	}

	// new auth service
	auth_service := service.NewAuthService(storage)

	// new api server
	api := NewAPIServer(":3000", auth_service)

	// run api server
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
