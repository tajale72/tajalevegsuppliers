CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,   -- Name of the customer
    customer_type VARCHAR(100) NOT NULL,   -- Party Palace, School, Store, etc.
    contact_info VARCHAR(255),             -- Optional: Contact details like phone or email
    created_at TIMESTAMP DEFAULT NOW()     -- Timestamp of customer creation
);


CREATE TABLE customer_payments (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id), -- Link to the customer table
    payment_amount NUMERIC(10, 2) NOT NULL,   -- Payment made by the customer
    payment_date DATE NOT NULL,               -- Date of payment
    created_at TIMESTAMP DEFAULT NOW()        -- Timestamp of payment entry
);