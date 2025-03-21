CREATE TABLE notification_templates (
    type TEXT PRIMARY KEY,
    template TEXT NOT NULL UNIQUE
);