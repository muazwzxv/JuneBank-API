package user

import (
	"account-service/module/user/app"
	"account-service/module/user/repository"
	"log"

	"github.com/jmoiron/sqlx"
)

func New(log log.Logger, db *sqlx.DB) *app.Module {
	userRepo := repository.NewUserRepository(db)

	return app.NewModule(
		log,
		userRepo,
	)
}
