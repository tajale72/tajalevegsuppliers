package db

import (
	"encoding/json"
	"fmt"
	"log"

	"myoneapp/model"
)

func (dbClient *DBClient) UpdateBill(data []byte) error {
	var bill model.Request
	err := json.Unmarshal(data, &bill)
	if err != nil {
		log.Println("Error unmarshalling the bill data:", err)
		return fmt.Errorf("error unmarshalling the bill data: %w", err)
	}
	bill.SellerName = "तजले भेज सप्लायर्स"
	bill.SellerPanNum = "६०१०८६४८९"
	// Marshal the bill.Items into a JSON string
	itemsJSON, err := json.Marshal(bill.Items)
	if err != nil {
		log.Println("Error marshaling products:", err)
		return fmt.Errorf("error marshaling products: %w", err)
	}

	// Prepare the SQL update statement
	query := `
		UPDATE Bill_Details
		SET 
			bill_number = $1,
			bill_date = $2,
			bill_total_amount = $3,
			seller_name = $4,
			seller_pan_num = $5,
			customer_name = $6,
			customer_location = $7,
			customer_phone_number = $8,
			customer_pan_container = $9,
			items = $10
		WHERE bill_number = $11;
	`

	// Execute the update query
	_, err = dbClient.DB.Exec(
		query,
		bill.BillNumber,
		bill.BillDate,
		bill.BillTotalAmount,
		bill.SellerName,
		bill.SellerPanNum,
		bill.CustomerName,
		bill.CustomerLocation,
		bill.CustomerPhoneNumber,
		bill.CustomerPanContainer,
		string(itemsJSON), // Convert JSON to string for storage
		bill.BillNumber,
	)

	if err != nil {
		log.Println("Error updating bill in the database:", err)
		return fmt.Errorf("error updating bill in the database: %w", err)
	}

	return nil
}
