package main

import (
	"backend/container"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := &container.Config{
		AccessSecret: os.Getenv("JWT_ACCESS_SECRET"),
	}

	container := container.NewContainer(config)

	if err := container.Build(); err != nil {
		log.Fatal(err)
	}

	if err := container.Server().Listen(); err != nil {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
