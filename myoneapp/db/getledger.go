package db

import (
	"fmt"
	"log"

	"myoneapp/model"
)

// GetLedgerEntries retrieves all ledger entries from the database.
func (dbClient *DBClient) GetLedgerEntries() ([]model.LedgerEntry, error) {
	rows, err := dbClient.DB.Query("SELECT * FROM ledger_entries;")
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, fmt.Errorf("error querying database: %w", err)
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
