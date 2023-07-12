package com.sap.icn.ledgerclientsample.ledgerclient;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.security.SignatureException;
import java.util.List;

import static com.sap.icn.ledgerclientsample.TestConstants.*;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

@SpringBootTest
class LedgerClientTest {

    @Autowired
    LedgerClient cut;
    @Test
    void addDocument() {
        AddDocumentResponse addDocumentResponse = cut.addDocument(KEY_PAIR, LEDGER_DOCUMENT);
        assertNotNull(addDocumentResponse.getTimestamp());
        assertNotNull(addDocumentResponse.getTransactionId());
    }

    @Test
    void addDocuments() {
        List<AddDocumentResponse> addDocumentsResponse = cut.addDocuments(KEY_PAIR, List.of(LEDGER_DOCUMENT));
        assertEquals(1,addDocumentsResponse.size());
        assertNotNull(addDocumentsResponse.get(0).getTimestamp());
        assertNotNull(addDocumentsResponse.get(0).getTransactionId());
    }

}