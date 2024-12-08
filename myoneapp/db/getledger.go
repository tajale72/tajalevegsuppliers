package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"

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
func (dbClient *DBClient) GetDashBoardData() (*model.Dashboard, error) {
	// Check if DB connection exists
	if dbClient.DB == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	// SQL query
	query := `
	WITH total_sales_and_expenses AS (
		SELECT 
			SUM(debit) AS profit -- Total sum of debit as profit
		FROM 
			ledger_entries
	),
	total_sales AS (
		SELECT 
			SUM(CAST(bill_total_amount AS DECIMAL)) AS total_amount
		FROM 
			bill_details
	),
	top_selling_vegetables AS (
		SELECT 
			vegetable_name
		FROM 
			dailyvegetablesales
		GROUP BY 
			vegetable_name
		ORDER BY 
			SUM(quantity_sold) DESC
		LIMIT 3
	)
	SELECT 
		tse.profit,
		fta.total_amount AS total_sales,
		ARRAY(
			SELECT vegetable_name
			FROM top_selling_vegetables
		) AS top_selling_vegetables
	FROM 
		total_sales_and_expenses tse
	CROSS JOIN 
		total_sales fta;
	`

	// Execute the query
	rows, err := dbClient.DB.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	// Initialize the Dashboard struct
	var dashboardData model.Dashboard

	// Fetch the result
	if rows.Next() {
		var topSellingVegetables []string
		if err := rows.Scan(
			&dashboardData.Profit,
			&dashboardData.TotalSales,
			pq.Array(&topSellingVegetables),
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		// Assign the array to the struct
		dashboardData.TopSellingVegetables = topSellingVegetables
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return &dashboardData, nil
}
