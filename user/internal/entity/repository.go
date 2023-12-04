package entity

type UserRepository interface {
	RegisterUser(user *User) error
	ListUsers() ([]User, error)
	FindUserById(id string) (*User, error)
	FindUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}
