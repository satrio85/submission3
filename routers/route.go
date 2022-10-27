package routers

import (
	"submission-3/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("views/*.html")
	router.GET("/", controllers.GetMain)
	// router.GET("data")

	return router
}
