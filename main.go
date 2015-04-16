package main

import (
//	"database/sql"
	"dorm/orm"
//	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
//	db, err := sql.Open("mysql", "root:@(localhost:3306)/test")
//	if nil != err {
//		log.Println(err)
//		return
//	}
//	defer db.Close()

//	u := &orm.Users{"1234", "elvindu", 30, 1}
//	insertSql, err := orm.InsertSql(u)
//	if nil != err {
//		log.Println(err)
//		return
//	}

//	_, err = db.Query(insertSql)
//	if nil != err {
//		log.Println(err)
//		return
//	}

    u := orm.Users{}
    orm.SelectOne(&u,"SELECT * FROM users")
    log.Println(u)
}
