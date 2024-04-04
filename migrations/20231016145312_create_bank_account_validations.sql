-- +goose Up
CREATE TABLE IF NOT EXISTS bank_account_validations (
    id                          serial                              PRIMARY KEY,
    beneficiary_account_number  varchar(255)                        NOT NULL,
    beneficiary_account_name    varchar(255)                        NOT NULL,
    created_at                  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at                  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
    );

-- +goose Down
DROP TABLE IF EXISTS bank_account_validations;
