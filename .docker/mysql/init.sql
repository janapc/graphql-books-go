CREATE DATABASE
    IF NOT EXISTS `graphql_books` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE IF NOT EXISTS books(
    id VARCHAR(150) NOT NULL,
    title VARCHAR(100),
    description TEXT(255),
    author_id VARCHAR(100),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS authors(
    id VARCHAR(150) NOT NULL,
    name VARCHAR(100),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY(id)
);
 
CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(150) NOT NULL,
    email VARCHAR(150) NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    is_deleted BOOLEAN GENERATED ALWAYS AS (IF(deleted_at IS NULL, 1, NULL)) VIRTUAL,
    PRIMARY KEY (id),
    CONSTRAINT email_is_deleted UNIQUE (email, is_deleted)
);
CREATE INDEX idx_deleted_at ON users(deleted_at);
