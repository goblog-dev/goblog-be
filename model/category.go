package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/michaelwp/goblog/entities"
	"log"
	"strings"
)

type CategoryModel interface {
	CreateCategory(ctx context.Context, category *entities.Category) (result sql.Result, err error)
	GetCategoryList(ctx context.Context, where *Where) (categoryList []*entities.Category, err error)
	FindCategory(ctx context.Context, where *Where) (category *entities.Category, err error)
	UpdateCategory(ctx context.Context, category *entities.Category) (result sql.Result, err error)
	DeleteCategory(ctx context.Context, categoryId int64) (result sql.Result, err error)
}

func NewCategoryModel(db *sql.DB) CategoryModel {
	return &PostgresRepository{db}
}

func (postgres *PostgresRepository) CreateCategory(ctx context.Context, category *entities.Category) (
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
	categoryList []*entities.Category, err error) {

	where = ValidateWhere(where)
	queryScript := `
		SELECT	id
				, name
		     	, created_by
		    	, created_at
		     	, updated_by
		     
		    	, updated_at
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

	categoryList = make([]*entities.Category, 0)

	for rows.Next() {
		category := new(entities.Category)

		err = rows.Scan(
			&category.Id,
			&category.Name,
			&category.CreatedBy,
			&category.CreatedAt,
			&category.UpdatedBy,

			&category.UpdatedAt,
		)

		if err != nil {
			return
		}

		categoryList = append(categoryList, category)
	}

	return
}

func (postgres *PostgresRepository) FindCategory(ctx context.Context, where *Where) (
	category *entities.Category, err error) {

	where = ValidateWhere(where)
	queryScript := `
		SELECT	id
				, name
		     	, created_by
		    	, created_at
		     	, updated_by
		     
		    	, updated_at
		FROM 	categories
	`

	query := fmt.Sprintf("%s %s", queryScript, where.Parameter)
	row := postgres.DB.QueryRowContext(ctx, query, where.Values...)

	category = new(entities.Category)
	err = row.Scan(
		&category.Id,
		&category.Name,
		&category.CreatedBy,
		&category.CreatedAt,
		&category.UpdatedBy,

		&category.UpdatedAt,
	)

	if err != nil {
		return
	}

	return
}

func (postgres *PostgresRepository) UpdateCategory(ctx context.Context, category *entities.Category) (
	result sql.Result, err error) {

	queryScript := `
		UPDATE 	categories SET 
		    	name = $1
				, updated_by = $2
		WHERE 	id = $3
		`

	return postgres.DB.ExecContext(ctx, queryScript,
		strings.ToLower(category.Name),
		category.UpdatedBy,
		category.Id,
	)
}

func (postgres *PostgresRepository) DeleteCategory(ctx context.Context, categoryId int64) (
	result sql.Result, err error) {

	queryScript := `DELETE FROM categories WHERE id = $1`
	return postgres.DB.ExecContext(ctx, queryScript, categoryId)
}
