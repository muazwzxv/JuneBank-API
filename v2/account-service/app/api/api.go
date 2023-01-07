package api

import "account-service/app/api/user"

type RestHandlers struct {
	User IUserAPI
}

func New() RestHandlers {
	return RestHandlers{
		User: user.New(),
	}
}
