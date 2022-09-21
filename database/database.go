package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func Main() {
	fmt.Println("> attempting connection to db")
	db := openConnection()
	defer db.Close()

	/*
		stmt := "INSERT INTO test (num, str) VALUES (4, 'test');"
		fmt.Println("> attempting execution of: " + stmt)
		_, err := db.Exec(stmt)
		if err != nil {
			fmt.Println("> failed execution of: " + stmt)
			panic(err)
		}
	*/

	stmt1 := "SELECT * FROM test;"
	fmt.Println("> attempting execution of: " + stmt1)
	rows, err := db.Query(stmt1)
	if err != nil {
		fmt.Println("> failed execution of: " + stmt1)
		panic(err)
	}

	var id int
	var num int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &num, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, "-", num, "-", name)
	}

	fmt.Println("> closing connection to db")
}

func openConnection() *sql.DB {
	host := os.Getenv("dbhost")
	port, _ := strconv.Atoi(os.Getenv("dbport"))
	user := os.Getenv("dbuser")
	password := os.Getenv("dbpassword")
	dbname := os.Getenv("dbname")

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("> failed connecting to db")
		panic(err)
	}
	fmt.Println("> succeded connecting to db")
	return db
}
