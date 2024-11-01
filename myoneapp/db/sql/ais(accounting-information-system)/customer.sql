-- Create the customers table
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    customer_location VARCHAR(255),
    customer_phone_number VARCHAR(20),
    customer_pan_container VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add customer_id to the bill_details table and set up a foreign key
ALTER TABLE bill_details
ADD COLUMN customer_id INT;

ALTER TABLE bill_details
ADD CONSTRAINT fk_bill_customer
FOREIGN KEY (customer_id) REFERENCES customers(id);

-- Add customer_id to the ledger_entries table and set up a foreign key
ALTER TABLE ledger_entries
ADD COLUMN customer_id INT;

ALTER TABLE ledger_entries
ADD CONSTRAINT fk_ledger_customer
FOREIGN KEY (customer_id) REFERENCES customers(id);

-- Insert unique customer names from bill_details into customers table
INSERT INTO customers (customer_name, customer_location, customer_phone_number, customer_pan_container)
SELECT DISTINCT 
    customer_name, 
    customer_location, 
    customer_phone_number, 
    customer_pan_container
FROM bill_details
WHERE customer_name IS NOT NULL
  AND customer_name NOT IN (SELECT customer_name FROM customers);

-- Update customer_id in bill_details after populating the customers table
UPDATE bill_details bd
SET customer_id = c.id
FROM customers c
WHERE bd.customer_name = c.customer_name;


INSERT INTO ledger_entries (date, account, billNumber, debit, credit, balance_amount)
SELECT 
    bill_date AS date,
    customer_name AS account,
    bill_number AS billNumber,
    0.00 AS debit,
    bill_total_amount::DECIMAL(10, 2) AS credit,
    bill_total_amount::DECIMAL(10, 2) AS balance_amount
FROM 
    bill_details;
