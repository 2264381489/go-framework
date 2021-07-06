package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
func main() {

	db,err:=sql.Open("mysql","root:a13840419132@tcp(127.0.0.1:3306)/seckill")
	if err != nil{
		panic(err)
	}
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}

}
