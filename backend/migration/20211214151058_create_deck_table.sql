-- +goose Up
-- +goose StatementBegin
CREATE TABLE decks (
    id serial NOT NULL,
    uuid varchar(255),
    company_name varchar(255),
    images text[],
    created_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX deck_cp_idx ON decks (company_name);
CREATE UNIQUE INDEX deck_uuid_idx ON decks (uuid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE decks
-- +goose StatementEnd
