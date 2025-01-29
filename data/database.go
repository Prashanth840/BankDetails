package data

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func DbConnect() (err error) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "@Pachi840",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "prashanth",
	}
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err = Db.Ping(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("MySQL init done")
	return nil
}
