package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Custom type for handling string-to-float64 conversion
type Float64 float64

type LedgerEntry struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	Date          string    `json:"date"`          // Date of the transaction
	Account       string    `json:"account"`       // Account associated with the entry
	BillNumber    string    `json:"billNumber"`    // Account associated with the entry
	Debit         Float64   `json:"debit"`         // Debit amount (custom type)
	Credit        Float64   `json:"credit"`        // Credit amount (custom type)
	BalanceAmount Float64   `json:"balanceAmount"` // Remaining balance (custom type)
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type SaleLedger struct {
	ID               int       `json:"id" gorm:"primaryKey"`
	BillID           int       `json:"bill_id"`
	BillNumber       string    `json:"bill_number"`
	BillDate         time.Time `json:"bill_date"`
	CustomerName     string    `json:"customer_name"`
	CustomerPhone    string    `json:"customer_phone_number"`
	TotalAmount      float64   `json:"total_amount"`
	PaidAmount       float64   `json:"paid_amount" gorm:"default:0.00"`
	RemainingBalance float64   `json:"remaining_balance" gorm:"->"` // Calculated field
	PaymentStatus    string    `json:"payment_status" gorm:"type:enum('Pending','Partially Paid','Paid');default:'Pending'"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type PurchaseLedger struct {
	ID               int       `json:"id" gorm:"primaryKey"`
	BillID           int       `json:"bill_id"`
	BillNumber       string    `json:"bill_number"`
	BillDate         time.Time `json:"bill_date"`
	SupplierName     string    `json:"supplier_name"`
	TotalAmount      float64   `json:"total_amount"`
	PaidAmount       float64   `json:"paid_amount" gorm:"default:0.00"`
	RemainingBalance float64   `json:"remaining_balance" gorm:"->"` // Calculated field
	PaymentStatus    string    `json:"payment_status" gorm:"type:enum('Pending','Partially Paid','Paid');default:'Pending'"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// UnmarshalJSON handles the conversion of string to float64
func (f *Float64) UnmarshalJSON(b []byte) error {
	var s string
	// Unmarshal the field into a string first
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("Float64 should be a string, got %s", b)
	}

	// Convert string to float64
	floatValue, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("error converting string to float64: %w", err)
	}

	*f = Float64(floatValue) // Assign the converted value to the custom type
	return nil
}
