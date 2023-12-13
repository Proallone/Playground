// models/setup.go

package db

import (
	"context"
	"example/web-service-gin/models"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var REDIS *redis.Client

const (
	dockerPGConnString = "DOCKER_PG_CONN_STRING"
	localPGConnString  = "LOCAL_PG_CONN_STRING"
	environment        = "ENVIROMENT"
	dockerRedisAddr    = "DOCKER_REDIS_ADDR"
	localRedisAddr     = "LOCAL_REDIS_ADDR"
	redisPassword      = "REDIS_PASSWORD"
)

func ConnectDatabases() {

	ConnectToPG()
	ConnectToRedis()
}

func ConnectToPG() {
	var dsn string

	if os.Getenv(environment) == "docker" {
		dsn = os.Getenv(dockerPGConnString)
	} else {
		dsn = os.Getenv(localPGConnString)
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to postgres with error:", err)
		return
	}

	err = database.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println("Failed to automigrate postgres schema with error: ", err)
		return
	}

	fmt.Println("Server connected to postgres")
	DB = database
}

func ConnectToRedis() {

	redis_pass := os.Getenv(redisPassword)
	var redis_addr string

	if os.Getenv(environment) == "docker" {
		redis_addr = os.Getenv(dockerRedisAddr)
	} else {
		redis_addr = os.Getenv(localRedisAddr)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: redis_pass,
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println("An error occured during redis connection: ", err)
		return
	}

	fmt.Println("Server connected to redis")

	REDIS = rdb
}
