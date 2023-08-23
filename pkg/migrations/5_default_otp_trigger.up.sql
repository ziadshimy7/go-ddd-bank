CREATE TRIGGER add_default_otp AFTER INSERT ON users
FOR EACH ROW
BEGIN
  INSERT INTO otp (user_id, otp_secret, otp_auth_url, otp_verified, otp_enabled, created_at, updated_at)
  VALUES (NEW.id, '', '', 0, 0, NOW(), NOW());
END;