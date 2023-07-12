import {hashClientData} from "../src/lib/hash";

it("should hash client data", () => {
    const hash = hashClientData({
        id:        "key1",
        originId:  "origin",
        document:  "{\"id\": 1, \"name\": \"Captain Jack Sparrow\"}",
        timestamp: "2020-10-29 11:05:37.0000000",
        publicKey: "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFckFUSFVPaVdUMmtBS2EwRlRsOVVLZFNGL3VocQo0Mi9NemRENlMwVDhPU1Rtd2pKS05yb1BMNWFaOE5YOWoyWm1xTmdKRUNQL250RlJ4d0RDWEJmYjd3PT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==",
    })

    expect(hash.toString('hex')).toEqual('24e693092614fb0d786164980b4f445bb20aea72d72ff6f79179f7ab53367f4a')
})

