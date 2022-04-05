package service

import "junebank/repository"

type accountService struct {
	accountRepository repository.AccountRepository
}
type AccountService interface {
	GetAll()
	GetByID(id uint)
	Create()
	DeleteByID(id uint)
}

func InitializeAccountService(repository repository.AccountRepository) AccountService {
	return &accountService{accountRepository: repository}
}

func (a *accountService) GetAll() {

}

func (a *accountService) GetByID(id uint) {

}

func (a *accountService) Create() {

}

func (a *accountService) DeleteByID(id uint) {

}
