package postgres

type Adapter struct{}

func New() *Adapter {
	return &Adapter{}
}
