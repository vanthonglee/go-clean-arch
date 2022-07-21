package domain

import (
	"context"
	"time"
)

// Author ...
type Author struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthorRepository represent the author's repository contract
type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (Author, error)
}
