package main

import (
	"encoding/json"
	"fmt"
	"github.com/evoila/osb-credential-agent/credhub"
	"io/ioutil"
	"net/http"
)

func GenerateAndInterpolateCredentials(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var meta Meta
	err := json.Unmarshal(body, &meta)
	if err != nil {
		panic(err)
	}

	values, err := Service.CreateCredentials()
	if err != nil {
		ErrHandler(w, "Script failed to create Credentials!", err)
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
