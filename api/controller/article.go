package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/model"
	"github.com/michaelwp/goblog/model/article"
	"net/http"
	"strconv"
)

type ArticleController interface {
	CreateArticle(c *gin.Context)
	GetArticleList(c *gin.Context)
	UpdateArticle(c *gin.Context)
	GetArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)
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

	userId, err := GetCurrentUserIdLoggedIn(c)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "article.create.error"

		c.JSON(http.StatusBadRequest, response)
		return
	}

	articleRequest.UserId = userId
	articleRequest.CreatedBy = userId

	_, err = article.CreateArticle(c, a.Config.Postgres, &articleRequest)
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

	articleList, err := article.GetArticleList(c, a.Postgres, nil)
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

	var articleRequest article.Article
	err := c.ShouldBindJSON(&articleRequest)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "article.update.error"

		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = a.UpdateCurrentArticle(c, &articleRequest)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "article.update.error"

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(200, response)
}

func (a articleController) UpdateCurrentArticle(ctx context.Context, articleRequest *article.Article) (err error) {
	userId, err := GetCurrentUserIdLoggedIn(ctx)
	if err != nil {
		return
	}

	articleRequest.UserId = userId
	articleRequest.UpdatedBy = &userId

	_, err = article.UpdateArticle(ctx, a.Postgres, articleRequest)
	if err != nil {
		return
	}

	return
}

func (a articleController) GetArticle(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "article successfully retrieved",
		Translate: "article.get.success",
	}

	articleId := c.Param("id")
	currArticle, response, httpStatus, err := a.FindCurrentArticle(c, response, articleId)
	if err != nil {
		response.Message = err.Error()
		c.JSON(httpStatus, response)
		return
	}

	response.Data = currArticle
	c.JSON(200, response)
}

func (a articleController) FindCurrentArticle(ctx context.Context, resp *Response, articleId string) (
	currArticle *article.Article, response *Response, httpStatus int, err error) {

	response.Status = ERROR
	response.Translate = "article.get.error"
	httpStatus = http.StatusInternalServerError

	articleIdInt, err := strconv.ParseInt(articleId, 10, 64)
	if err != nil {
		return
	}

	response = resp
	where := &model.Where{
		Parameter: "WHERE id=$1",
		Values:    []any{articleIdInt},
	}

	currArticle, err = article.FindArticle(ctx, a.Postgres, where)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Translate = "article.not.found"
			httpStatus = http.StatusNotFound
		}

		return
	}

	return
}

func (a articleController) DeleteArticle(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "article successfully deleted",
		Translate: "article.delete.success",
	}

	articleId := c.Request.URL.Query().Get("id")
	articleIdInt, err := strconv.ParseInt(articleId, 10, 64)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "article.delete.error"

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = article.DeleteArticle(c, a.Postgres, articleIdInt)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "article.delete.error"

		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(200, response)
}
