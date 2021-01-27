package main

import "github.com/dchest/uniuri"

type DummyCredentialHandler struct{
	name string
}

func (h DummyCredentialHandler) DeleteCredentials(i interface{}) error {
	// Remove the Credentials from the Service
	return nil
}

func (h DummyCredentialHandler) CreateCredentials() (map[string]interface{}, error) {
	values := map[string]interface{}{
		"datebase": uniuri.New(),
		"username": uniuri.New(),
		"password": uniuri.New(),
	}
	//Inject the credentials into the service here
	return values, nil
}