import * as hdb from '@sap/hana-client'
import {LedgerEntry} from "./model";


// connects to the hana / hdi container where the Ledger table is located
// reads the entry with the given transaction id
export function readEntryFromHana(transactionId: string) {
    const host = process.env.HANA_HOST
    if (!host) {
        throw new Error('HANA_HOST environment variable not set')
    }
    const hanaUser = process.env.HANA_USER
    if (!hanaUser) {
        throw new Error('HANA_USER environment variable not set')
    }
    const hanaPassword = process.env.HANA_PASSWORD
    if (!hanaPassword) {
        throw new Error('HANA_PASSWORD environment variable not set')
    }
    const lederTableName = 'Ledger'

    const conn = hdb.createConnection()

    return new Promise<LedgerEntry>((resolve, reject) => {
        conn.connect({
            host: host,
            port: 443,
            user: hanaUser,
            password: hanaPassword,
        }, function (err) {
            if (err) reject(err)

            // do not use this query in production, it may lead to sql injection
            conn.exec(`SELECT *
                       FROM "${hanaUser}"."${lederTableName}"
                       WHERE TRANSACTIONID = ?`, [transactionId], (err, res: any) => {
                if (err) reject(err)

                if (res.length === 0) {
                    reject(new Error('No entry found'))
                } else {
                    resolve({
                        SequenceNumber: res[0].SEQUENCENUMBER,
                        id: res[0].ID,
                        originId: res[0].ORIGIN_ID,
                        document: res[0].DOCUMENT,
                        timestamp: res[0].TIMESTAMP,
                        publicKey: res[0].PUBLICKEY,
                        hash: res[0].HASH,
                        signature: res[0].SIGNATURE,
                        TransactionId: res[0].TRANSACTIONID,
                        TenantId: res[0].TENANTID,
                        NodePublicKey: res[0].NODEPUBLICKEY,
                        NodeTimestamp: res[0].NODETIMESTAMP,
                        NodeHash: res[0].NODEHASH,
                        NodeSignature: res[0].NODESIGNATURE,
                    })
                }
            })
        })
    })
}