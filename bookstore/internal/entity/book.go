package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AuthorId    string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewBook(title, description, authorId string) (*Book, error) {
	err := validateBook(title, description, authorId)
	if err != nil {
		return nil, err
	}
	return &Book{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		AuthorId:    authorId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func validateBook(title, description, authorId string) error {
	if title == "" {
		return errors.New("the parameter title is mandatory")
	}

	if description == "" {
		return errors.New("the parameter description is mandatory")
	}

	if authorId == "" {
		return errors.New("the parameter author_id is mandatory")
	}

	return nil
}
