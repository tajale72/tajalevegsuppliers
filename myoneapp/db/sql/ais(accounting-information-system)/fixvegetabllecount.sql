
INSERT INTO dailyvegetablesales (vegetable_name, sale_date, quantity_sold, rate, total_amount, created_at)
VALUES 
('Garlic - लसुन (Lasun)', '2024-10-14', 48.00, 200.00, 10105.00, '2024-10-14 11:00:26'),
('Cardinier Leaves - कार्डिनियर पात (Cardinier Leaves)', '2024-10-14', 71.50, 120.00, 10010.00, '2024-10-14 11:00:26');



-- Delete multiple entries using OR statements
DELETE FROM dailyvegetablesales
WHERE (vegetable_name = 'Zucchini - जुचिनी (Zucchini)';
        OR vegetable_name = 'Cardinier - कार्डिनियर (Cardinier)'
       OR vegetable_name = 'Cucumber (काक्रो)'
       OR vegetable_name = 'Cucumber (काक्रो)30'
       OR vegetable_name = 'Dhaniya (धनियाँ) and lasun'
       OR vegetable_name = 'Garlic (लसुन)'
       OR vegetable_name = 'Ginger (अदुवा) garlic'
       OR vegetable_name = 'Ginger (अदुवा)5'
       OR vegetable_name = 'Lapsi (लप्सी )'
       OR vegetable_name = 'Lapsi Mada (लप्सी माडा)');





UPDATE dailyvegetablesales
SET quantity_sold = 71.50, rate = 140.00, total_amount = 10010.00
WHERE vegetable_name = 'Cardinier Leaves - कार्डिनियर पात (Cardinier Leaves)';


-- Update "Cucumber - काक्रो (Kakro)"
UPDATE dailyvegetablesales
SET quantity_sold = 1793.50, rate = 60.00, total_amount = 113200.00
WHERE vegetable_name = 'Cucumber - काक्रो (Kakro)';

-- Update "Dhaniya (धनियाँ)"
UPDATE dailyvegetablesales
SET quantity_sold = 15.50, rate = 500.00, total_amount = 4240.00
WHERE vegetable_name = 'Dhaniya (धनियाँ)';

-- Update "Garlic (लसुन)"
UPDATE dailyvegetablesales
SET quantity_sold = 48.00, rate = 150.00, total_amount = 10105.00
WHERE vegetable_name = 'Garlic - लसुन (Lasun)';

-- Update "Ginger (अदुवा)"
UPDATE dailyvegetablesales
SET quantity_sold = 255.80, rate = 170.00, total_amount = 46676.00
WHERE vegetable_name = 'Ginger (अदुवा)';

-- Update "Lapsi Mada - लप्सी माडा (Lapsi Mada)"
UPDATE dailyvegetablesales
SET quantity_sold = 130.00, rate = 70.00, total_amount = 10235.00
WHERE vegetable_name = 'Lapsi Mada - लप्सी माडा (Lapsi Mada)';
