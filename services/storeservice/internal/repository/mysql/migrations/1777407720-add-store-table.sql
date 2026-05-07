-- +migrate Up

CREATE TABLE `stores` (
                        `id`                  INT AUTO_INCREMENT PRIMARY KEY,
                        `user_id`             INT NOT NULL,
                        `name`                VARCHAR(255) NOT NULL,

                        `description`         TEXT NOT NULL,
                        `logo_url`            VARCHAR(255) NOT NULL DEFAULT '',

                        `street`              VARCHAR(255) NOT NULL DEFAULT '',
                        `city`                VARCHAR(100) NOT NULL DEFAULT '',
                        `province`            VARCHAR(100) NOT NULL DEFAULT '',
                        `postal_code`         VARCHAR(20) NOT NULL DEFAULT '',
                        `address_description` TEXT NOT NULL,

                        `phone_number`        VARCHAR(50) NOT NULL DEFAULT '',

                        `is_active`           BOOLEAN NOT NULL DEFAULT TRUE,

                        `created_at`          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `updated_at`          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


UPDATE stores SET description = '' WHERE description IS NULL;
UPDATE stores SET address_description = '' WHERE address_description IS NULL;

-- +migrate Down
DROP TABLE IF EXISTS stores;
