CREATE DATABASE tajalevegsuppliers;

CREATE TABLE Bill_Details (
    id SERIAL PRIMARY KEY,
    bill_name VARCHAR(255),
    bill_date DATE,
    bill_place VARCHAR(255),
    products JSONB,
    bill_number VARCHAR(255),
    customer_pan_num VARCHAR(255),
    customer_phonenumber VARCHAR(255),
    bill_total_amount DECIMAL
);
