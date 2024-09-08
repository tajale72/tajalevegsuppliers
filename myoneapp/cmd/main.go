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

	r.POST("/billtransaction", dbClient.Submit)
	r.GET("/billtransaction", dbClient.GetProductsDetails)
	r.GET("/billtransaction/:id", dbClient.GetProductsDetailsByID)
	r.DELETE("/billtransaction/:id", dbClient.DeleteProductsDetailsByID)
	r.PUT("billtransaction/:id", dbClient.UpdateBill)

	r.POST("/ledger", dbClient.Ledger)
	r.GET("/ledger", dbClient.GetLedgerEntries)

	r.GET("/getBillNumber", dbClient.GetBillNumber)
	r.GET("/vegetablecount", dbClient.GetVegetableCount)

	r.Run()
}

func (dbClient *DBClient) GetLedgerEntries(c *gin.Context) {
	lisofProducts, err := dbClient.DB.GetLedgerEntries()
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

func (dbClient *DBClient) Ledger(c *gin.Context) {
	body := c.Request.Body
	data, _ := io.ReadAll(body)

	err := dbClient.DB.CreateLedger(data)
	if err != nil {
		log.Println("Error CreateTable: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, "Successfully created the bill")
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

func (dbClient *DBClient) DeleteProductsDetailsByID(c *gin.Context) {
	id := c.Param("id") // Assuming you get the ID from the request URL

	// Convert id to int
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	deleteResponse, err := dbClient.DB.DeleteProductByID(productID)
	if err != nil {
		log.Println("Error getting product details: ", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, deleteResponse)
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

func GetProductsDetails(c *gin.Context) {
	c.JSON(http.StatusAccepted, "succes from 8080")
}
