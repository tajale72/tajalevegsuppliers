package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"myoneapp/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "" // Empty string for no username
	password = "" // Empty string for no password
	dbname   = "tajalevegsuppliers"
	sslmode  = "disable"
)

type DBClient struct {
	DB *sql.DB
}

func GetDBConnection() (*DBClient, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", user, password, host, port, dbname, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres db: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &DBClient{DB: db}, nil
}

func (dbClient *DBClient) Close() {
	if dbClient.DB != nil {
		dbClient.DB.Close()
	}
}

func (dbClient *DBClient) CreateTable(data []byte) error {
	// Unmarshal JSON data into the request model
	var req model.Request
	if err := json.Unmarshal(data, &req); err != nil {
		log.Println("Unmarshal error:", err)
		return fmt.Errorf("error unmarshaling request: %w", err)
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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9, $10::jsonb);
`

	fmt.Println("bill amount", req.BillTotalAmount)

	// Execute the SQL query with specific values
	_, err = dbClient.DB.Exec(sqlStatement, req.BillNumber, req.BillDate, req.BillTotalAmount, req.SellerName, req.SellerPanNum, req.CustomerName, req.CustomerLocation, req.CustomerPhoneNumber, req.CustomerPanContainer, itemsJSON)
	if err != nil {
		log.Println("Exec statement error:", err)
		return fmt.Errorf("error executing SQL statement: %w", err)
	}

	return nil
}

func (dbClient *DBClient) GetProducts() ([]model.Request, error) {
	rows, err := dbClient.DB.Query("SELECT * FROM Bill_Details;")
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var listOfBills []model.Request
	for rows.Next() {
		var bill model.Request
		var productsJSON sql.RawBytes // Use sql.RawBytes to handle JSONB data

		// Update the Scan call to match the number of columns
		if err := rows.Scan(
			&bill.ID,
			&bill.BillNumber,
			&bill.BillDate,
			&bill.BillTotalAmount,
			&bill.SellerName,
			&bill.SellerPanNum,
			&bill.CustomerName,
			&bill.CustomerLocation,
			&bill.CustomerPhoneNumber,
			&bill.CustomerPanContainer,
			&productsJSON,
		); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		// Unmarshal the JSONB data into the Items field of the bill
		if err := json.Unmarshal(productsJSON, &bill.Items); err != nil {
			log.Println("Error unmarshaling products:", err)
			return nil, fmt.Errorf("error unmarshaling products: %w", err)
		}

		listOfBills = append(listOfBills, bill)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return listOfBills, nil
}

func (dbClient *DBClient) GetProductByID(id int) (model.Request, error) {
	var bill model.Request
	var itemsJSON string

	// Use a pointer for itemsJSON
	err := dbClient.DB.QueryRow("SELECT * FROM Bill_Details WHERE id = $1;", id).
		Scan(&bill.ID, &bill.BillNumber, &bill.BillDate, &bill.BillTotalAmount, &bill.SellerName, &bill.SellerPanNum, &bill.CustomerName, &bill.CustomerLocation, &bill.CustomerPhoneNumber, &bill.CustomerPanContainer, &itemsJSON)
	if err != nil {
		log.Println("Error getting data from the database:", err)
		return bill, fmt.Errorf("Error getting data from the database: %w", err)
	}

	// Unmarshal the JSON string into the Products field of the bill
	if err := json.Unmarshal([]byte(itemsJSON), &bill.Items); err != nil {
		log.Println("Error unmarshaling products:", err)
		return bill, fmt.Errorf("error unmarshaling products: %w", err)
	}

	return bill, nil
}

func (dbClient *DBClient) UpdateUser(id int, newUsername, newEmail string) error {
	_, err := dbClient.DB.Exec("UPDATE users SET username=$1, email=$2 WHERE id=$3", newUsername, newEmail, id)
	if err != nil {
		log.Println("Error updating user:", err)
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (dbClient *DBClient) DeleteUser(id int) error {
	_, err := dbClient.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		log.Println("Error deleting user:", err)
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

type User struct {
	ID       int
	Username string
	Email    string
}
