package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files from the "static" directory
	router.Static("/dashboard", "./")

	router.GET("/", Welcome)

	router.Run(":3000")
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Home Page")
}
