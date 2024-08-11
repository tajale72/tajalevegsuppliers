package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"myoneapp/model"
)

func (dbClient *DBClient) GetProducts() ([]model.Request, error) {
	rows, err := dbClient.DB.Query("SELECT * FROM Bill_Details;")
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var listOfBills []model.Request
	for rows.Next() {
		var bill model.Request
		var productsJSON sql.RawBytes // Use sql.RawBytes to handle JSONB data

		// Update the Scan call to match the number of columns
		if err := rows.Scan(
			&bill.ID,
			&bill.BillNumber,
			&bill.BillDate,
			&bill.BillTotalAmount,
			&bill.SellerName,
			&bill.SellerPanNum,
			&bill.CustomerName,
			&bill.CustomerLocation,
			&bill.CustomerPhoneNumber,
			&bill.CustomerPanContainer,
			&productsJSON,
		); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		// Unmarshal the JSONB data into the Items field of the bill
		if err := json.Unmarshal(productsJSON, &bill.Items); err != nil {
			log.Println("Error unmarshaling products:", err)
			return nil, fmt.Errorf("error unmarshaling products: %w", err)
		}

		listOfBills = append(listOfBills, bill)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return listOfBills, nil
}
