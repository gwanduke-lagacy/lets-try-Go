package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	var email string

	// 트랜잭션 시작
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() //중간에 에러시 롤백

	rows, err := db.Query("SELECT email, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 트랜젝션 커밋
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&name, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name, email)
	}
}
