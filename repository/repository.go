package repository

import (
	"context"
	"database/sql"
	"github.com/azinudinachzab/br-be-test/model"
	"time"
)

type PgRepository struct {
	dbCore *sql.DB
}

func NewPgRepository(dbCore *sql.DB) *PgRepository {
	return &PgRepository{
		dbCore: dbCore,
	}
}

func (p *PgRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tx, err := p.dbCore.BeginTx(ctx, opts)
	if err != nil {
		return &sql.Tx{}, err
	}
	return tx, err
}

func (p *PgRepository) CommitTx(tx *sql.Tx) error {
	return tx.Commit()
}

func (p *PgRepository) RollbackTx(tx *sql.Tx) error {
	return tx.Rollback()
}

type Repository interface {
	StoreAccountValidation(ctx context.Context, accNumber, accName string) error
	GetAccountValidationData(ctx context.Context, accNumber string) (string, time.Time, error)

	StoreTransfers(ctx context.Context, req model.DoTransferReq, bankTrxID, status string) (uint64, error)
	UpdateTransfer(ctx context.Context, bankTrxID, status string) error

	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	CommitTx(tx *sql.Tx) error
	RollbackTx(tx *sql.Tx) error
}
