package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/model"
	"github.com/michaelwp/goblog/model/user"
	"github.com/michaelwp/goblog/tool"
	"net/http"
	"strconv"
	"time"
)

type LoginCredential struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type AuthorizationController interface {
	Login(*gin.Context)
	Logout(*gin.Context)
}

type authorizationController struct {
	*Config
}

func NewAuthorizationController(c *Config) AuthorizationController {
	return &authorizationController{c}
}

func (a authorizationController) Login(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "user successfully login",
		Translate: "user.success.login",
	}

	var loginCredential LoginCredential
	err := c.ShouldBindJSON(&loginCredential)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "user.error.login"

		c.JSON(http.StatusUnauthorized, response)
		return
	}

	userId, token, err := a.LoginProcess(c, &loginCredential)
	if err != nil {
		translate := "user.error.login"
		if errors.Is(err, sql.ErrNoRows) {
			translate = "email.or.password.not.found"
		}

		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = translate

		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err = a.RedisClient.Set(c, strconv.FormatInt(userId, 10), token, 24*time.Hour).Err()
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "user.error.login"

		c.JSON(http.StatusUnauthorized, response)
		return
	}

	response.Data = map[string]interface{}{"token": token}
	c.JSON(http.StatusAccepted, response)
}

func (a authorizationController) LoginProcess(ctx context.Context, cred *LoginCredential) (
	userId int64, token string, err error) {

	where := &model.Where{
		Parameter: "WHERE email = $1 AND status = $2",
		Values:    []interface{}{cred.Email, user.ACTIVE},
	}

	currUser, err := user.FindUser(ctx, a.Postgres, where)
	if err != nil {
		return
	}

	err = tool.CompareHashAndPassword([]byte(currUser.Password), []byte(cred.Password))
	if err != nil {
		return
	}

	token, err = tool.GenerateJWT(currUser.Id)
	if err != nil {
		return
	}

	userId = currUser.Id
	return
}

func (a authorizationController) Logout(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "user successfully logout",
		Translate: "user.success.logout",
	}

	userId, err := GetCurrentUserIdLoggedIn(c)
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "user.error.logout"

		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err = a.RedisClient.Del(c, strconv.FormatInt(userId, 10)).Err()
	if err != nil {
		response.Status = ERROR
		response.Message = err.Error()
		response.Translate = "user.error.logout"

		c.JSON(http.StatusUnauthorized, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
