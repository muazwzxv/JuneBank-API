package ports

import (
	"account-service/app/pkg/core/domain"
	"net/http"
)

type IUserRepository interface {
	Save(domain.CreateUser) error
	GetByID(id uint64) (*domain.User, error)
}

type IUserService interface {
	Create(domain.CreateUser) (*domain.User, error)
	Get(id uint64) (*domain.User, error)
}

type IUserHandlers interface {
	Get(w http.ResponseWriter, r *http.Request)
}
