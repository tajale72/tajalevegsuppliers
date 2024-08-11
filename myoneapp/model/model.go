package model

import (
	"encoding/json"
)

// Item represents a single item in the bill.
type Item struct {
	Kra     string `json:"kra"`
	Sa      string `json:"sa"`
	Sankhya string `json:"sankhya"`
	Dar     string `json:"dar"`
	Rakam   string `json:"rakam"`
}

type Request struct {
	ID                   int             `json:"id"`
	BillNumber           string          `json:"billNumber" validate:"required"`
	BillDate             string          `json:"billDate" validate:"required"`
	BillTotalAmount      string          `json:"billTotalAmount" validate:"required"`
	SellerName           string          `json:"sellerName" validate:"required"`
	SellerPanNum         string          `json:"sellerPanNum" validate:"required"`
	CustomerName         string          `json:"customerName" validate:"required"`
	CustomerLocation     string          `json:"customerLocation" validate:"required"`
	CustomerPhoneNumber  string          `json:"customerPhoneNumber" validate:"required"`
	CustomerPanContainer string          `json:"customerPanContainer" validate:"required"`
	Items                json.RawMessage `json:"items" validate:"required"` // Use RawMessage to handle JSON data
}

type Result struct {
	BillNumber int64 `json:"billNumber"`
}
