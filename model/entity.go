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

type Article struct {
	Id          int64      `json:"id,omitempty" db:"id"`
	UserId      int64      `json:"user_id" db:"user_id"`
	CategoryId  int64      `json:"category_id" db:"category_id"`
	Content     string     `json:"content" db:"content"`
	Title       string     `json:"title" db:"title"`
	Tags        *string    `json:"tags,omitempty" db:"tags"`
	Description *string    `json:"description" db:"description"`
	Image       *string    `json:"image,omitempty" db:"image"`
	CreatedBy   int64      `json:"created_by" db:"created_by"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy   *int64     `json:"updated_by,omitempty" db:"updated_by"`
}

type Category struct {
	Id        int64      `json:"id,omitempty" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedBy int64      `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy *int64     `json:"updated_by,omitempty" db:"updated_by"`
}

type User struct {
	Id        int64      `json:"id,omitempty"  db:"id"`
	Name      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"password"`
	Online    Status     `json:"online" db:"online"`
	Active    Status     `json:"active" db:"active"`
	Avatar    *string    `json:"avatar,omitempty" db:"avatar"`
	Page      *string    `json:"page,omitempty" db:"page"`
	CreatedBy int64      `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy *int64     `json:"updated_by,omitempty" db:"updated_by"`
}
