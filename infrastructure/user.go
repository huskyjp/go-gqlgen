package infrastructure

import (
	"context"
	"fmt"
	db "go-gqlgen/database"
	domain "go-gqlgen/domain/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

// make accessible to database
type UserRepository struct {
	DB *db.DB
}

func NewUserRepository(db *db.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (udb *UserRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	flow, err := udb.DB.Pool.Begin(ctx)
	if err != nil {
		return domain.User{}, fmt.Errorf("error happend when starting query transaction flow: %v", err)
	}
	// rollback once error happens so we can keep constant database state
	defer flow.Rollback(ctx)

	user, err = createUser(ctx, flow, user)
	if err != nil {
		return domain.User{}, err
	}

	err = flow.Commit(ctx)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func createUser(ctx context.Context, flow pgx.Tx, user domain.User) (domain.User, error) {
	q := `INSERT INTO users (email, username, password) 
				VALUES ($1, $2, $3)
				RETURNING *;`

	u := domain.User{}

	if err := pgxscan.Get(ctx, flow, &u, q, user.Email, user.Username, user.Password); err != nil {
		return domain.User{}, err
	}

	return u, nil
}

// TODO: Implement user generation
func (udb *UserRepository) GenerateUser(ctx context.Context, user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

// function finds user by username
func (udb *UserRepository) GetByUserName(ctx context.Context, username string) (domain.User, error) {
	q := `SELECT * FROM users
				WHERE username = $1 LIMIT 1;`

	u := domain.User{}

	// user will be assigned automatically
	err := pgxscan.Get(ctx, udb.DB.Pool, &u, q)
	if err != nil {
		return domain.User{}, fmt.Errorf("error happend when GET not found error: %v", err)
	}

	return u, nil

}

func (udb *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	q := `SELECT * FROM users
				WHERE email = $1 LIMIT 1;`

	u := domain.User{}

	// user will be assigned automatically
	err := pgxscan.Get(ctx, udb.DB.Pool, &u, q)
	if err != nil {
		return domain.User{}, fmt.Errorf("error happend when GET not found error: %v", err)
	}

	return u, nil
}
