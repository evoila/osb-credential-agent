package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/auth"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials/values"
	"code.cloudfoundry.org/credhub-cli/credhub/server"
	"github.com/evoila/osb-credential-agent/config"
)

func TestConnection(agentConfig *config.Config) (*server.Info, error) {
	client, err := credhubLogin(agentConfig)
	if err != nil {
		panic(err)
	}
	return client.Info()


}

func credhubLogin(config *config.Config) (*credhub.CredHub, error) {

	if config.Uaa.UaaCert != "" {
		credhub.CaCerts(config.Uaa.UaaCert)
	}

	if config.CredhubCert != "" {
		credhub.CaCerts(config.CredhubCert)
	}

	return credhub.New(config.CredhubEndpoint, credhub.SkipTLSValidation(config.SlipSSLValidation),
		credhub.Auth(auth.UaaClientCredentials(config.Uaa.ClientName, config.Uaa.ClientSecret)))
}

func SetCredentials(bindingId string, values values.JSON, config *config.Config) (credentials.Credential, error) {
	credHub, err := credhubLogin(config)
	if err != nil {
		return credentials.Credential{}, err
	}
	name := buildCredHubRef(bindingId, config)
	return credHub.SetCredential(name, "JSON", values)
}

func GetCredentials(bindingId string, config *config.Config) (credentials.JSON, error) {
	credHub, err := credhubLogin(config)
	if err != nil {
		return credentials.JSON{}, err
	}
	name := buildCredHubRef(bindingId, config)

	return credHub.GetLatestJSON(name)
}

func DeleteCredentials(bindingId string, config *config.Config) error {
	credHub, err := credhubLogin(config)
	if err != nil {
		panic(err)
	}
	name := buildCredHubRef(bindingId, config)
	return credHub.Delete(name)
}

func buildCredHubRef(bindingId string, config *config.Config) string {
	return "c/" + config.ClientIdentifier + "/" + config.ServiceIdentifier + "/" + bindingId + "/credentials-json"
}
