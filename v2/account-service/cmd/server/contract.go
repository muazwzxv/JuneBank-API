package server

type IServer interface {
	Start() error
	Setup() error
	Routes() error
	Shutdown() error
}
