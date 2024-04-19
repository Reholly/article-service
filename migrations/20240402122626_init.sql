-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS articles(
                                       id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                       title VARCHAR(255) NOT NULL,
                                       text VARCHAR(255) NOT NULL,
                                       publication_date TIME NOT NULL,
                                       author_username VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS tags(
                                   id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                   title VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS article_tag_pairs(
                                                article_id BIGINT NOT NULL,
                                                tag_id BIGINT NOT NULL
);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE articles
DROP TABLE tags