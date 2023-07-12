package com.sap.icn.ledgerclientsample.ledgerclient;

import lombok.Getter;
import lombok.NoArgsConstructor;

@NoArgsConstructor
@Getter
public class AddDocumentResponse {
    private String transactionId;
    private String timestamp;
}
