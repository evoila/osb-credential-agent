package main

import (
	"github.com/evoila/osb-credential-agent/config"
	"github.com/evoila/osb-credential-agent/credhub"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Meta struct {
	BindingGuid string `json:"binding-guid"`
}
type Response struct {
	CredhubRef string `json:"credhub-ref"`
}

var Config *config.Config
var Service ServiceHandler

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/credentials", GenerateAndInterpolateCredentials).Methods("PUT")
	router.Path("/credentials").Queries("binding-id", "{binding-id}").HandlerFunc(DeleteCredentials).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8082", router))
}

func ErrHandler(w http.ResponseWriter, response string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
}

func main() {
	path := os.Args[1]

	Service = DummyCredentialHandler{"Test"}

	Config = config.ReadConfig(path)
	credhub.TestConnection(Config)
	handleRequests()
}
