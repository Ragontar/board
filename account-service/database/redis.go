package database

import (
	"account-service/dry"
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v9"
)

var (
	redis_addr     string
	redis_password string
)

const (
	REDIS_TOKENS = iota
	REDIS_TELEGRAM_SESSIONS
)

var (
	redisTokenStorage   *redis.Client
	redisSessionStorage *redis.Client
)

func NewRedisClient(db int) *redis.Client {
	err := lookupRedisEnv()
	if err != nil {
		panic(fmt.Sprintf("[REDIS DB %v]: Cannot initialize parameters: %v", db, err))
	}
	log.Println("[REDIS]: parameters initiated")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: redis_password,
		DB:       db,
	})
	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	log.Printf("[REDIS DB %v]: %v\n", db, "connection established")

	return rdb
}

func GetTokenStorage() *redis.Client {
	if redisTokenStorage == nil {
		redisTokenStorage = NewRedisClient(REDIS_TOKENS)
	}
	return redisTokenStorage
}

func GetSessionStorage() *redis.Client {
	if redisSessionStorage == nil {
		redisSessionStorage = NewRedisClient(REDIS_TELEGRAM_SESSIONS)
	}
	return redisSessionStorage
}

func lookupRedisEnv() error {
	redis_addr = dry.LookupOrPanic("REDIS_ADDR")

	redis_password = dry.LookupOrPanic("REDIS_PASSWORD")

	log.Printf(
		"[REDIS]: connection parameters are set!\nADDR: %s\nPASSWORD: %s\n",
		redis_addr,
		redis_password,
	)

	return nil
}
