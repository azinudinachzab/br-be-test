-- +goose Up
CREATE TABLE IF NOT EXISTS transfers (
    id                          serial                              PRIMARY KEY,
    source_account_number       varchar(255)                        NOT NULL,
    beneficiary_account_number  varchar(255)                        NOT NULL,
    beneficiary_account_name    varchar(255)                        NOT NULL,
    amount                      decimal(18,2)                       NOT NULL,
    bank_trx_id                 varchar(255)                        NOT NULL,
    status                      varchar(25)                         NOT NULL,
    created_at                  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at                  timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS transfers;
