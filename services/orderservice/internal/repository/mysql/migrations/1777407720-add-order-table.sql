-- +migrate Up

CREATE TABLE `orders` (
                          `id`           INT PRIMARY KEY AUTO_INCREMENT,
                          `checkout_id`  VARCHAR(255) NOT NULL UNIQUE,
                          `user_id`      INT NOT NULL,
                          `status`       VARCHAR(50) NOT NULL,
                          `created_at`   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE `orders`;
