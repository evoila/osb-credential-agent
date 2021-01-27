package main

type ServiceHandler interface {
	CreateCredentials()(map[string]interface{}, error)
	DeleteCredentials(map[string]interface{}) error
}
