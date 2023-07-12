package com.sap.icn.ledgerclientsample;

import com.sap.icn.ledgerclientsample.keyutils.KeyPairFactory;
import com.sap.icn.ledgerclientsample.ledgerclient.LedgerDocument;

import java.security.KeyPair;
import java.time.LocalDateTime;

public class TestConstants  {
    public static final KeyPair KEY_PAIR = KeyPairFactory.provideKeyPair();
    public static final String ID = "id";
    public static final String ORIGIN_ID = "originId";
    public static final String DOCUMENT = "{\"property\":\"value\"}";
    public static final LedgerDocument LEDGER_DOCUMENT = LedgerDocument.builder()
            .id(ID).originId(ORIGIN_ID).document(DOCUMENT).build();

}
