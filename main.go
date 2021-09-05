package main

import (
	"context"
	"fmt"
	"go-gqlgen/config"
	pg "go-gqlgen/database"
	"log"
)

func main() {
	ctx := context.Background()

	config := config.New()

	db := pg.New(ctx, config)

	if err := db.Migration(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started")
}
