package entities

import "time"

type Category struct {
	Id        int64      `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	CreatedBy int64      `json:"created_by,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *int64     `json:"updated_by,omitempty"`
}
