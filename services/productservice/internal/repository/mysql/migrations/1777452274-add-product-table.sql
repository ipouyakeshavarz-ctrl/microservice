-- +migrate Up

CREATE TABLE `products` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `store_id` INT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `category` ENUM('Sport', 'Electronics', 'Fashion', 'Home', 'Beauty') NOT NULL,
    `price` DECIMAL(12,2) NOT NULL DEFAULT 0.00,
    `stock` INT NOT NULL DEFAULT 0,
    `sku` VARCHAR(100) UNIQUE,
    `image_url` VARCHAR(500),
    `is_active` BOOLEAN DEFAULT TRUE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE `products`;
