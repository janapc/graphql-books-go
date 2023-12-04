package entity

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewUser(email, password string) (*User, error) {
	err := ValidateNewUser(email, password)
	if err != nil {
		return nil, err
	}
	hash, err := GeneratePasswordUser(password)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}, nil
}

func ValidateNewUser(email, password string) error {
	if email == "" {
		return errors.New("the email is mandatory")
	}

	if password == "" {
		return errors.New("the password is mandatory")
	}

	err := ValidateEmailUser(email)
	if err != nil {
		return err
	}

	err = ValidatePasswordUser(password)
	if err != nil {
		return err
	}

	return nil
}

func ValidateEmailUser(email string) error {
	isEmailValid, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
	if err != nil || !isEmailValid {
		return errors.New("the email is invalid")
	}
	return nil
}

func ValidatePasswordUser(password string) error {
	if len(password) < 8 {
		return errors.New("password must have greater than 8 characters")
	}
	return nil
}

func GeneratePasswordUser(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}
