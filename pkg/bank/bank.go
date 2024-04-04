package bank

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/azinudinachzab/br-be-test/model"
	"net/http"
)

type (
	BankPkg interface {
		AccountValidation(ctx context.Context, req model.BankAccountValidationReq) (model.BankAccountValidationRes, error)
		DoTransfer(ctx context.Context, req model.DoTransferReq) (model.DoTransferRes, error)
	}

	Bank struct {
		conf   model.Configuration
		client *http.Client
	}
)

func NewBankPkg(cfg model.Configuration) BankPkg {
	return &Bank{
		conf:   cfg,
		client: &http.Client{},
	}
}

func (b *Bank) AccountValidation(ctx context.Context, req model.BankAccountValidationReq) (model.BankAccountValidationRes, error) {
	r, err := json.Marshal(req)
	if err != nil {
		return model.BankAccountValidationRes{}, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, b.conf.BankAPIURL+"/account-validations", bytes.NewBuffer(r))
	if err != nil {
		return model.BankAccountValidationRes{}, err
	}
	resp, err := b.client.Do(httpReq)
	if err != nil {
		return model.BankAccountValidationRes{}, err
	}

	defer resp.Body.Close()
	var res model.BankAccountValidationRes

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return model.BankAccountValidationRes{}, err
	}

	return model.BankAccountValidationRes{
		BeneficiaryAccountName: res.BeneficiaryAccountName,
	}, nil
}

func (b *Bank) DoTransfer(ctx context.Context, req model.DoTransferReq) (model.DoTransferRes, error) {
	r, err := json.Marshal(req)
	if err != nil {
		return model.DoTransferRes{}, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, b.conf.BankAPIURL+"/transfers", bytes.NewBuffer(r))
	if err != nil {
		return model.DoTransferRes{}, err
	}
	resp, err := b.client.Do(httpReq)
	if err != nil {
		return model.DoTransferRes{}, err
	}

	defer resp.Body.Close()
	var res struct {
		TrxID  string `json:"transaction_id"`
		Status bool   `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return model.DoTransferRes{}, err
	}
	flag := ""
	if res.Status == true {
		flag = "pending"
	} else {
		flag = "failed"
	}
	return model.DoTransferRes{
		TrxID:  res.TrxID,
		Status: flag,
	}, nil
}
