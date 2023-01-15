package api

type userApp interface {
	CreateUser()
}

type API struct {
	user userApp
}

func New(user userApp) *API {
	return &API{
		user: user,
	}
}
