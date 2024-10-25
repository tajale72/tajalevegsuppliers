package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"dashboard/cmd/add"
)

func main() {
	// Call the function to add numbers from the file
	total, err := add.AddNumbersFromFile("add.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Print the total
	fmt.Printf("Total: %.2f\n", total)

	router := gin.Default()

	// Serve static files from the "static" directory
	router.Static("/dashboard/static", "./static")

	router.GET("/", Welcome)

	router.Run(":3000")
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Home Page")
}
