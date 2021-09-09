migrate:
	migrate -source file://database/migration \
			-database postgres://postgres:postgres@127.0.0.1:5432/go_graphql?sslmode=disable up

rollback:
	migrate -source file://database/migration \
			-database postgres://postgres:postgres@127.0.0.1:5432/go_graphql?sslmode=disable down

drop:
	migrate -source file://database/migration \
			-database postgres://postgres:postgres@127.0.0.1:5432/go_graphql?sslmode=disable drop

migration:
	migrate create -ext sql -dir database/migration go_graphql

run:
	go run main.go

gqlgen:
	go run github.com/99designs/gqlgen