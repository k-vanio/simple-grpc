package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name, description string) (*Category, error) {
	id := uuid.New().String()
	if _, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description); err != nil {
		return nil, err
	}

	c.ID = id
	c.Name = name
	c.Description = description

	return c, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description from categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	return categories, nil
}

func (c *Category) FindByCourseId(courseId string) (*Category, error) {
	category := &Category{}
	query := "SELECT c.id, c.name, c.description from categories c JOIN courses co ON c.id = co.category_id where co.id = $1"

	err := c.db.QueryRow(query, courseId).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}

	return category, nil
}
