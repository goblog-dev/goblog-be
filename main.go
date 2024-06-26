package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/michaelwp/goblog/api"
	"github.com/michaelwp/goblog/api/controller"
	"github.com/michaelwp/goblog/config"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(os.Getenv(gin.EnvGinMode))
}

func main() {
	r := gin.Default()

	api.NewRouter(r, SetupConfig())

	log.Println("server started on port " + os.Getenv("APP_SERVER_PORT"))
	err := r.Run(os.Getenv("APP_SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}

func SetupConfig() (config *controller.Config) {
	postgres, err := SetupPostgres()
	if err != nil {
		log.Fatal("error connect to postgres", err)
	}

	client, err := SetupRedis()
	if err != nil {
		log.Fatal("error connect to redis", err)
	}

	config = &controller.Config{
		Postgres:    postgres,
		RedisClient: client,
	}

	return
}

func SetupPostgres() (postgres *sql.DB, err error) {
	configDb := &config.PostgresDBConfig{
		Host:    os.Getenv("POSTGRES_DB_HOST"),
		Port:    os.Getenv("POSTGRES_DB_PORT"),
		User:    os.Getenv("POSTGRES_DB_USER"),
		Pass:    os.Getenv("POSTGRES_DB_PASS"),
		Name:    os.Getenv("POSTGRES_DB_NAME"),
		SslMode: os.Getenv("POSTGRES_DB_SSL_MODE"),
	}

	return configDb.Connect()
}

func SetupRedis() (client *redis.Client, err error) {
	configRedis := &config.RedisDBConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
		Pass: os.Getenv("REDIS_PASSWORD"),
		DB:   os.Getenv("REDIS_DB"),
		User: os.Getenv("REDIS_USER"),
	}

	return configRedis.ConnectWithString()
}

func SetupStaticFile(r *gin.Engine) {
	var fileSystem http.FileSystem
	fileSystem = http.Dir(os.Getenv("APP_CLIENT_FILE"))

	r.Group("/")
	r.NoRoute(func(c *gin.Context) {
		http.FileServer(fileSystem).ServeHTTP(c.Writer, c.Request)
	})
}
