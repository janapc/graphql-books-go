package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAuthor(name string) (*Author, error) {
	if name == "" {
		return nil, errors.New("the parameter name is mandatory")
	}
	return &Author{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
