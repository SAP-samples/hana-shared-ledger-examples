import crypto from 'crypto';
import {ClientData, LedgerEntry, SignedClientData} from "./model";
import {prepareClientDataHash, hashClientData, hashLedgerEntry, prepareLedgerEntryHash} from "./hash";

type PemKeyPair = {
    privateKey: string;
    publicKey: string;
}

// generate a new ec key pair (prime256v1) in PEM format base64 encoded
export function generatePemKeyPairBase64(): PemKeyPair {
    // generate a new ec key pair (prime256v1) in PEM format
    const ecdh = crypto.generateKeyPairSync('ec', {
        namedCurve: 'prime256v1',// same as 'p256' or 'secp256r1'  may also use 'secp256k1' that is ethereum curve
        publicKeyEncoding: {type: 'spki', format: 'pem'},
        privateKeyEncoding: {type: 'sec1', format: 'pem'} // same format used by openssl
    });

    return {
        privateKey: Buffer.from(ecdh.privateKey).toString('base64'),
        publicKey: Buffer.from(ecdh.publicKey).toString('base64')
    }
}

// sign the client data (this also sets the calculated hash)
// see: concatenateStrings to learn how the data is concatenated for hashing
export function signClientData(clientData: ClientData, privateKeyPem: string): SignedClientData {
    const hash = hashClientData(clientData)
    const signature = crypto.sign('sha256', Buffer.from(prepareClientDataHash(clientData)), privateKeyPem);

    // return the signed client data
    // hash is returned as hex string
    // signature is returned as base64 string
    return {
        ...clientData,
        hash: hash.toString('hex'),
        signature: signature.toString('base64'),
    }
}

// verify the signature of the client data
// also recalculates the hash, so we do not rely on the hash value in the client data
export function verifyClientData(clientData: SignedClientData): boolean {
    if (!clientData.signature || !clientData.publicKey) {
        return false;
    }

    const hash = hashClientData(clientData)
    if (hash.toString('hex') !== clientData.hash) {
        return false;
    }

    return crypto.verify('sha256',
        Buffer.from(prepareClientDataHash(clientData)),
        Buffer.from(clientData.publicKey, 'base64'),
        Buffer.from(clientData.signature, 'base64')
    )
}

export function verifyLedgerEntry(ledgerEntry: LedgerEntry): boolean {
    const clientDataOk = verifyClientData(ledgerEntry);
    if (!clientDataOk) {
        return false;
    }

    const hash = hashLedgerEntry(ledgerEntry)
    if (hash.toString('hex') !== ledgerEntry.NodeHash) {
        return false;
    }

    // TODO: you may also add additional checks on the timestamp, if it is within a certain range, or bigger than the previous one, etc.

    return crypto.verify('sha256',
        Buffer.from(prepareLedgerEntryHash(ledgerEntry)),
        Buffer.from(ledgerEntry.NodePublicKey, 'base64'),
        Buffer.from(ledgerEntry.NodeSignature, 'base64')
    )
}




