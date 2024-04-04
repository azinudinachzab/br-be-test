package service

import (
	"context"
	"github.com/azinudinachzab/br-be-test/model"
	"github.com/azinudinachzab/br-be-test/pkg/errs"
	"log"
	"strconv"
	"time"
)

func (s *AppService) DoTransfer(ctx context.Context, req model.DoTransferReq) (model.DoTransferRes, error) {

	// validate request
	if err := s.validator.Struct(req); err != nil {
		err = errs.New(model.ECodeBadRequest, "invalid request")
		return model.DoTransferRes{}, err
	}

	// check validation
	name, validatedAt, err := s.repo.GetAccountValidationData(ctx, req.BeneficiaryAccountNumber)
	if err != nil {
		log.Println(err)
		err = errs.New(model.ECodeInternal, "error when check validation data")
		return model.DoTransferRes{}, err
	}

	if time.Now().Sub(validatedAt) > 5*time.Minute {
		vld, err := s.BankAccountValidation(ctx, model.BankAccountValidationReq{
			BeneficiaryAccountNumber: req.BeneficiaryAccountNumber,
			BankName:                 req.BankName,
		})
		if err != nil {
			log.Println(err)
			err = errs.New(model.ECodeInternal, "error when check validation data")
			return model.DoTransferRes{}, err
		}
		name = vld.BeneficiaryAccountName
	}
	req.BeneficiaryAccountName = name

	// call bank api
	res, err := s.BankPkg.DoTransfer(ctx, req)
	if err != nil {
		log.Println(err)
		err = errs.New(model.ECodeInternal, "error when call bank api")
		return model.DoTransferRes{}, err
	}

	// store transfer to db
	id, err := s.repo.StoreTransfers(ctx, req, res.TrxID, res.Status)
	if err != nil {
		log.Println(err)
		err = errs.New(model.ECodeInternal, "error when store transfer data")
		return model.DoTransferRes{}, err
	}

	return model.DoTransferRes{
		TrxID:  strconv.FormatUint(id, 10),
		Status: res.Status,
	}, nil
}

func (s *AppService) UpdateTransferStatus(ctx context.Context, req model.TransferCallbackReq) (model.TransferCallbackRes, error) {
	// validate request
	if err := s.validator.Struct(req); err != nil {
		err = errs.New(model.ECodeBadRequest, "invalid request")
		return model.TransferCallbackRes{}, err
	}

	status := "success"
	if !req.Status {
		status = "failed"
	}
	if err := s.repo.UpdateTransfer(ctx, req.TrxID, status); err != nil {
		log.Println(err)
		err = errs.New(model.ECodeInternal, "error when update transfer data")
		return model.TransferCallbackRes{}, err
	}

	return model.TransferCallbackRes{
		Status: status,
	}, nil
}
