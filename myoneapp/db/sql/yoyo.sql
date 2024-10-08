CREATE TABLE Category (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE Product (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock_quantity INT NOT NULL,
    category_id INT REFERENCES Category(category_id) ON DELETE SET NULL
);

CREATE TABLE User (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Order (
    order_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES User(user_id) ON DELETE SET NULL,
    order_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    total_amount DECIMAL(10, 2) NOT NULL
);

CREATE TABLE OrderItem (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES Order(order_id) ON DELETE CASCADE,
    product_id INT REFERENCES Product(product_id),
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE Review (
    review_id SERIAL PRIMARY KEY,
    product_id INT REFERENCES Product(product_id) ON DELETE CASCADE,
    user_id INT REFERENCES User(user_id) ON DELETE SET NULL,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Payment (
    payment_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES Order(order_id) ON DELETE SET NULL,
    payment_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    amount DECIMAL(10, 2) NOT NULL,
    payment_status VARCHAR(50) NOT NULL
);


CREATE TABLE bills (
    id SERIAL PRIMARY KEY,
    seller_name TEXT NOT NULL,
    seller_pan_num TEXT NOT NULL,
    customer_name TEXT NOT NULL,
    customer_location TEXT NOT NULL,
    customer_phone_number TEXT NOT NULL,
    customer_pan_container TEXT NOT NULL,
    bill_number TEXT NOT NULL,
    bill_date DATE NOT NULL,
    items JSONB NOT NULL
);
