package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"myoneapp/db"
	d "myoneapp/db"
	"myoneapp/imageprocessing"
	"myoneapp/model"
)

type DBClient struct {
	DB db.Database
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	// Serve static files from the "static" directory
	r.Static("/myoneapp/static", "./static")

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

	r.GET("/customer", dbClient.GetCustomers)
	r.GET("/ledger:customerName", dbClient.GetCustomers)

	r.POST("/upload", UploadImageHandler)

	r.GET("/dashboard", dbClient.GetDashBoard)

	r.Run()
}

func (dbClient *DBClient) GetDashBoard(c *gin.Context) {

	dashboardData, err := dbClient.DB.GetDashBoardData()
	if err != nil {
		log.Println("Error getting dashboard data: ", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Send JSON response
	c.JSON(http.StatusOK, dashboardData)
}

func UploadImageHandler(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Error retrieving the file")
		return
	}

	// Example: Save the file to a temporary location
	tempFilePath := "/tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		c.String(http.StatusInternalServerError, "Error saving the file")
		return
	}

	imageprocessing.SendImageToPythonService(tempFilePath)

	// Send a success response with the file location
	c.String(http.StatusOK, fmt.Sprintf("File uploaded successfully: %s", tempFilePath))
}

func (dbClient *DBClient) GetLedgerEntries(c *gin.Context) {
	// Get the customerName from the query parameter
	searchQuery := c.Query("customerName")

	// Trim leading and trailing spaces
	searchQuery = strings.TrimSpace(searchQuery)

	var lisofLedgerEntries []model.LedgerEntry
	var err error
	// Check if a search query is provided
	if searchQuery != "" {
		// If a search query is provided, get filtered products based on the search term
		lisofLedgerEntries, err = dbClient.DB.GetCustomerLedgerByName(searchQuery)
		if err != nil {
			log.Println("Error getting lisofLedgerEntries by name: ", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	} else {
		// If no search query is provided, get all products
		lisofLedgerEntries, err = dbClient.DB.GetLedgerEntries()
		if err != nil {
			log.Println("Error getting lisofLedgerEntries: ", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusAccepted, lisofLedgerEntries)
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

	// Retrieve the 'search' parameter from the query string
	searchQuery := c.Query("search")

	fmt.Printf("searchq", searchQuery)

	var listOfProducts []model.Request
	var err error

	// Check if a search query is provided
	if searchQuery != "" {
		// If a search query is provided, get filtered products based on the search term
		listOfProducts, err = dbClient.DB.GetProductsBySearch(searchQuery)
	} else {
		// If no search query is provided, get all products
		listOfProducts, err = dbClient.DB.GetProducts()
	}

	if err != nil {
		log.Println("Error getting products: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	// Calculate the final total amount
	var finalTotalAmount float64
	for _, product := range listOfProducts {
		amount, err := strconv.ParseFloat(product.BillTotalAmount, 64)
		if err == nil {
			finalTotalAmount += amount
		}
	}
	if len(listOfProducts) > 0 {
		listOfProducts[0].FinalTotalAmount = finalTotalAmount
	}

	c.JSON(http.StatusOK, listOfProducts)
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

func (dbClient *DBClient) GetCustomers(c *gin.Context) {
	lisofCustomers, err := dbClient.DB.GetCustomers()
	if err != nil {
		log.Println("Error getting products: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, lisofCustomers)
}

func (dbClient *DBClient) GetCustomerLedgerByName(c *gin.Context) {
	// Get the customerName from the query parameter
	customerName := c.Query("customerName")

	// Trim leading and trailing spaces
	customerName = strings.TrimSpace(customerName)
	if customerName == "" {
		log.Println("Customer name is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer name cannot be empty"})
		return
	}

	fmt.Println("Customer Name:", customerName)

	// Query the customer ledger using the trimmed customerName
	customerLedger, err := dbClient.DB.GetCustomerLedgerByName(customerName)
	if err != nil {
		log.Println("Error getting customer ledger: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customerLedger)
}
