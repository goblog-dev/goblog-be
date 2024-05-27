package category

import "time"

type Category struct {
	Id        int64      `json:"id,omitempty" form:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
