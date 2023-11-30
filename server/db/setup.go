// models/setup.go

package db

import (
	"context"
	"example/web-service-gin/models"
	"fmt"
	"os"

	sredis "github.com/gin-contrib/sessions/redis"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var REDIS *redis.Client
var SESSION_STORE sredis.Store

func ConnectDatabases() {

	ConnectToPG()
	ConnectToRedis()
}

func ConnectToPG() {
	var dsn string
	if os.Getenv("ENVIROMENT") == "docker" {
		dsn = os.Getenv("DOCKER_PG_CONN_STRING")
	} else {
		dsn = os.Getenv("LOCAL_PG_CONN_STRING")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	fmt.Println("Postgres is up")
	DB = database
}

func ConnectToRedis() {
	var redis_addr string
	redis_pass := os.Getenv("REDIS_PASSWORD")

	if os.Getenv("ENVIROMENT") == "docker" {
		redis_addr = os.Getenv("DOCKER_REDIS_ADDR")
	} else {
		redis_addr = os.Getenv("LOCAL_REDIS_ADDR")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: redis_pass,
		DB:       0,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("An error occured during redis connection: ", err)
		return
	}

	fmt.Println("Redis is up and says:", pong)

	REDIS = rdb
}

func CreateSessionStore() {
	store, err := sredis.NewStore(1, "tcp", "localhost:6379", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", []byte("secret"))
	if err != nil {
		fmt.Println("Błąd podczas inicjalizacji magazynu sesji:", err)
		return
	}
	SESSION_STORE = store
}
