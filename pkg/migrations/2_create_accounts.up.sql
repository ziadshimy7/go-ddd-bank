CREATE TABLE IF NOT EXISTS `accounts` (
  `accounts_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `account_number` varchar(255) NOT NULL,
  `expenses` float DEFAULT NULL,
  `income` float DEFAULT NULL,
  `balance` float DEFAULT NULL,
  `card_number` varchar(19) DEFAULT NULL,
  `expiration_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`accounts_id`),
  UNIQUE KEY `accounts_id_UNIQUE` (`accounts_id`),
  UNIQUE KEY `account_number_UNIQUE` (`account_number`),
  KEY `user_id_idx` (`user_id`),
  CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;