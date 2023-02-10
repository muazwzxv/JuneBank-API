package ports

import (
	"account-service/app/pkg/core/domain"
	"net/http"
)

type IUserRepository interface {
	Save(domain.CreateUser) (*domain.User, error)
	GetByID(id uint64) (*domain.User, error)
	Update(id uint64, update domain.CreateUser) (*domain.User, error)
	DeleteByID(id uint64) (*domain.User, error)
	IsUserExist(id uint64) (bool, error)
}

type IUserService interface {
	Create(domain.CreateUser) error
	Get(id uint64) (*domain.User, error)
	Update(id uint64, update domain.CreateUser) (*domain.User, error)
	Delete(id uint64) (*domain.User, error)
	IsUserExist(id uint64) (bool, error)
}

type IUserHandlers interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
