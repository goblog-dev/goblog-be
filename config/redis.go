package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisDBConfig struct {
	Host, Port, Pass, DB string
}

func (r *RedisDBConfig) Connect() (client *redis.Client) {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
		Password: r.Pass,
		DB:       0,
	})
}
