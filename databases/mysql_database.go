package databases

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:qwerty@tcp(localhost:3306)/fatih_be?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}
