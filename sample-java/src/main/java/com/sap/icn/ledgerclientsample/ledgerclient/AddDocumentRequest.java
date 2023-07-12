/**
 * Represents the body for the POST request of the LT API
 */
package com.sap.icn.ledgerclientsample.ledgerclient;

import lombok.Builder;
import lombok.Getter;

@Builder
@Getter
public class AddDocumentRequest {
    private String document;
    private String hash;
    private String id;
    private String originId;
    private String signature;
    private String timestamp;
    private String publicKey;

    @Override
    public String toString() {
        return "AddDocumentRequest{" +
                "document=" + document +
                ", hash='" + hash + '\'' +
                ", id='" + id + '\'' +
                ", originId='" + originId + '\'' +
                ", signature='" + signature + '\'' +
                ", timestamp='" + timestamp + '\'' +
                ", publicKey='" + publicKey + '\'' +
                '}';
    }
}
