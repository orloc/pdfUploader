-- +goose Up
-- +goose StatementBegin
CREATE TABLE decks (
    id serial NOT NULL,
    company_name varchar(255),
    images text[],
    created_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY(id)
);
CREATE UNIQUE INDEX deck_cp_idx ON decks (company_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE decks
-- +goose StatementEnd
