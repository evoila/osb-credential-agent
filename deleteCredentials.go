package main

import (
	"fmt"
	"github.com/evoila/osb-credential-agent/credhub"
	"net/http"
)


func DeleteCredentials(w http.ResponseWriter, r *http.Request) {
	bindingId := r.FormValue("binding-id")
	credentials, err := credhub.GetCredentials(bindingId, Config)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		return
	}

	fmt.Printf("%+v\n", credentials)
	err = Service.DeleteCredentials(credentials.Value)
	if err != nil {
		ErrHandler(w, "Failed to remove credentials from Service!", err)
		return
	}

	err = credhub.DeleteCredentials(bindingId, Config)
	if err != nil {
		ErrHandler(w, "could not delete credentials", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
