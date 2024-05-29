package user

import "time"

type Status int

const (
	INACTIVE int = iota
	ACTIVE
)

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
