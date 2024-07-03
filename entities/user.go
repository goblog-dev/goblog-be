package entities

import "time"

type Status bool

type User struct {
	Id        int64      `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Password  string     `json:"password,omitempty"`
	Online    Status     `json:"online,omitempty"`
	Active    Status     `json:"active,omitempty"`
	Avatar    *string    `json:"avatar,omitempty"`
	Page      *string    `json:"page,omitempty"`
	CreatedBy int64      `json:"created_by,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *int64     `json:"updated_by,omitempty"`
}
