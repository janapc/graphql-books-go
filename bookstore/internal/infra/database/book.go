package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/janapc/graphql-books-go/bookstore/internal/entity"
)

type BookDatabase struct {
	BookDB *sql.DB
}

func NewBookDatabase(db *sql.DB) *BookDatabase {
	return &BookDatabase{BookDB: db}
}

func (b *BookDatabase) RegisterBook(book *entity.Book) error {
	stmt, err := b.BookDB.Prepare(`INSERT INTO books (id, title, description, author_id, created_at, updated_at) VALUES (?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.ID, book.Title, book.Description, book.AuthorId, book.CreatedAt, book.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookDatabase) ListBooks() ([]entity.Book, error) {
	rows, err := b.BookDB.Query(`SELECT id, title, description, author_id, created_at, updated_at FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []entity.Book
	for rows.Next() {
		var id, title, description, author_id string
		var createdAt, updatedAt time.Time
		err := rows.Scan(&id, &title, &description, &author_id, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, entity.Book{
			ID:          id,
			Title:       title,
			Description: description,
			AuthorId:    author_id,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}
	return books, nil
}

func (b *BookDatabase) FindBooksByAuthorId(authorID string) ([]entity.Book, error) {
	fmt.Println(authorID)
	rows, err := b.BookDB.Query("SELECT id, title, description, author_id, created_at, updated_at FROM books WHERE author_id = ?", authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := []entity.Book{}
	for rows.Next() {
		var id, title, description, authorId string
		var createdAt, updatedAt time.Time
		err := rows.Scan(&id, &title, &description, &authorId, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, entity.Book{
			ID:          id,
			Title:       title,
			Description: description,
			AuthorId:    authorId,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}
	return books, nil
}

func (b *BookDatabase) FindBookById(id string) (*entity.Book, error) {
	stmt, err := b.BookDB.Prepare(`SELECT id, title, description, author_id, created_at, updated_at FROM books WHERE id = ?`)
	if err != nil {
		return nil, errors.New("book is not found")
	}
	defer stmt.Close()
	var book entity.Book
	err = stmt.QueryRow(id).Scan(&book.ID, &book.Title, &book.Description, &book.AuthorId, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return nil, errors.New("book is not found")
	}
	return &book, nil
}

func (b *BookDatabase) DeleteBook(id string) error {
	_, err := b.FindBookById(id)
	if err != nil {
		return err
	}
	stmt, err := b.BookDB.Prepare("DELETE FROM books WHERE id = ?")
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

func (b *BookDatabase) UpdateBook(book *entity.Book) error {
	stmt, err := b.BookDB.Prepare("UPDATE books SET title = ?, description = ?, author_id = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.Title, book.Description, book.AuthorId, book.UpdatedAt, book.ID)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookDatabase) SearchBooksByTitle(title string) ([]entity.Book, error) {
	fullTitleQuery := "%" + title + "%"
	rows, err := b.BookDB.Query("SELECT id, title, description, author_id, created_at, updated_at FROM books WHERE title LIKE ?", fullTitleQuery)
	if err != nil {
		return nil, errors.New("book is not found")
	}
	defer rows.Close()
	var books []entity.Book
	for rows.Next() {
		var id, title, description, authorId string
		var createdAt, updatedAt time.Time
		err := rows.Scan(&id, &title, &description, &authorId, &createdAt, &updatedAt)
		if err != nil {
			return nil, errors.New("book is not found")
		}
		books = append(books, entity.Book{
			ID:          id,
			Title:       title,
			Description: description,
			AuthorId:    authorId,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}
	return books, nil
}
