package controller

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/model/article"
	"net/http"
)

type ArticleController interface {
	CreateArticle(c *gin.Context)
	GetArticleList(c *gin.Context)
	UpdateArticle(c *gin.Context)
	GetArticle(c *gin.Context)
}

type articleController struct {
	*Config
}

func NewArticleController(c *Config) ArticleController {
	return &articleController{c}
}

func (a articleController) CreateArticle(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "article successfully created",
		Translate: "article.create.success",
	}

	var articleRequest article.Article
	err := c.ShouldBindJSON(&articleRequest)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "article.create.error"

		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = article.CreateArticle(c, a.Config.Postgres, &articleRequest)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "article.create.error"

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (a articleController) GetArticleList(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "article successfully retrieved",
		Translate: "article.get.success",
	}

	where := "WHERE deleted_at IS NULL"
	articleList, err := article.GetArticleList(c, a.Postgres, where, nil)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "article.get.error"

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = articleList
	c.JSON(200, response)
}

func (a articleController) UpdateArticle(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "article successfully updated",
		Translate: "article.update.success",
	}

	c.JSON(200, response)
}

func (a articleController) GetArticle(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "article successfully retrieved",
		Translate: "article.get.success",
	}

	articleId := c.Param("id")
	where := "WHERE id=$1"
	value := []any{articleId}

	currArticle, err := article.FindArticle(c, a.Postgres, where, value)
	if err != nil {
		translate := "article.get.error"
		httpStatus := http.StatusInternalServerError

		if errors.Is(err, sql.ErrNoRows) {
			translate = "article.not.found"
			httpStatus = http.StatusNotFound
		}

		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = translate

		c.JSON(httpStatus, response)
		return
	}

	response.Data = currArticle
	c.JSON(200, response)
}
