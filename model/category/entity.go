package category

import "time"

type Category struct {
	Id        int64      `json:"id,omitempty"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedBy int64      `json:"created_by"`
	UpdatedBy *int64     `json:"updated_by,omitempty"`
}
