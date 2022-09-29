package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func SetupDatabase() {
	// info stored in environment
	host := os.Getenv("dbhost")
	port, _ := strconv.Atoi(os.Getenv("dbport"))
	user := os.Getenv("dbuser")
	password := os.Getenv("dbpassword")
	dbname := os.Getenv("dbname")

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("> failed connecting to db")
		panic(err)
	}

	DB.SetMaxOpenConns(3)
	DB.SetMaxIdleConns(3)
	DB.SetConnMaxLifetime(60 * time.Second)
}
