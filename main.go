package main

import (
	"fmt"
	"github.com/evoila/osb-credential-agent/config"
	"github.com/evoila/osb-credential-agent/credhub"
	"github.com/evoila/osb-credential-agent/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Meta struct {
	BindingGuid string `json:"binding-guid"`
}
type Response struct {
	CredhubRef string `json:"credhub-ref"`
}

var Config *config.Config
var Service services.ServiceHandler

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/credentials", GenerateAndInterpolateCredentials).Methods("PUT")
	router.Path("/credentials").Queries("binding-id", "{binding-id}").HandlerFunc(DeleteCredentials).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Config.Port), router))
}

func ErrHandler(w http.ResponseWriter, response string, error error) {
	fmt.Print(error)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(response))
//	panic(err)
}

func main() {
	path := os.Args[1]
	Config = config.ReadConfig(path)
	Service = &Config.MongoDB

	credhub.TestConnection(Config)
	handleRequests()
}
