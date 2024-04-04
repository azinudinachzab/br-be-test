package repository

import (
	"context"
	"github.com/azinudinachzab/br-be-test/model"
)

func (p *PgRepository) StoreTransfers(ctx context.Context, req model.DoTransferReq, bankTrxID, status string) (uint64, error) {
	q := `INSERT INTO transfers (source_account_number, beneficiary_account_number, amount, beneficiary_account_name,
                       bank_trx_id, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	var lastinsertedid uint64
	if err := p.dbCore.QueryRowContext(ctx, q, &req.SourceAccountNumber, &req.BeneficiaryAccountNumber, &req.Amount,
		&req.BeneficiaryAccountName, bankTrxID, status).Scan(&lastinsertedid); err != nil {
		return 0, err
	}

	return lastinsertedid, nil
}

func (p *PgRepository) UpdateTransfer(ctx context.Context, bankTrxID, status string) error {
	q := `UPDATE transfers SET status=$1 WHERE bank_trx_id=$2;`
	if _, err := p.dbCore.ExecContext(ctx, q, &status, &bankTrxID); err != nil {
		return err
	}

	return nil
}
