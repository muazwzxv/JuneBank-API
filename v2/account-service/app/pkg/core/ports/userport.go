package ports

import (
	"account-service/app/pkg/core/domain"
	"net/http"
)

type IUserRepository interface {
	Save(domain.CreateUser) error
	GetByID(id uint64) (*domain.User, error)
	Update(id uint64, update domain.CreateUser) (*domain.User, error)
}

type IUserService interface {
	Create(domain.CreateUser) error
	Get(id uint64) (*domain.User, error)
	Update(id uint64, update domain.CreateUser) (*domain.User, error)
}

type IUserHandlers interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}
