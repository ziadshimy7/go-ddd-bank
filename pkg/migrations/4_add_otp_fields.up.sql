ALTER TABLE `users` ADD otp_enabled TINYINT(1) DEFAULT 0;
ALTER TABLE `users` ADD otp_verified TINYINT(1) DEFAULT 0;
ALTER TABLE `users` ADD otp_secret VARCHAR(50) DEFAULT '';
ALTER TABLE `users` ADD otp_auth_url VARCHAR(255) DEFAULT '';