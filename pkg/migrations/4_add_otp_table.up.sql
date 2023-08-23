CREATE TABLE IF NOT EXISTS `otp` (
  `otp_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `otp_secret` varchar(60) NOT NULL DEFAULT '',
  `otp_auth_url` varchar(60) NOT NULL DEFAULT '',
  `otp_verified` TINYINT(1) DEFAULT 0,
  `otp_enabled` TINYINT(1) DEFAULT 0,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`otp_id`),
  KEY `user_id_idx` (`user_id`),
  CONSTRAINT `fk_otp_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;