package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	helper "github.com/vamsikartik01/Ethanhunt/helpers"
)

var db *sql.DB

func InitConnection() error {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", helper.Config.Mysql.Username, helper.Config.Mysql.Password, helper.Config.Mysql.Host, helper.Config.Mysql.Port, helper.Config.Mysql.Database)
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println("Connection to Db failed with error : ", err)
		return err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Connection unsuccessfull with error :", err)
		return err
	}
	log.Println("Successfullt connected to Mysql db.")
	return nil
}

func CloseConnection() {
	db.Close()
}
