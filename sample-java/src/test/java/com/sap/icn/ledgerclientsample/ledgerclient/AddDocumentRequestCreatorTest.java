package com.sap.icn.ledgerclientsample.ledgerclient;

import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.security.SignatureException;
import java.security.spec.InvalidKeySpecException;
import java.util.HexFormat;

import static org.junit.jupiter.api.Assertions.*;

class AddDocumentRequestCreatorTest {
    private static final String DATA= "somedata";
    private static final String HASH = "87d149cb424c0387656f211d2589fb5b1e16229921309e98588419ccca8a7362";

    private static final AddDocumentRequest DOCUMENT = AddDocumentRequest.builder()
            .id("key1")
            .originId("origin")
            .document("{\"id\": 1, \"name\": \"Captain Jack Sparrow\"}")
            .timestamp("2020-10-29 11:05:37.0000000")
            .publicKey("LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFL2ZRSlJ3ZGxzQXpuaFBaSWxFSnViQmY1eWlsKwpsNHZmNUdtSWdQVG96NUxQS2xsaDYyWW13aHFUOG9UaXM0Tk9yQjdNa2JvajNBZzhETUxqSEFlRjFBPT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==")
            .hash("b722df4f7e15be62116b3358a2e05bf3ad2cab396bd5f677f5324eb5a0d5e5e3")
            .signature("MEYCIQCNWAwKPq2ijoSKsKiexBxVMKrKtTesAo0C0pe+YltfTAIhAKpBVjyuDGRFDK6fQceVtQBngTj1dgEAlYeTDlXy69cQ").build();

    @Test
    void calculateHash() throws NoSuchAlgorithmException {
        String hash = AddDocumentRequestCreator.calculateHash(DATA.getBytes());
        assertEquals(HASH, hash);
    }

    @Test
    void verifySignature() throws NoSuchAlgorithmException, SignatureException, IOException, InvalidKeySpecException, InvalidKeyException {
        assertTrue(AddDocumentRequestCreator.verify(DOCUMENT));
    }
}