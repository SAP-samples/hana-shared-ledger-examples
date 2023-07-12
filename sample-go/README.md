# Shared Ledger on SAP HANA Example

## Prerequisites

- Installation of GO 1.19 or newer
- Generated Private/Public Keypair
- Embedded Cross-Company Workflow service instance setup on SAP BTP

## Installation

```
go mod tidy
```

### Generate Private / Public Keypair with OpenSSL

```
openssl ecparam -name secp256r1 -genkey -noout -out private-key.pem
```

```
openssl ec -in private-key.pem -pubout | base64
```


## Run Local

Set the following environment variables according to the eccwf service key on SAP BTP.

```
BASE_URL
CLIENT_ID
CLIENT_SECRET
TOKEN_URL
```

Set the following environment variables according to the HANA Schema service key on SAP BTP.

```
HANA_HOST
HANA_USER
HANA_PASSWORD
```

In addition, set the Private / Public Keypair generated in the previous step into the environment
```
PRIVATE_KEY
PUBLIC_KEY
```

Run the app!
```
go run main.go
```