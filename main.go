package main

import (
	db "gin_server/database"
	"gin_server/routers"
)

func main() {
	defer db.SqlDB.Close()
	router := routers.InitRouter()
	router.Run(":8999")}