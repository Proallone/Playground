// models/setup.go

package db

import (
	"fmt"
	"os"
)

const (
	dockerPGConnString = "DOCKER_PG_CONN_STRING"
	localPGConnString  = "LOCAL_PG_CONN_STRING"
	environment        = "ENVIROMENT"
	dockerRedisAddr    = "DOCKER_REDIS_ADDR"
	localRedisAddr     = "LOCAL_REDIS_ADDR"
	redisPassword      = "REDIS_PASSWORD"
)

func ConnectDatabases() error {
	var pgConnString string

	if os.Getenv(environment) == "docker" {
		pgConnString = os.Getenv(dockerPGConnString)
	} else {
		pgConnString = os.Getenv(localPGConnString)
	}

	db, err := ConnectToPG(pgConnString)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	var redisAddr string
	redisPass := os.Getenv(redisPassword)

	if os.Getenv(environment) == "docker" {
		redisAddr = os.Getenv(dockerRedisAddr)
	} else {
		redisAddr = os.Getenv(localRedisAddr)
	}

	redisClient, err := ConnectToRedis(redisAddr, redisPass)
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	DB = db
	REDIS = redisClient

	return nil
}
