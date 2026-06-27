-- ==========================================
-- Database
-- ==========================================

CREATE DATABASE graphql_demo;

-- Connect to the database
-- \c graphql_demo


-- ==========================================
-- Users
-- ==========================================

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);


-- ==========================================
-- Posts
-- ==========================================

CREATE TABLE posts (
    id UUID PRIMARY KEY,

    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,

    author_id UUID NOT NULL,

    CONSTRAINT fk_posts_author
        FOREIGN KEY (author_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);


-- ==========================================
-- Comments
-- ==========================================

CREATE TABLE comments (
    id UUID PRIMARY KEY,

    text TEXT NOT NULL,

    author_id UUID NOT NULL,
    post_id UUID NOT NULL,

    CONSTRAINT fk_comments_author
        FOREIGN KEY (author_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_comments_post
        FOREIGN KEY (post_id)
        REFERENCES posts(id)
        ON DELETE CASCADE
);


-- ==========================================
-- Todos
-- ==========================================

CREATE TABLE todos (
    id UUID PRIMARY KEY,

    text TEXT NOT NULL,

    done BOOLEAN NOT NULL DEFAULT FALSE,

    user_id UUID NOT NULL,

    CONSTRAINT fk_todos_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);


-- ==========================================
-- Indexes
-- ==========================================

CREATE INDEX idx_posts_author_id
    ON posts(author_id);

CREATE INDEX idx_comments_author_id
    ON comments(author_id);

CREATE INDEX idx_comments_post_id
    ON comments(post_id);

CREATE INDEX idx_todos_user_id
    ON todos(user_id);