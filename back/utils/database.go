// database connection
package utils

import (
	"log"
	"fmt"

)

// Substitute your DB details
const (
	host     = "localhost"
	port     = 5432
	user     = "username"  
	password = "password"
	dbname   = "postgres"
)

func GetConnection() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	log.Println("DB Connection established...")
	return db
}