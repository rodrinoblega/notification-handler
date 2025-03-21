CREATE TABLE notifications (
       id TEXT PRIMARY KEY,
       user_action_id INT NOT NULL REFERENCES user_actions(id) ON DELETE CASCADE,
       type TEXT NOT NULL REFERENCES notification_templates(type),
       status TEXT NOT NULL,
       content TEXT NOT NULL,
       retries INT NOT NULL DEFAULT 0,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);