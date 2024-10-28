package model

import (
	"time"
)

// Customer represents the customers table structure
type Customer struct {
	ID                   int       `json:"id"`                     // Primary key
	CustomerName         string    `json:"customer_name"`          // Customer name
	CustomerLocation     string    `json:"customer_location"`      // Customer location
	CustomerPhoneNumber  string    `json:"customer_phone_number"`  // Customer phone number
	CustomerPanContainer string    `json:"customer_pan_container"` // Customer PAN container
	CreatedAt            time.Time `json:"created_at"`             // Timestamp of record creation
}
