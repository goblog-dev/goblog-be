package category

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func CreateCategory(ctx context.Context, postgres *sql.DB, category *Category) (err error) {
	queryScript := `INSERT INTO categories (name) VALUES ($1)`
	_, err = postgres.ExecContext(ctx, queryScript, category.Name)
	return
}

func GetCategoryList(ctx context.Context, postgres *sql.DB, where string, value []any) (categoryList []*Category, err error) {
	queryScript := `
		SELECT	id
				, name
		    	, created_at
		    	, updated_at
		FROM 	categories
	`

	query := fmt.Sprintf("%s %s", queryScript, where)
	rows, err := postgres.QueryContext(ctx, query, value...)
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
		)

		if err != nil {
			return
		}

		categoryList = append(categoryList, category)
	}

	return
}

func FindCategory(ctx context.Context, postgres *sql.DB, where string, value []any) (category *Category, err error) {
	queryScript := `
		SELECT	id
				, name
		    	, created_at
		    	, updated_at
		FROM 	categories
	`

	query := fmt.Sprintf("%s %s", queryScript, where)
	row := postgres.QueryRowContext(ctx, query, value...)

	category = new(Category)
	err = row.Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return
	}

	return
}
