package repository

import (
	"context"
	"database/sql"
	"time"
)

func (p *PgRepository) StoreAccountValidation(ctx context.Context, accNumber, accName string) error {
	q := `INSERT INTO bank_account_validations (beneficiary_account_number, beneficiary_account_name) VALUES ($1, $2);`
	if _, err := p.dbCore.ExecContext(ctx, q, &accNumber, &accName); err != nil {
		return err
	}

	return nil
}

func (p *PgRepository) GetAccountValidationData(ctx context.Context, accNumber string) (string, time.Time, error) {
	q := `SELECT beneficiary_account_name, created_at FROM bank_account_validations WHERE beneficiary_account_number=$1;`
	var (
		accName   sql.NullString
		createdAt sql.NullTime
	)
	if err := p.dbCore.QueryRowContext(ctx, q, accNumber).Scan(&accName, &createdAt); err != nil {
		return "", time.Time{}, err
	}

	return accName.String, createdAt.Time, nil
}
