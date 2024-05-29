package article

import "time"

type Article struct {
	Id           int64      `json:"id,omitempty"`
	UserId       int64      `json:"user_id"`
	UserName     string     `json:"user_name"`
	CategoryId   int64      `json:"category_id"`
	CategoryName string     `json:"category_name"`
	Title        string     `json:"title"`
	Content      string     `json:"content"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	CreatedBy    int64      `json:"created_by"`
	UpdatedBy    *int64     `json:"updated_by,omitempty"`
}
