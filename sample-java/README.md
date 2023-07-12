# ledger-client-sample-java

## Prerequisites

- Embedded Cross-Company Workflow service instance setup on SAP BTP

## Run Local

Set the following application.yaml entries according to the eccwf service key on SAP BTP.

```
ledger-service-key:
    ledgerAPI: 
    clientid:
    clientsecret
    url:
```

Set the following application.yaml entries according to the HANA Schema service key on SAP BTP.

```
hana-service-key:
    clientid:
    clientsecret
    url:
    host:
    port:
    user: 
    password: 
    schema: 
    driver:
```

## Write Transactions
Run the LedgerClient tests.
## Read Transactions
Run the HanaClient tests.

