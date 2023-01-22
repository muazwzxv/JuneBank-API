package user

import (
	"account-service/app/pkg/core/domain"
	"fmt"
)

func (r *userRepo) GetByID(id uint64) (*domain.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	var userdata user

	err := r.db.QueryRowx(
		query,
		id,
	).StructScan(&userdata)
	if err != nil {
		return &domain.User{}, err
	}

	return userdata.toDomain(), nil
}

func (r *userRepo) Save(data domain.CreateUser) error {
	query := `
		INSERT INTO 
			users (first_name, last_name, email)
			VALUES ($1, $2, $3)`

	_, err := r.db.Exec(
		query,
		data.FirstName,
		data.LastName,
		data.Email,
	)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	return nil
}
