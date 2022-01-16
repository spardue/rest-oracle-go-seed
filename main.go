package main

import (
	"database/sql"

	"rest-oracle-go-seed/person"

	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
)

func main() {

	db, err := sql.Open("godror", "goseed/password1234@192.168.1.93:1521/orclcdb")
	if err != nil {
		panic("Cannot connect to the DB: " + err.Error())
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic("Cannot ping the DB: " + err.Error())
	}

	r := gin.Default()

	person.Controller(r, db)
	r.Run()
}
