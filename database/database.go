package db

import (
	"context"
	"fmt"
	"go-gqlgen/config"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Pool   *pgxpool.Pool
	config *config.Config
}

func New(ctx context.Context, config *config.Config) *DB {
	dbConfig, err := pgxpool.ParseConfig(config.Database.URL)
	if err != nil {
		log.Fatalf("can't parse postgres config: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, dbConfig)
	if err != nil {
		log.Fatalf(("error connection to database: %v"), err)
	}

	db := &DB{Pool: pool}

	db.Open(ctx)
	return db
}

func (db *DB) Open(ctx context.Context) {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Fatal("can't open postgres: %")
	}

	log.Println("connected to postgres")
}

func (db *DB) Migration() error {

	// _, b, _, _ := runtime.Caller(0)

	// migrationPath := fmt.Sprintf("file:///database/migration", path.Dir(b))

	m, err := migrate.New("file://database/migration", "postgres://postgres:postgres@127.0.0.1:5432/go_graphql?sslmode=disable") //db.config.Database.URL)
	println(m)
	if err != nil {
		return fmt.Errorf("error happened when migration %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error when migration up: %v", err)
	}

	log.Println("migration completed!")
	return err
}

func (db *DB) Close() {
	db.Pool.Close()
}
