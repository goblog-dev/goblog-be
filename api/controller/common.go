package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/redis/go-redis/v9"
)

const (
	ERROR   = "error"
	SUCCESS = "success"
)

type Response struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Translate string `json:"translate"`
	Data      any    `json:"data,omitempty"`
	HttpCode  int    `json:"http_code,omitempty"`
}

type Config struct {
	Postgres    *sql.DB
	RedisClient *redis.Client
}

func GetCurrentUserIdLoggedIn(ctx context.Context) (userId int64, err error) {
	userIdCtx := ctx.Value("user_id")
	if userIdCtx == nil {
		err = errors.New("user not found")
		return
	}

	userId, ok := userIdCtx.(int64)
	if !ok {
		err = errors.New("user not found")
		return
	}

	return
}
