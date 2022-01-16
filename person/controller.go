package person

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.Engine, db *sql.DB) {
	r.GET("/person", func(c *gin.Context) {
		people := getAll(db)
		c.JSON(http.StatusOK, people)
	})

	r.GET("/person/:id", func(c *gin.Context) {
		strId := c.Param("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			panic("Cannot convert ID parameter: " + err.Error())
		}
		person := getById(db, id)
		c.JSON(http.StatusOK, person)
	})
}
