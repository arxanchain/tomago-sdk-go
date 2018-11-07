
# Status
[![Build Status](https://travis-ci.org/arxanchain/tomago-sdk-go.svg?branch=master)](https://travis-ci.org/arxanchain/tomago-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/arxanchain/tomago-sdk-go)](https://goreportcard.com/report/github.com/arxanchain/tomago-sdk-go)
[![GoDoc](https://godoc.org/github.com/arxanchain/tomago-sdk-go?status.svg)](https://godoc.org/github.com/arxanchain/tomago-sdk-go)

# tomago-sdk-go

Tomago is a project code name, which is used to invoke the Smartcontract which is
deployed by yourself. You need not care about how the backend blockchain runs or 
the unintelligible techniques, such as consensus, endorsement and decentralization. 
Simply use the SDK we provide to implement your business logics, we will handle 
the caching, tagging, compressing, encrypting and high availability.

This SDK enables Go developers to develop applications that interact with the
SmartContract which is deployed out of the box or by yourself in the ArxanChain
BaaS Platform via Tomago.

# Usage

## Install

Run following command to download Go SDK

```code
go get github.com/arxanchain/tomago-sdk-go/api
```

## Request APIKey and download certificates

Before using the Tomago SDK Client, you need to request an ApiKey and download
the certificates from ArxanChain BaaS ChainConsole for communicating with 
Tomago service via HTTPS protocol.

The certificates include:

* The root CA (ca.crt). You can download it from the ArxanChain BaaS ChainConsole
  -> System Management -> Root CA Certs Management
* The private key and cert of the client user (`ApiKey.key` and `ApiKey.pem`), and 
  they are compressed into a ZIP-file. You can download it when you create an API
  Certificate.

## Create Tomago Client

```code
tlsConfig := &api.TLSCfonfig{
    CAFile: "/path/to/tls/ca/cert",
    KeyFile: "/path/to/tls/user/key",
    CertFile: "/path/to/tls/user/cert",
}

config := &api.Config{
    Address:    "https://<API-Gateway-DomainName>:<Port>",
    ApiKey: "6fD9G0QpM1516158053",
    TLSConfig: tlsConfig,
}

tomagoClient, err := NewTomagoClient(config)
```

* **Address** - The BaaS Service DomainName with `https` prefix
* **ApiKey** - The API access key on `ChainConsole` management page
* **TLSConfig** - The real TLS configuration.

## Invoke the SmartContract deployed by yourself

After you publish a chaincode into production blockchain environment, you can
use following APIs to invoke or query your chaincode.

* Invoke a chaincode function

  ```code
  chaincodeClient = tomagoClient.GetBlockchainClient()

  reqBody := &structs.PayloadWithTags{
      Payload: &structs.ChaincodeRequest{
          ChaincodeID: chaincodeID,
          Args:        []string{"invoke", "a", "b", "1"},
      },
  }

  //set http header
  header := http.Header{}
  header.Set("Channel-Id", "mychannel")

  //do invoke blockchain
  resp, err := chaincodeClient.Invoke(header, reqBody)
  ```

  - Channel-Id: This is your private blockchain ID, you can get this ID from
    the System Admin. If the chaincode is published on the public blockchain,
    you need not set this in http header.
  - ChaincodeID: The name of the chaincode when you publish it into the
    production blockchain environment
  - Args: The first element of this array is the function name that you want to
    invoke, the rest are the arguments of the function

* Query a chaincode function

  ```code
  chaincodeClient = tomagoClient.GetBlockchainClient()

  reqBody := &structs.PayloadWithTags{
      Payload: &structs.ChaincodeRequest{
          ChaincodeID: chaincodeID,
          Args:        []string{"query", "a"},
      },
  }

  //set http header
  header := http.Header{}
  header.Set("Channel-Id", "mychannel")

  //do query blockchain
  resp, err := chaincodeClient.Query(header, reqBody)
  ```
