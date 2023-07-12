package com.sap.icn.ledgerclientsample.hanaclient;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class LedgerEntry {
    private String id;
    private String originId;
    private String document;
    private String timestamp;
    private String publicKey;
    private String transactionId;
}
