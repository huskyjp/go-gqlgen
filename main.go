package main

import (
	"context"
	"fmt"
	"go-gqlgen/config"
	pg "go-gqlgen/database"
	"go-gqlgen/graph/generated"
	graph "go-gqlgen/graph/resolver"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	ctx := context.Background()

	config := config.New()

	db := pg.New(ctx, config)

	if err := db.Migration(); err != nil {
		log.Fatal(err)
	}

	// setting router with chi

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Timeout(time.Second * 60))

	r.Handle("/", playground.Handler("GraphQL Golang", "/query"))
	r.Handle("/query", handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	))

	fmt.Println("GraphQL Server is About to Running....")

}
