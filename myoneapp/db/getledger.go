package db

import (
	"database/sql"
	"fmt"
	"log"

	"myoneapp/model"
)

// GetLedgerEntries retrieves all ledger entries from the database.
func (dbClient *DBClient) GetLedgerEntries() ([]model.LedgerEntry, error) {
	var rows *sql.Rows
	var err error

	// Try querying the local DB first
	if dbClient.DB != nil {
		rows, err = dbClient.DB.Query("SELECT * FROM ledger_entries;")
		if err != nil {
			log.Printf("Error querying local DB: %v", err)
		} else {
			log.Println("Queried local DB successfully")
		}
	}

	// If local DB query fails or local DB is not connected, try AivenDB
	if rows == nil && dbClient.AivenDB != nil {
		rows, err = dbClient.AivenDB.Query("SELECT * FROM ledger_entries;")
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

	fmt.Println("listOfEntries", listOfEntries)
	return listOfEntries, nil
}
