package graph

import domain "go-gqlgen/domain/interface"

// go:generate go run github.com/99designs/gqlgen

// injection with outside of GraphQL such as DB
type Resolver struct {
	// access to domain interface so we can access to the actual function
	AuthRepository domain.AuthRepository
}
