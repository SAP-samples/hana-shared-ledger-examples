package model

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hana-shared-ledger-sample/pkg/crypto"
)

// ClientDataBatch allows client data batch write requests to improve overall write performance.
type ClientDataBatch []ClientData

// BatchResponse is returned as response on a successful batch write request.
// The Response holds the Transactions responses in the same order it was received in the BatchEntries.
type BatchResponse []Transaction

// ClientData holds the user data that is completely created and signed on client side.
// After signing, the data cannot be modified without invalidating the hash or the signature.
// ClientData will be written into the ledger after verified (signature, hash, timestamp) and enriched with
// additional information by the server side ledger service (node-signature, node-timestamp, etc.)
type ClientData struct {
	ID        string `json:"id"`
	OriginId  string `json:"originId"`
	Document  string `json:"document"`
	Timestamp string `json:"timestamp"`
	PublicKey string `json:"publicKey"`
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
}

// SignHash signs the calculated hash of the client data (using the byte representation, not the hex encoded field).
// Sets the Hash field (hex) and the Sig field (asn.1 base64).
func (e *ClientData) SignHash(privateKey interface{}) error {
	hash := e.CalculateHash()

	sig, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return fmt.Errorf("could not sign client data: %w", err)
	}

	e.Hash = hex.EncodeToString(hash)
	e.Signature = base64.StdEncoding.EncodeToString(sig)
	return nil
}

// VerifySignature verifies the signature using the provided public key and the calculated hash.
func (e *ClientData) VerifySignature() error {
	hash := e.CalculateHash()

	sig, err := crypto.NewSigFromBase64(e.Signature)
	if err != nil {
		return fmt.Errorf("could not get signature: %w", err)
	}

	pubKey, err := crypto.LoadPublicKeyFromBase64(e.PublicKey)
	if err != nil {
		return fmt.Errorf("could not get public key: %w", err)
	}

	err = crypto.Verify(pubKey, hash, sig)
	if err != nil {
		return fmt.Errorf("verification of client data failed: %w", err)
	}

	return nil
}

// CalculateHash returns the sha256 hash of the entry.
func (e *ClientData) CalculateHash() []byte {
	// Seperators are needed!
	// otherwise i.e H(ID = "" and OriginId = "1") would equal H(ID = "1" and OriginID = "")
	return crypto.HashStrings(
		base64.StdEncoding.EncodeToString([]byte(e.ID)),
		".",
		base64.StdEncoding.EncodeToString([]byte(e.OriginId)),
		".",
		base64.StdEncoding.EncodeToString([]byte(e.Document)),
		".",
		base64.StdEncoding.EncodeToString([]byte(e.Timestamp)),
		".",
		base64.StdEncoding.EncodeToString([]byte(e.PublicKey)),
	)
}

// ClientDataResponse is a client data struct enriched with information after it was successfully written to the ledger.
type ClientDataResponse struct {
	ClientData
	Transaction Transaction `json:"transaction"`
}

type Transaction struct {
	TransactionID string `json:"transactionId"`
	Timestamp     string `json:"timestamp"`
}

// LedgerEntry holds the client data enriched with additional fields provided by the ledger service.
type LedgerEntry struct {
	SequenceNumber int64 `json:"sequenceNumber"`
	ClientData
	TransactionID string `json:"transactionId"`
	TenantID      string `json:"tenantId"`
	NodePublicKey string `json:"nodePublicKey"`
	NodeTimestamp string `json:"nodeTimestamp"`
	NodeHash      string `json:"nodeHash"`
	NodeSignature string `json:"nodeSignature"`
}

// SignHash signs the calculated hash of the LedgerEntry (using the byte representation, not the hex encoded field).
// Sets the NodeHash field (hex) and the NodeSignature field (asn.1 base64).
// NodeHash is calculated: sha256(entry-hash, transactionID, nodeTimestamp, nodePubKey)
func (e *LedgerEntry) SignHash(privateKey interface{}) error {
	hash := e.CalculateHash()

	sig, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return fmt.Errorf("could not sign entry: %w", err)
	}

	e.NodeHash = hex.EncodeToString(hash)
	e.NodeSignature = base64.StdEncoding.EncodeToString(sig)
	return nil
}

// VerifySignature verifies the signature of the ledger entry using the public key and the hash.
func (e *LedgerEntry) VerifySignature() error {
	nodeHash := e.CalculateHash()

	nodeSig, err := crypto.NewSigFromBase64(e.NodeSignature)
	if err != nil {
		return fmt.Errorf("could not get node signature: %w", err)
	}

	nodePubKey, err := crypto.LoadPublicKeyFromBase64(e.NodePublicKey)
	if err != nil {
		return fmt.Errorf("could not get node public key: %w", err)
	}

	err = crypto.Verify(nodePubKey, nodeHash, nodeSig)
	if err != nil {
		return fmt.Errorf("verification of ledger entry failed: %w", err)
	}

	return nil
}

// CalculateHash calculates the hash of the LedgerEntry using the hash of the user entry and the additional fields.
// Returns the sha256 hash.
func (e *LedgerEntry) CalculateHash() []byte {
	return crypto.HashStrings(
		base64.StdEncoding.EncodeToString([]byte(e.Hash)),
		".",
		base64.StdEncoding.EncodeToString([]byte(e.TenantID)),
		".",
		base64.StdEncoding.EncodeToString([]byte(e.TransactionID)),
		".",
		base64.StdEncoding.EncodeToString([]byte(e.NodeTimestamp)),
		".",
		base64.StdEncoding.EncodeToString([]byte(e.NodePublicKey)),
	)
}

// CommittedLedgerEntry contains a ledger entry that is committed to the ledger.
// This means it may already be mined inside a block, but not necessarily.
// If the transaction is not yet mined, BlockTransactionID is empty, otherwise it holds the ID of the block.
type CommittedLedgerEntry struct {
	LedgerEntry
	BlockTransactionID string `json:"blockTransactionId"`
}
