package credhub

import (
	"code.cloudfoundry.org/credhub-cli/credhub"
	"code.cloudfoundry.org/credhub-cli/credhub/auth"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials"
	"code.cloudfoundry.org/credhub-cli/credhub/credentials/values"
	"fmt"
	"osb-credential-agent/config"
)

func TestConnection(agentConfig *config.Config) {
	client, err := credhubLogin(agentConfig)
	if err != nil {
		panic(err)
	}
	info, err := client.Info()

	fmt.Printf("%+v\n", info)

}

func credhubLogin(config *config.Config) (*credhub.CredHub, error) {
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

func GetCredentials(bindingId string, config *config.Config) (credentials.Credential, error) {
	credHub, err := credhubLogin(config)
	if err != nil {
		return credentials.Credential{}, err
	}
	name := buildCredHubRef(bindingId, config)
	return credHub.GetLatestVersion(name)
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
