package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"myoneapp/model"
)

func (dbClient *DBClient) GetVegetableCount() ([]model.Item, error) {

	dbClient.updateVegetableData()

	var rows *sql.Rows
	var err error

	// Try querying AivenDB first if it's connected
	if dbClient.AivenDB != nil {
		rows, err = dbClient.AivenDB.Query("SELECT * FROM dailyvegetablesales ORDER BY vegetable_name;")
		if err != nil {
			log.Printf("Error querying AivenDB: %v", err)
		} else {
			log.Println("Queried AivenDB successfully")
		}
	}

	// If AivenDB is not available or the query failed, try the local DB
	if rows == nil && dbClient.DB != nil {
		rows, err = dbClient.DB.Query("SELECT * FROM dailyvegetablesales ORDER BY vegetable_name;")
		if err != nil {
			log.Printf("Error querying local DB: %v", err)
			return nil, fmt.Errorf("error querying local DB: %w", err)
		}
		log.Println("Queried local DB successfully")
	} else if rows == nil {
		// If both queries failed
		return nil, fmt.Errorf("no available database to query from")
	}

	defer rows.Close()

	var vegetableSales []model.Item
	for rows.Next() {
		var sale model.Item

		if err := rows.Scan(
			&sale.ID,
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

// Define a mapping of old data to new data with sale_date
var vegetableDataUpdates = []struct {
	OldName  string
	NewName  string
	Quantity float64
	BillNum  int
	Rate     float64
	TotalAmt float64
	SaleDate string // Add the sale_date field
}{
	{"Tomato (टमाटर)", "Tomato - टमाटर (Tamatar)", 323.80, 5593, 90.00, 21856.00, "2024-10-14"},
	{"bodi", "Bodi - बोडी (Bodi)", 124.00, 5707, 60.00, 7440.00, "2024-10-14"},
	{"brinjal (bhanta)", "Brinjal - भाँटा (Bhanta)", 2.50, 5693, 80.00, 200.00, "2024-10-14"},
	{"cacumber", "Cucumber - काक्रो (Kakro)", 1120.50, 5723, 60.00, 67230.00, "2024-10-14"},
	{"cardinier", "Cardinier - कार्डिनियर (Cardinier)", 20.60, 5722, 140.00, 2884.00, "2024-10-14"},
	{"cardinier leavs", "Cardinier Leaves - कार्डिनियर पात (Cardinier Leaves)", 50.90, 5704, 140.00, 7126.00, "2024-10-14"},
	{"ghiraula", "Ghiraula - घिरौंला (Ghiraula)", 107.00, 5698, 40.00, 4280.00, "2024-10-14"},
	{"hariyo lasun", "Spring Garlic - हरियो लसुन (Spring Garlic)", 1.00, 5661, 200.00, 200.00, "2024-10-14"},
	{"iskus", "Iskus - इस्कुस (Iskus)", 300.00, 5711, 35.00, 10500.00, "2024-10-14"},
	{"karela", "Bitter Gourd - करेले (Karela)", 90.50, 5712, 60.00, 5430.00, "2024-10-14"},
	{"kino", "Small Tomato - सानो टमाटर (Sano Tamatar)", 2465.00, 5743, 100.00, 246500.00, "2024-10-14"},
	{"lapsi", "Lapsi Mada - लप्सी माडा (Lapsi Mada)", 61.50, 5720, 70.00, 4305.00, "2024-10-14"},
	{"lasun", "Garlic - लसुन (Lasun)", 0.50, 5670, 200.00, 100.00, "2024-10-14"},
	{"lemon", "Lemon - कागती (Lemon)", 83.80, 5716, 150.00, 12570.00, "2024-10-14"},
	{"naspati", "Pear - नासपाती (Naspati)", 206.00, 5701, 75.00, 15450.00, "2024-10-14"},
	{"pakucha saag", "Pakucha Saag - पक्कुचा साग (Pakucha Saag)", 25.50, 5674, 220.00, 5460.00, "2024-10-14"},
	{"palak sag", "Spinach - पालुङ्गो (Palungo)", 277.50, 5741, 65.00, 18037.50, "2024-10-14"},
	{"parwal", "Parwal - परवल (Parwal)", 320.00, 5724, 60.00, 19200.00, "2024-10-14"},
	{"saag", "Saag - साग (Saag)", 1.00, 5675, 150.00, 150.00, "2024-10-14"},
	{"spring garlic", "Spring Garlic - हरियो लसुन (Spring Garlic)", 120.70, 5733, 120.00, 14484.00, "2024-10-14"},
	{"spring ginger", "Spring Ginger - हरियो अदुवा (Spring Ginger)", 104.30, 5709, 120.00, 12516.00, "2024-10-14"},
}

func (dbClient *DBClient) updateVegetableData() error {
	for _, update := range vegetableDataUpdates {
		// Convert sale_date to time.Time format
		saleDate, err := time.Parse("2006-01-02", update.SaleDate)
		if err != nil {
			log.Printf("Invalid sale_date format for %s: %v", update.OldName, err)
			continue
		}

		query := `
			INSERT INTO dailyvegetablesales (
				vegetable_name,
				sale_date,
				quantity_sold,
				rate,
				total_amount,
				created_at
			) VALUES ($1, $2, $3, $4, $5, NOW())
			ON CONFLICT (vegetable_name, sale_date) 
			DO UPDATE SET 
				vegetable_name = EXCLUDED.vegetable_name,
				quantity_sold = dailyvegetablesales.quantity_sold + EXCLUDED.quantity_sold,
				rate = EXCLUDED.rate,
				total_amount = dailyvegetablesales.total_amount + EXCLUDED.total_amount;
		`

		// Execute the upsert query
		result, err := dbClient.DB.Exec(query, update.NewName, saleDate, update.Quantity, update.Rate, update.TotalAmt)
		if err != nil {
			log.Printf("Failed to upsert %s to %s: %v", update.OldName, update.NewName, err)
			return err
		}

		// Log the number of rows affected
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("Failed to get rows affected for %s to %s: %v", update.OldName, update.NewName, err)
			return err
		}

		if rowsAffected == 0 {
			log.Printf("No rows matched or inserted for %s on date %s", update.OldName, update.SaleDate)
		} else {
			log.Printf("Upserted %d rows: %s to %s", rowsAffected, update.OldName, update.NewName)
		}

		// Delete the old record after the update
		deleteQuery := `
			DELETE FROM dailyvegetablesales
			WHERE vegetable_name = $1;
		`

		// Execute the delete query
		_, err = dbClient.DB.Exec(deleteQuery, update.OldName)
		if err != nil {
			log.Printf("Failed to delete %s: %v", update.OldName, err)
			return err
		}
		log.Printf("Deleted old record: %s", update.OldName)

	}

	return nil
}

// func step2() {
// 	var vegetableDataUpdates = []struct {
// 		OldName  string
// 		NewName  string
// 		Quantity float64
// 		Rate     float64
// 		TotalAmt float64
// 		SaleDate string // Sale date for matching
// 	}{
// 		{"Apple (स्याउ)", "Apple - स्याउ (Syaau)", 1904.60, 270.00, 514302.00, "2024-10-23"},
// 		{"Banana (केरा)", "Banana - केरा (Kera)", 1902.20, 85.00, 227513.00, "2024-10-23"},
// 		{"Beans (सिमी)", "Beans - सिमी (Simi)", 283.50, 200.00, 20810.00, "2024-10-23"},
// 		{"Beet (चकुन्दर)", "Beet - चकुन्दर (Chakundar)", 2.00, 100.00, 200.00, "2024-10-23"},
// 		{"Bitter Gourd - करेले (Karela)", "Bitter Gourd - करेले (Karela)", 90.50, 60.00, 5430.00, "2024-10-23"},
// 		{"Bodi - बोडी (Bodi)", "Bodi - बोडी (Bodi)", 124.00, 60.00, 7440.00, "2024-10-23"},
// 		{"Brinjal - भाँटा (Bhanta)", "Brinjal - भाँटा (Bhanta)", 2.50, 80.00, 200.00, "2024-10-23"},
// 		{"Broccoli (ब्रोकोली)", "Broccoli - ब्रोकोली (Broccoli)", 379.00, 100.00, 39100.00, "2024-10-23"},
// 		{"Cabbage (बन्दगोभी)", "Cabbage - बन्दगोभी (Bandgobhi)", 764.00, 40.00, 30810.00, "2024-10-23"},
// 		{"Capsicum (क्याप्सिकम)", "Capsicum - क्याप्सिकम (Capsicum)", 321.20, 140.00, 34200.00, "2024-10-23"},
// 		{"Cardinier Leaves - कार्डिनियर पात (Cardinier Leaves)", "Cardinier Leaves - कार्डिनियर पात (Cardinier Leaves)", 71.50, 140.00, 10010.00, "2024-10-23"},
// 		{"Carrot (गाजर)", "Carrot - गाजर (Gajar)", 1094.00, 120.00, 85160.00, "2024-10-23"},
// 		{"Cauliflower (काउली)", "Cauliflower - काउली (Cauliflower)", 973.50, 100.00, 75527.50, "2024-10-23"},
// 		{"Celery (अजमोदा)", "Celery - सेलेरी (Celery)", 1.00, 100.00, 100.00, "2024-10-23"},
// 		{"Chili (खुर्सानी)", "Chili - खुर्सानी (Khursani)", 50.50, 140.00, 5095.00, "2024-10-23"},
// 		{"Chyau Patey (च्याउ पाते)", "Chyau Patey - च्याउ पाते (Chyau Patey)", 388.00, 160.00, 69780.00, "2024-10-23"},
// 		{"Cucumber (काक्रो)", "Cucumber - काक्रो (Kakro)", 643.00, 60.00, 43270.00, "2024-10-23"},
// 		{"Dhaniya (धनियाँ)", "Dhaniya - धनियाँ (Dhaniya)", 13.50, 500.00, 3540.00, "2024-10-23"},
// 		{"Garlic (लसुन)", "Garlic - लसुन (Lasun)", 47.50, 150.00, 10005.00, "2024-10-23"},
// 		{"Ghiraula - घिरौंला (Ghiraula)", "Ghiraula - घिरौंला (Ghiraula)", 107.00, 40.00, 4280.00, "2024-10-23"},
// 		{"Ginger (अदुवा)", "Ginger - अदुवा (Aduwa)", 246.80, 170.00, 44156.00, "2024-10-23"},
// 		{"Golo Chyau (गोलो च्याउ)", "Golo Chyau - गोलो च्याउ (Golo Chyau)", 20.00, 550.00, 9500.00, "2024-10-23"},
// 		{"Grapes (अंगुर)", "Grapes - अंगुर (Angur)", 2.00, 600.00, 1200.00, "2024-10-23"},
// 		{"Green Chillies (हरियो खुर्सानी)", "Green Chillies - हरियो खुर्सानी (Hariyo Khursani)", 52.50, 90.00, 5065.00, "2024-10-23"},
// 		{"Hariyo Simi (हरियो सिमी)", "Hariyo Simi - हरियो सिमी (Hariyo Simi)", 4.00, 220.00, 880.00, "2024-10-23"},
// 		{"Iskus - इस्कुस (Iskus)", "Iskus - इस्कुस (Iskus)", 300.00, 35.00, 10500.00, "2024-10-23"},
// 		{"Kagati (कागती)", "Kagati - कागती (Kagati)", 25.00, 280.00, 6880.00, "2024-10-23"},
// 		{"Lapsi (लप्सी )", "Lapsi - लप्सी (Lapsi)", 59.00, 70.00, 4130.00, "2024-10-23"},
// 		{"Lapsi Mada (लप्सी माडा)", "Lapsi Mada - लप्सी माडा (Lapsi Mada)", 9.50, 200.00, 1800.00, "2024-10-23"},
// 		{"Lemon - कागती (Lemon)", "Lemon - कागती (Lemon)", 83.80, 150.00, 12570.00, "2024-10-23"},
// 		{"Mushroom (च्याउ)", "Mushroom - च्याउ (Chyau)", 62.00, 500.00, 15250.00, "2024-10-23"},
// 		{"Onion (प्याज)", "Onion - प्याज (Pyaj)", 312.00, 100.00, 33675.00, "2024-10-23"},
// 		{"Pakucha Saag - पक्कुचा साग (Pakucha Saag)", "Pakucha Saag - पक्कुचा साग (Pakucha Saag)", 25.50, 220.00, 5460.00, "2024-10-23"},
// 		{"Parwal - परवल (Parwal)", "Parwal - परवल (Parwal)", 320.00, 60.00, 19200.00, "2024-10-23"},
// 		{"Pear - नासपाती (Naspati)", "Pear - नासपाती (Naspati)", 206.00, 75.00, 15450.00, "2024-10-23"},
// 		{"Peas (केराउ)", "Peas - केराउ (Keraau)", 1.00, 120.00, 120.00, "2024-10-23"},
// 		{"Potato (आलु)", "Potato - आलु (Aalu)", 9084.30, 48.00, 457276.40, "2024-10-23"},
// 		{"Pudina (पुदिना)", "Pudina - पुदिना (Pudina)", 7.50, 1000.00, 1200.00, "2024-10-23"},
// 		{"Pumpkin (फर्सी)", "Pumpkin - फर्सी (Pharsi)", 204.00, 70.00, 7280.00, "2024-10-23"},
// 		{"Radish (मूला)", "Radish - मूला (Mula)", 1194.00, 45.00, 48340.00, "2024-10-23"},
// 		{"Ramtoria (रामतोरिया)", "Ramtoria - रामतोरिया (Ramtoria)", 457.00, 60.00, 27420.00, "2024-10-23"},
// 		{"Rayo Saag (रायो साग)", "Rayo Saag - रायो साग (Rayo Saag)", 583.00, 250.00, 64270.00, "2024-10-23"},
// 		{"Saag - साग (Saag)", "Saag - साग (Saag)", 1.00, 150.00, 150.00, "2024-10-23"},
// 		{"Salad Patta (सलाद पत्ता)", "Salad Patta - सलाद पत्ता (Salad Patta)", 4.00, 700.00, 2600.00, "2024-10-23"},
// 		{"Small Tomato (सानो टमाटर)", "Small Tomato - सानो टमाटर (Sano Tamatar)", 116.50, 90.00, 9555.00, "2024-10-23"},
// 		{"Spinach - पालुङ्गो (Palungo)", "Spinach - पालुङ्गो (Palungo)", 277.50, 65.00, 18037.50, "2024-10-23"},
// 		{"Spring Garlic - हरियो लसुन (Spring Garlic)", "Spring Garlic - हरियो लसुन (Spring Garlic)", 121.70, 120.00, 14684.00, "2024-10-23"},
// 		{"Spring Ginger - हरियो अदुवा (Spring Ginger)", "Spring Ginger - हरियो अदुवा (Spring Ginger)", 104.30, 120.00, 12516.00, "2024-10-23"},
// 		{"Tomato - टमाटर (Tamatar)", "Tomato - टमाटर (Tamatar)", 323.80, 90.00, 21856.00, "2024-10-23"},
// 	}

// }
