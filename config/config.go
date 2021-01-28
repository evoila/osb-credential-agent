package config

import (
	"github.com/evoila/osb-credential-agent/services"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Uaa               Uaa                            `yaml:"uaa"`
	CredhubEndpoint   string                         `yaml:"credhub_endpoint"`
	SlipSSLValidation bool                           `yaml:"skip_ssl_validation"`
	ClientIdentifier  string                         `yaml:"client_identifier"`
	ServiceIdentifier string                         `yaml:"service_identifier"`
	Port              int                            `yaml:"port"`
	MongoDB           services.MongoDBServiceHandler `yaml:"mongodb"`
}

type Uaa struct {
	ClientName   string `yaml:"client_name"`
	ClientSecret string `yaml:"client_secret"`
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
