package router

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"

	"challenge-gosolve/backend/app/service"
	"challenge-gosolve/backend/app/utils"
)

type Router struct {
	se service.SearcherInterface
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
		return
	}
	index, err := route.se.FindIndex(value)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"value": value, "status": "value not_found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"value": value, "index": index})
	}
}

func SetupRouter(se service.SearcherInterface) *gin.Engine {
	router := Router{se}
	r := gin.Default()
	r.Use(ginlogrus.Logger(utils.Log), gin.Recovery())
	if os.Getenv("DEV") == "true" {
		r.Use(cors.Default())
	}
	r.GET("/health", router.getRouteHealthcheck)
	r.GET("/endpoint/:value", router.getRouteFindValue)
	return r
}
