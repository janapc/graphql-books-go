package database

import (
	"database/sql"
	"time"

	"github.com/janapc/graphql-books-go/user/internal/entity"
)

type UserDatabase struct {
	UserDB *sql.DB
}

func NewUserDatabase(db *sql.DB) *UserDatabase {
	return &UserDatabase{
		UserDB: db,
	}
}

func (u *UserDatabase) RegisterUser(user *entity.User) error {
	stmt, err := u.UserDB.Prepare("INSERT INTO users(id, email, password, created_at, updated_at, deleted_at) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDatabase) ListUsers() ([]entity.User, error) {
	rows, err := u.UserDB.Query("SELECT id, email, created_at, updated_at, deleted_at FROM users WHERE is_deleted IS NOT NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserDatabase) FindUserById(id string) (*entity.User, error) {
	stmt, err := u.UserDB.Prepare("SELECT id, email, password, created_at, updated_at, deleted_at FROM users WHERE id = ? AND is_deleted IS NOT NULL")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user entity.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDatabase) FindUserByEmail(email string) (*entity.User, error) {
	stmt, err := u.UserDB.Prepare("SELECT id, email, password, created_at, updated_at, deleted_at FROM users WHERE email = ? AND is_deleted IS NOT NULL")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user entity.User
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDatabase) UpdateUser(user *entity.User) error {
	stmt, err := u.UserDB.Prepare("UPDATE users SET email = ?, password = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Email, user.Password, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDatabase) DeleteUser(id string) error {
	deletedAt := time.Now()
	_, err := u.FindUserById(id)
	if err != nil {
		return err
	}
	stmt, err := u.UserDB.Prepare(`UPDATE users SET deleted_at = ?  WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(deletedAt, id)
	if err != nil {
		return err
	}
	return nil
}
