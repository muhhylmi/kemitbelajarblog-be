-- 003_analytics_and_shares.sql
-- Replace likes with shares and add views for analytics

ALTER TABLE posts DROP COLUMN likes;
ALTER TABLE posts ADD COLUMN shares INTEGER NOT NULL DEFAULT 0;
ALTER TABLE posts ADD COLUMN views INTEGER NOT NULL DEFAULT 0;
