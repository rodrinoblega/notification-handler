CREATE TABLE notifications (
       id TEXT PRIMARY KEY,
       user_id TEXT NOT NULL,
       content TEXT NOT NULL,
       status TEXT NOT NULL,
       retries INT NOT NULL DEFAULT 0,
       created_at TIMESTAMP NOT NULL,
       updated_at TIMESTAMP NOT NULL
);