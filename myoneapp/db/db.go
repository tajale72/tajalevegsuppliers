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
	var req model.Request
	if err := json.Unmarshal(data, &req); err != nil {
		log.Println("Unmarshal", err)
		return fmt.Errorf("error unmarshaling request: %w", err)
	}

	// Convert the products slice to a JSON string
	productsJSON, err := json.Marshal(req.Products)
	if err != nil {
		log.Println("Marshal", err)
		return fmt.Errorf("error marshaling products: %w", err)
	}

	// SQL statement with placeholders
	sqlStatement := `
    INSERT INTO Bill_Details (bill_name, bill_date, bill_place, bill_number, customer_phonenumber, customer_pan_num, products, bill_total_amount)
    VALUES ($1, $2, $3, $4::jsonb, $5, $6, $7, $8);`

	// Execute the SQL query with specific values
	_, err = dbClient.DB.Exec(sqlStatement, req.BillName, req.BillDate, req.BillPlace, req.BillNumber, req.CustomerPhonenumber, req.CustomerPanNum, productsJSON, req.BillTotalAmount)
	if err != nil {
		log.Println("error Exec statement", err)
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
		var productsJSON string
		if err := rows.Scan(&bill.ID, &bill.BillName, &bill.BillDate, &bill.BillPlace, &productsJSON, &bill.BillNumber, &bill.CustomerPanNum, &bill.CustomerPhonenumber, &bill.BillTotalAmount); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		// Unmarshal the JSON string into the Products field of the bill
		if err := json.Unmarshal([]byte(productsJSON), &bill.Products); err != nil {
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
	var productsJSON string

	err := dbClient.DB.QueryRow("SELECT * FROM Bill_Details WHERE id = $1;", id).
		Scan(&bill.ID, &bill.BillName, &bill.BillDate, &bill.BillPlace, &productsJSON, &bill.BillNumber, &bill.CustomerPanNum, &bill.CustomerPhonenumber, &bill.BillTotalAmount)
	if err != nil {
		log.Println("Error querying database:", err)
		return bill, fmt.Errorf("error querying database: %w", err)
	}

	// Unmarshal the JSON string into the Products field of the bill
	if err := json.Unmarshal([]byte(productsJSON), &bill.Products); err != nil {
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
