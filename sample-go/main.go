package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SAP/go-hdb/driver"
	"hana-shared-ledger-sample/pkg/crypto"
	"hana-shared-ledger-sample/pkg/httpclient"
	"hana-shared-ledger-sample/pkg/model"
	"hana-shared-ledger-sample/pkg/times"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	// Get ECCWF Service Instance Credentials from environment
	baseUrl := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	tokenURL := os.Getenv("TOKEN_URL")
	hanaHost := os.Getenv("HANA_HOST")
	hanaUser := os.Getenv("HANA_USER")
	hanaPassword := os.Getenv("HANA_PASSWORD")
	ledgerTableName := "Ledger"

	privateKey := os.Getenv("PRIVATE_KEY")
	publicKey := os.Getenv("PUBLIC_KEY")

	/**
	 * Write data to the ledger
	 */

	// sample data - key/value
	e := model.ClientData{
		ID:        "someDocumentID",
		OriginId:  "someOriginID",
		Document:  "{\"someKey\": \"someValue\"}",
		Timestamp: times.ToString(time.Now().UTC()),
		PublicKey: publicKey,
	}

	// load private key
	pk, err := crypto.LoadPrivateKeyFromBase64(privateKey)
	if err != nil {
		panic(err)
	}
	// sign client data
	if err := e.SignHash(pk); err != nil {
		panic(err)
	}
	// send request to shared ledger API
	client := httpclient.NewOAuthClient(baseUrl, clientID, clientSecret, tokenURL)
	r := client.NewRequest(context.Background())
	r.SetBody(e)
	r.SetDoNotParseResponse(true)
	restyResponse, err := r.Post("/documents")
	if err != nil {
		panic(err)
	}
	resp := restyResponse.RawResponse
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, resp.Body)
		if err != nil {
			panic(err)
		}
		panic(errors.New(buf.String()))
	}
	var tx model.Transaction
	err = json.NewDecoder(resp.Body).Decode(&tx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully written to the ledger. Transaction: %s \n", tx.TransactionID)

	/**
	 * Read data from the ledger
	 */

	// connect to hana db
	connector := driver.NewBasicAuthConnector(fmt.Sprintf("%s:%v", hanaHost, "443"), hanaUser, hanaPassword)
	// use "user" as default schema
	connector.SetDefaultSchema(hanaUser)
	connector.SetTLSConfig(&tls.Config{
		RootCAs:            nil,
		ServerName:         hanaHost,
		InsecureSkipVerify: false,
	})
	db := sql.OpenDB(connector)
	if err := db.PingContext(context.Background()); err != nil {
		panic(err)
	}
	rows, err := db.Query(fmt.Sprintf("SELECT TRANSACTIONID FROM \"%s\".\"%s\" WHERE TRANSACTIONID=?", hanaUser, ledgerTableName), tx.TransactionID)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for rows.Next() {
		var row model.LedgerEntry
		if err := rows.Scan(&row.TransactionID); err != nil {
			panic(err)
		}
		fmt.Println("Transaction info: " + row.TransactionID)
	}

	/**
	 * Batch write to ledger
	 */
	// TODO batch implementation open

}
