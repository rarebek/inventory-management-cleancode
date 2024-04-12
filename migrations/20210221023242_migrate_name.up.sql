CREATE TABLE products (
                          id TEXT PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          category_id TEXT REFERENCES categories(id),
                          unit_price DECIMAL(10,2),
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
                            id TEXT PRIMARY KEY,
                            name VARCHAR(255) NOT NULL,
                            description TEXT,
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE transactions (
                              id TEXT PRIMARY KEY,
                              transaction_type VARCHAR(255),
                              product_id TEXT REFERENCES products(id),
                              quantity INTEGER,
                              unit_price DECIMAL(10,2),
                              transaction_date TIMESTAMP WITH TIME ZONE,
                              created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);



-- Mock data to insert
INSERT INTO categories (name, description) VALUES ('Electronics', 'Electronic devices and accessories');
INSERT INTO categories (name, description) VALUES ('Clothing', 'A variety of clothing items for all genders');
INSERT INTO categories (name, description) VALUES ('Books', 'Fiction, non-fiction, educational, and children\'s books');
INSERT INTO INTO categories (name, description) VALUES ('Home & Kitchen', 'Household items, appliances, and kitchenware');
INSERT INTO categories (name, description) VALUES ('Sports & Outdoors', 'Equipment and apparel for various sports and outdoor activities');

INSERT INTO products (name, description, category_id, unit_price)
VALUES ('Laptop', 'High-performance laptop for work and entertainment', (SELECT id FROM categories WHERE name = 'Electronics'), 799.99);

INSERT INTO products (name, description, category_id, unit_price)
VALUES ('T-Shirt', 'Comfortable and stylish T-shirt', (SELECT id FROM categories WHERE name = 'Clothing'), 19.99);

INSERT INTO products (name, description, category_id, unit_price)
VALUES ('Novel', 'A captivating story for book lovers', (SELECT id FROM categories WHERE name = 'Books'), 14.95);

INSERT INTO products (name, description, category_id, unit_price)
VALUES ('Coffee Maker', 'Brews delicious coffee in minutes', (SELECT id FROM categories WHERE name = 'Home & Kitchen'), 49.99);

INSERT INTO products (name, description, category_id, unit_price)
VALUES ('Running Shoes', 'Lightweight and supportive running shoes', (SELECT id FROM categories WHERE name = 'Sports & Outdoors'), 89.99);
