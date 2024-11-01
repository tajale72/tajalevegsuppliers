package db

import (
	"database/sql"
	"fmt"
	"log"

	"myoneapp/model"
)

func (dbClient *DBClient) GetCustomerLedgerByName(account string) ([]model.LedgerEntry, error) {
	var rows *sql.Rows
	var err error

	// Try querying the local DB first
	if dbClient.DB != nil {
		// Add the % wildcard to the search term for partial matching
		account = "%" + account + "%"

		// Use the modified account in the query
		query := `SELECT * FROM ledger_entries WHERE account ILIKE $1`
		rows, err := dbClient.DB.Query(query, account)
		if err != nil {
			log.Printf("Error querying local DB: %v", err)
		} else {
			log.Println("Queried local DB successfully")
			defer rows.Close() // Ensure rows are closed if query is successful
			return scanLedgerEntries(rows)
		}
	}

	// If local DB query fails or local DB is not connected, try AivenDB
	if dbClient.AivenDB != nil {
		rows, err = dbClient.AivenDB.Query(`SELECT * FROM ledger_entries WHERE account = $1;`, account)
		if err != nil {
			log.Printf("Error querying AivenDB: %v", err)
			return nil, fmt.Errorf("error querying AivenDB: %w", err)
		}
		log.Println("Queried AivenDB successfully")
		defer rows.Close() // Ensure rows are closed if query is successful
		return scanLedgerEntries(rows)
	}

	// If both queries failed or no database is connected
	return nil, fmt.Errorf("no available database to query from")
}

// Helper function to scan rows into a list of LedgerEntry
func scanLedgerEntries(rows *sql.Rows) ([]model.LedgerEntry, error) {
	var listOfEntries []model.LedgerEntry

	for rows.Next() {
		var entry model.LedgerEntry

		// Update the Scan call to match the number of columns in ledger_entries
		if err := rows.Scan(
			&entry.ID,
			&entry.Date,
			&entry.Account,
			&entry.BillNumber,
			&entry.Debit,
			&entry.Credit,
			&entry.BalanceAmount,
			&entry.CreatedAt,
			&entry.CustomerID,
		); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		listOfEntries = append(listOfEntries, entry)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return listOfEntries, nil
}

// func (dbClient *DBClient) DeleteProductByID(id int) (sql.Result, error) {
// 	var result sql.Result
// 	var err error

// 	// Try deleting from the local DB first
// 	if dbClient.DB != nil {
// 		result, err = dbClient.DB.Exec("DELETE FROM Bill_Details WHERE id = $1;", id)
// 		if err != nil {
// 			log.Printf("Error deleting data from local DB: %v", err)
// 		} else {
// 			log.Println("Deleted data from local DB successfully")
// 			return result, nil
// 		}
// 	}

// 	// If the local DB query fails or local DB is not connected, try AivenDB
// 	if dbClient.AivenDB != nil {
// 		result, err = dbClient.AivenDB.Exec("DELETE FROM Bill_Details WHERE id = $1;", id)
// 		if err != nil {
// 			log.Printf("Error deleting data from AivenDB: %v", err)
// 			return nil, fmt.Errorf("error deleting data from AivenDB: %w", err)
// 		}
// 		log.Println("Deleted data from AivenDB successfully")
// 		return result, nil
// 	}

// 	// If both databases failed to delete
// 	return nil, fmt.Errorf("no available database to delete from")
// }
