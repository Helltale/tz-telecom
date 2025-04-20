CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INTEGER NOT NULL,
    is_married BOOLEAN,
    password TEXT NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    description TEXT,
    tags TEXT[],
    quantity INTEGER NOT NULL,
    price NUMERIC NOT NULL DEFAULT 0
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE order_items (
    order_id INTEGER NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES products(id),
    quantity INTEGER NOT NULL,
    price NUMERIC NOT NULL
);
