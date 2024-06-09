package model

import (
	"database/sql"
	"time"
)

type Status bool

const (
	INACTIVE = false
	ACTIVE   = true
)

type Where struct {
	Parameter string
	Values    []any
	Order     string
	Limit     string
}

type PostgresRepository struct {
	DB *sql.DB
}

type ArticleExtend struct {
	UserName     string `json:"user_name"`
	CategoryName string `json:"category_name"`
}

type Article struct {
	Id         int64      `json:"id,omitempty"`
	UserId     int64      `json:"user_id"`
	CategoryId int64      `json:"category_id"`
	Content    string     `json:"content"`
	Title      string     `json:"title"`
	Tag        *[]string  `json:"tag"`
	CreatedBy  int64      `json:"created_by"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	UpdatedBy  *int64     `json:"updated_by,omitempty"`
}

type ArticleWithExtend struct {
	Article
	ArticleExtend
}

type Category struct {
	Id        int64      `json:"id,omitempty"`
	Name      string     `json:"name"`
	CreatedBy int64      `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *int64     `json:"updated_by,omitempty"`
}

type User struct {
	Id        int64      `json:"id,omitempty"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Online    Status     `json:"online"`
	Active    Status     `json:"active"`
	Avatar    *string    `json:"avatar,omitempty"`
	CreatedBy int64      `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *int64     `json:"updated_by,omitempty"`
}
