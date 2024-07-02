package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
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

	lowerTags := strings.ToLower(*article.Tags)
	cleanTags := strings.Replace(lowerTags, " ", "", -1)

	queryScript := `
		INSERT INTO articles (
			user_id
			, category_id
			, content
			, title
			, tags
			
			, description
			, image
			, created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	return postgres.DB.ExecContext(ctx, queryScript,
		article.UserId,
		article.CategoryId,
		article.Content,
		article.Title,
		cleanTags,

		article.Description,
		article.Image,
		article.CreatedBy,
	)
}

func (postgres *PostgresRepository) GetArticleList(ctx context.Context, where *Where) (
	articleWithExtendList []*ArticleWithExtend, err error) {

	where = ValidateWhere(where)

	queryScript := `
		SELECT	a.id
		    	, a.title
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
			&articleWithExtend.Title,
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
		     	, a.content
		    	, a.title
		     	
		     	, a.tags
		     	, a.created_by
		    	, a.created_at
		     	, a.updated_by
		    	, a.updated_at
				
				, u.name AS user_name
				, c.name AS category_name
				, u.page
				, a.description
				, a.image
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
		&articleWithExtend.Content,
		&articleWithExtend.Title,

		&articleWithExtend.Tags,
		&articleWithExtend.CreatedBy,
		&articleWithExtend.CreatedAt,
		&articleWithExtend.UpdatedBy,
		&articleWithExtend.UpdatedAt,

		&articleWithExtend.UserName,
		&articleWithExtend.CategoryName,
		&articleWithExtend.Page,
		&articleWithExtend.Description,
		&articleWithExtend.Image,
	)

	if err != nil {
		return
	}

	return
}

func (postgres *PostgresRepository) UpdateArticle(ctx context.Context, article *Article) (
	result sql.Result, err error) {

	lowerTags := strings.ToLower(*article.Tags)
	cleanTags := strings.Replace(lowerTags, " ", "", -1)

	queryScript := `
		UPDATE 	articles SET 
		        user_id = $1
				, category_id = $2
		        , content = $3
				, title = $4
		        , tags = $5
		        
		        , description = $6
		        , image = $7
		        , updated_by = $8
		WHERE 	id = $9
		`

	return postgres.DB.ExecContext(ctx, queryScript,
		article.UserId,
		article.CategoryId,
		article.Content,
		article.Title,
		cleanTags,

		article.Description,
		article.Image,
		article.UpdatedBy,
		article.Id,
	)
}

func (postgres *PostgresRepository) DeleteArticle(ctx context.Context, articleId int64) (
	result sql.Result, err error) {

	queryScript := `DELETE FROM articles WHERE id = $1`
	return postgres.DB.ExecContext(ctx, queryScript, articleId)
}
