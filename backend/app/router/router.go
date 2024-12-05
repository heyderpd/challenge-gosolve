package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"challenge-gosolve/backend/app/service"
)

type Router struct {
	db service.DatabaseInterface
}

type RouterInterface interface {
	getRouteHealthcheck(c *gin.Context)
	getRouteFindIndex(c *gin.Context)
}

func (route *Router) getRouteHealthcheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func (route *Router) getRouteFindIndex(c *gin.Context) {
	indexStr := c.Params.ByName("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"index": index, "status": "index invalid"})
		return;
	}
	value, err := route.db.FindIndex(int(index))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"index": index, "status": "no value"})
	} else {
		c.JSON(http.StatusOK, gin.H{"index": index, "value": value})
	}
}

func SetupRouter(db service.DatabaseInterface) *gin.Engine {
	router := Router{ db }
	r := gin.Default()
	r.GET("/health", router.getRouteHealthcheck)
	r.GET("/endpoint/:index", router.getRouteFindIndex)
	return r
}
