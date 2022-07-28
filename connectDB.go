package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

const (
	userName = "root"
	password = "1234567890"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "go"
)

func GetDataBase() *sql.DB {
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Println(path)
	DB, _ := sql.Open("mysql", path)
	if DB == nil {
		log.Fatal("連線失敗！")
		return nil
	}
	DB.SetConnMaxLifetime(10)
	DB.SetMaxIdleConns(5)
	if err := DB.Ping(); err != nil{
		log.Fatal("opon database fail")
		return nil
	}
	return DB
}