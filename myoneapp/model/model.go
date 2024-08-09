package model

import (
	"encoding/json"
	"time"
)

// BillDetails represents the structure of the Vegetable_Bill.Bill_Details table
type BillDetails struct {
	ID          int       `json:"id"`
	BillName    string    `json:"bill_name"`
	BillDate    time.Time `json:"bill_date"`
	ProductName string    `json:"product_name"`
	Quantity    float64   `json:"quantity"`
	RatePerKg   float64   `json:"rate_per_kg"`
	Price       float64   `json:"price"`
}

type Product struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Rate     int    `json:"rate"`
	Price    int    `json:"price"`
}

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
	BillNumber           string          `json:"billNumber"`
	BillDate             string          `json:"billDate"`
	BillTotalAmount      string          `json:"billTotalAmount"`
	SellerName           string          `json:"sellerName"`
	SellerPanNum         string          `json:"sellerPanNum"`
	CustomerName         string          `json:"customerName"`
	CustomerLocation     string          `json:"customerLocation"`
	CustomerPhoneNumber  string          `json:"customerPhoneNumber"`
	CustomerPanContainer string          `json:"customerPanContainer"`
	Items                json.RawMessage `json:"items"` // Use RawMessage to handle JSON data
}

// type CustomerInfo struct {
// 	CustomerLocation    string `json:"customerLocation"`
// 	CustomerPanNum      string `json:"customerPanNumber"`
// 	CustomerPhonenumber string `json:"customerPhoneNumber"`
// }
