DROP DATABASE IF EXISTS `product_variant_service`;
CREATE DATABASE `product_variant_service`;

USE `product_variant_service`;

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
                           `id` int(11) NOT NULL AUTO_INCREMENT,
                           `name` varchar(500) NOT NULL,
                           `brand_name` varchar(500) NOT NULL,
                           `details` varchar(500) NOT NULL,
                           `image_url` varchar(500) NOT NULL,
                           `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           `deleted_at` timestamp NULL DEFAULT NULL,
                           PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `variants`;
CREATE TABLE `variants` (
                            `id` int(11) NOT NULL AUTO_INCREMENT,
                            `product_id` int(11) NOT NULL,
                            `variant_name` varchar(500) NOT NULL,
                            `variant_details` varchar(500) NOT NULL,
                            `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            FOREIGN KEY (product_id) REFERENCES products(id)
) DEFAULT CHARSET=utf8;