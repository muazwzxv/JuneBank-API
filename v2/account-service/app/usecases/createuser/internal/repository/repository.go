package repository

import "errors"

func (r *repository) CreateUserData() error {
	query := `
		INSERT INTO 
		    users (first_name, last_name, email)
			values ($1, $2, $3)`

	// TODO: write a create user struct to pass into this function
	res := r.db.MustExec(query, "arg1", "arg2", "arg3")
	if rowsEffected, err := res.RowsAffected(); err != nil {
		if rowsEffected == 0 {
			return errors.New("0 rows effected")
		}
	}

	return nil
}
