/*
 * @author: D023226
 */

package com.sap.icn.ledgerclientsample.ledgerclient;

import com.sap.icn.ledgerclientsample.keyutils.KeyPairFactory;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.security.*;
import java.security.spec.InvalidKeySpecException;
import java.time.Instant;
import java.time.ZoneId;
import java.time.ZoneOffset;
import java.time.format.DateTimeFormatter;
import java.util.Base64;
import java.util.HexFormat;


public class AddDocumentRequestCreator {
    private static final Logger log = LoggerFactory.getLogger(AddDocumentRequestCreator.class);

    private static String createSignature(byte[] toBeHashed, PrivateKey privateKey) throws InvalidKeyException, SignatureException, NoSuchAlgorithmException {
        Signature signature = Signature.getInstance("SHA256withECDSA");
        signature.initSign(privateKey);
        signature.update(toBeHashed);
        byte[] signed = signature.sign();
        return toBase64(signed);
    }

    private static boolean verifySignature(String signatureBase64, String publicKeyBase64, byte[] toBeHashed) throws InvalidKeyException, SignatureException, NoSuchAlgorithmException, IOException, InvalidKeySpecException {
        String pemPublicKey = new String(Base64.getDecoder().decode(publicKeyBase64));
        PublicKey publicKey = KeyPairFactory.pem2PublicKey(pemPublicKey);
        Signature signatureVerify = Signature.getInstance("SHA256withECDSA");
        signatureVerify.initVerify(publicKey);
        signatureVerify.update(toBeHashed);
        return signatureVerify.verify(Base64.getDecoder().decode(signatureBase64));
    }


    public static AddDocumentRequest createRequest(KeyPair keyPair, LedgerDocument ledgerDocument) {
        try {
            String timestamp = getCurrentTimestamp();
            String publicKey = KeyPairFactory.publicKeyToPemBase64(keyPair.getPublic());
            byte[] toBeHashed = provideToBeHashed(ledgerDocument, timestamp, publicKey);
            String hash = calculateHash(toBeHashed);
            String signature = createSignature(toBeHashed, keyPair.getPrivate());
            AddDocumentRequest request = AddDocumentRequest.builder()
                    .document(ledgerDocument.getDocument())
                    .hash(hash)
                    .id(ledgerDocument.getId())
                    .originId(ledgerDocument.getOriginId())
                    .signature(signature)
                    .timestamp(timestamp)
                    .publicKey(publicKey).build();
            log.info(request.toString());
            return request;
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    public static String getCurrentTimestamp() {
        DateTimeFormatter DATE_TIME_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss.SSSSSSS")
                .withZone(ZoneId.from(ZoneOffset.UTC));
        return DATE_TIME_FORMATTER.format(Instant.now());
    }

    public static String toBase64(String s) {
        return toBase64(s.getBytes(StandardCharsets.UTF_8));
    }

    public static String toBase64(byte[] b) {
        return Base64.getEncoder().encodeToString(b);
    }

    public static byte[] provideToBeHashed(LedgerDocument ledgerDocument, String timestamp, String publicKey) {
        return (toBase64(ledgerDocument.getId()) + '.'
                + toBase64(ledgerDocument.getOriginId()) + '.'
                + toBase64(ledgerDocument.getDocument()) + '.'
                + toBase64(timestamp) + '.'
                + toBase64(publicKey)).getBytes();
    }

    public static String calculateHash(byte[] toBeHashed) throws NoSuchAlgorithmException {
        MessageDigest digest = MessageDigest.getInstance("SHA-256");
        return HexFormat.of().formatHex(digest.digest(toBeHashed));
    }

    public static boolean verify(AddDocumentRequest document) throws NoSuchAlgorithmException, SignatureException, IOException, InvalidKeySpecException, InvalidKeyException {
        byte[] toBeHashed = provideToBeHashed(LedgerDocument.builder()
                .id(document.getId())
                .originId(document.getOriginId())
                .document(document.getDocument()).build(), document.getTimestamp(), document.getPublicKey());
        String hash = calculateHash(toBeHashed);
        if (!hash.equals(document.getHash())) return false;
        return verifySignature(document.getSignature(), document.getPublicKey(),toBeHashed);
    }
}
