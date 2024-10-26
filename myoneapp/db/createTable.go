package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"myoneapp/model"
)

func (dbClient *DBClient) CreateTable(data []byte) error {

	req, err := ValidateRequest(data)
	if err != nil {
		log.Println("error validating the feild:", err)
		return fmt.Errorf("error marshaling items: %w", err)
	}

	// Marshal items slice to JSON
	itemsJSON, err := json.Marshal(req.Items)
	if err != nil {
		log.Println("Marshal error:", err)
		return fmt.Errorf("error marshaling items: %w", err)
	}

	// SQL statement with placeholders
	sqlStatement := `
	INSERT INTO Bill_Details (
    bill_number, 
    bill_date,
    bill_total_amount, 
    total_debit,
    total_credit,
    seller_name, 
    seller_pan_num, 
    customer_name, 
    customer_location, 
    customer_phone_number, 
    customer_pan_container, 
    items
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12::jsonb);`

	// Execute the SQL query with specific values
	_, err = dbClient.DB.Exec(sqlStatement,
		req.BillNumber,
		req.BillDate,
		req.BillTotalAmount,
		req.TotalDebit,  // Add total_debit
		req.TotalCredit, // Add total_credit
		req.SellerName,
		req.SellerPanNum,
		req.CustomerName,
		req.CustomerLocation,
		req.CustomerPhoneNumber,
		req.CustomerPanContainer,
		itemsJSON)
	if err != nil {
		log.Println("Exec statement error:", err)
		return fmt.Errorf("error executing SQL statement: %w", err)
	}

	if dbClient.AivenDB != nil {
		// Execute the SQL query with specific values
		_, err = dbClient.AivenDB.Exec(sqlStatement,
			req.BillNumber,
			req.BillDate,
			req.BillTotalAmount,
			req.TotalDebit,  // Add total_debit
			req.TotalCredit, // Add total_credit
			req.SellerName,
			req.SellerPanNum,
			req.CustomerName,
			req.CustomerLocation,
			req.CustomerPhoneNumber,
			req.CustomerPanContainer,
			itemsJSON)
		if err != nil {
			log.Println("Exec statement error:", err)
			return fmt.Errorf("error executing SQL statement: %w", err)
		}
	}

	dbClient.CalucaltedVegTableQuantitySold(itemsJSON, req.BillNumber)

	return nil
}

func (dbClient *DBClient) CalucaltedVegTableQuantitySold(data []byte, billNumber string) error {
	log.Println("items", string(data))

	var items []model.Item
	err := json.Unmarshal(data, &items)
	if err != nil {
		log.Println("Error unmarshalling the items:", err)
		return fmt.Errorf("error unmarshalling items: %w", err)
	}

	// Map to aggregate items by VegetableName
	aggregatedItems := make(map[string]model.Item)
	for _, item := range items {
		if item.VegetableName == "" {
			log.Println("Skipping item due to missing vegetable name:", item)
			continue
		}

		// Convert QuantitySold, Rate, and Amount from string to float64
		quantitySold, err := strconv.ParseFloat(item.QuantitySold, 64)
		if err != nil {
			log.Printf("Invalid quantity_sold value: %v. Using 0 instead.", item.QuantitySold)
			quantitySold = 0.00
		}

		rate, err := strconv.ParseFloat(item.Rate, 64)
		if err != nil {
			log.Printf("Invalid rate value: %v. Using 0.00 instead.", item.Rate)
			rate = 0.00
		}

		amount, err := strconv.ParseFloat(item.Amount, 64)
		if err != nil {
			log.Printf("Invalid amount value: %v. Using 0.00 instead.", item.Amount)
			amount = 0.00
		}

		// Aggregate by VegetableName
		if existingItem, found := aggregatedItems[item.VegetableName]; found {
			// Sum the quantities and amounts
			existingQuantity, _ := strconv.ParseFloat(existingItem.QuantitySold, 64)
			existingAmount, _ := strconv.ParseFloat(existingItem.Amount, 64)

			aggregatedItems[item.VegetableName] = model.Item{
				VegetableName: item.VegetableName,
				QuantitySold:  fmt.Sprintf("%.2f", existingQuantity+quantitySold),
				Rate:          fmt.Sprintf("%.2f", rate), // Use the latest rate
				Amount:        fmt.Sprintf("%.2f", existingAmount+amount),
				CreatedAt:     time.Now(),
			}
		} else {
			// First time adding this item
			aggregatedItems[item.VegetableName] = model.Item{
				VegetableName: item.VegetableName,
				QuantitySold:  fmt.Sprintf("%.2f", quantitySold),
				Rate:          fmt.Sprintf("%.2f", rate),
				Amount:        fmt.Sprintf("%.2f", amount),
				CreatedAt:     time.Now(),
			}
		}
	}

	query := `
        INSERT INTO dailyvegetablesales (
            vegetable_name, 
            sale_date, 
            quantity_sold, 
            rate, 
            total_amount, 
            created_at
        ) VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (vegetable_name, sale_date) 
        DO UPDATE SET 
            quantity_sold = dailyvegetablesales.quantity_sold + EXCLUDED.quantity_sold,
            rate = EXCLUDED.rate,
            total_amount = dailyvegetablesales.total_amount + EXCLUDED.total_amount,
            created_at = EXCLUDED.created_at;`

	// Function to execute the query on a given DB
	executeQuery := func(db *sql.DB) error {
		for _, v := range aggregatedItems {
			// Log the values before executing the query for debugging
			log.Printf("Inserting/updating vegetable: %v, QuantitySold: %v, Rate: %v, Amount: %v", v.VegetableName, v.QuantitySold, v.Rate, v.Amount)

			quantitySold, _ := strconv.ParseFloat(v.QuantitySold, 64)
			rate, _ := strconv.ParseFloat(v.Rate, 64)
			amount, _ := strconv.ParseFloat(v.Amount, 64)

			_, err = db.Exec(
				query,
				v.VegetableName,
				time.Now(),
				quantitySold,
				rate,
				amount,
				time.Now(),
			)
			if err != nil {
				log.Println("Error updating item in the database:", err)
				return fmt.Errorf("error updating item in the database: %w", err)
			}
		}
		return nil
	}

	// Execute the query on the local DB
	if err := executeQuery(dbClient.DB); err != nil {
		return err
	}

	// Execute the query on the Aiven DB if available
	if dbClient.AivenDB != nil {
		if err := executeQuery(dbClient.AivenDB); err != nil {
			log.Println("Error updating item in the Aiven database:", err)
			return fmt.Errorf("error updating item in the Aiven database: %w", err)
		}
	}

	return nil
}
