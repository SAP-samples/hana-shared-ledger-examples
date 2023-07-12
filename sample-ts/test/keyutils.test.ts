import {ClientData, LedgerEntry, SignedClientData} from "../src/lib/model";
import {generatePemKeyPairBase64, signClientData, verifyClientData} from "../src/lib/key";

it('should correctly sign and verify an entry', () => {
    const pemKeyPair = generatePemKeyPairBase64();

    const clientData: ClientData = {
        id: "key1",
        originId: "origin",
        document: "{\"id\": 1, \"name\": \"Captain Jack Sparrow\"}",
        timestamp: "2020-10-29 11:05:37.0000000",
        publicKey: pemKeyPair.publicKey,
    }

    const signedClientData = signClientData(clientData, Buffer.from(pemKeyPair.privateKey, 'base64').toString('utf-8'));

    expect(verifyClientData(signedClientData)).toBeTruthy()
})

it('should correctly verify a ledger entry', () => {
    const clientData: SignedClientData = {
        id: "key1",
        originId: "origin",
        document: "{\"id\": 1, \"name\": \"Captain Jack Sparrow\"}",
        timestamp: "2020-10-29 11:05:37.0000000",
        publicKey: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFL2ZRSlJ3ZGxzQXpuaFBaSWxFSnViQmY1eWlsKwpsNHZmNUdtSWdQVG96NUxQS2xsaDYyWW13aHFUOG9UaXM0Tk9yQjdNa2JvajNBZzhETUxqSEFlRjFBPT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==",
        hash: "b722df4f7e15be62116b3358a2e05bf3ad2cab396bd5f677f5324eb5a0d5e5e3",
        signature: "MEYCIQCNWAwKPq2ijoSKsKiexBxVMKrKtTesAo0C0pe+YltfTAIhAKpBVjyuDGRFDK6fQceVtQBngTj1dgEAlYeTDlXy69cQ",
    }

    expect(verifyClientData(clientData)).toBeTruthy()

    const ledgerEntry: LedgerEntry = {
        ...clientData,
        SequenceNumber: 0,
        TransactionId: "12345",
        TenantId: "tenant123",
        NodePublicKey: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFMmN6dnNSN1dXRm1NVkRmYzRBU2pydnM1bWd1QQpCNU1sKzdWclpldmQ5U09iZmZCdzFTVGZJRjBoRzRmUURPMUczc2Y0R29KdEpqMmNkOTRiejc2Z0Z3PT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==",
        NodeTimestamp: "2020-08-11 10:11:13.1234567",
        NodeHash: "9dcf34a951fce4a754bbe283f3b2c0285dbc20050f16de359c9a84fcb3e6fdff",
        NodeSignature: "MEQCIF/SPbysK0xLGaye5i5A3v9idvJBBoXuzHSNyIRyo+zAAiBZStAFo2hUDYpf1aeLbvaIaNiwifqAb0D67lsa33tv0g=="
    }

    expect(verifyClientData(ledgerEntry)).toBeTruthy()
})