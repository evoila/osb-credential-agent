package services

type ServiceHandler interface {
	CreateCredentials()(map[string]interface{}, error)
	DeleteCredentials(json map[string]interface{}) error
}
