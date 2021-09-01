package model

type Countries struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}
