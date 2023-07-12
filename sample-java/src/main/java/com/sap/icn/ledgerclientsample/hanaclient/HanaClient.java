package com.sap.icn.ledgerclientsample.hanaclient;

import com.sap.icn.ledgerclientsample.HanaServiceKey;
import com.sap.icn.ledgerclientsample.LedgerServiceKey;
import org.springframework.stereotype.Component;

import java.sql.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Properties;

@Component
public class HanaClient {

    private final String jdbcUri;
    private final String ledgerTable;
    private final String driver;
    private final String user;
    private final String password;

    public HanaClient(HanaServiceKey serviceKey) {
        this.jdbcUri = String.format("jdbc:sap://%s:%s?encrypt=true", serviceKey.getHost(), serviceKey.getPort());
        //log.info("create " + config + "," + jdbcUri);
        this.ledgerTable = "\"Ledger\"";
        this.driver = serviceKey.getDriver();
        this.user = serviceKey.getUser();
        this.password = serviceKey.getPassword();
        setDriver();
    }


    public LedgerEntry getLedgerEntryByTransactionId(String transactionId) throws SQLException {
        Connection connection = getConnection();
        Statement stmt = connection.createStatement();
        String query = "SELECT ID,ORIGIN_ID,DOCUMENT,TIMESTAMP,PUBLICKEY,TRANSACTIONID from "
                + ledgerTable + " WHERE TRANSACTIONID='%s'".formatted(transactionId);
        ResultSet resultSet = stmt.executeQuery(query);
        if (resultSet.next()) {
            return toLedgerEntry(resultSet);
        }
        if (!connection.isClosed()) connection.close();
        throw new RuntimeException("No Ledger Entry found");
    }


    public List<LedgerEntry> getLedgerEntries() throws SQLException {
        List<LedgerEntry> ledgerEntries = new ArrayList<>();
        Connection connection = getConnection();
        Statement stmt = connection.createStatement();
        String query = "SELECT ID,ORIGIN_ID,DOCUMENT,TIMESTAMP,PUBLICKEY,TRANSACTIONID from "
                + ledgerTable + " ORDER BY TIMESTAMP DESC";
        ResultSet resultSet = stmt.executeQuery(query);
        while (resultSet.next()) {
            ledgerEntries.add(toLedgerEntry(resultSet));
        }
        if (!connection.isClosed()) connection.close();
        return ledgerEntries;
    }

    private LedgerEntry toLedgerEntry(ResultSet resultSet) throws SQLException {
        LedgerEntry ledgerEntry = new LedgerEntry();
        ledgerEntry.setId(resultSet.getString(1));
        ledgerEntry.setOriginId(resultSet.getString(2));
        ledgerEntry.setDocument(resultSet.getString(3));
        ledgerEntry.setTimestamp(resultSet.getString(4));
        ledgerEntry.setPublicKey(resultSet.getString(5));
        ledgerEntry.setTransactionId(resultSet.getString(6));
        return ledgerEntry;
    }

    private void setDriver() {
        try {
            Class.forName(driver);
        } catch (ClassNotFoundException e) {
            throw new RuntimeException(e);
        }
    }

    private Connection getConnection() throws SQLException {
        Properties connectionInfo = new Properties();
        connectionInfo.put("user", user);
        connectionInfo.put("password", password);
        return DriverManager.getConnection(jdbcUri, connectionInfo);
    }

}
