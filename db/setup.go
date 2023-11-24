// models/setup.go

package db

import (
	"context"
	"example/web-service-gin/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var dsn string

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}

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

	ConnectToRedis()
	DB = database
}

func ConnectToPG() {

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
	fmt.Println("Redis answered:", pong)

	if err != nil {
		fmt.Println("Błąd podczas połączenia z Redisem:", err)
		return
	}

	err = rdb.Set(context.Background(), "TEST", "DUPA", 0).Err()
	if err != nil {
		fmt.Println("Błąd podczas wpisania do Redisa:", err)
	}

	val, err := rdb.Get(context.Background(), "TEST").Result()
	if err != nil {
		fmt.Println("Błąd podczas wczytania z Redisa:", err)
		return
	}
	fmt.Printf("Wartość dla klucza %s: %s", "TEST", val)
}
