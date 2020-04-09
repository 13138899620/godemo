package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//abc
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World yyyy!! ")
	})

	111
	router.Run(":3554")
}
