package repository

import "account-service/module/user/domain"

func (r *UserRepository) Create(item *domain.CreateUser) error {
	query := `
		INSERT INTO 
			users (first_name, last_name, email)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(
		query,
		item.FirstName,
		item.LastName,
		item.Email,
	)

	if err != nil {
		return err
	}

	return nil
}
