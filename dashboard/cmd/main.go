package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Serve static files from the "static" directory
	router.Static("/dashboard", "./")

	router.Run(":3000")
}
