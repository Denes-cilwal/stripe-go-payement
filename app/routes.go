package app

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {

	router := gin.Default()
	r := router.Group("/stripe")
	{
		r.POST("/charge")
	}

	return router
}
