package controller

import "database/sql"

const (
	ERROR   = "error"
	SUCCESS = "success"
)

type Response struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Translate string `json:"translate"`
	Data      any    `json:"data,omitempty"`
}

type Config struct {
	Postgres *sql.DB
}
