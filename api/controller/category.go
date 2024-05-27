package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/model/category"
	"net/http"
)

type CategoryController interface {
	CreateCategory(c *gin.Context)
	GetCategoryList(c *gin.Context)
	UpdateCategory(c *gin.Context)
	GetCategory(c *gin.Context)
}

type categoryController struct {
	*Config
}

func NewCategoryController(c *Config) CategoryController {
	return &categoryController{c}
}

func (g categoryController) CreateCategory(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "category successfully created",
		Translate: "category.create.success",
	}

	var categoryRequest category.Category
	err := c.ShouldBindJSON(&categoryRequest)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "category.create.error"

		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = g.InsertCategory(c, &categoryRequest)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "category.create.error"

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (g categoryController) InsertCategory(ctx context.Context, categoryRequest *category.Category) (err error) {
	value := []any{categoryRequest.Name}
	where := "WHERE name=$1"

	currCategory, err := category.FindCategory(ctx, g.Postgres, where, value)
	if !errors.Is(err, sql.ErrNoRows) && err != nil {
		return
	}

	if currCategory != nil && currCategory.Name != "" {
		return errors.New("category already registered")
	}

	err = category.CreateCategory(ctx, g.Config.Postgres, categoryRequest)
	if err != nil {
		return
	}

	return
}

func (g categoryController) GetCategoryList(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "category successfully retrieved",
		Translate: "category.get.success",
	}

	categoryList, err := category.GetCategoryList(c, g.Postgres, "", nil)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "category.get.error"

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = categoryList
	c.JSON(200, response)
}

func (g categoryController) UpdateCategory(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "category successfully updated",
		Translate: "category.update.success",
	}

	c.JSON(200, response)
}

func (g categoryController) GetCategory(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "category successfully retrieved",
		Translate: "category.get.success",
	}

	categoryId := c.Param("id")
	where := "WHERE id=$1"
	value := []any{categoryId}

	currCategory, err := category.FindCategory(c, g.Postgres, where, value)
	if err != nil {
		translate := "category.get.error"
		httpStatus := http.StatusInternalServerError

		if errors.Is(err, sql.ErrNoRows) {
			translate = "category.not.found"
			httpStatus = http.StatusNotFound
		}

		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = translate

		c.JSON(httpStatus, response)
		return
	}

	response.Data = currCategory
	c.JSON(200, response)
}
