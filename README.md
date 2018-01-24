
# Status
[![Build Status](https://travis-ci.org/arxanchain/tomago-sdk-go.svg?branch=master)](https://travis-ci.org/arxanchain/tomago-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/arxanchain/tomago-sdk-go)](https://goreportcard.com/report/github.com/arxanchain/tomago-sdk-go)
[![GoDoc](https://godoc.org/github.com/arxanchain/tomago-sdk-go?status.svg)](https://godoc.org/github.com/arxanchain/tomago-sdk-go)

# tomago-sdk-go

Tomago is a project code name, which is used to wrap SmartContract invocation
from the business point of view, including APIs for asset owner (entity)
management, digital assets, etc. You need not care about how the backend
blockchain runs or the unintelligible techniques, such as consensus, endorsement
and decentralization. Simply use the SDK we provide to implement your business
logics, we will handle the caching, tagging, compressing, encrypting and high
availability.

We also provide a way from this SDK to invoke the SmartContract, a.k.a.
Chaincode, which is deployed by yourself.

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
the certificates from ArxanChain BaaS ChainConsole for data encryption and
signing. This will help ensure the data cannot be tampered with or illegally
accessed to even if the client communicates with Tomago service via HTTPS.

The certificates include:

* The public key of ArxanChain BaaS Platform (server.crt) which is used to
  encrypt the data sent to Tomago service. You can download it from the
  ArxanChain BaaS ChainConsole -> System Management -> API Certs Management
* The private key of the client user (`ApiKey`.key) which is used to sign the
  data. You can download it when you create an API Certificate.

## Create Tomago Client

```code
cryptoConfig := &api.CryptoConfig{
    Enable:         true,
    CertsStorePath: "./certs",
}

config := &api.Config{
    Address:    "http://127.0.0.1:8003",
    ApiKey: "6fD9G0QpM1516158053",
    CryptoCfg cryptoConfig,
}

tomagoClient, err := NewTomagoClient(config)
```

* Enable - Whether encrypted communication is enabled between the client and
  the server
* CertsStorePath - The directory where the public key and the private key are
  kept, e.g. `server.crt` and `6fD9G0QpM1516158053.key`.
* Address - The Tomago Service API URL, we support http and https both
* ApiKey - We assume the ApiKey for your request is "6fD9G0QpM1516158053"

## Issue digital assets

* Before issuing a digital asset, we need to create the owner of this asset
  first. Here we call that an entity.

  ```code
  entityClient := tomagoClient.GetEntityClient()
  reqBody := &structs.EntityBody{
      Id:           id,
      Metadata: `{
          "name": "Army",
          "type": 1,
          "order_id": "678997"
      }`,
  }

  //set Channel Id
  header := http.Header{}
  header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

  //do create asset
  resp, err := entityClient.CreateEntity(header, reqBody)

  ```

  - CallbackUrl: Asynchronous event notification which will notify if the
    request succeeded or failed
  - Id: The entity ID, Tomago will generated a UUID if you do not provide this

* Issue a digital asset

  ```code
  assetClient = tomagoClient.GetAssetClient()

  //request body & response body
  reqBody := &structs.AssetBody{
      Id:           id,
      Name:         "bill001",
      Owner:        "did:axn:8uQhQMGzWxR8vw5P3UWH1j",
      Metadata: `{
          "name": "Army",
          "type": 1,
          "order_id": "678997"
      }`,
  }
  //set http header
  header := http.Header{}
  header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

  //do create asset
  resp, err := assetClient.CreateAsset(header, reqBody)

  ```

  - CallbackUrl: Asynchronous event notification which will notify if the
    request succeeded or failed
  - Id: The asset ID, Tomago will generated a UUID if you do not provide this
  - Owner: The Entity ID of the asset owner

* Update a digital asset

  ```code
  assetClient = tomagoClient.GetAssetClient()

  //request body & response body
  reqBody := &structs.AssetBody{
      Name:       "bill002",
      ParentId:   "xxxxxx",
      Metadata: `{
          "name": "Army",
          "type": 2,
          "order_id": "678997"
      }`,
  }
  //set http header
  header := http.Header{}
  header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

  //do update asset
  resp, err := assetClient.UpdateAsset(header, "did:axn:f9zG01pp1516158024xxe1", reqBody)
  ```

  - We only allow updates of the Name, ParentId and Metadata of asset

* Transfer a digital asset

  ```code
  assetClient = tomagoClient.GetAssetClient()

  reqBody := &structs.TransferAssetBody{
      From:   "did:axn:ss",
      To:     "did:axn:xx",
      Assets: []string{""},
      Fees:   nil,
  }
  //set http header
  header := http.Header{}
  header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

  //do transfer asset
  resp, err := assetClient.TransferAsset(header, reqBody)
  ```

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

## Colored Coin

Deprecated, colored coin will be moved into wallet-sdk-go
