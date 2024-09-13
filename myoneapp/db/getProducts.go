package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"myoneapp/model"
)

func (dbClient *DBClient) GetProducts() ([]model.Request, error) {
	var rows *sql.Rows
	var err error

	// Always try to query the local DB first
	if dbClient.DB != nil {
		rows, err = dbClient.DB.Query("SELECT * FROM Bill_Details;")
		if err != nil {
			log.Printf("Error querying local DB: %v", err)
		} else {
			log.Println("Queried local DB successfully")
		}
	}

	// If the local DB query fails or local DB is not connected, try AivenDB
	if rows == nil && dbClient.AivenDB != nil {
		rows, err = dbClient.AivenDB.Query("SELECT * FROM Bill_Details;")
		if err != nil {
			log.Printf("Error querying AivenDB: %v", err)
			return nil, fmt.Errorf("error querying AivenDB: %w", err)
		}
		log.Println("Queried AivenDB successfully")
	} else if rows == nil {
		// If both queries failed or no database is connected
		return nil, fmt.Errorf("no available database to query from")
	}

	defer rows.Close()

	var listOfBills []model.Request
	for rows.Next() {
		var bill model.Request
		var productsJSON sql.RawBytes // Use sql.RawBytes to handle JSONB data

		// Scan the columns from the result set
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
