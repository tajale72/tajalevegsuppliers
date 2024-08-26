package db

import (
	"fmt"
	"log"

	"myoneapp/model"
)

func (dbClient *DBClient) GetVegetableCount() ([]model.Item, error) {
	rows, err := dbClient.DB.Query("SELECT * FROM dailyvegetablesales;")
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var vegetableSales []model.Item
	for rows.Next() {
		var sale model.Item

		if err := rows.Scan(
			&sale.ID,
			&sale.BillNumber,
			&sale.VegetableName,
			&sale.CreatedAt,
			&sale.QuantitySold,
			&sale.Rate,
			&sale.Amount,
			&sale.CreatedAt,
		); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		vegetableSales = append(vegetableSales, sale)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return vegetableSales, nil
}
