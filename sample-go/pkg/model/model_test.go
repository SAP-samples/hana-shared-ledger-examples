package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"hana-shared-ledger-sample/pkg/crypto"
	"testing"
)

func TestEntry_SignEntryAndVerify(t *testing.T) {
	check := assert.New(t)

	publicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFckFUSFVPaVdUMmtBS2EwRlRsOVVLZFNGL3VocQo0Mi9NemRENlMwVDhPU1Rtd2pKS05yb1BMNWFaOE5YOWoyWm1xTmdKRUNQL250RlJ4d0RDWEJmYjd3PT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
	privateKey := "LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSUt4NHhPNk8yNGk0ZkRaRlR2bXIwUzRvQWNhZzJPcU5Lb05SMFJGU0pJZjVvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFckFUSFVPaVdUMmtBS2EwRlRsOVVLZFNGL3VocTQyL016ZEQ2UzBUOE9TVG13akpLTnJvUApMNWFaOE5YOWoyWm1xTmdKRUNQL250RlJ4d0RDWEJmYjd3PT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo="
	priv, _ := crypto.LoadPrivateKeyFromBase64(privateKey)

	entry := &ClientData{
		ID:        "key1",
		OriginId:  "origin",
		Document:  "{\"id\": 1, \"name\": \"Captain Jack Sparrow\"}",
		Timestamp: "2020-10-29 11:05:37.0000000",
		PublicKey: publicKey,
		Hash:      "",
		Signature: "",
	}

	err := entry.SignHash(priv)
	check.NoError(err)
	check.Equal("24e693092614fb0d786164980b4f445bb20aea72d72ff6f79179f7ab53367f4a", entry.Hash)
	check.NotEmpty(entry.Signature)

	check.NoError(entry.VerifySignature())
}

func TestClientData_VerifySignature_SignedWithP256k1(t *testing.T) {
	check := assert.New(t)

	// this entry is signed with an secp256k1 (ethereum) private key
	entry := &ClientData{
		ID:        "myid",
		OriginId:  "origin",
		Document:  "example",
		Timestamp: "2020-10-29 11:05:37.0000000",
		PublicKey: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFK0tYdHZkMWRsVkpSS2tabGZWTkhYVE93R2hYZgpyK3BaSTZpbmE0MW0vUFlwWjJVaXNMWXE5aytxaU9sWkFrOC9FZDZtR0pwOFBPODg3SUlBUUxxY0tnPT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==",
		Hash:      "8aad7f66e60820c947d1a9e56a89751df6b74c8182f142147ed8a1dcf8cd78f2",
		Signature: "MEYCIQDCWytG9KucK39Rilkw4Xwe4lH82Va2BZ7qVoeBAsGjxwIhALUD+ibYcXXP1AzkfjMfYOuXfbI0qzSmESie1eetYgnf",
	}

	check.NoError(entry.VerifySignature())
}

func TestLTEntry_SignEntryAndVerify(t *testing.T) {
	check := assert.New(t)

	testEntry := ClientData{
		ID:        "key1",
		OriginId:  "origin",
		Document:  "{\"id\": 1, \"name\": \"Captain Jack Sparrow\"}",
		Timestamp: "2020-10-29 11:05:37.0000000",
		PublicKey: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFL2ZRSlJ3ZGxzQXpuaFBaSWxFSnViQmY1eWlsKwpsNHZmNUdtSWdQVG96NUxQS2xsaDYyWW13aHFUOG9UaXM0Tk9yQjdNa2JvajNBZzhETUxqSEFlRjFBPT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==",
		Hash:      "b722df4f7e15be62116b3358a2e05bf3ad2cab396bd5f677f5324eb5a0d5e5e3",
		Signature: "MEYCIQCNWAwKPq2ijoSKsKiexBxVMKrKtTesAo0C0pe+YltfTAIhAKpBVjyuDGRFDK6fQceVtQBngTj1dgEAlYeTDlXy69cQ",
	}

	check.NoError(testEntry.VerifySignature())

	nodePubKeyEnc := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFMmN6dnNSN1dXRm1NVkRmYzRBU2pydnM1bWd1QQpCNU1sKzdWclpldmQ5U09iZmZCdzFTVGZJRjBoRzRmUURPMUczc2Y0R29KdEpqMmNkOTRiejc2Z0Z3PT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
	nodePrivKeyEnc := "LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSUxpZElvWVBDNTRPRXNXajg0azlBeEJDOWs5czlpWWlPTFdaZGhDK2tyZThvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFMmN6dnNSN1dXRm1NVkRmYzRBU2pydnM1bWd1QUI1TWwrN1ZyWmV2ZDlTT2JmZkJ3MVNUZgpJRjBoRzRmUURPMUczc2Y0R29KdEpqMmNkOTRiejc2Z0Z3PT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo="
	nodePrivKey, _ := crypto.LoadPrivateKeyFromBase64(nodePrivKeyEnc)

	ltEntry := &LedgerEntry{
		ClientData:    testEntry,
		TenantID:      "tenant123",
		NodePublicKey: nodePubKeyEnc,
		NodeTimestamp: "2020-08-11 10:11:13.1234567",
		TransactionID: "12345",
	}

	err := ltEntry.SignHash(nodePrivKey)
	check.NoError(err)
	check.NotEmpty(ltEntry.NodeHash)
	check.NotEmpty(ltEntry.NodeSignature)

	err = ltEntry.VerifySignature()
	check.NoError(err)
}

func TestBatch_Format(t *testing.T) {
	check := assert.New(t)
	batchEntries := ClientDataBatch{
		{
			ID:        "id1",
			OriginId:  "id2",
			Document:  "doc1",
			Timestamp: "2006-01-02 15:04:05.0000000",
			PublicKey: "pub1",
			Hash:      "hash1",
			Signature: "sig1",
		},
		{
			ID:        "id2",
			OriginId:  "id3",
			Document:  "doc2",
			Timestamp: "2006-01-02 15:04:06.0000000",
			PublicKey: "pub2",
			Hash:      "hash2",
			Signature: "sig2",
		},
	}

	jsonData, _ := json.Marshal(batchEntries)
	check.Equal("[{\"id\":\"id1\",\"originId\":\"id2\",\"document\":\"doc1\",\"timestamp\":\"2006-01-02 15:04:05.0000000\",\"publicKey\":\"pub1\",\"hash\":\"hash1\",\"signature\":\"sig1\"},{\"id\":\"id2\",\"originId\":\"id3\",\"document\":\"doc2\",\"timestamp\":\"2006-01-02 15:04:06.0000000\",\"publicKey\":\"pub2\",\"hash\":\"hash2\",\"signature\":\"sig2\"}]", string(jsonData))

	batchResponses := BatchResponse{
		Transaction{
			TransactionID: "txid1",
			Timestamp:     "2006-01-02 15:04:07.0000000",
		},
		Transaction{
			TransactionID: "txid2",
			Timestamp:     "2006-01-02 15:04:07.2000000",
		},
	}
	jsonData, _ = json.Marshal(batchResponses)
	check.Equal(`[{"transactionId":"txid1","timestamp":"2006-01-02 15:04:07.0000000"},{"transactionId":"txid2","timestamp":"2006-01-02 15:04:07.2000000"}]`, string(jsonData))
}
