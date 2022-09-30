package database

import (
	"account-service/dry"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	db_addr     string
	db_database string
	db_user     string
	db_password string
)

// var connectionConfig pgx.ConnConfig
var db *pgxpool.Pool

func newDBConnection() *pgxpool.Pool {
	err := lookupDBEnv()
	if err != nil {
		panic(fmt.Sprintf("[DATABASE]: Cannot establish DB connection: %v", err))
	}

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s",
		db_user,
		db_password,
		db_addr,
		db_database,
	)
	log.Println("[DATABASE]: dsn initialized")

	db, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(fmt.Sprintf("[DATABASE]: Cannot establish DB connection: %v", err))
	}
	log.Println("[DATABASE]: connection established")

	return db
}

func GetDB() *pgxpool.Pool {
	if db == nil {
		db = newDBConnection()
	}
	return db
}

func lookupDBEnv() error {
	db_addr = dry.LookupOrPanic("DB_ADDR")

	db_database = dry.LookupOrPanic("DB_DATABASE")

	db_user = dry.LookupOrPanic("DB_USER")

	db_password = dry.LookupOrPanic("DB_PASSWORD")

	log.Printf(
		"[DATABASE]: DSN is set!\nADDR: %s\nDATABASE: %s\nUSER: %s\nPASSWORD: %s\n",
		db_addr,
		db_database,
		db_user,
		db_password,
	)

	return nil
}
