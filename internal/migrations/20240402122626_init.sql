-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS article(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    text VARCHAR(255) NOT NULL,
    publication_date TIME NOT NULL,
    author_username VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS tag(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS article_tag(
    article_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL
);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE article;
DROP TABLE tag;
DROP TABLE article_tag;