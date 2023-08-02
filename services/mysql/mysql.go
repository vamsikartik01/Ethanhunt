package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitConnection() {
	const (
		host     = "localhost"
		port     = "3300"
		user     = "root"
		password = "password"
		dbname   = "jack_db"
	)

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println("Connection to postgres failed with error : ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Println("Connection unsuccessfull with error :", err)
	}
	log.Println("Successfullt connected to Mysql db.")

	rows, err := db.Query("select * from Accounts")
	if err != nil {
		log.Fatalf("Error retriving data: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id          int
			name        string
			email       string
			pass        string
			saltid      int
			dateCreated string
			passDate    string
		)
		err := rows.Scan(&id, &name, &email, &pass, &saltid, &dateCreated, &passDate)
		if err != nil {
			log.Fatalf("error", err)
		}
		log.Println(id, name, email, pass, saltid, dateCreated, passDate)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("", err)
	}
	_ = GetRooms()
}
