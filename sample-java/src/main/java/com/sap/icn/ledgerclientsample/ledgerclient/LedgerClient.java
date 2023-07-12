package com.sap.icn.ledgerclientsample.ledgerclient;

import com.sap.icn.ledgerclientsample.LedgerServiceKey;
import org.cloudfoundry.identity.client.UaaContext;
import org.cloudfoundry.identity.client.UaaContextFactory;
import org.cloudfoundry.identity.client.token.GrantType;
import org.cloudfoundry.identity.client.token.TokenRequest;
import org.cloudfoundry.identity.uaa.oauth.token.CompositeAccessToken;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Component;
import org.springframework.web.reactive.function.BodyInserters;
import org.springframework.web.reactive.function.client.ClientResponse;
import org.springframework.web.reactive.function.client.ExchangeFilterFunction;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Mono;

import java.net.URI;
import java.security.KeyPair;
import java.util.List;
import java.util.Objects;

import static org.springframework.http.HttpHeaders.AUTHORIZATION;
import static org.springframework.http.HttpHeaders.CONTENT_TYPE;

@Component
public class LedgerClient {

    private static final String DOCUMENTS_PATH = "/documents";
    private static final String DOCUMENTS_BATCH_PATH = DOCUMENTS_PATH + "/batch";
    private final String authenticationUrl;
    private final String clientId;
    private final String clientSecret;
    private final String ledgerApi;

    public LedgerClient(LedgerServiceKey ledgerServiceKey) {
        this.authenticationUrl = ledgerServiceKey.getUrl();
        this.clientId = ledgerServiceKey.getClientId();
        this.clientSecret = ledgerServiceKey.getClientSecret();
        this.ledgerApi = ledgerServiceKey.getLedgerAPI();
    }

    public AddDocumentResponse addDocument(KeyPair keyPair, LedgerDocument ledgerDocument) {
        AddDocumentRequest addDocumentRequest = AddDocumentRequestCreator.createRequest(keyPair, ledgerDocument);
        WebClient.ResponseSpec response = provideWebClient(ledgerApi)
                .post()
                .uri(uriBuilder -> uriBuilder.path(DOCUMENTS_PATH).build())
                .body(BodyInserters.fromValue(addDocumentRequest))
                .retrieve();
        return response.bodyToMono(AddDocumentResponse.class).block();
    }

    public List<AddDocumentResponse> addDocuments(KeyPair keyPair, List<LedgerDocument> ledgerDocuments) {
        List<AddDocumentRequest> addDocumentsRequest = ledgerDocuments.stream()
                .map(d -> AddDocumentRequestCreator.createRequest(keyPair, d)).toList();
        WebClient.ResponseSpec response = provideWebClient(ledgerApi)
                .post()
                .uri(uriBuilder -> uriBuilder.path(DOCUMENTS_BATCH_PATH).build())
                .body(BodyInserters.fromValue(addDocumentsRequest))
                .retrieve();
        return List.of(
                Objects.requireNonNull(response.bodyToMono(AddDocumentResponse[].class).block()));
    }
    public WebClient provideWebClient(String baseUrl) {
        ExchangeFilterFunction errorResponseFilter = ExchangeFilterFunction
                .ofResponseProcessor(LedgerClient::exchangeFilterResponseProcessor);
        return WebClient.builder().baseUrl(baseUrl)
                .defaultHeader(AUTHORIZATION,bearer())
                .defaultHeader(CONTENT_TYPE,MediaType.APPLICATION_JSON.toString())
                .filter(errorResponseFilter).build();
    }

    private static Mono<ClientResponse> exchangeFilterResponseProcessor(ClientResponse response) {
        HttpStatusCode status = response.statusCode();
        if (!status.is2xxSuccessful()) {
            return response.bodyToMono(String.class)
                    .flatMap(body -> Mono.error(new RuntimeException(body+response)));
        }
        return Mono.just(response);
    }

    private String getToken(URI authenticationUrl, String clientId, String clientSecret) {
        UaaContextFactory factory = UaaContextFactory.factory(authenticationUrl).authorizePath("/oauth/authorize")
                .tokenPath("/oauth/token");
        TokenRequest tokenRequest = factory.tokenRequest();
        tokenRequest.setGrantType(GrantType.CLIENT_CREDENTIALS);
        tokenRequest.setClientId(clientId);
        tokenRequest.setClientSecret(clientSecret);
        UaaContext xsuaaContext = factory.authenticate(tokenRequest);
        CompositeAccessToken accessToken = xsuaaContext.getToken();
        return accessToken.getValue();
    }

    private String bearer() {
        try {
            return "Bearer " + getToken(new URI(authenticationUrl), clientId, clientSecret);
        } catch (Exception e) {
            throw new RuntimeException("Error getting the bearer token", e);
        }
    }
}
