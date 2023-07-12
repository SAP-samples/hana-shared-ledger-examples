package com.sap.icn.ledgerclientsample.ledgerclient;

import lombok.Builder;
import lombok.Getter;

@Getter
@Builder
public class LedgerDocument {
    private String id;
    private String originId;
    private String document;
}
