package com.sap.icn.ledgerclientsample;

import lombok.Getter;
import lombok.Setter;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;


@Configuration
@ConfigurationProperties(prefix = "ledger-service-key")
@Getter
@Setter
public class LedgerServiceKey {
    private String clientId;
    private String clientSecret;
    private String ledgerAPI;
    private String url;
}
