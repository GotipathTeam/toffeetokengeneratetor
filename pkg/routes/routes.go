package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gotipath/pkg/controllers"
)

func BaseRouteSetup() {
	router := gin.Default()

	v1 := router.Group("/api")

	v1.POST("/", controllers.GetToken)

	router.Run(":8080")

}
