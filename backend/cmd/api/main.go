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

	c := container.NewContainer(config)

	if err := c.Build(); err != nil {
		log.Fatal(err)
	}

	if err := c.Server().Listen(); err != nil {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
