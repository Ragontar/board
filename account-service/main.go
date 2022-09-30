package main

import (
	"account-service/crud"
	"account-service/database"
	"account-service/server"
	"os"
)

var DEV bool = false

func main() {
	if DEV {
		setDevEnv()
	}
	database.GetDB()
	database.GetTokenStorage()
	database.GetSessionStorage()
	crud.Init()

	if err := server.Run("0.0.0.0:9000"); err != nil {
		panic(err)
	}
}

func setDevEnv() {
	// Postgres
	err := os.Setenv("DB_ADDR", "localhost:8081")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("DB_DATABASE", "postgres")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("DB_USER", "postgres")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("DB_PASSWORD", "postgres")
	if err != nil {
		panic(err)
	}

	//Redis
	err = os.Setenv("REDIS_ADDR", "localhost:8090")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("REDIS_PASSWORD", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("REDIS_DB", "0")
	if err != nil {
		panic(err)
	}
}
