package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisDBConfig struct {
	User, Host, Port, Pass, DB string
}

func (r *RedisDBConfig) Connect() (client *redis.Client) {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
		Password: r.Pass,
		DB:       0,
	})
}

func (r *RedisDBConfig) ConnectWithString() (client *redis.Client, err error) {
	connectionString := fmt.Sprintf("redis://%s:%s@%s:%s/0", r.User, r.Pass, r.Host, r.Port)

	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}

	return redis.NewClient(opt), nil
}
