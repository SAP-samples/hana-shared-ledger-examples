/*
 * @author: D023226
 */

package com.sap.icn.ledgerclientsample.keyutils;


import org.bouncycastle.util.io.pem.PemObject;
import org.bouncycastle.util.io.pem.PemReader;
import org.bouncycastle.util.io.pem.PemWriter;

import java.io.IOException;
import java.io.StringReader;
import java.io.StringWriter;
import java.security.*;
import java.security.spec.InvalidKeySpecException;
import java.security.spec.X509EncodedKeySpec;
import java.util.Base64;

/**
 * Provides a keypair to be used for inserting data in HANATABA
 */
public class KeyPairFactory {

    public static PublicKey pem2PublicKey(String publicPem) throws NoSuchAlgorithmException, IOException, InvalidKeySpecException {
        KeyFactory factory = KeyFactory.getInstance("EC");
        StringReader publicKeyReader = new StringReader(publicPem);
        PemObject pemPublicKeyObject = new PemReader(publicKeyReader).readPemObject();
        X509EncodedKeySpec publicKeySpec = new X509EncodedKeySpec(pemPublicKeyObject.getContent());
        return factory.generatePublic(publicKeySpec);
    }
    public static KeyPair provideKeyPair() {
        try {
            KeyPairGenerator keyPair = KeyPairGenerator.getInstance("EC");

            SecureRandom random = SecureRandom.getInstance("SHA1PRNG");
            keyPair.initialize(256, random);
            return keyPair.generateKeyPair();
        } catch (Exception e) {
            throw new RuntimeException("keypair could not be created", e);
        }
    }

    public static String publicKeyToPemBase64(PublicKey publicKey) {
        PemObject pemPublic = new PemObject("PUBLIC KEY", publicKey.getEncoded());
        String publicKeyPem = write(pemPublic);
        return Base64.getEncoder().encodeToString(publicKeyPem.getBytes());
    }

    private static String write(PemObject pemObject) {
        StringWriter writer = new StringWriter();
        PemWriter pemWriter = new PemWriter(writer);
        try {
            pemWriter.writeObject(pemObject);
            pemWriter.close();
            return writer.toString().replaceAll("\\r", "");
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

}
