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

func GetDBConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", user, password, host, port, dbname, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("error conneting to postgres db", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("error from pinging", err)
		return nil, err
	}

	return db, nil
}

func CreateTable(db *sql.DB, data []byte) error {
	var req model.Request
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Println("Unmarshal", err)
		return err
	}

	// Convert the products slice to a JSON string
	productsJSON, err := json.Marshal(req.Products)
	if err != nil {
		log.Println("Marshal", err)
		return err
	}

	// SQL statement with placeholders
	sqlStatement := `
    INSERT INTO Bill_Details (bill_name, bill_date, bill_place, bill_number, customer_phonenumber, customer_pan_num, products, bill_total_amount)
VALUES ($1, $2, $3, $4::jsonb, $5, $6, $7, $8);`

	// Execute the SQL query with specific values
	_, err = db.Exec(sqlStatement, req.BillName, req.BillDate, req.BillPlace, req.BillNumber, req.CustomerPhonenumber, req.CustomerPanNum, productsJSON, req.BillTotalAmount)
	if err != nil {
		log.Println("error Exec statement", err)
		return err
	}

	return nil
}

func GetProducts(db *sql.DB) ([]model.Request, error) {
	rows, err := db.Query("SELECT * FROM Bill_Details;")
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	var listOfBills []model.Request
	for rows.Next() {
		var bill model.Request
		var productsJSON string
		if err := rows.Scan(&bill.ID, &bill.BillName, &bill.BillDate, &bill.BillPlace, &productsJSON, &bill.BillNumber, &bill.CustomerPanNum, &bill.CustomerPhonenumber, &bill.BillTotalAmount); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		// Unmarshal the JSON string into the Products field of the bill
		if err := json.Unmarshal([]byte(productsJSON), &bill.Products); err != nil {
			log.Println("Error unmarshaling products:", err)
			return nil, err
		}
		listOfBills = append(listOfBills, bill)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, err
	}

	return listOfBills, nil
}

func GetProductByID(db *sql.DB, id int) (model.Request, error) {
	var bill model.Request
	var productsJSON string

	err := db.QueryRow("SELECT * FROM Bill_Details WHERE id = $1;", id).
		Scan(&bill.ID, &bill.BillName, &bill.BillDate, &bill.BillPlace, &productsJSON, &bill.BillNumber, &bill.CustomerPanNum, &bill.CustomerPhonenumber, &bill.BillTotalAmount)
	if err != nil {
		log.Println("Error querying database:", err)
		return bill, err
	}

	// Unmarshal the JSON string into the Products field of the bill
	if err := json.Unmarshal([]byte(productsJSON), &bill.Products); err != nil {
		log.Println("Error unmarshaling products:", err)
		return bill, err
	}

	return bill, nil
}

func updateUser(db *sql.DB, id int, newUsername, newEmail string) error {
	_, err := db.Exec("UPDATE users SET username=$1, email=$2 WHERE id=$3", newUsername, newEmail, id)
	log.Println("")
	return err
}

func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	log.Println("")
	return err
}

type User struct {
	ID       int
	Username string
	Email    string
}

// func main() {
// 	db, err := getDBConnection()
// 	if err != nil {
// 		log.Fatal("Error connecting to the database: ", err)
// 	}
// 	defer db.Close()

// 	// Example usage
// 	err = createUser(db, "john_doe", "john@example.com")
// 	if err != nil {
// 		log.Fatal("Error creating user: ", err)
// 	}

// 	users, err := getUsers(db)
// 	if err != nil {
// 		log.Fatal("Error fetching users: ", err)
// 	}

// 	fmt.Println("Users:")
// 	for _, user := range users {
// 		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
// 	}

// 	err = updateUser(db, 1, "john_smith", "john.smith@example.com")
// 	if err != nil {
// 		log.Fatal("Error updating user: ", err)
// 	}

// 	err = deleteUser(db, 1)
// 	if err != nil {
// 		log.Fatal("Error deleting user: ", err)
// 	}
// }
