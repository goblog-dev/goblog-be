package article

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/michaelwp/goblog/model"
	"log"
	"time"
)

func CreateArticle(ctx context.Context, postgres *sql.DB, article *Article) (result sql.Result, err error) {
	queryScript := `
		INSERT INTO articles (
			user_id
			, category_id
			, title
			, content
			, created_by
		) VALUES ($1, $2, $3, $4, $5)
	`

	return postgres.ExecContext(ctx, queryScript,
		article.UserId,
		article.CategoryId,
		article.Title,
		article.Content,
		article.CreatedBy,
	)
}

func GetArticleList(ctx context.Context, postgres *sql.DB, where *model.Where) (articleList []*Article, err error) {
	where = model.ValidateWhere(where)

	queryScript := `
		SELECT	a.id
				, a.user_id
				, a.category_id
		    	, a.title
		     	, a.content
		     
		    	, a.created_at
		    	, a.updated_at
				, a.created_by
				, a.updated_by
				, u.name AS user_name
		     
				, c.name AS category_name
		FROM 	articles a
				JOIN users u ON a.user_id = u.id
				JOIN categories c ON a.category_id = c.id
	`

	query := fmt.Sprintf("%s %s", queryScript, where.Parameter)
	rows, err := postgres.QueryContext(ctx, query, where.Values...)
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
			&article.CreatedBy,
			&article.UpdatedBy,
			&article.UserName,

			&article.CategoryName,
		)

		if err != nil {
			return
		}

		articleList = append(articleList, article)
	}

	return
}

func FindArticle(ctx context.Context, postgres *sql.DB, where *model.Where) (article *Article, err error) {
	where = model.ValidateWhere(where)

	queryScript := `
		SELECT	a.id
				, a.user_id
				, a.category_id
		    	, a.title
		     	, a.content
		     
		    	, a.created_at
		    	, a.updated_at
				, a.created_by
				, a.updated_by
				, u.name AS user_name
		     
				, c.name AS category_name
		FROM 	articles a
				JOIN users u ON a.user_id = u.id
				JOIN categories c ON a.category_id = c.id
	`

	query := fmt.Sprintf("%s %s", queryScript, where.Parameter)
	row := postgres.QueryRowContext(ctx, query, where.Values...)

	article = new(Article)
	err = row.Scan(
		&article.Id,
		&article.UserId,
		&article.CategoryId,
		&article.Title,
		&article.Content,

		&article.CreatedAt,
		&article.UpdatedAt,
		&article.CreatedBy,
		&article.UpdatedBy,
		&article.UserName,

		&article.CategoryName,
	)

	if err != nil {
		return
	}

	return
}

func UpdateArticle(ctx context.Context, postgres *sql.DB, article *Article) (result sql.Result, err error) {
	queryScript := `
		UPDATE 	articles SET 
		    	user_id = $1
				, category_id = $2
				, title = $3
				, content = $4
				, updated_at = $5
				
		        , updated_by = $6
		WHERE 	id = $7
		`

	return postgres.ExecContext(ctx, queryScript,
		article.UserId,
		article.CategoryId,
		article.Title,
		article.Content,
		time.Now(),

		article.UpdatedBy,
		article.Id,
	)
}

func DeleteArticle(ctx context.Context, postgres *sql.DB, articleId int64) (result sql.Result, err error) {
	queryScript := `DELETE FROM articles WHERE id = $1`
	return postgres.ExecContext(ctx, queryScript, articleId)
}
