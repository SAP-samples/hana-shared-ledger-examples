import {ClientData} from "./lib/model";
import {currentTime} from "./lib/time";
import {generatePemKeyPairBase64, signClientData, verifyLedgerEntry} from "./lib/key";
import {readEntryFromHana} from "./lib/hana";
import {writeDataToLedger} from "./lib/httpclient";


async function main() {
    // generate a new ec key pair (secp256r1) in PEM format base64 encoded
    const keyPairPemBase64 = generatePemKeyPairBase64();

    // prepare the client data to be written to the ledger
    const clientData: ClientData = {
        id: "key1",
        originId: "origin",
        document: "{\"id\": 1, \"name\": \"Captain Jack Sparrow\"}",
        timestamp: currentTime(), // make sure to use utc here, otherwise the ledger server will reject the entry
        publicKey: keyPairPemBase64.publicKey,
    }

    // sign the client data (also calculates the hash)
    const signedClientData = signClientData(clientData, Buffer.from(keyPairPemBase64.privateKey, 'base64').toString('utf-8'));

    console.log("Signed client data:\n" + JSON.stringify(signedClientData, null, 2))

    // write data to ledger using the api and get the transaction id
    const txId = await writeDataToLedger(signedClientData)

    console.log("\n\nTransaction id: " + txId)

    // read entry directly from hana using the transaction id
    const entry = await readEntryFromHana(txId)

    console.log("\n\nEntry from hana:\n" + JSON.stringify(entry, null, 2), "\n")

    if (!verifyLedgerEntry(entry)) {
        console.error("Signature verification failed")
    } else {
        console.log("Signature verification successful")
    }
}


main()