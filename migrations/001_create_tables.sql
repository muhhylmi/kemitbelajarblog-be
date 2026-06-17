-- 001_create_tables.sql
-- Creates the core tables for the Kemitbelajar blog

CREATE TABLE IF NOT EXISTS authors (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name       VARCHAR(255) NOT NULL,
    avatar     TEXT NOT NULL DEFAULT '',
    role       VARCHAR(100) NOT NULL DEFAULT 'Author',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS posts (
    id           VARCHAR(255) PRIMARY KEY,
    title        VARCHAR(500) NOT NULL,
    summary      TEXT NOT NULL DEFAULT '',
    content      TEXT NOT NULL DEFAULT '',
    category     VARCHAR(100) NOT NULL DEFAULT '',
    author_id    UUID NOT NULL REFERENCES authors(id),
    read_time    VARCHAR(50) NOT NULL DEFAULT '',
    published_at VARCHAR(100) NOT NULL DEFAULT '',
    image        TEXT NOT NULL DEFAULT '',
    image_alt    TEXT NOT NULL DEFAULT '',
    likes        INTEGER NOT NULL DEFAULT 0,
    is_featured  BOOLEAN NOT NULL DEFAULT FALSE,
    status       VARCHAR(20) NOT NULL DEFAULT 'draft' CHECK (status IN ('published', 'draft')),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_posts_status ON posts(status);
CREATE INDEX IF NOT EXISTS idx_posts_category ON posts(category);
CREATE INDEX IF NOT EXISTS idx_posts_author_id ON posts(author_id);
