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
	var err error

	// Try querying the local DB first
	if dbClient.DB != nil {
		err = dbClient.DB.QueryRow("SELECT * FROM Bill_Details WHERE id = $1;", id).
			Scan(&bill.ID, &bill.BillNumber, &bill.BillDate, &bill.BillTotalAmount, &bill.SellerName, &bill.SellerPanNum, &bill.CustomerName, &bill.CustomerLocation, &bill.CustomerPhoneNumber, &bill.CustomerPanContainer, &itemsJSON, &bill.TotalDebit, &bill.TotalCredit)
		if err != nil {
			log.Printf("Error querying local DB: %v", err)
		} else {
			log.Println("Queried local DB successfully")
		}
	}

	// If local DB query fails or local DB is not connected, try AivenDB
	if err != nil && dbClient.AivenDB != nil {
		err = dbClient.AivenDB.QueryRow("SELECT * FROM Bill_Details WHERE id = $1;", id).
			Scan(&bill.ID, &bill.BillNumber, &bill.BillDate, &bill.BillTotalAmount, &bill.SellerName, &bill.SellerPanNum, &bill.CustomerName, &bill.CustomerLocation, &bill.CustomerPhoneNumber, &bill.CustomerPanContainer, &itemsJSON, &bill.TotalDebit, &bill.TotalCredit)
		if err != nil {
			log.Printf("Error querying AivenDB: %v", err)
			return bill, fmt.Errorf("error querying AivenDB: %w", err)
		}
		log.Println("Queried AivenDB successfully")
	} else if err != nil {
		// If both queries failed or no database is connected
		return bill, fmt.Errorf("no available database to query from")
	}

	// Unmarshal the JSON string into the Products field of the bill
	if err := json.Unmarshal([]byte(itemsJSON), &bill.Items); err != nil {
		log.Println("Error unmarshaling products:", err)
		return bill, fmt.Errorf("error unmarshaling products: %w", err)
	}

	return bill, nil
}

func (dbClient *DBClient) DeleteProductByID(id int) (sql.Result, error) {
	var result sql.Result
	var err error

	// Try deleting from the local DB first
	if dbClient.DB != nil {
		result, err = dbClient.DB.Exec("DELETE FROM Bill_Details WHERE id = $1;", id)
		if err != nil {
			log.Printf("Error deleting data from local DB: %v", err)
		} else {
			log.Println("Deleted data from local DB successfully")
			return result, nil
		}
	}

	// If the local DB query fails or local DB is not connected, try AivenDB
	if dbClient.AivenDB != nil {
		result, err = dbClient.AivenDB.Exec("DELETE FROM Bill_Details WHERE id = $1;", id)
		if err != nil {
			log.Printf("Error deleting data from AivenDB: %v", err)
			return nil, fmt.Errorf("error deleting data from AivenDB: %w", err)
		}
		log.Println("Deleted data from AivenDB successfully")
		return result, nil
	}

	// If both databases failed to delete
	return nil, fmt.Errorf("no available database to delete from")
}
