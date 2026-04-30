-- +migrate Up

CREATE TABLE `stores` (
                          `id`           INT PRIMARY KEY AUTO_INCREMENT,
                          `user_id`      INT NOT NULL,
                          `name`         VARCHAR(255) NOT NULL,
                          `description`  TEXT,
                          `logo_url`     VARCHAR(255),
                          `street`       VARCHAR(255),
                          `city`         VARCHAR(100),
                          `province`     VARCHAR(100),
                          `postal_code`  VARCHAR(20),
                          `phone_number` VARCHAR(50),
                          `is_active`    BOOLEAN DEFAULT TRUE,
                          `created_at`   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          `updated_at`   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE `stores`;
