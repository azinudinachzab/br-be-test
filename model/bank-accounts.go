package model

type (
	BankAccountValidationReq struct {
		BeneficiaryAccountNumber string `json:"beneficiary_account_number" validate:"required"`
		BankName                 string `json:"bank_name" validate:"required"`
	}

	BankAccountValidationRes struct {
		BeneficiaryAccountName string `json:"beneficiary_account_name"`
	}
)
