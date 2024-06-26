package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/model"
	"github.com/michaelwp/goblog/tool"
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

	var articleRequest model.Article
	err := c.ShouldBindJSON(&articleRequest)
	if err != nil {
		response.Status = ERROR
		response.Message = tool.PrintLog("create_article:", err).Error()
		response.Translate = "article.create.error"

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId, err := GetCurrentUserIdLoggedIn(c)
	if err != nil {
		response.Status = ERROR
		response.Message = tool.PrintLog("create_article:", err).Error()
		response.Translate = "article.create.error"

		c.JSON(http.StatusBadRequest, response)
		return
	}

	articleRequest.UserId = userId
	articleRequest.CreatedBy = userId

	articleModel := model.NewArticleModel(a.Config.Postgres)
	_, err = articleModel.CreateArticle(c, &articleRequest)
	if err != nil {
		response.Status = ERROR
		response.Message = tool.PrintLog("create_article:", err).Error()
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

	articleModel := model.NewArticleModel(a.Config.Postgres)

	where := &model.Where{
		Order: "ORDER BY a.id DESC",
	}

	articleList, err := articleModel.GetArticleList(c, where)
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

	var articleRequest model.Article
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

func (a articleController) UpdateCurrentArticle(ctx context.Context, articleRequest *model.Article) (err error) {
	userId, err := GetCurrentUserIdLoggedIn(ctx)
	if err != nil {
		return
	}

	articleRequest.UserId = userId
	articleRequest.UpdatedBy = &userId

	articleModel := model.NewArticleModel(a.Config.Postgres)
	_, err = articleModel.UpdateArticle(ctx, articleRequest)
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
	currArticle *model.ArticleWithExtend, response *Response, httpStatus int, err error) {

	articleIdInt, err := strconv.ParseInt(articleId, 10, 64)
	if err != nil {
		response.Message = tool.PrintLog("FindCurrentArticle", err).Error()
		return
	}

	response = resp
	where := &model.Where{
		Parameter: "WHERE a.id=$1",
		Values:    []any{articleIdInt},
	}

	articleModel := model.NewArticleModel(a.Config.Postgres)
	currArticle, err = articleModel.FindArticle(ctx, where)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Message = tool.PrintLog("FindCurrentArticle", err).Error()
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

	articleModel := model.NewArticleModel(a.Config.Postgres)
	_, err = articleModel.DeleteArticle(c, articleIdInt)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "article.delete.error"

		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(200, response)
}
