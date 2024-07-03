package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/entities"
	"github.com/michaelwp/goblog/model"
	"github.com/michaelwp/goblog/tool"
	"net/http"
)

type UserController interface {
	CreateUser(c *gin.Context)
	GetUserList(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUser(c *gin.Context)
}

type userController struct {
	*Config
}

func NewUserController(c *Config) UserController {
	return &userController{c}
}

func (u userController) CreateUser(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "user successfully created",
		Translate: "user.create.success",
	}

	var userRequest entities.User
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "user.create.error"

		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = u.InsertUser(c, &userRequest)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "user.create.error"

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (u userController) InsertUser(ctx context.Context, userRequest *entities.User) (err error) {
	where := &model.Where{
		Parameter: "WHERE email=$1",
		Values:    []any{userRequest.Email},
	}

	userModel := model.NewUserModel(u.Config.Postgres)
	currUser, err := userModel.FindUser(ctx, where)
	if !errors.Is(err, sql.ErrNoRows) && err != nil {
		return
	}

	if currUser != nil && currUser.Email != "" {
		return errors.New("email already registered")
	}

	hash, err := tool.GenerateHash([]byte(userRequest.Password))
	if err != nil {
		return
	}

	userRequest.Password = string(hash)
	_, err = userModel.CreateUser(ctx, userRequest)
	if err != nil {
		return
	}

	return
}

func (u userController) GetUserList(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "user successfully retrieved",
		Translate: "user.get.success",
	}

	userModel := model.NewUserModel(u.Config.Postgres)
	userList, err := userModel.GetUserList(c, nil)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "user.get.error"

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = userList
	c.JSON(200, response)
}

func (u userController) UpdateUser(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "user successfully updated",
		Translate: "user.update.success",
	}

	c.JSON(200, response)
}

func (u userController) GetUser(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "user successfully retrieved",
		Translate: "user.get.success",
	}

	userId := c.Param("id")
	where := &model.Where{
		Parameter: "WHERE id=$1",
		Values:    []any{userId},
	}

	userModel := model.NewUserModel(u.Config.Postgres)
	currUser, err := userModel.FindUser(c, where)
	if err != nil {
		translate := "user.get.error"
		httpStatus := http.StatusInternalServerError

		if errors.Is(err, sql.ErrNoRows) {
			translate = "user.not.found"
			httpStatus = http.StatusNotFound
		}

		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = translate

		c.JSON(httpStatus, response)
		return
	}

	response.Data = currUser
	c.JSON(200, response)
}
