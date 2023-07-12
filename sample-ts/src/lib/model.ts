export type ClientData = {
    id: string;
    originId: string;
    document: string;
    timestamp: string; // must be in format "YYYY-MM-DD HH:mm:ss.SSSSSSS";
    publicKey: string;
}

export type SignedClientData = ClientData & {
    hash: string;
    signature: string;
}

export type LedgerEntry = SignedClientData & {
    SequenceNumber: number;
    TransactionId: string;
    TenantId: string;
    NodePublicKey: string;
    NodeTimestamp: string;
    NodeHash: string;
    NodeSignature: string;
}