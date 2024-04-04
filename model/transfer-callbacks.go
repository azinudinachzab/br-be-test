package model

type (
	TransferCallbackReq struct {
		TrxID  string `json:"transaction_id" validate:"required"`
		Status bool   `json:"status" validate:"required"`
	}

	TransferCallbackRes struct {
		Status string `json:"status"`
	}
)
