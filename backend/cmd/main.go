package main

import (
	"log"

	"github.com/mishaRomanov/maglev/internal/config"
	"github.com/mishaRomanov/maglev/internal/service"
	"github.com/mishaRomanov/maglev/internal/storage"
)

func main() {
	log.Println("Service starting..")

	cfg := config.NewFromEnv()

	redisClient := storage.NewRedisInstance(
		cfg.Port,
		cfg.Password,
		cfg.DB,
	)

	// Create and run our backend.
	if fatalErr := service.New(cfg, redisClient).Run(); fatalErr != nil {
		log.Fatal(fatalErr)
	}

}
