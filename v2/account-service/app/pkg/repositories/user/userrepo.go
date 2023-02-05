package user

import (
	"account-service/app/pkg/core/domain"
	"fmt"
)

func (r *userRepo) GetByID(id uint64) (*domain.User, error) {
	query := `
		SELECT 
			* 
		FROM 
			users 
		WHERE 
			id = $1 
		AND 
			deleted_at = NULL`

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
			users 
				(first_name, last_name, email)
			VALUES 
				($1, $2, $3)`

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

func (r *userRepo) Update(id uint64, toUpdate domain.CreateUser) (*domain.User, error) {
	query := `
		UPDATE 
			users
		SET 
			first_name = $1,
			last_name = $2,
			email = $3,
			updated_at = NOW()
		WHERE 
			id = $4
		RETURNING *
	`

	var userdata user
	err := r.db.QueryRowx(
		query,
		toUpdate.FirstName,
		toUpdate.LastName,
		toUpdate.Email,
		id,
	).StructScan(&userdata)

	if err != nil {
		return &domain.User{}, err
	}

	return userdata.toDomain(), nil
}

func (r *userRepo) DeleteByID(id uint64) (*domain.User, error) {

	query := `
		UPDATE 
			users
		SET 
			deleted_at = NOW()
		WHERE
			id = $1
		RETURNING *
	`

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

func (r *userRepo) IsUserExist(id uint64) (bool, error) {

	isExistQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND deleted_at = NULL)`

	var isExist bool

	err := r.db.QueryRowx(
		isExistQuery,
		id,
	).Scan(&isExist)
	if err != nil {
		return isExist, fmt.Errorf("error querying user exists id=%d, %v", id, err)
	}

	return isExist, nil
}
