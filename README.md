# go-secured-api

Secured api go template

- - - - 

## Necessary Technology Versions

Technology  | Version
------------- | -------------
Go | go1.14.3 linux/amd64
Docker | 18.09.6
docker-compose | 1.24.1

## Necessary System Configurations

Add to __/etc/hosts__ the hostname _keycloak_ to _127.0.0.1_ IP

    $ vim /etc/hosts
    
Should be  something like this:

    127.0.0.1   localhost keycloak

## Running Server

To run the chat server we create a docker container for it

    $ docker-compose up -d

## Get Token

    $ TOKEN=$(curl -s -X POST "http://keycloak:8080/auth/realms/go/protocol/openid-connect/token" -H "Content-Type: application/x-www-form-urlencoded" -d "grant_type=password" -d "username=user1" -d "password=user1" -d "client_id=go-sso" | jq -r '.access_token') && echo $TOKEN

## Configurations

### Client Environment Variables

| Name | Description | Default |
| ---- | ----------- | ------- |
| PORT | Server Port | 9090 |
| MONGO_URI | Mongo DB Connection URI | mongodb://mongo:mongo@mongo:27017 |
| DATABASE | Mongo DB Application Database Name | go-template |