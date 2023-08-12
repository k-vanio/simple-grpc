package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryId string) (*Course, error) {
	id := uuid.New().String()
	query := "INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)"

	if _, err := c.db.Exec(query, id, name, description, categoryId); err != nil {
		return nil, err
	}

	c.ID = id
	c.Name = name
	c.Description = description
	c.CategoryId = categoryId

	return c, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id  from courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []Course{}
	for rows.Next() {
		var id, name, description, category_id string
		if err := rows.Scan(&id, &name, &description, &category_id); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: category_id})
	}

	return courses, nil
}

func (c *Course) FindByCategoryId(categoryId string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id  from courses where category_id = $1", categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []Course{}
	for rows.Next() {
		var id, name, description, category_id string
		if err := rows.Scan(&id, &name, &description, &category_id); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: category_id})
	}

	return courses, nil
}
