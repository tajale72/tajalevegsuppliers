package model

import (
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

type Request struct {
	ID                  int       `json:"id"`
	BillName            string    `json:"billName"`
	BillPlace           string    `json:"billPlace"`
	BillDate            string    `json:"billDate"`
	Products            []Product `json:"products"`
	BillNumber          string    `json:"billNumber"`
	CustomerPanNum      string    `json:"customerPanNumber"`
	CustomerPhonenumber string    `json:"customerPhoneNumber"`
	BillTotalAmount     float64   `json:"billTotalAmount"`
}
