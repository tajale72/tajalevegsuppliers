-- General Ledger Accounts Table
CREATE TABLE GL_Accounts (
    account_id SERIAL PRIMARY KEY,
    account_name VARCHAR(255) NOT NULL,
    account_type VARCHAR(50) NOT NULL, -- Examples: 'Asset', 'Liability', 'Equity', 'Revenue', 'Expense'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- General Ledger Transactions Table
CREATE TABLE GL_Transactions (
    transaction_id SERIAL PRIMARY KEY,
    transaction_date DATE NOT NULL,
    description VARCHAR(255),
    debit_account_id INT NOT NULL REFERENCES GL_Accounts(account_id),
    credit_account_id INT NOT NULL REFERENCES GL_Accounts(account_id),
    amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Customers Table
CREATE TABLE Customers (
    customer_id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    contact_number VARCHAR(15),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Suppliers Table
CREATE TABLE Suppliers (
    supplier_id SERIAL PRIMARY KEY,
    supplier_name VARCHAR(255) NOT NULL,
    contact_number VARCHAR(15),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Products Table
CREATE TABLE Products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    category VARCHAR(255),
    unit_price DECIMAL(10, 2) NOT NULL,
    stock_quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Sales Invoices Table
CREATE TABLE Sales_Invoices (
    invoice_id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL REFERENCES Customers(customer_id),
    invoice_date DATE NOT NULL,
    total_amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Sales Invoice Details Table
CREATE TABLE Sales_Invoice_Details (
    invoice_detail_id SERIAL PRIMARY KEY,
    invoice_id INT NOT NULL REFERENCES Sales_Invoices(invoice_id),
    product_id INT NOT NULL REFERENCES Products(product_id),
    quantity INT NOT NULL,
    price_per_unit DECIMAL(10, 2) NOT NULL,
    total_price DECIMAL(15, 2) NOT NULL
);

-- Purchase Orders Table
CREATE TABLE Purchase_Orders (
    purchase_order_id SERIAL PRIMARY KEY,
    supplier_id INT NOT NULL REFERENCES Suppliers(supplier_id),
    order_date DATE NOT NULL,
    total_amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Purchase Order Details Table
CREATE TABLE Purchase_Order_Details (
    purchase_order_detail_id SERIAL PRIMARY KEY,
    purchase_order_id INT NOT NULL REFERENCES Purchase_Orders(purchase_order_id),
    product_id INT NOT NULL REFERENCES Products(product_id),
    quantity INT NOT NULL,
    price_per_unit DECIMAL(10, 2) NOT NULL,
    total_price DECIMAL(15, 2) NOT NULL
);

-- Payments Table
CREATE TABLE Payments (
    payment_id SERIAL PRIMARY KEY,
    payment_date DATE NOT NULL,
    supplier_id INT NOT NULL REFERENCES Suppliers(supplier_id),
    amount_paid DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Receipts Table
CREATE TABLE Receipts (
    receipt_id SERIAL PRIMARY KEY,
    receipt_date DATE NOT NULL,
    customer_id INT NOT NULL REFERENCES Customers(customer_id),
    amount_received DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inventory Adjustments Table
CREATE TABLE Inventory_Adjustments (
    adjustment_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES Products(product_id),
    adjustment_date DATE NOT NULL,
    adjustment_type VARCHAR(50) NOT NULL, -- Examples: 'Increase', 'Decrease'
    quantity INT NOT NULL,
    reason TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Expenses Table
CREATE TABLE Expenses (
    expense_id SERIAL PRIMARY KEY,
    expense_date DATE NOT NULL,
    description TEXT,
    amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Reports Table
CREATE TABLE Reports (
    report_id SERIAL PRIMARY KEY,
    report_type VARCHAR(255) NOT NULL, -- Examples: 'Profit and Loss', 'Balance Sheet'
    generated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
