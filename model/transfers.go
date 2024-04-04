package model

type (
	DoTransferReq struct {
		SourceAccountNumber      string  `json:"source_account_number" validate:"required"`
		BeneficiaryAccountNumber string  `json:"beneficiary_account_number" validate:"required"`
		BankName                 string  `json:"bank_name" validate:"required"`
		Amount                   float64 `json:"amount" validate:"gt=0,required"`
		BeneficiaryAccountName   string
	}

	DoTransferRes struct {
		TrxID  string `json:"transaction_id"`
		Status string `json:"status"`
	}
)
