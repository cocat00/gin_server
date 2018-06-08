package database

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DB_NAME	    = "root:070417@(127.0.0.1:3306)/test?parseTime=true"
	_SQLITES_DRIVER = "mysql"
)

var SqlDB *sql.DB

func init()  {
	var err error
	SqlDB, err = sql.Open(_SQLITES_DRIVER, _DB_NAME)
	if err != nil {
		log.Fatalln(err)
	}

	SqlDB.SetMaxIdleConns(20)
	SqlDB.SetMaxOpenConns(20)

	if err := SqlDB.Ping(); err != nil {
		log.Fatalln(err)
	}

}