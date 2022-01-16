package main

import (
	"database/sql"
	"os"

	"rest-oracle-go-seed/person"

	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
)

func main() {

	// Setup database connection
	dataSourceName := os.Getenv("DATA_SOURCE_NAME")
	db, err := sql.Open("godror", dataSourceName)
	if err != nil {
		panic("Cannot connect to the DB: " + err.Error())
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic("Cannot ping the DB: " + err.Error())
	}

	// Inject GIN and the database connection into the controller(s).
	r := gin.Default()
	person.Controller(r, db)
	r.Run()
}
