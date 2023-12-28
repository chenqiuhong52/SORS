package main

import (
	"SORS/biz/router"
	"SORS/dal/db"
)

func main() {
	db.InitMySQL()
	router.InitRouter()
}
