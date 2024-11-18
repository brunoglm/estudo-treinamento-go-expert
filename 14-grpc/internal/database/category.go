package database

import (
	"context"
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

func (c *Category) Create(ctx context.Context, name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.ExecContext(ctx, "INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll(ctx context.Context) ([]Category, error) {
	rows, err := c.db.QueryContext(ctx, "SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *Category) FindByCourseID(ctx context.Context, courseID string) (Category, error) {
	var category Category
	err := c.db.QueryRowContext(ctx, "SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", courseID).
		Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return Category{}, err
	}
	return category, nil
}

func (c *Category) Find(ctx context.Context, id string) (Category, error) {
	var category Category
	err := c.db.QueryRowContext(ctx, "SELECT id, name, description FROM categories WHERE id = $1", id).
		Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return Category{}, err
	}
	return category, nil
}
