package db

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	// replace with your actual package path

	"myoneapp/model"
)

func (dbClient *DBClient) CreateLedger(data []byte) error {
	// Log incoming data
	fmt.Println("Incoming JSON Data:", string(data))

	var ledgerData model.LedgerEntry

	// Unmarshal the incoming JSON data into the LedgerEntry struct
	err := json.Unmarshal(data, &ledgerData)
	if err != nil {
		log.Println("Error unmarshalling data:", err)
		return fmt.Errorf("error parsing input data: %w", err)
	}

	// Parse the date from string to time.Time for SQL (including time if needed)
	date, err := time.Parse("2006-01-02", ledgerData.Date)
	if err != nil {
		log.Println("Error parsing date:", err)
		return fmt.Errorf("error parsing date: %w", err)
	}

	// Log the parsed ledgerData
	fmt.Println("Parsed Ledger Data:", ledgerData)

	// SQL statement for inserting into ledger_entries
	sqlStatement := `
	INSERT INTO ledger_entries (
		date,
		account,
		billNumber,
		debit,
		credit,
		balance_amount
		)
		VALUES ($1, $2, $3, $4, $5, $6);`

	// Execute the SQL query with values from ledgerData
	_, err = dbClient.DB.Exec(sqlStatement,
		date,
		ledgerData.Account,
		ledgerData.BillNumber,
		ledgerData.Debit,
		ledgerData.Credit,
		ledgerData.BalanceAmount)
	if err != nil {
		log.Println("Exec statement error:", err)
		return fmt.Errorf("error executing SQL statement: %w", err)
	}

	if dbClient.AivenDB != nil {
		// Execute the SQL query with values from ledgerData
		_, err = dbClient.AivenDB.Exec(sqlStatement,
			date,
			ledgerData.Account,
			ledgerData.BillNumber,
			ledgerData.Debit,
			ledgerData.Credit,
			ledgerData.BalanceAmount)
		if err != nil {
			log.Println("Exec statement error:", err)
			return fmt.Errorf("error executing SQL statement: %w", err)
		}
	}

	return nil
}
