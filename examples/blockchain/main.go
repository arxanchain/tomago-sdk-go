package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	restapi "github.com/arxanchain/sdk-go-common/rest/api"
	commstructs "github.com/arxanchain/sdk-go-common/structs"
	structs "github.com/arxanchain/sdk-go-common/structs/tomago"
	tomagoapi "github.com/arxanchain/tomago-sdk-go/api"
)

func main() {
	var err error
	var tomagoClient structs.ITomagoClient
	var bcClient structs.IBlockchainClient
	var payload structs.PayloadWithTags
	var ccResp *structs.ChaincodeResponse
	var txnResp *structs.TransactionResponse

	// Create tomago client
	config := restapi.Config{
		Address:     "https://remotehost.com:443",
		ApiKey:      "Mc31HANHp1541501752",
		CallbackUrl: "http://172.16.12.21:18091/v1/test",
		TLSConfig: restapi.TLSConfig{
			CAFile:   "./ca.crt",
			KeyFile:  "./Mc31HANHp1541501752.key",
			CertFile: "./Mc31HANHp1541501752.pem",
		},
	}

	tomagoClient, err = tomagoapi.NewTomagoClient(&config)
	if err != nil {
		log.Printf("New tomago client fail: %v\n", err)
		return
	}
	log.Printf("New tomago client succ\n")

	bcClient = tomagoClient.GetBlockchainClient()

	// Build request header
	header := http.Header{}
	header.Add(commstructs.ChannelIdHeader, "pubchain")

	// Invoke
	invokeBodyBytes, err := ioutil.ReadFile("./invoke.json")
	if err != nil {
		log.Printf("Read transfer payload fail: %v", err)
		return
	}
	log.Printf("Invoke payload:\n%s\n", invokeBodyBytes)

	err = json.Unmarshal(invokeBodyBytes, &payload)
	if err != nil {
		log.Printf("Unmarshal invoke body fail: %v", err)
		return
	}

	ccResp, err = bcClient.Invoke(header, &payload)
	if err != nil {
		log.Printf("Invoke fail: %v", err)
		return
	}

	log.Printf("Invoke succ. Response: %+v", ccResp)

	txID := ccResp.Message

	// Query
	queryBodyBytes, err := ioutil.ReadFile("./query.json")
	if err != nil {
		log.Printf("Read transfer payload fail: %v", err)
		return
	}
	log.Printf("Invoke payload:\n%s\n", queryBodyBytes)

	err = json.Unmarshal(queryBodyBytes, &payload)
	if err != nil {
		log.Printf("Unmarshal query body fail: %v", err)
		return
	}

	ccResp, err = bcClient.Query(header, &payload)
	if err != nil {
		log.Printf("Invoke fail: %v", err)
		return
	}

	log.Printf("Query succ. Response: %+v", ccResp)

	// Query Txn
	txnResp, err = bcClient.QueryTxn(header, txID)
	if err != nil {
		log.Printf("Query transaction (%v) fail: %v", txID, err)
		return
	}

	log.Printf("Query transaction (%v) succ. Response: %+v", txID, txnResp)
}
