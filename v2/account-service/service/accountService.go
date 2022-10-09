package service

import (
	"github.com/gofiber/fiber/v2"
	"junebank/v2/account-service/application/util"
	"junebank/v2/account-service/constants"
	"junebank/v2/account-service/entity"
	"junebank/v2/account-service/repository"
	"junebank/v2/account-service/request"
)

type accountService struct {
	accountRepository repository.IAccountRepository
}

type IAccountService interface {
	Create(ctx *fiber.Ctx, req *request.CreateAccount) error
}

func CreateAccountService(repository repository.IAccountRepository) IAccountService {
	return &accountService{accountRepository: repository}
}

func (s *accountService) Create(ctx *fiber.Ctx, req *request.CreateAccount) error {

	account := &entity.Account{
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Occupation: req.Occupation,
		DOB:        util.GetDOB(req.Year, req.Month, req.Day),
	}

	// Validate the payload
	err := account.Validate()
	if err != nil {
		return util.BadRequest(ctx, constants.BAD_REQUEST, err)
	}

	// check existing email
	if s.accountRepository.IsExistByEmail(account.Email) {
		return util.BadRequest(ctx, constants.EMAIL_IS_USED, nil)
	}

	// check existing phone
	if s.accountRepository.IsExistByPhone(account.Phone) {
		return util.BadRequest(ctx, constants.PHONE_IS_USED, nil)
	}

	// store in db
	if err := s.accountRepository.Create(account); err != nil {
		return util.BadRequest(ctx, constants.SOMETHING_WRONG_HAPPENED, err)
	}

	//TODO: submit event account created

	return util.Created(ctx, constants.ACCOUNT_CREATED, account)
}
