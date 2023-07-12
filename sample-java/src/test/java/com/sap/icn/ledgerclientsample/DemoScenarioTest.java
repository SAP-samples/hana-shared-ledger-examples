package com.sap.icn.ledgerclientsample;

import com.sap.icn.ledgerclientsample.hanaclient.HanaClient;
import com.sap.icn.ledgerclientsample.hanaclient.LedgerEntry;
import com.sap.icn.ledgerclientsample.ledgerclient.AddDocumentResponse;
import com.sap.icn.ledgerclientsample.ledgerclient.LedgerClient;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.security.SignatureException;
import java.sql.SQLException;

import static com.sap.icn.ledgerclientsample.TestConstants.KEY_PAIR;
import static com.sap.icn.ledgerclientsample.TestConstants.LEDGER_DOCUMENT;
import static org.junit.jupiter.api.Assertions.assertEquals;


@SpringBootTest
class DemoScenarioTest {
    @Autowired
    LedgerClient ledgerClient;
    @Autowired
    HanaClient hanaClient;

    @Test
    void e2e() throws SQLException {
        AddDocumentResponse addDocumentResponse = ledgerClient.addDocument(KEY_PAIR, LEDGER_DOCUMENT);
        LedgerEntry ledgerEntry = hanaClient.getLedgerEntryByTransactionId(addDocumentResponse.getTransactionId());
        assertEquals(LEDGER_DOCUMENT.getDocument(), ledgerEntry.getDocument());
    }

}