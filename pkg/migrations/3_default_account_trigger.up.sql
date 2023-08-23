CREATE TRIGGER add_default_account AFTER INSERT ON users
FOR EACH ROW
BEGIN
  DECLARE account_number VARCHAR(10);
  DECLARE expiration_date DATE;
  DECLARE card_number VARCHAR(19);


  SET account_number = SUBSTRING(MD5(RAND()), 1, 10);
  SET expiration_date = DATE_ADD(CURDATE(), INTERVAL 3 YEAR);
  SET card_number = CONCAT(
      SUBSTRING(MD5(RAND()), 1, 4), ' ', 
      SUBSTRING(MD5(RAND()), 5, 4), ' ', 
      SUBSTRING(MD5(RAND()), 9, 4), ' ', 
      SUBSTRING(MD5(RAND()), 13, 4)
  );

  INSERT INTO accounts (user_id, account_number, expenses, income, balance,card_number,expiration_date,created_at)
  VALUES (NEW.id, account_number, 0.0, 0.0, 0.0,card_number, expiration_date,NOW());
END;