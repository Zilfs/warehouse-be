CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    sku VARCHAR(50) UNIQUE,
    name VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    description TEXT,
    stock INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);