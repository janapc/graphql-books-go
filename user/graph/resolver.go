package graph

import "github.com/janapc/graphql-books-go/user/internal/entity"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserDB entity.UserRepository
}
