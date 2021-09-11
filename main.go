package main

import (
	"context"
	"fmt"
	"go-gqlgen/config"
	pg "go-gqlgen/database"
	"go-gqlgen/graph/generated"
	graph "go-gqlgen/graph/resolver"
	"go-gqlgen/infrastructure"
	usecase "go-gqlgen/usecase"
	"log"
	"net/http"
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

	// call repository from infrastructure
	userRepo := infrastructure.NewUserRepository(db) // *infrastructure.UserRepository
	// call auth service from domain
	authUsecase := usecase.NewUsecaseAuthImpl(userRepo)
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
				Resolvers: &graph.Resolver{
					AuthRepository: authUsecase,
				},
			},
		),
	))

	fmt.Println("GraphQL Server is About to Running....")
	log.Fatal(http.ListenAndServe(":8080", r))

}
