package config

import (
	"errors"
	"github.com/evoila/osb-credential-agent/services"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type Config struct {
	Uaa               Uaa                            `yaml:"uaa"`
	CredhubEndpoint   string                         `yaml:"credhub_endpoint"`
	CredhubCert       string                         `yaml:"credhub_cert"`
	SlipSSLValidation bool                           `yaml:"skip_ssl_validation"`
	ClientIdentifier  string                         `yaml:"client_identifier"`
	ServiceIdentifier string                         `yaml:"service_identifier"`
	Port              int                            `yaml:"port"`
	ServiceHandler    string                         `yaml:"service_handler"`
	MongoDB           services.MongoDBServiceHandler `yaml:"mongodb"`
}

type Uaa struct {
	ClientName   string `yaml:"client_name"`
	ClientSecret string `yaml:"client_secret"`
	UaaCert      string `yaml:"uaa_cert"`
}

func ReadConfig(path string) *Config {
	agentConfig := Config{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &agentConfig)
	if err != nil {
		panic(err)
	}

	return &agentConfig
}

func (c Config) SelectServiceHandler() services.ServiceHandler {
	switch strings.ToLower(c.ServiceHandler) {
	case "dummy":
		return services.DummyCredentialHandler{}
	case "mongodb":
		return &c.MongoDB
	default:
		panic(errors.New("no valid service_handler provided. select dummy or mongoDB"))
	}
}
