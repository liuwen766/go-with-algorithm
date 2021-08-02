package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DSN:Data Source Name
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面
}
