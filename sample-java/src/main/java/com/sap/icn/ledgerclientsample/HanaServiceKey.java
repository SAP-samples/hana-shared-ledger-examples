package com.sap.icn.ledgerclientsample;

import lombok.Getter;
import lombok.Setter;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;


@Configuration
@ConfigurationProperties(prefix = "hana-service-key")
@Getter
@Setter
public class HanaServiceKey {
    private String clientId;
    private String clientSecret;
    private String url;
    private String host;
    private String port;
    private String user;
    private String password;
    private String driver;
}
