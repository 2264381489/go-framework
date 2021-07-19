package main

import (
	"database/sql"
	"github.com/apsdehal/go-logger"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log, err := logger.New("test", 1, os.Stdout)
	if err != nil {
		panic(err) // Check for error
	}
	db, err := sql.Open("mysql", "root:a13840419132@/seckill")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 测试是否能访问
	err = db.Ping()
	if err != nil {
		log.Error(err.Error())
	}
	_, err = db.Exec("DROP TABLE IF EXISTS User;")
	if err != nil {
		log.Error(err.Error())
	}
	_, err = db.Exec("create table user(name varchar(64)  null comment '名字',id  int not null comment '用户账号');")
	if err != nil {
		log.Error(err.Error())
	}
	_, err = db.Exec("insert into user values ('yan',10086)")
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("insert into")
	}

}
