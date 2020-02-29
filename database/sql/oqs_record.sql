CREATE TABLE IF NOT EXISTS `oqs_records` (
  id INT UNSIGNED,
  challenge CHAR(2),
  real_url VARCHAR(255),
  PRIMARY KEY(id)
)