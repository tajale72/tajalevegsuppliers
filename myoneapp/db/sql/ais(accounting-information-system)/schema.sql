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

CREATE TABLE ledger_entries (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    account VARCHAR(255) NOT NULL,
    billNumber VARCHAR(255) NOT NULL,
    debit DECIMAL(10, 2) DEFAULT 0.00,
    credit DECIMAL(10, 2) DEFAULT 0.00,
    balance_amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sale_ledger (
    id SERIAL PRIMARY KEY,
    bill_id INT REFERENCES bill_details(id) ON DELETE CASCADE,
    bill_number VARCHAR(255) NOT NULL,
    bill_date DATE NOT NULL,
    customer_name VARCHAR(255) NOT NULL,
    customer_phone_number VARCHAR(20) NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    paid_amount DECIMAL(10, 2) DEFAULT 0.00,
    remaining_balance DECIMAL(10, 2) GENERATED ALWAYS AS (total_amount - paid_amount) STORED,
    payment_status ENUM('Pending', 'Partially Paid', 'Paid') DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE purchase_ledger (
    id SERIAL PRIMARY KEY,
    bill_id INT REFERENCES bill_details(id) ON DELETE CASCADE,
    bill_number VARCHAR(255) NOT NULL,
    bill_date DATE NOT NULL,
    supplier_name VARCHAR(255) NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    paid_amount DECIMAL(10, 2) DEFAULT 0.00,
    remaining_balance DECIMAL(10, 2) GENERATED ALWAYS AS (total_amount - paid_amount) STORED,
    payment_status ENUM('Pending', 'Partially Paid', 'Paid') DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE balance_sheet (
    id SERIAL PRIMARY KEY,
    account_type VARCHAR(50) NOT NULL,    -- Asset, Liability, Equity
    account_name VARCHAR(255) NOT NULL,   -- Cash, Inventory, Payables
    debit NUMERIC(10, 2),                 -- Debit amount for transactions
    credit NUMERIC(10, 2),                -- Credit amount for transactions
    balance NUMERIC(10, 2) NOT NULL,      -- Account balance
    transaction_date DATE NOT NULL,       -- Date of the transaction
    created_at TIMESTAMP DEFAULT NOW()    -- Timestamp of creation
);
