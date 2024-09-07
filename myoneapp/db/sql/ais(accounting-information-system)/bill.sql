CREATE DATABASE tajalevegsuppliers;

-- Create the bill_details table with items stored as JSONB
CREATE TABLE bill_details (
    id SERIAL PRIMARY KEY,
    bill_number VARCHAR(255) NOT NULL,
    bill_date DATE NOT NULL,
    bill_total_amount  VARCHAR(255) NOT NULL,
    seller_name VARCHAR(255) NOT NULL,
    seller_pan_num VARCHAR(255) NOT NULL,
    customer_name VARCHAR(255) NOT NULL,
    customer_location VARCHAR(255) NOT NULL,
    customer_phone_number VARCHAR(20) NOT NULL,
    customer_pan_container VARCHAR(255) NOT NULL,
    items JSONB -- Store items as JSONB
);

CREATE TABLE DailyVegetableSales (
    id SERIAL PRIMARY KEY,
    vegetable_name VARCHAR(255) NOT NULL,  -- Name of the vegetable sold
    sale_date DATE NOT NULL,  -- Date of the sale
    quantity_sold INT NOT NULL,  -- Quantity of vegetables sold
    rate NUMERIC(10, 2) NOT NULL,  -- Rate per unit of the vegetable
    total_amount NUMERIC(10, 2) NOT NULL,  -- Total amount for the sale
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Record creation timestamp
    UNIQUE (vegetable_name, sale_date)  -- Updated unique constraint
);

