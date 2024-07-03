package model

import "database/sql"

func ValidateWhere(where *Where) (whereNew *Where) {
	whereNew = where

	if whereNew == nil {
		whereNew = &Where{
			Parameter: "",
			Values:    []any{},
			Order:     "",
			Limit:     "",
		}
	}

	return
}

type PostgresRepository struct {
	DB *sql.DB
}

const (
	INACTIVE = false
	ACTIVE   = true
)

type Where struct {
	Parameter string
	Values    []any
	Order     string
	Limit     string
}
