package com.sap.icn.ledgerclientsample.hanaclient;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.sql.SQLException;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest
class HanaClientTest {

    @Autowired
    HanaClient cut;

    @Test
    void getLedgerEntries() throws SQLException {
        List<LedgerEntry> ledgerEntries = cut.getLedgerEntries();
        assertNotNull(ledgerEntries);
    }

    @Test
    void getLedgerEntryByTransactionHash() throws SQLException {
        String TRANSACTION_ID = getTransactionId();
        LedgerEntry ledgerEntry = cut.getLedgerEntryByTransactionId(TRANSACTION_ID);
        assertEquals(TRANSACTION_ID, ledgerEntry.getTransactionId());
    }

    private String getTransactionId() throws SQLException {
        return cut.getLedgerEntries().get(0).getTransactionId();
    }


}