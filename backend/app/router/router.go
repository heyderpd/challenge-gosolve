package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	db map[string]string
}

type RouterInterface interface {
	getRouteHealthcheck(c *gin.Context)
	getRouteFindIndex(c *gin.Context)
}

func (route *Router) getRouteHealthcheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func (route *Router) getRouteFindIndex(c *gin.Context) {
	index := c.Params.ByName("index")
	value, ok := route.db[index]
	if ok {
		c.JSON(http.StatusOK, gin.H{"index": index, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"index": index, "status": "no value"})
	}
}

func SetupRouter() *gin.Engine {
	var db = make(map[string]string)
	router := Router{ db }
	r := gin.Default()
	r.GET("/health", router.getRouteHealthcheck)
	r.GET("/endpoint/:index", router.getRouteFindIndex)
	return r
}
