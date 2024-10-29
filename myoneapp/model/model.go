package model

import (
	"encoding/json"
	"time"
)

// Item represents a single item in the bill.
type Item struct {
	ID            int       `json:"id"`
	BillNumber    string    `json:"billNumber" validate:"required"`
	VegetableName string    `json:"sa"`      // विवरण (Description)
	QuantitySold  string    `json:"sankhya"` // मात्रा (Quantity)
	Rate          string    `json:"dar"`     // दर (Rate)
	Amount        string    `json:"rakam"`   // रकम (Amount)
	CreatedAt     time.Time `json:"created_at"`
}

type Request struct {
	ID                   string          `json:"id"`
	BillNumber           string          `json:"billNumber" validate:"required"`
	BillDate             string          `json:"billDate" validate:"required"`
	BillTotalAmount      string          `json:"billTotalAmount" validate:"required"`
	SellerName           string          `json:"sellerName" validate:"required"`
	SellerPanNum         string          `json:"sellerPanNum" validate:"required"`
	CustomerID           int             `json:"customerID" validate:"required"`
	CustomerName         string          `json:"customerName" validate:"required"`
	CustomerLocation     string          `json:"customerLocation" validate:"required"`
	CustomerPhoneNumber  string          `json:"customerPhoneNumber" validate:"required"`
	CustomerPanContainer string          `json:"customerPanContainer" validate:"required"`
	TotalDebit           string          `json:"totalDebit"`
	TotalCredit          string          `json:"totalCredit"`
	FinalTotalAmount     float64         `json:"finalTotalAmount"`
	Items                json.RawMessage `json:"items" validate:"required"` // Use RawMessage to handle JSON data
}

type Result struct {
	BillNumber int64 `json:"billNumber"`
}
