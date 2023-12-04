package database

import (
	"database/sql"
	"time"

	"github.com/janapc/graphql-books-go/bookstore/internal/entity"
)

type AuthorDatabase struct {
	AuthorDB *sql.DB
}

func NewAuthorDatabase(db *sql.DB) *AuthorDatabase {
	return &AuthorDatabase{AuthorDB: db}
}

func (a *AuthorDatabase) RegisterAuthor(author *entity.Author) error {
	stmt, err := a.AuthorDB.Prepare(`INSERT INTO authors (id, name, created_at, updated_at) VALUES (?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(author.ID, author.Name, author.CreatedAt, author.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthorDatabase) ListAuthors() ([]entity.Author, error) {
	rows, err := a.AuthorDB.Query(`SELECT id, name, created_at, updated_at FROM authors`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var authors []entity.Author
	for rows.Next() {
		var id, name string
		var createdAt, updatedAt time.Time
		err := rows.Scan(&id, &name, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		authors = append(authors, entity.Author{
			ID:        id,
			Name:      name,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return authors, nil
}

func (a *AuthorDatabase) FindAuthorByBookId(bookId string) (*entity.Author, error) {
	var id, name string
	var created_at, update_at time.Time
	err := a.AuthorDB.QueryRow("SELECT a.id, a.name, a.created_at, a.updated_at FROM authors a JOIN books b ON a.id = b.author_id WHERE b.id= ?", bookId).Scan(&id, &name, &created_at, &update_at)
	if err != nil {
		return nil, err
	}
	return &entity.Author{
		ID:        id,
		Name:      name,
		CreatedAt: created_at,
		UpdatedAt: update_at,
	}, nil
}

func (a *AuthorDatabase) FindAuthorById(id string) (*entity.Author, error) {
	var author entity.Author
	err := a.AuthorDB.QueryRow("SELECT id, name, created_at, updated_at FROM authors WHERE id = ?", id).Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (a *AuthorDatabase) SearchAuthorsByName(name string) ([]entity.Author, error) {
	fullNameQuery := "%" + name + "%"
	rows, err := a.AuthorDB.Query("SELECT id, name, created_at, updated_at FROM authors WHERE name LIKE ?", fullNameQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var authors []entity.Author
	for rows.Next() {
		var id, name string
		var createdAt, updateAt time.Time
		err := rows.Scan(&id, &name, &createdAt, &updateAt)
		if err != nil {
			return nil, err
		}
		authors = append(authors, entity.Author{
			ID:        id,
			Name:      name,
			CreatedAt: createdAt,
			UpdatedAt: updateAt,
		})
	}
	return authors, nil
}

func (a *AuthorDatabase) UpdateAuthor(author *entity.Author) error {
	stmt, err := a.AuthorDB.Prepare("UPDATE authors SET name = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(author.Name, author.UpdatedAt, author.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthorDatabase) DeleteAuthor(id string) error {
	_, err := a.FindAuthorById(id)
	if err != nil {
		return err
	}
	stmt, err := a.AuthorDB.Prepare(`DELETE FROM authors WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
