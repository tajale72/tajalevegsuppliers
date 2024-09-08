package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"myoneapp/model"
)

func (dbClient *DBClient) GetProductByID(id int) (model.Request, error) {
	var bill model.Request
	var itemsJSON string

	// Use a pointer for itemsJSON
	err := dbClient.DB.QueryRow("SELECT * FROM Bill_Details WHERE id = $1;", id).
		Scan(&bill.ID, &bill.BillNumber, &bill.BillDate, &bill.BillTotalAmount, &bill.SellerName, &bill.SellerPanNum, &bill.CustomerName, &bill.CustomerLocation, &bill.CustomerPhoneNumber, &bill.CustomerPanContainer, &itemsJSON)
	if err != nil {
		log.Println("Error getting data from the database:", err)
		return bill, fmt.Errorf("Error getting data from the database: %w", err)
	}

	// Unmarshal the JSON string into the Products field of the bill
	if err := json.Unmarshal([]byte(itemsJSON), &bill.Items); err != nil {
		log.Println("Error unmarshaling products:", err)
		return bill, fmt.Errorf("error unmarshaling products: %w", err)
	}

	return bill, nil
}

// DeleteProductByID deletes a product by its ID and returns the deleted product's details.
func (dbClient *DBClient) DeleteProductByID(id int) (sql.Result, error) {

	// Delete the bill from the database
	result, err := dbClient.DB.Exec("DELETE FROM Bill_Details WHERE id = $1;", id)
	if err != nil {
		log.Println("Error deleting data from the database:", err)
		return nil, fmt.Errorf("error deleting data from the database: %w", err)
	}

	return result, nil
}