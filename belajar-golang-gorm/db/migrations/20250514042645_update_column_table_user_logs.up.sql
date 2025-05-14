ALTER TABLE user_logs
  MODIFY created_at BIGINT NOT NULL;

ALTER TABLE user_logs
  MODIFY updated_at BIGINT NOT NULL;
