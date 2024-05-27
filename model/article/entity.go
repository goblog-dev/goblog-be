package article

import "time"

type Article struct {
	Id         int64      `json:"id,omitempty" form:"id"`
	UserId     int64      `json:"user_id"`
	CategoryId int64      `json:"category_id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	CreatedAt  time.Time  `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
