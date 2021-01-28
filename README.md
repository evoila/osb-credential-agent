# osb-credential-agent

Secure Service Delivery API

1. [Introduction](#introduction)
4. [Configuration](#configuration)
5. [Supported Services](#supported_services)
6. [Endpoints](#endpoints)

### Introduction
This project is small REST-API meant to be deployed with a service in a service broker ecosystem. 
It creates credentials, sets them up in a service and saves them in Credhub for credhub interpolation.
Read [here](https://github.com/cloudfoundry-incubator/credhub/blob/master/docs/secure-service-credentials.md)
for more information in secure service delivery.

### Configuration

For usage configure a UAA-endpoint and a Credhub instance in a config.yml file.
Set the path to the config file as an argument to the config file.

```yaml
uaa:
  uaa_endpoint: "https://35.196.32.64:8443"
  client_name: "credhub_client"
  client_secret: "secret"
credhub_endpoint: "https://localhost:9000"
skip_ssl_validation: true
client_identifier: osb-service-broker
service_identifier: service-id

port: 8082

service_handler: mongoDB

mongodb:
  username: admin
  password: adminPassword
  hosts:
    - localhost:27017
  database: serviceIdentifier
```

### Endpoints

To create a binding your Service Broker needs to fire the following REST call. 
```
 curl 'http://path/credentials' -i -X PUT \
-H 'Content-Type: application/json' \
-H 'Authorization: Bearer [some-token]' \
-d '{
"binding-id": "your-id"
}'
```
The agent will create the credentials like implemented in the selected service_handler
and save them in credhub. The agent will respond with: 
```json
{"credhub-ref":"/c/osb-service-broker/service-id/your-id/credentials-json"}
```

To delete a binding call 
```curl http://path/credentials?binding-id=your-id -i -X DELETE```

The agent will remove the credentials from your service and credhub.



