package db

import (
	"context"
	"fmt"
	"myoneapp/model"
)

func (dbClient *DBClient) GetLastBillNumber() (model.Result, error) {
	var result model.Result

	// Query to count the number of rows in the bills table
	query := "SELECT COUNT(*) FROM bill_details"

	// Execute the query
	err := dbClient.DB.QueryRowContext(context.Background(), query).Scan(&result.BillNumber)
	if err != nil {
		return result, fmt.Errorf("failed to count records: %v", err)
	}

	// Increment the BillNumber by 1
	result.BillNumber += 1

	return result, nil
}
