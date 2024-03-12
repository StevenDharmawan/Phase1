package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/Toko_Buku")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to Database")
}
