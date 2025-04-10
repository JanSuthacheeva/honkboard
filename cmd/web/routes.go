package main


import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("ui/html/pages/*")

	router.GET("/", app.home)

	return router
}

