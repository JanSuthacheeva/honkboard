package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func (app *application) home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Hello",
	})
}
