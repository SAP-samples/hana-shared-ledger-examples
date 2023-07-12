import * as crypto from 'crypto';
import {ClientData, LedgerEntry} from "./model";

// calculate the sha256 hash of multiple Buffers
function hash256(...args: Buffer[]): Buffer {
    const hash = crypto.createHash('sha256');
    for (const arg of args) {
        hash.update(arg);
    }
    return hash.digest();
}

// calculate the sha256 hash of multiple strings
function hashStrings(...args: string[]): Buffer {
    const buffers = args.map(arg => Buffer.from(arg));
    return hash256(...buffers);
}

// calculate the sha256 hash of a ClientData object
export function hashClientData(clientData: ClientData): Buffer {
    return hashStrings(prepareClientDataHash(clientData));
}

export function hashLedgerEntry(ledgerEntry: LedgerEntry): Buffer {
    return hashStrings(prepareLedgerEntryHash(ledgerEntry));
}


// concatenate the ClientData object to a single string to be hashed
// strings are separated with a dot to prevent collisions
// encode strings to base64 to prevent the strings containing dots
// this function is used because crypto.sign() does the sha256 hashing internally
export function prepareClientDataHash(clientData: ClientData): string {
    return Buffer.from(clientData.id || "").toString('base64') +
        "." +
        Buffer.from(clientData.originId || "").toString('base64') +
        "." +
        Buffer.from(clientData.document || "").toString('base64') +
        "." +
        Buffer.from(clientData.timestamp || "").toString('base64') +
        "." +
        Buffer.from(clientData.publicKey || "").toString('base64');
}

// concatenate the relevant fields of the LedgerEntry object to a single string to be hashed
// strings are separated with a dot to prevent collisions
// encode strings to base64 to prevent the strings containing dots
// this function is used because crypto.sign() does the sha256 hashing internally
export function prepareLedgerEntryHash(ledgerEntry: LedgerEntry): string {
    return Buffer.from(ledgerEntry.hash || "").toString('base64') +
        "." +
        Buffer.from(ledgerEntry.TenantId || "").toString('base64') +
        "." +
        Buffer.from(ledgerEntry.TransactionId || "").toString('base64') +
        "." +
        Buffer.from(ledgerEntry.NodeTimestamp || "").toString('base64') +
        "." +
        Buffer.from(ledgerEntry.NodePublicKey || "").toString('base64');
}