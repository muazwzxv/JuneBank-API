package createuser

import (
	"account-service/app/usecases"
	"account-service/app/usecases/createuser/internal/repository"
	"github.com/jmoiron/sqlx"
)

type UseCase struct {
	repo repository.IRepository
}

func New(db *sqlx.DB) usecases.CreateUserUseCase {
	return UseCase{repository.New(db)}
}
