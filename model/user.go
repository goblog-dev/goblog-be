package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type UserModel interface {
	CreateUser(ctx context.Context, user *User) (result sql.Result, err error)
	GetUserList(ctx context.Context, where *Where) (userList []*User, err error)
	FindUser(ctx context.Context, where *Where) (user *User, err error)
	UpdateOnlineStatus(ctx context.Context, user *User) (result sql.Result, err error)
	DeleteUser(ctx context.Context, userId int64) (result sql.Result, err error)
}

func NewUserModel(db *sql.DB) UserModel {
	return &PostgresRepository{db}
}

func (postgres *PostgresRepository) CreateUser(ctx context.Context, user *User) (result sql.Result, err error) {
	queryScript := `
		INSERT INTO users (
		    name
			, email
			, password
			, created_by
			, page
		) VALUES ($1, $2, $3, $4, $5)
	`

	return postgres.DB.ExecContext(ctx, queryScript,
		strings.ToLower(user.Name),
		strings.ToLower(user.Email),
		user.Password,
		user.CreatedBy,
		user.Page,
	)
}

func (postgres *PostgresRepository) GetUserList(ctx context.Context, where *Where) (userList []*User, err error) {
	where = ValidateWhere(where)

	queryScript := `
		SELECT	id
		     	, name
				, email
		     	, password
		    	, online
		     	
		     	, active
		     	, avatar
		     	, created_by
				, created_at
		    	, updated_at
		
				, updated_by
				, page
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
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Online,

			&user.Active,
			&user.Avatar,
			&user.CreatedBy,
			&user.CreatedAt,
			&user.UpdatedAt,

			&user.UpdatedBy,
			&user.Page,
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
		     	, name
				, email
		     	, password
		    	, online
		     	
		     	, active
		     	, avatar
		     	, created_by
				, created_at
		    	, updated_at
		
				, updated_by
				, page
		FROM 	users
	`

	query := fmt.Sprintf("%s %s", queryScript, where.Parameter)
	row := postgres.DB.QueryRowContext(ctx, query, where.Values...)

	user = new(User)
	err = row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Online,

		&user.Active,
		&user.Avatar,
		&user.CreatedBy,
		&user.CreatedAt,
		&user.UpdatedAt,

		&user.UpdatedBy,
		&user.Page,
	)

	if err != nil {
		return
	}

	return
}

func (postgres *PostgresRepository) UpdateOnlineStatus(ctx context.Context, user *User) (result sql.Result, err error) {
	queryScript := `
		UPDATE 	users SET 
		        online = $1
				, updated_by = $2      
		WHERE 	id = $3
		`

	return postgres.DB.ExecContext(ctx, queryScript,
		user.Online,
		user.UpdatedBy,
		user.Id,
	)
}

func (postgres *PostgresRepository) DeleteUser(ctx context.Context, userId int64) (result sql.Result, err error) {
	queryScript := `DELETE FROM users WHERE id = $1`
	return postgres.DB.ExecContext(ctx, queryScript, userId)
}
