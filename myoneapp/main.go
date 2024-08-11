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
	mongoDB "myoneapp/mongo"
)

// Define a database interface or type from myoneapp/db package
type Database interface {
	CreateTable(data []byte) error
	GetProducts() ([]model.Request, error)
	GetProductByID(id int) (model.Request, error)
	GetLastBillNumber() (model.Result, error)
	// Add other database-related methods if needed
}

type DBClient struct {
	DB Database
}

func (dbClient *DBClient) Submit(c *gin.Context) {
	body := c.Request.Body
	data, _ := io.ReadAll(body)
	mongoClient, err := mongoDB.GetDBConnection()
	if err != nil {
		log.Println("Error CreateTable: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = mongoClient.CreateTable(data)
	if err != nil {
		log.Println("Error CreateTable: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = dbClient.DB.CreateTable(data)
	if err != nil {
		log.Println("Error CreateTable: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, "Successfully created the bill")
}

func (dbClient *DBClient) GetProductsDetails(c *gin.Context) {
	lisofProducts, err := dbClient.DB.GetProducts()
	if err != nil {
		log.Println("Error getting products: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	//lisofProducts, err := dbClient.DB.GetProducts()
	if err != nil {
		log.Println("Error getting products: ", err)
		c.JSON(http.StatusInternalServerError, err)
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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, productDetails)
}

func (dbClient *DBClient) GetBillNumber(c *gin.Context) {
	id, err := dbClient.DB.GetLastBillNumber()
	if err != nil {
		log.Println("Error getting product details: ", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Println("bill number", id)

	c.JSON(http.StatusOK, id)
}

type BillNumber struct {
	billNumber string `json:"billNumber"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	db, err := d.GetDBConnection()
	if err != nil {
		log.Println("Error getting products: ", err)
	}
	dbClient := &DBClient{
		DB: db, // Assuming GetDBConnection returns a Database object
	}
	r.GET("/", GetProductsDetails)

	r.POST("/submit", dbClient.Submit)
	r.GET("/products", dbClient.GetProductsDetails)
	r.GET("/products/:id", dbClient.GetProductsDetailsByID)
	r.GET("/getBillNumber", dbClient.GetBillNumber)

	r.Run()
}

func GetProductsDetails(c *gin.Context) {
	c.JSON(http.StatusAccepted, "succes from 8080")
}
