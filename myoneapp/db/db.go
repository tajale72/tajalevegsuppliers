package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"myoneapp/db/aiven"
	"myoneapp/model"
)

// Define a database interface or type from myoneapp/db package
type Database interface {
	CreateTable(data []byte) error
	CreateLedger(data []byte) error
	GetLedgerEntries() ([]model.LedgerEntry, error)

	GetProducts() ([]model.Request, error)
	GetProductsBySearch(searchQuery string) ([]model.Request, error)
	GetVegetableCount() ([]model.Item, error)
	GetProductByID(id int) (model.Request, error)
	DeleteProductByID(id int) (sql.Result, error)

	GetLastBillNumber() (model.Result, error)
	UpdateBill(data []byte) error

	GetCustomers() ([]model.Customer, error)
	GetCustomerLedgerByName(account string) ([]model.LedgerEntry, error)
	// Add other database-related methods if needed

	GetDashBoardData() (*model.Dashboard, error)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "romit" // Empty string for no username
	password = "admin" // Empty string for no password
	dbname   = "tajalevegsuppliers"
	sslmode  = "disable"
)

type DBClient struct {
	DB      *sql.DB
	AivenDB *sql.DB
}

func GetDBConnection() (*DBClient, error) {
	// Try to connect to AivenDB first
	var aivendb, localdb *sql.DB
	var err error

	aivendb, err = aiven.ConnectDBAivenPostgres()
	if err != nil {
		log.Printf("Error connecting to AivenDB: %v", err)
	} else {
		log.Println("Connected to AivenDB successfully")
	}

	// Now attempt to connect to the local DB
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", user, password, host, port, dbname, sslmode)
	localdb, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to local postgres DB: %w", err)
	}

	// Ping the local DB to ensure it's connected
	if err := localdb.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging local database: %w", err)
	}

	log.Println("Connected to local DB successfully")

	// If both AivenDB and local DB are connected, return both
	if aivendb != nil && localdb != nil {
		log.Println("Both AivenDB and local DB connections successful")
		return &DBClient{DB: localdb, AivenDB: aivendb}, nil
	}

	// If only local DB is connected, return that
	if aivendb == nil {
		log.Println("AivenDB connection failed, using only local DB")
		return &DBClient{DB: localdb, AivenDB: nil}, nil
	}

	// If AivenDB is connected but local DB is not
	if err != nil {
		log.Println("Local DB connection failed, using only AivenDB")
		return &DBClient{DB: nil, AivenDB: aivendb}, nil
	}

	// Return an error if both connections fail
	return nil, fmt.Errorf("both AivenDB and local DB connections failed")
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
