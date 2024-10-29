CREATE TABLE IF NOT EXISTS farms (
    id BIGSERIAL PRIMARY KEY,
    farmer_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    geo_loc VARCHAR(255),
    size VARCHAR(50),
    crop_types VARCHAR(255),
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (farmer_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS applications (
    id SERIAL PRIMARY KEY,
    farmer_id INT NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    farm_id INT NOT NULL REFERENCES farms(id) ON DELETE SET NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'approved', 'rejected', 'under_review')),
    rejection_reason TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
