package vehicle

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.Engine, db *sql.DB) {
	r.GET("/vehicle", func(c *gin.Context) {
		c.JSON(http.StatusOK, 1)
	})

	r.GET("/vehicle/:id", func(c *gin.Context) {
		strId := c.Param("id")
		id, err := strconv.Atoi(strId)
		if err != nil {
			panic("Cannot convert ID parameter: " + err.Error())
		}
		c.JSON(http.StatusOK, id)
	})
}
