# Shared Ledger on SAP HANA Example

## Disclaimer

This is a sample application. It is not intended to be used in production.
It should demonstrate how to use the Shared Ledger with typescript i.e.:
- ClientData format (e.g. timestamp, encodings, ...)
- correct hashing of the client data
- creating keypairs and signing client data
- write data to ledger using oauth2 and ledger service endpoint
- read data directly from hana and verify it locally

The example mainly is intended to provide code snippets as a starting point for your own implementation.
Also have a look at the tests for more details.

## Prerequisites

- node 16
- Embedded Cross-Company Workflow service instance setup on SAP BTP

## Installation

```
npm install
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


Run the app!
```
npm start
```

This should create a sample entry that is written using the ledger service api.
In the next step the entry is read directly from the hana database and verified locally.
