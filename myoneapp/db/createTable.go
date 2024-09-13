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
		seller_name, 
		seller_pan_num, 
		customer_name, 
		customer_location, 
		customer_phone_number, 
		customer_pan_container, 
		items
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9, $10::jsonb);`

	// Execute the SQL query with specific values
	_, err = dbClient.DB.Exec(sqlStatement,
		req.BillNumber,
		req.BillDate,
		req.BillTotalAmount,
		req.SellerName,
		req.SellerPanNum,
		req.CustomerName,
		req.CustomerLocation,
		req.CustomerPhoneNumber,
		req.CustomerPanContainer, itemsJSON)
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
			req.SellerName,
			req.SellerPanNum,
			req.CustomerName,
			req.CustomerLocation,
			req.CustomerPhoneNumber,
			req.CustomerPanContainer, itemsJSON)
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
		for _, v := range items {
			if v.VegetableName == "" {
				log.Println("Skipping item due to missing vegetable name:", v)
				continue
			}

			createdAt := time.Now()

			// Convert QuantitySold, Rate, and Amount from string to float64
			quantitySold, err := strconv.ParseFloat(v.QuantitySold, 64)
			if err != nil {
				log.Printf("Invalid quantity_sold value: %v. Using 0 instead.", v.QuantitySold)
				quantitySold = 0.00
			}

			rate, err := strconv.ParseFloat(v.Rate, 64)
			if err != nil {
				log.Printf("Invalid rate value: %v. Using 0.00 instead.", v.Rate)
				rate = 0.00
			}

			amount, err := strconv.ParseFloat(v.Amount, 64)
			if err != nil {
				log.Printf("Invalid amount value: %v. Using 0.00 instead.", v.Amount)
				amount = 0.00
			}

			// Log the values before executing the query for debugging
			log.Printf("Inserting/updating vegetable: %v, QuantitySold: %v, Rate: %v, Amount: %v", v.VegetableName, quantitySold, rate, amount)

			_, err = db.Exec(
				query,
				v.VegetableName,
				createdAt,
				quantitySold,
				rate,
				amount,
				createdAt,
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

// func (dbClient *DBClient) CreateCustomer(req model.Request) error {
// 	_, err := dbClient.DB.Exec(
// 		query,
// 		v.VegetableName,
// 		createdAt,
// 		quantitySold,
// 		rate,
// 		amount,
// 		createdAt,
// 	)
// 	if err != nil {
// 		log.Println("Error updating item in the database:", err)
// 		return fmt.Errorf("error updating item in the database: %w", err)
// 	}
// }
