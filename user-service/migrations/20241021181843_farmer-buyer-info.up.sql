CREATE TABLE IF NOT EXISTS farmer_info (
    farmer_id INT PRIMARY KEY,
    rating FLOAT DEFAULT 0,
    experience INT,
    bio TEXT,
    FOREIGN KEY (farmer_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS buyer_info (
    buyer_id INT PRIMARY KEY,
    delivery_address VARCHAR(255),
    payment_preferences VARCHAR(255),
    FOREIGN KEY (buyer_id) REFERENCES users (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS socials (
    user_id INT PRIMARY KEY,
    platform VARCHAR(255) NOT NULL,
    account_url VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS documents (
    id BIGSERIAL,
    user_id INT,
    document_type VARCHAR(255) NOT NULL,
    document_url VARCHAR(255) NOT NULL,
    PRIMARY KEY (id, user_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

