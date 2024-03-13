# work-handler
 Keeps track of available jobs and offers functionallity to take a job

## Setup 

Create an `.env` file, and enter the following.
```
ADDRESS=<address>
PORT=<port>

SYSTEM_NAME=<system name>

MONGO_DB_CONNECTION_STRING=<uri to connect to a mongoDB instance>


SERVER_MODE=<what security the server uses>

SERVICE_REGISTRY_ADDRESS=<service registry address>
SERVICE_REGISTRY_PORT=<service registry port>
SERVICE_REGISTRY_IMPLEMENTATION=<service registry implementation>

CERT_FILE_PATH=<path to cert .pem file>
KEY_FILE_PATH=<path to key .pem file>
TRUSTSTORE_FILE_PATH=<path to truststore .pem file>
AUTHENTICATION_INFO=<authentication info>

```