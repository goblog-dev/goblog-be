package entities

import "time"

type Article struct {
	Id          int64      `json:"id,omitempty"`
	UserId      int64      `json:"user_id,omitempty"`
	CategoryId  int64      `json:"category_id,omitempty"`
	Content     string     `json:"content,omitempty"`
	Title       string     `json:"title,omitempty"`
	Tags        *string    `json:"tags,omitempty"`
	Description *string    `json:"description,omitempty"`
	Image       *string    `json:"image,omitempty"`
	CreatedBy   int64      `json:"created_by,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	UpdatedBy   *int64     `json:"updated_by,omitempty"`
}
