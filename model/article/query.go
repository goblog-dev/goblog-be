package article

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func CreateArticle(ctx context.Context, postgres *sql.DB, article *Article) (err error) {
	queryScript := `
		INSERT INTO articles (
			user_id
			, category_id
			, title
			, content
		) VALUES ($1, $2, $3, $4)
	`

	_, err = postgres.ExecContext(
		ctx,
		queryScript,
		article.UserId,
		article.CategoryId,
		article.Title,
		article.Content,
	)

	return
}

func GetArticleList(ctx context.Context, postgres *sql.DB, where string, value []any) (articleList []*Article, err error) {
	queryScript := `
		SELECT	id
				, user_id
				, category_id
		    	, title
		     	, content
		     
		    	, created_at
		    	, updated_at
				, deleted_at
		FROM 	articles
	`

	query := fmt.Sprintf("%s %s", queryScript, where)
	rows, err := postgres.QueryContext(ctx, query, value...)
	if err != nil {
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("error closing get article list rows:", err)
		}
	}(rows)

	articleList = make([]*Article, 0)

	for rows.Next() {
		article := new(Article)

		err = rows.Scan(
			&article.Id,
			&article.UserId,
			&article.CategoryId,
			&article.Title,
			&article.Content,

			&article.CreatedAt,
			&article.UpdatedAt,
			&article.DeletedAt,
		)

		if err != nil {
			return
		}

		articleList = append(articleList, article)
	}

	return
}

func FindArticle(ctx context.Context, postgres *sql.DB, where string, value []any) (article *Article, err error) {
	queryScript := `
		SELECT	id
				, user_id
				, category_id
		    	, title
		     	, content
		     
		    	, created_at
		    	, updated_at
				, deleted_at
		FROM 	articles
	`

	query := fmt.Sprintf("%s %s", queryScript, where)
	row := postgres.QueryRowContext(ctx, query, value...)

	article = new(Article)
	err = row.Scan(
		&article.Id,
		&article.UserId,
		&article.CategoryId,
		&article.Title,
		&article.Content,

		&article.CreatedAt,
		&article.UpdatedAt,
		&article.DeletedAt,
	)

	if err != nil {
		return
	}

	return
}
