package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

type UserModel interface {
	CreateUser(ctx context.Context, user *User) (result sql.Result, err error)
	GetUserList(ctx context.Context, where *Where) (userList []*User, err error)
	FindUser(ctx context.Context, where *Where) (user *User, err error)
	UpdateUser(ctx context.Context, user *User) (result sql.Result, err error)
	DeleteUser(ctx context.Context, userId int64) (result sql.Result, err error)
}

func NewUserModel(db *sql.DB) UserModel {
	return &PostgresRepository{db}
}

func (postgres *PostgresRepository) CreateUser(ctx context.Context, user *User) (result sql.Result, err error) {
	queryScript := `
		INSERT INTO users (
			email
			, password
			, name
			, status
			, created_by
		) VALUES ($1, $2, $3, $4, $5)
	`

	return postgres.DB.ExecContext(ctx, queryScript,
		strings.ToLower(user.Email),
		user.Password,
		strings.ToLower(user.Name),
		user.Status,
		user.CreatedBy,
	)
}

func (postgres *PostgresRepository) GetUserList(ctx context.Context, where *Where) (userList []*User, err error) {
	where = ValidateWhere(where)

	queryScript := `
		SELECT	id
				, email
		     	, password
				, name
		    	, status
		     
		    	, created_at
		    	, updated_at
				, created_by
				, updated_by
		FROM 	users
	`

	query := fmt.Sprintf("%s %s", queryScript, where.Parameter)
	rows, err := postgres.DB.QueryContext(ctx, query, where.Values...)
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
			&user.Password,
			&user.Name,
			&user.Status,

			&user.CreatedAt,
			&user.UpdatedAt,
			&user.CreatedBy,
			&user.UpdatedBy,
		)

		if err != nil {
			return
		}

		userList = append(userList, user)
	}

	return
}

func (postgres *PostgresRepository) FindUser(ctx context.Context, where *Where) (user *User, err error) {
	where = ValidateWhere(where)
	queryScript := `
		SELECT	id
				, email
		     	, password
				, name
		    	, status
		     
		    	, created_at
		    	, updated_at
				, created_by
				, updated_by
		FROM 	users
	`

	query := fmt.Sprintf("%s %s", queryScript, where.Parameter)
	row := postgres.DB.QueryRowContext(ctx, query, where.Values...)

	user = new(User)
	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Status,

		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
	)

	if err != nil {
		return
	}

	return
}

func (postgres *PostgresRepository) UpdateUser(ctx context.Context, user *User) (result sql.Result, err error) {
	queryScript := `
		UPDATE 	users SET 
		    	email = $1
				, password = $2
		        , name = $3
		        , status = $4
				, updated_by = $5
		              
		        , updated_at = $6
		WHERE 	id = $7
		`

	return postgres.DB.ExecContext(ctx, queryScript,
		strings.ToLower(user.Email),
		user.Password,
		strings.ToLower(user.Name),
		user.Status,
		user.UpdatedBy,

		time.Now(),
		user.Id,
	)
}

func (postgres *PostgresRepository) DeleteUser(ctx context.Context, userId int64) (result sql.Result, err error) {
	queryScript := `DELETE FROM users WHERE id = $1`
	return postgres.DB.ExecContext(ctx, queryScript, userId)
}