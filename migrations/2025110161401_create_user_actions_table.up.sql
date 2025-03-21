CREATE TABLE user_actions (
    id SERIAL PRIMARY KEY,
      user_id VARCHAR(255) NOT NULL ,
      action_type VARCHAR(50) NOT NULL,
    amount NUMERIC(20,6) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);