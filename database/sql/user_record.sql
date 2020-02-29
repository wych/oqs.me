CREATE TABLE IF NOT EXISTS `user_records` (
  id INT UNSIGNED AUTO_INCREMENT,
  ip VARCHAR(20),
  ua VARCHAR(255),
  forwarded_for VARCHAR(255),
  created_at TIMESTAMP,
  oqs_record_id INT UNSIGNED,
  FOREIGN KEY(oqs_record_id) REFERENCES oqs_records(id),
  PRIMARY KEY(id)
)