package category

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/michaelwp/goblog/model"
	"log"
	"strings"
	"time"
)

func CreateCategory(ctx context.Context, postgres *sql.DB, category *Category) (result sql.Result, err error) {
	queryScript := `
		INSERT INTO categories (
			name
			, created_by
		) VALUES ($1, $2)
	`

	return postgres.ExecContext(ctx, queryScript,
		strings.ToLower(category.Name),
		category.CreatedBy,
	)
}

func GetCategoryList(ctx context.Context, postgres *sql.DB, where *model.Where) (categoryList []*Category, err error) {
	where = model.ValidateWhere(where)

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
	rows, err := postgres.QueryContext(ctx, query, where.Values...)
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

func FindCategory(ctx context.Context, postgres *sql.DB, where *model.Where) (category *Category, err error) {
	where = model.ValidateWhere(where)

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
	row := postgres.QueryRowContext(ctx, query, where.Values...)

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

func UpdateCategory(ctx context.Context, postgres *sql.DB, category *Category) (result sql.Result, err error) {
	queryScript := `
		UPDATE 	categories SET 
		    	name = $1
				, updated_at = $2
				, updated_by = $3
		WHERE 	id = $4
		`

	return postgres.ExecContext(ctx, queryScript,
		strings.ToLower(category.Name),
		time.Now(),
		category.UpdatedBy,
		category.Id,
	)
}

func DeleteCategory(ctx context.Context, postgres *sql.DB, categoryId int64) (result sql.Result, err error) {
	queryScript := `DELETE FROM categories WHERE id = $1`
	return postgres.ExecContext(ctx, queryScript, categoryId)
}
