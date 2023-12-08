package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	d "myoneapp/db"
	"myoneapp/model"
)

// Define a database interface or type from myoneapp/db package
type Database interface {
	CreateTable(data []byte) error
	GetProducts() ([]model.Request, error)
	GetProductByID(id int) (model.Request, error)
	// Add other database-related methods if needed
}

type DBClient struct {
	DB Database
}

func (dbClient *DBClient) Submit(c *gin.Context) {
	body := c.Request.Body
	data, _ := io.ReadAll(body)

	err := dbClient.DB.CreateTable(data)
	if err != nil {
		log.Println("Error CreateTable: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusAccepted, "Successfully created the bill")
}

func (dbClient *DBClient) GetProductsDetails(c *gin.Context) {
	lisofProducts, err := dbClient.DB.GetProducts()
	if err != nil {
		log.Println("Error getting products: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusAccepted, lisofProducts)
}

func (dbClient *DBClient) GetProductsDetailsByID(c *gin.Context) {
	id := c.Param("id") // Assuming you get the ID from the request URL

	// Convert id to int
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	productDetails, err := dbClient.DB.GetProductByID(productID)
	if err != nil {
		log.Println("Error getting product details: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, productDetails)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	db, _ := d.GetDBConnection()
	dbClient := &DBClient{
		DB: db, // Assuming GetDBConnection returns a Database object
	}

	r.POST("/submit", dbClient.Submit)
	r.GET("/products", dbClient.GetProductsDetails)
	r.GET("/products/:id", dbClient.GetProductsDetailsByID)

	r.Run()
}
