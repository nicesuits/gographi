CREATE DATABASE bank;

CREATE TABLE bank.accounts (
  id INT PRIMARY KEY,
  balance DECIMAL
);

INSERT INTO bank.accounts VALUES (1, 1000.50);`