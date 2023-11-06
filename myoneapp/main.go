package main

import (
	"fmt"
	"io"
	"log"
	d "myoneapp/db"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/submit", Submit)
	r.GET("/products", GetProductsDetails)
	r.GET("/products/:id", GetProductsDetailsByID)
	r.Run()
}

func Submit(c *gin.Context) {
	body := c.Request.Body
	data, _ := io.ReadAll(body)

	db, err := d.GetDBConnection()
	if err != nil {
		log.Println("Error connecting to the database: ", err)
	}
	defer db.Close()

	err = d.CreateTable(db, data)
	if err != nil {
		log.Println("Error CreateTable: ", err)
	}
	c.JSON(http.StatusAccepted, "succefully created the bill")
}

func GetProductsDetails(c *gin.Context) {

	db, err := d.GetDBConnection()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	lisofProducts, err := d.GetProducts(db)
	if err != nil {
		log.Println("Error getting products: ", err)
	}

	fmt.Println("lisofProducts:", lisofProducts)
	c.JSON(http.StatusAccepted, lisofProducts)
}

func GetProductsDetailsByID(c *gin.Context) {
	id := c.Param("id") // Assuming you get the ID from the request URL

	// Convert id to int
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	db, err := d.GetDBConnection()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	productDetails, err := d.GetProductByID(db, productID)
	if err != nil {
		log.Println("Error getting product details: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, productDetails)
}
