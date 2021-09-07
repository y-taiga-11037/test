DROP DATABASE IF EXISTS `mdtd_bootcamp`;
CREATE DATABASE `mdtd_bootcamp`;

use `mdtd_bootcamp`;

DROP TABLE IF EXISTS `shopping`;
CREATE TABLE shopping (
    shopping_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    shopping_day DATE NOT NULL UNIQUE KEY,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME 
);

DROP TABLE IF EXISTS `product`;
CREATE TABLE product (
    shopping_product_id INT AUTO_INCREMENT NOT NULL ,
    shopping_id INT NOT NULL ,
    product_name VARCHAR(100) NOT NULL ,
    price INT NOT NULL ,
    quantity INT NOT NULL ,
    purchase_flag TINYINT NOT NULL DEFAULT 0 ,
    PRIMARY KEY (shopping_product_id) ,
    UNIQUE KEY (shopping_id, product_name) ,
    FOREIGN KEY (shopping_id) REFERENCES shopping(shopping_id)
);

INSERT INTO shopping (shopping_id, shopping_day) VALUES (1, "2021-09-01");

INSERT INTO product (shopping_product_id, shopping_id, product_name, price, quantity) VALUES (1, 1, 'にんじん', 190, 1);




INSERT INTO shopping (shopping_id, shopping_day) VALUES (2, "2021-09-07");
INSERT INTO product (shopping_product_id, shopping_id, product_name, price, quantity) VALUES (2, 2, 'たまねぎ', 300, 1);
