package db

import (
	"encoding/json"
	"fmt"
	"log"
)

func (dbClient *DBClient) CreateTable(data []byte) error {
	req, err := ValidateRequest(data)
	if err != nil {
		log.Println("error validating the feild:", err)
		return fmt.Errorf("error marshaling items: %w", err)
	}

	// Marshal items slice to JSON
	itemsJSON, err := json.Marshal(req.Items)
	if err != nil {
		log.Println("Marshal error:", err)
		return fmt.Errorf("error marshaling items: %w", err)
	}

	// SQL statement with placeholders
	sqlStatement := `
INSERT INTO Bill_Details (
	bill_number, 
	bill_date,
	bill_total_amount, 
	seller_name, 
	seller_pan_num, 
	customer_name, 
	customer_location, 
	customer_phone_number, 
	customer_pan_container, 
	items
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9, $10::jsonb);
`

	fmt.Println("bill amount", req.BillTotalAmount)

	// Execute the SQL query with specific values
	_, err = dbClient.DB.Exec(sqlStatement, req.BillNumber, req.BillDate, req.BillTotalAmount, req.SellerName, req.SellerPanNum, req.CustomerName, req.CustomerLocation, req.CustomerPhoneNumber, req.CustomerPanContainer, itemsJSON)
	if err != nil {
		log.Println("Exec statement error:", err)
		return fmt.Errorf("error executing SQL statement: %w", err)
	}

	return nil
}
