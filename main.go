package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Mysql in Go")

	db, err := sql.Open("mysql", "admin:phpmyadmin@tcp(127.0.0.1:3306)/news_site")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to MySQL")
}
