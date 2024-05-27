package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelwp/goblog/model/user"
	"github.com/michaelwp/goblog/tool"
	"log"
	"net/http"
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

	token, err := a.LoginProcess(c, &loginCredential)
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

	response.Data = map[string]interface{}{"token": token}
	c.JSON(http.StatusAccepted, response)
}

func (a authorizationController) LoginProcess(ctx context.Context, cred *LoginCredential) (token string, err error) {
	where := "WHERE email = $1"
	value := []interface{}{cred.Email}

	currUser, err := user.FindUser(ctx, a.Postgres, where, value)
	if err != nil {
		return
	}

	log.Println("currUser:", currUser)
	log.Println("hash:", currUser.Password)

	err = tool.CompareHashAndPassword([]byte(currUser.Password), []byte(cred.Password))
	if err != nil {
		return
	}

	token, err = tool.GenerateJWT(currUser.Id)
	if err != nil {
		return
	}

	return
}

func (a authorizationController) Logout(c *gin.Context) {
	response := &Response{
		Status:    SUCCESS,
		Message:   "user successfully logout",
		Translate: "user.success.logout",
	}

	c.JSON(http.StatusOK, response)
}
