package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

type CategoryModel interface {
	CreateCategory(ctx context.Context, category *Category) (result sql.Result, err error)
	GetCategoryList(ctx context.Context, where *Where) (categoryList []*Category, err error)
	FindCategory(ctx context.Context, where *Where) (category *Category, err error)
	UpdateCategory(ctx context.Context, category *Category) (result sql.Result, err error)
	DeleteCategory(ctx context.Context, categoryId int64) (result sql.Result, err error)
}

func NewCategoryModel(db *sql.DB) CategoryModel {
	return &PostgresRepository{db}
}

func (postgres *PostgresRepository) CreateCategory(ctx context.Context, category *Category) (
	result sql.Result, err error) {

	queryScript := `
		INSERT INTO categories (
			name
			, created_by
		) VALUES ($1, $2)
	`

	return postgres.DB.ExecContext(ctx, queryScript,
		strings.ToLower(category.Name),
		category.CreatedBy,
	)
}

func (postgres *PostgresRepository) GetCategoryList(ctx context.Context, where *Where) (
	categoryList []*Category, err error) {

	where = ValidateWhere(where)
	queryScript := `
		SELECT	id
				, name
		    	, created_at
		    	, updated_at
				, created_by
		     
				, updated_by
		FROM 	categories
	`

	query := fmt.Sprintf("%s %s", queryScript, where.Parameter)
	rows, err := postgres.DB.QueryContext(ctx, query, where.Values...)
	if err != nil {
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("error closing get category list rows:", err)
		}
	}(rows)

	categoryList = make([]*Category, 0)

	for rows.Next() {
		category := new(Category)

		err = rows.Scan(
			&category.Id,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.CreatedBy,

			&category.UpdatedBy,
		)

		if err != nil {
			return
		}

		categoryList = append(categoryList, category)
	}

	return
}

func (postgres *PostgresRepository) FindCategory(ctx context.Context, where *Where) (
	category *Category, err error) {

	where = ValidateWhere(where)
	queryScript := `
		SELECT	id
				, name
		    	, created_at
		    	, updated_at
				, created_by
		
				, updated_by
		FROM 	categories
	`

	query := fmt.Sprintf("%s %s", queryScript, where.Parameter)
	row := postgres.DB.QueryRowContext(ctx, query, where.Values...)

	category = new(Category)
	err = row.Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.CreatedBy,

		&category.UpdatedBy,
	)

	if err != nil {
		return
	}

	return
}

func (postgres *PostgresRepository) UpdateCategory(ctx context.Context, category *Category) (
	result sql.Result, err error) {

	queryScript := `
		UPDATE 	categories SET 
		    	name = $1
				, updated_at = $2
				, updated_by = $3
		WHERE 	id = $4
		`

	return postgres.DB.ExecContext(ctx, queryScript,
		strings.ToLower(category.Name),
		time.Now(),
		category.UpdatedBy,
		category.Id,
	)
}

func (postgres *PostgresRepository) DeleteCategory(ctx context.Context, categoryId int64) (
	result sql.Result, err error) {

	queryScript := `DELETE FROM categories WHERE id = $1`
	return postgres.DB.ExecContext(ctx, queryScript, categoryId)
}
