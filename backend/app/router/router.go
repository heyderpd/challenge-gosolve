package router

import (
	"net/http"
	"strconv"

  "github.com/toorop/gin-logrus"
	"github.com/gin-gonic/gin"

	"challenge-gosolve/backend/app/service"
	"challenge-gosolve/backend/app/utils"
)

type Router struct {
	db service.DatabaseInterface
}

type RouterInterface interface {
	getRouteHealthcheck(c *gin.Context)
	getRouteFindValue(c *gin.Context)
}

func (route *Router) getRouteHealthcheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func (route *Router) getRouteFindValue(c *gin.Context) {
	valueStr := c.Params.ByName("value")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"value": value, "status": "value invalid"})
		return;
	}
	index, err := route.db.FindIndex(value)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"value": value, "status": "value not_found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"value": value, "index": index})
	}
}

func SetupRouter(db service.DatabaseInterface) *gin.Engine {
	router := Router{ db }
	r := gin.Default()
	r.Use(ginlogrus.Logger(utils.Log), gin.Recovery())
	r.GET("/health", router.getRouteHealthcheck)
	r.GET("/endpoint/:value", router.getRouteFindValue)
	return r
}
