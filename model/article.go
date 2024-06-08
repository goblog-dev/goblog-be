package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type ArticleModel interface {
	CreateArticle(ctx context.Context, article *Article) (result sql.Result, err error)
	GetArticleList(ctx context.Context, where *Where) (articleList []*ArticleWithExtend, err error)
	FindArticle(ctx context.Context, where *Where) (article *ArticleWithExtend, err error)
	UpdateArticle(ctx context.Context, article *Article) (result sql.Result, err error)
	DeleteArticle(ctx context.Context, articleId int64) (result sql.Result, err error)
}

func NewArticleModel(db *sql.DB) ArticleModel {
	return &PostgresRepository{db}
}

func (postgres *PostgresRepository) CreateArticle(ctx context.Context, article *Article) (
	result sql.Result, err error) {

	queryScript := `
		INSERT INTO articles (
			user_id
			, category_id
			, title
			, content
			, created_by
		) VALUES ($1, $2, $3, $4, $5)
	`

	return postgres.DB.ExecContext(ctx, queryScript,
		article.UserId,
		article.CategoryId,
		article.Title,
		article.Content,
		article.CreatedBy,
	)
}

func (postgres *PostgresRepository) GetArticleList(ctx context.Context, where *Where) (
	articleWithExtendList []*ArticleWithExtend, err error) {

	where = ValidateWhere(where)

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

	query := fmt.Sprintf("%s %s %s %s", queryScript, where.Parameter, where.Order, where.Limit)
	rows, err := postgres.DB.QueryContext(ctx, query, where.Values...)
	if err != nil {
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("error closing get article list rows:", err)
		}
	}(rows)

	articleWithExtendList = make([]*ArticleWithExtend, 0)

	for rows.Next() {
		articleWithExtend := new(ArticleWithExtend)

		err = rows.Scan(
			&articleWithExtend.Id,
			&articleWithExtend.UserId,
			&articleWithExtend.CategoryId,
			&articleWithExtend.Title,
			&articleWithExtend.Content,

			&articleWithExtend.CreatedAt,
			&articleWithExtend.UpdatedAt,
			&articleWithExtend.CreatedBy,
			&articleWithExtend.UpdatedBy,
			&articleWithExtend.UserName,

			&articleWithExtend.CategoryName,
		)

		if err != nil {
			return
		}

		articleWithExtendList = append(articleWithExtendList, articleWithExtend)
	}

	return
}

func (postgres *PostgresRepository) FindArticle(ctx context.Context, where *Where) (
	articleWithExtend *ArticleWithExtend, err error) {

	where = ValidateWhere(where)

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
	row := postgres.DB.QueryRowContext(ctx, query, where.Values...)

	articleWithExtend = new(ArticleWithExtend)
	err = row.Scan(
		&articleWithExtend.Id,
		&articleWithExtend.UserId,
		&articleWithExtend.CategoryId,
		&articleWithExtend.Title,
		&articleWithExtend.Content,

		&articleWithExtend.CreatedAt,
		&articleWithExtend.UpdatedAt,
		&articleWithExtend.CreatedBy,
		&articleWithExtend.UpdatedBy,
		&articleWithExtend.UserName,

		&articleWithExtend.CategoryName,
	)

	if err != nil {
		return
	}

	return
}

func (postgres *PostgresRepository) UpdateArticle(ctx context.Context, article *Article) (
	result sql.Result, err error) {

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

	return postgres.DB.ExecContext(ctx, queryScript,
		article.UserId,
		article.CategoryId,
		article.Title,
		article.Content,
		time.Now(),

		article.UpdatedBy,
		article.Id,
	)
}

func (postgres *PostgresRepository) DeleteArticle(ctx context.Context, articleId int64) (
	result sql.Result, err error) {

	queryScript := `DELETE FROM articles WHERE id = $1`
	return postgres.DB.ExecContext(ctx, queryScript, articleId)
}
