package domain

import (
	"time"
)

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type UserRepository interface {
// 	GetByUsername(ctx context.Context, username string) (User, error)
// 	GetByEmail(ctx context.Context, email string) (User, error)
// }
