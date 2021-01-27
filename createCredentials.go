package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"osb-credential-agent/credhub"
)

func GenerateAndInterpolateCredentials(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var meta Meta
	err := json.Unmarshal(body, &meta)
	if err != nil {
		panic(err)
	}

	//values, err := Service.CreateCredentials()
	//	if err != nil {
	//		ErrHandler(w, "Script failed to create Credentials!", err)
	//	}

	values := map[string]interface{}{
		"username": "test",
		"password": "test",
	}

	credentials, err := credhub.SetCredentials(meta.BindingGuid, values, Config)
	if err != nil {
		ErrHandler(w, "Failed to save Credentials in Credhub", err)
	}

	fmt.Printf("%+v\n", credentials)

	response := Response{
		CredhubRef: credentials.Name,
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}
