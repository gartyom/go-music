package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gartyom/go-music/config"
	_ "github.com/lib/pq"
)

func Connect(c *config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.DbHost, c.DbPort, c.DbUser, c.DbPassword, c.DbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database succesfully connected!")

	return db
}
