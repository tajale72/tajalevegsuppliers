package db

import (
	"database/sql"
	"fmt"
	"log"

	"myoneapp/model"
)

func (dbClient *DBClient) GetCustomers() ([]model.Customer, error) {

	//dbClient.updateVegetableData()

	var rows *sql.Rows
	var err error

	// Try querying AivenDB first if it's connected
	if dbClient.AivenDB != nil {
		rows, err = dbClient.AivenDB.Query("SELECT * FROM customers ORDER BY customer_name ASC;")
		if err != nil {
			log.Printf("Error querying AivenDB: %v", err)
		} else {
			log.Println("Queried AivenDB successfully")
		}
	}

	// If AivenDB is not available or the query failed, try the local DB
	if rows == nil && dbClient.DB != nil {
		rows, err = dbClient.DB.Query("SELECT * FROM customers ORDER BY customer_name ASC;")
		if err != nil {
			log.Printf("Error querying local DB: %v", err)
			return nil, fmt.Errorf("error querying local DB: %w", err)
		}
		log.Println("Queried local DB successfully")
	} else if rows == nil {
		// If both queries failed
		return nil, fmt.Errorf("no available database to query from")
	}

	defer rows.Close()

	var lisofCustomers []model.Customer
	for rows.Next() {
		var list model.Customer

		if err := rows.Scan(
			&list.ID,
			&list.CustomerName,
			&list.CustomerLocation,
			&list.CustomerPhoneNumber,
			&list.CustomerPanContainer,
			&list.CreatedAt,
		); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		lisofCustomers = append(lisofCustomers, list)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return lisofCustomers, nil
}
