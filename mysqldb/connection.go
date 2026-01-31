package mysqldb

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	var db *sql.DB
	cfg := mysql.NewConfig()
	cfg.User = "teste"
	cfg.Passwd = "Teste123@"
	cfg.Net = "tcp"
	cfg.Addr = "localhost:3306"
	cfg.DBName = "super_market"
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db

}
