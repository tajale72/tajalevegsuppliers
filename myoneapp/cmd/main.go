package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"myoneapp/db"
	d "myoneapp/db"
)

type DBClient struct {
	DB db.Database
}

func (dbClient *DBClient) Submit(c *gin.Context) {
	body := c.Request.Body
	data, _ := io.ReadAll(body)
	// mongoClient, err := mongoDB.GetDBConnection()
	// if err != nil {
	// 	log.Println("Error CreateTable: ", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	return
	// }

	// err = mongoClient.CreateTable(data)
	// if err != nil {
	// 	log.Println("Error CreateTable: ", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	err := dbClient.DB.CreateTable(data)
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

	// //lisofProducts, err := dbClient.DB.GetProducts()
	// if err != nil {
	// 	log.Println("Error getting products: ", err)
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

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

func (dbClient *DBClient) UpdateBill(c *gin.Context) {
	fmt.Println("UpdateBill....")
	body := c.Request.Body
	data, _ := io.ReadAll(body)
	err := dbClient.DB.UpdateBill(data)
	if err != nil {
		log.Println("Error getting product details: ", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "bill updated")
}

func (dbClient *DBClient) GetVegetableCount(c *gin.Context) {
	lisofVegtableCount, err := dbClient.DB.GetVegetableCount()
	if err != nil {
		log.Println("Error getting products: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// //lisofProducts, err := dbClient.DB.GetProducts()
	// if err != nil {
	// 	log.Println("Error getting products: ", err)
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	c.JSON(http.StatusAccepted, lisofVegtableCount)
}

type BillNumber struct {
	billNumber string `json:"billNumber"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	// Middleware to log requests
	r.Use(func(c *gin.Context) {
		ip := c.ClientIP()                                                 // Get the client's IP address
		userAgent := c.Request.UserAgent()                                 // Get the user agent
		log.Printf("Request from IP: %s, User-Agent: %s\n", ip, userAgent) // Log the details

		c.Next() // Call the next handler
	})

	db, err := d.GetDBConnection()
	if err != nil {
		log.Println("Error getting products: ", err)
	}
	dbClient := &DBClient{
		DB: db, // Assuming GetDBConnection returns a Database object
	}
	// Serve static files from the "static" directory
	r.Static("/myoneapp/static", "./static")

	r.GET("/", GetProductsDetails)

	r.POST("/submit", dbClient.Submit)
	r.GET("/products", dbClient.GetProductsDetails)
	r.GET("/products/:id", dbClient.GetProductsDetailsByID)
	r.PUT("updateBill/:id", dbClient.UpdateBill)
	r.GET("/getBillNumber", dbClient.GetBillNumber)
	r.GET("/vegetablecount", dbClient.GetVegetableCount)

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080

}

func GetProductsDetails(c *gin.Context) {
	c.JSON(http.StatusAccepted, "succes from 8080")
}
