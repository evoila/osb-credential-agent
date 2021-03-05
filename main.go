package main

import (
	"code.cloudfoundry.org/credhub-cli/credhub/server"
	"encoding/json"
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
type CreateResponse struct {
	CredhubRef string `json:"credhub-ref"`
}

type HealthResponse struct {
	Status      string      `json:"status"`
	CredhubInfo server.Info `json:"credhub_info"`
}

var Config *config.Config
var Service services.ServiceHandler

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/credentials", GenerateAndInterpolateCredentials).Methods(http.MethodPut, http.MethodPost)
	router.Path("/credentials").Queries("binding-id", "{binding-id}").HandlerFunc(DeleteCredentials).Methods(http.MethodDelete)
	router.Path("/health").Methods(http.MethodGet).HandlerFunc(GetHealth)
	router.Path("/shutdown").Methods(http.MethodPost).HandlerFunc(Shutdown)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Config.Port), router))
}

func Shutdown(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("Shutting down!")
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}

func GetHealth(w http.ResponseWriter, _ *http.Request) {
	info, err := credhub.TestConnection(Config)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusFailedDependency)

		_, err2 := w.Write([]byte("Failed to connect to credhub"))
		if err2 != nil {
			panic(err2)
		}
		return
	}

	response := HealthResponse{
		Status:      "ready",
		CredhubInfo: *info,
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func ErrHandler(w http.ResponseWriter, response string, error error) {
	fmt.Print(error)
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(response))
	if err != nil {
		panic(err)
	}
}

func main() {
	path := os.Args[1]
	Config = config.ReadConfig(path)
	Service = Config.SelectServiceHandler()
	handleRequests()
}
