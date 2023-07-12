import axios from 'axios';
import {SignedClientData} from "./model";

async function getAccessToken() {
    // get tokenEndpoint from environment variable
    let tokenUrl = process.env.TOKEN_URL
    const clientId = process.env.CLIENT_ID
    const clientSecret = process.env.CLIENT_SECRET
    if (!tokenUrl) {
        throw new Error('TOKEN_URL environment variable not set');
    }
    if (!clientId) {
        throw new Error('CLIENT_ID environment variable not set');
    }
    if (!clientSecret) {
        throw new Error('CLIENT_SECRET environment variable not set');
    }
    if (!tokenUrl.endsWith('/oauth/token')) {
        tokenUrl = `${tokenUrl}/oauth/token`;
    }

    try {
        const response = await axios({
            url: tokenUrl,
            method: 'post',
            params: {
                grant_type: 'client_credentials',
                client_id: clientId,
                client_secret: clientSecret,
            }
        })

        return response.data.access_token;
    } catch (err) {
        console.error('Failed to get access token:', err);
        throw err;
    }
}

export async function writeDataToLedger(signedClientData: SignedClientData) {
    const baseUrl = process.env.BASE_URL
    if (!baseUrl) {
        throw new Error('BASE_URL environment variable not set');
    }

    const token = await getAccessToken();

    try {

        const res = await axios({
            url: `${baseUrl}/documents`,
            method: 'post',
            headers: {
                'Authorization': `Bearer ${token}`,
            },
            data: signedClientData,
        });

        if (res.status !== 201) {
            throw new Error(`Failed to write data to ledger: ${res.statusText}`);
        }

        return res.data.transactionId;
    } catch (err) {
        console.error('Failed to write data to ledger:', err);
        throw err;
    }
}

