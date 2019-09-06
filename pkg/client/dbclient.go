package client

type DbClient interface {
	Execute () (string, error)
}