INSERT INTO products (id, description, tags, quantity) VALUES
(1, 'Mechanical Keyboard', ARRAY['tech', 'keyboard'], 100),
(2, 'Wireless Mouse', ARRAY['tech', 'mouse'], 200),
(3, 'USB-C Hub', ARRAY['tech', 'hub'], 50),
(4, 'Gaming Monitor', ARRAY['tech', 'monitor'], 10),
(5, 'Ergonomic Chair', ARRAY['furniture', 'office'], 5);

INSERT INTO product_price_history (product_id, price)
VALUES
(1, 99.99),
(2, 49.99),
(3, 79.99),
(4, 299.99),
(5, 199.99);
