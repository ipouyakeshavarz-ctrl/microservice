-- +migrate Up


CREATE TABLE `order_items` (
                               `id`          INT PRIMARY KEY AUTO_INCREMENT,
                               `order_id`    INT NOT NULL,
                               `product_id`  INT NOT NULL,
                               `quantity`    INT NOT NULL,
                               `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

                               CONSTRAINT `fk_order_items_order`
                                   FOREIGN KEY (`order_id`)
                                       REFERENCES `orders`(`id`)
                                       ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE `order_items`;
