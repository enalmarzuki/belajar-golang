CREATE TABLE user_like_product(
  user_id VARCHAR(100) NOT NULL,
  product_id VARCHAR(100) NOT NULL,
  PRIMARY KEY (user_id, product_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
) ENGINE = INNODB;
