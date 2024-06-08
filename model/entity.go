package model

import (
	"database/sql"
	"time"
)

type Status int

const (
	INACTIVE int = iota
	ACTIVE
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
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	CreatedBy  int64      `json:"created_by"`
	UpdatedBy  *int64     `json:"updated_by,omitempty"`
}

type ArticleWithExtend struct {
	Article
	ArticleExtend
}

type Category struct {
	Id        int64      `json:"id,omitempty"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedBy int64      `json:"created_by"`
	UpdatedBy *int64     `json:"updated_by,omitempty"`
}

type User struct {
	Id        int64      `json:"id,omitempty"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Name      string     `json:"name"`
	Status    Status     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedBy int64      `json:"created_by"`
	UpdatedBy *int64     `json:"updated_by,omitempty"`
}
