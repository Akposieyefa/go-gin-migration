package main

import (
	"log"

	"github.com/akposiyefa/go-gin-migration/cmd/api"
	"github.com/akposiyefa/go-gin-migration/config"
	"github.com/akposiyefa/go-gin-migration/internal"
)

func main() {
	internal.ConnectToDB()
	server := api.NewAPIServer(config.Envs.APP_PORT)
	if err := server.Run(); err != nil {
		log.Fatal("Falied to run api server...")
	}
}
