# osb-credential-agent

Secure Service Delivery API

1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
3. [Usage](#usage)

### Introduction
This project is small REST-API meant to be deployed with a service in a service broker ecosystem. 
It creates credentials, sets them up in a service and saves them in Credhub for credhub interpolation.
Read [here](https://github.com/cloudfoundry-incubator/credhub/blob/master/docs/secure-service-credentials.md)
for more information in secure service delivery.

### Prerequisites
A Service Broker and Credhub.

### Usage
1. Implement the interface ServiceHandler according to how credentials are handled in your service.
2. When binding to your service with your service-broker
   use the endpoints by this api accordingly. It only needs the binding-id.


