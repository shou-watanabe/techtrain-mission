package sql

import (
	"database/sql"
	"log"

	"go-gacha-server/src/config"
)

func NewDriver() *sql.DB {
	db, err := sql.Open("mysql", config.GetDbUri())
	if err != nil {
		log.Println("DB connect failed")
		panic(err)
	}

	log.Println("DB connect success")

	return db
}
