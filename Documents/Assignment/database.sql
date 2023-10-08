CREATE DATABASE IF NOT EXISTS Zocket;

USE Zocket;

CREATE TABLE IF NOT EXISTS Users (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    mobile VARCHAR(20),
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO Users (id, name, mobile, latitude, longitude)
VALUES
    (1, 'Aarav Patel', '1234567890', 37.7749, -122.4194),
    (2, 'Ananya Sharma', '0987654321', 40.7128, -74.0060),
    (3, 'Ravi Kumar', '9876543210', 28.6139, 77.2090),
    (4, 'Sunita Gupta', '8765432109', 19.0760, 72.8777),
    (5, 'Amit Singh', '7654321098', 12.9716, 77.5946),
    (6, 'Anita Sharma', '6543210987', 22.5726, 88.3639),
    (7, 'Pradeep Verma', '5432109876', 26.9124, 75.7873),
    (8, 'Shikha Patel', '4321098765', 17.3850, 78.4867),
    (9, 'Vikram Joshi', '3210987654', 18.5204, 73.8567),
    (10, 'Sapna Yadav', '2109876543', 23.2599, 77.4126),
    (50, 'Anushka Pandey', '9246789281',12.0022,12.56775);

CREATE TABLE IF NOT EXISTS Products (
    product_id INT PRIMARY KEY,
    product_name VARCHAR(255),
    product_description TEXT,
    product_images JSON,
    product_price DECIMAL(10, 2),
    compressed_product_images JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
