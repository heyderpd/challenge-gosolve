package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/endpoint/:index", func(c *gin.Context) {
		index := c.Params.ByName("index")
		value, ok := db[index]
		if ok {
			c.JSON(http.StatusOK, gin.H{"index": index, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"index": index, "status": "no value"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":3000")
}
