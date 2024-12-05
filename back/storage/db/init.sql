CREATE DATABASE IF NOT EXISTS `goticket`;

USE `goticket`;

CREATE TABLE IF NOT EXISTS `users` (
  `id` int unsigned AUTO_INCREMENT NOT NULL,
  `name` char(20) NOT NULL,
  `sex` enum('Male','Female') NOT NULL,
  `password` varchar(20) NOT NULL,
  `phone` char(15) NOT NULL,
  `create_date` datetime NOT NULL,
  `id_number` char(18) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `name_index` (`name`),
  KEY `phone_index` (`phone`),
  KEY `id_number_idex` (`id_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `admins` (
  `id` int unsigned NOT NULL,
  FOREIGN KEY (`id`) REFERENCES users(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `trains` (
  `id` smallint unsigned AUTO_INCREMENT NOT NULL,
  `train_type` enum('G','D','K') NOT NULL,
  `max_capacity` smallint unsigned NOT NULL,
  `seats` json NOT NULL,
  `avg_speed` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `routes` (
  `id` smallint unsigned AUTO_INCREMENT NOT NULL,
  `price_pk` decimal(9, 2)  NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `stations` (
  `id` smallint unsigned AUTO_INCREMENT NOT NULL,
  `name` varchar(32) NOT NULL,
  `postion` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `station_index` (`name`),
  KEY `postion_index` (`postion`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS `route_station` (
  `route_id` smallint unsigned NOT NULL,
  `station_id` smallint unsigned NOT NULL,
  `distance_from_start` double NOT NULL,
  PRIMARY KEY (`route_id`,`station_id`),
  FOREIGN KEY (`route_id`) REFERENCES routes(`id`),
  FOREIGN KEY (`station_id`) REFERENCES stations(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `train_numbers` (
  `id` int unsigned AUTO_INCREMENT NOT NULL,
  `code` varchar(4) NOT NULL,
  `status` enum('Online','Offline') NOT NULL,
  `available_seats` smallint unsigned NOT NULL,
  `start_time` datetime NOT NULL,
  `train_id` smallint unsigned NOT NULL,
  `route_id` smallint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`train_id`) REFERENCES trains(`id`),
  FOREIGN KEY (`route_id`) REFERENCES routes(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `tickets` (
  `id` int unsigned AUTO_INCREMENT NOT NULL,
  `start_station` varchar(32) NOT NULL,
  `end_station` varchar(32) NOT NULL,
  `departure_time` datetime NOT NULL,
  `arrival_time` datetime NOT NULL,
  `carriage` tinyint unsigned NOT NULL,
  `row` tinyint unsigned NOT NULL,
  `seat` tinyint unsigned NOT NULL,
  `status` enum('Used','Unused','Refunded') NOT NULL,
  `user_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES users(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `orders` (
  `id` int unsigned AUTO_INCREMENT NOT NULL,
  `pay_method` enum('WeChatPay','Alipay') NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `status` enum('Paid','Unpaid') NOT NULL,
  `user_id` int unsigned NOT NULL,
  `ticket_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES users(`id`),
  FOREIGN KEY (`ticket_id`) REFERENCES tickets(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
