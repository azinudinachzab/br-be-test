package service

import (
	"context"
	"github.com/azinudinachzab/br-be-test/model"
	"github.com/azinudinachzab/br-be-test/pkg/errs"
	"log"
)

func (s *AppService) BankAccountValidation(ctx context.Context, req model.BankAccountValidationReq) (
	model.BankAccountValidationRes, error) {

	// validate request
	if err := s.validator.Struct(req); err != nil {
		err = errs.New(model.ECodeBadRequest, "invalid request")
		return model.BankAccountValidationRes{}, err
	}

	// call bank api
	resp, err := s.BankPkg.AccountValidation(ctx, req)
	if err != nil {
		log.Println(err)
		err = errs.New(model.ECodeInternal, "error when call bank API")
		return model.BankAccountValidationRes{}, err
	}

	// store validation to db
	if err := s.repo.StoreAccountValidation(ctx, req.BeneficiaryAccountNumber, resp.BeneficiaryAccountName); err != nil {
		log.Println(err)
		err = errs.New(model.ECodeInternal, "error when store validation data")
		return model.BankAccountValidationRes{}, err
	}

	return model.BankAccountValidationRes{
		BeneficiaryAccountName: resp.BeneficiaryAccountName,
	}, nil
}
