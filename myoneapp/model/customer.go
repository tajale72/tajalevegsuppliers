package model

import (
	"time"
)

// Customer represents the customers table structure
type Customer struct {
	ID                   int       `json:"id"`                   // Primary key
	CustomerName         string    `json:"customerName"`         // Customer name
	CustomerLocation     string    `json:"customerLocation"`     // Customer location
	CustomerPhoneNumber  string    `json:"customerPhoneNumber"`  // Customer phone number
	CustomerPanContainer string    `json:"customerPanContainer"` // Customer PAN container
	CreatedAt            time.Time `json:"created_at"`           // Timestamp of record creation
}
