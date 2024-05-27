package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func CreateUser(ctx context.Context, postgres *sql.DB, user *User) (err error) {
	queryScript := `
		INSERT INTO users (
			email
			, password
			, name
			, status
		) VALUES ($1, $2, $3, $4)
	`

	_, err = postgres.ExecContext(
		ctx,
		queryScript,
		user.Email,
		user.Password,
		user.Name,
		user.Status,
	)

	return
}

func GetUserList(ctx context.Context, postgres *sql.DB, where string, value []any) (userList []*User, err error) {
	queryScript := `
		SELECT	id
				, email
				, name
		    	, status
		    	, created_at
		     
		    	, updated_at
		FROM 	users
	`

	query := fmt.Sprintf("%s %s", queryScript, where)
	rows, err := postgres.QueryContext(ctx, query, value...)
	if err != nil {
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("error closing get user list rows:", err)
		}
	}(rows)

	userList = make([]*User, 0)

	for rows.Next() {
		user := new(User)

		err = rows.Scan(
			&user.Id,
			&user.Email,
			&user.Name,
			&user.Status,
			&user.CreatedAt,

			&user.UpdatedAt,
		)

		if err != nil {
			return
		}

		userList = append(userList, user)
	}

	return
}

func FindUser(ctx context.Context, postgres *sql.DB, where string, value []any) (user *User, err error) {
	queryScript := `
		SELECT	id
				, email
		     	, password
				, name
				, status
		     
				, created_at
				, updated_at
		FROM 	users
	`

	query := fmt.Sprintf("%s %s", queryScript, where)
	row := postgres.QueryRowContext(ctx, query, value...)

	user = new(User)
	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Status,

		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return
	}

	return
}
