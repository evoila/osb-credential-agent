package main

import (
	"fmt"
	"net/http"
	"osb-credential-agent/credhub"
)

func DeleteCredentials(w http.ResponseWriter, r *http.Request) {
	bindingId := r.FormValue("binding-id")
	credentials, err := credhub.GetCredentials(bindingId, Config)
	if err != nil {
		ErrHandler(w, "could not retrieve credentials", err)
	}
	fmt.Printf("%+v\n", credentials)

	err = credhub.DeleteCredentials(bindingId, Config)
	if err != nil {
		ErrHandler(w, "could not delete credentials", err)
	}
}
