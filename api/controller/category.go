package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/model"
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

	var categoryRequest model.Category
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

func (g categoryController) InsertCategory(ctx context.Context, categoryRequest *model.Category) (err error) {
	where := &model.Where{
		Parameter: "WHERE name=$1",
		Values:    []any{categoryRequest.Name},
	}

	categoryModel := model.NewCategoryModel(g.Config.Postgres)
	currCategory, err := categoryModel.FindCategory(ctx, where)
	if !errors.Is(err, sql.ErrNoRows) && err != nil {
		return
	}

	if currCategory != nil && currCategory.Name != "" {
		return errors.New("category already registered")
	}

	categoryRequest.CreatedBy, err = GetCurrentUserIdLoggedIn(ctx)
	if err != nil {
		return
	}

	_, err = categoryModel.CreateCategory(ctx, categoryRequest)
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

	categoryModel := model.NewCategoryModel(g.Config.Postgres)
	categoryList, err := categoryModel.GetCategoryList(c, nil)
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
	where := &model.Where{
		Parameter: "WHERE id=$1",
		Values:    []any{categoryId},
	}

	categoryModel := model.NewCategoryModel(g.Config.Postgres)
	currCategory, err := categoryModel.FindCategory(c, where)
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
