package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresDBConfig struct {
	Host, Port, User, Pass, Name, SslMode string
}

func (db *PostgresDBConfig) Connect() (postgresDb *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Pass, db.Name, db.SslMode)

	sqlClient, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	sqlClient.SetMaxIdleConns(5)

	return sqlClient, nil
}
