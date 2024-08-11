package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"myoneapp/model"
)

// Define a database interface or type from myoneapp/db package
type Database interface {
	CreateTable(data []byte) error
	GetProducts() ([]model.Request, error)
	GetProductByID(id int) (model.Request, error)
	GetLastBillNumber() (model.Result, error)
	// Add other database-related methods if needed
}

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

// ValidateRequest checks if all required fields in the Request struct are present and valid.
func ValidateRequest(data []byte) (*model.Request, error) {
	// Unmarshal JSON data into the request model
	var req model.Request
	if err := json.Unmarshal(data, &req); err != nil {
		log.Println("Unmarshal error:", err)
		return nil, fmt.Errorf("error unmarshaling request: %w", err)
	}

	if req.BillNumber == "" {
		return nil, errors.New("BillNumber is required")
	}
	if req.BillDate == "" {
		return nil, errors.New("BillDate is required")
	}
	if req.CustomerName == "" {
		return nil, errors.New("CustomerName is required")
	}
	if req.CustomerLocation == "" {
		return nil, errors.New("CustomerLocation is required")
	}
	if req.CustomerPhoneNumber == "" {
		return nil, errors.New("CustomerPhoneNumber is required")
	}
	if req.CustomerPanContainer == "" {
		return nil, errors.New("CustomerPanContainer is required")
	}
	if req.BillTotalAmount == "" {
		return nil, errors.New("BillTotalAmount is required")
	}
	if req.SellerName == "" {
		return nil, errors.New("SellerName is required")
	}
	if req.SellerPanNum == "" {
		return nil, errors.New("SellerPanNum is required")
	}

	if len(req.Items) == 0 {
		return nil, errors.New("items are required")
	}
	return &req, nil
}
