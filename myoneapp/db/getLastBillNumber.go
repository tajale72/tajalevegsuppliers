package db

import (
	"context"
	"fmt"
	"log"

	"github.com/lib/pq"

	"myoneapp/model"
)

func (dbClient *DBClient) GetLastBillNumber() (model.Result, error) {
	var result model.Result

	// Query to count the number of rows in the bills table
	query := "SELECT COUNT(*) FROM bill_details"

	// Try querying the local DB first
	if dbClient.DB != nil {
		err := dbClient.DB.QueryRowContext(context.Background(), query).Scan(&result.BillNumber)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "42P01" {
				// Error code 42P01 indicates "table not found or empty"
				result.BillNumber = 1
				return result, nil
			}
			return result, fmt.Errorf("failed to count records in local DB: %v", err)
		}
		log.Println("Queried local DB successfully for bill count.")
	} else if dbClient.AivenDB != nil {
		// If local DB is unavailable, try querying AivenDB
		err := dbClient.AivenDB.QueryRowContext(context.Background(), query).Scan(&result.BillNumber)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "42P01" {
				// Error code 42P01 indicates "table not found or empty"
				result.BillNumber = 1
				return result, nil
			}
			return result, fmt.Errorf("failed to count records in AivenDB: %v", err)
		}
		log.Println("Queried AivenDB successfully for bill count.")
	} else {
		// If no database connection is available
		return result, fmt.Errorf("no available database to query from")
	}

	// Increment the BillNumber by 1
	result.BillNumber += 1

	return result, nil
}
