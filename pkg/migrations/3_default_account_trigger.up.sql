CREATE TRIGGER add_default_account AFTER INSERT ON users
FOR EACH ROW
BEGIN
  DECLARE account_number VARCHAR(10);
  SET account_number = SUBSTRING(MD5(RAND()), 1, 10);

  INSERT INTO accounts (user_id, account_number, expenses, income, balance)
  VALUES (NEW.id, account_number, 0.0, 0.0, 0.0);
END;