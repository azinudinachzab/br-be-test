package service

import (
	"context"
	"github.com/azinudinachzab/br-be-test/model"
	"github.com/azinudinachzab/br-be-test/pkg/bank"
	"github.com/azinudinachzab/br-be-test/repository"
	"github.com/go-playground/validator/v10"
)

type Dependency struct {
	Validator *validator.Validate
	Repo      repository.Repository
	Conf      model.Configuration
	BankPkg   bank.BankPkg
}

type AppService struct {
	validator *validator.Validate
	repo      repository.Repository
	conf      model.Configuration
	BankPkg   bank.BankPkg
}

func NewService(dep Dependency) Service {
	return &AppService{
		validator: dep.Validator,
		repo:      dep.Repo,
		conf:      dep.Conf,
		BankPkg:   dep.BankPkg,
	}
}

type Service interface {
	//Registration(ctx context.Context, req model.RegistrationRequest) error
	BankAccountValidation(ctx context.Context, req model.BankAccountValidationReq) (model.BankAccountValidationRes, error)
	DoTransfer(ctx context.Context, req model.DoTransferReq) (model.DoTransferRes, error)
	UpdateTransferStatus(ctx context.Context, req model.TransferCallbackReq) (model.TransferCallbackRes, error)
}
