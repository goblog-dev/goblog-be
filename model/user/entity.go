package user

import "time"

type User struct {
	Id        int64      `json:"id,omitempty" form:"id"`
	Email     string     `json:"email" form:"email"`
	Password  string     `json:"password,omitempty"`
	Name      string     `json:"name" form:"name"`
	Status    int        `json:"status" form:"status"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
