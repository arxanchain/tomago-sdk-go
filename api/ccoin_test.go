package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/arxanchain/sdk-go-common/rest/api"
	structs "github.com/arxanchain/sdk-go-common/structs/tomago"
	gock "gopkg.in/h2non/gock.v1"
)

var (
	ccoinClient structs.ICCoinClient
)

func initCcoinClient(t *testing.T) {
	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	config := &api.Config{
		Address:    "http://127.0.0.1:8003",
		ApiKey:     "xxxxxxxxxxxxx",
		HttpClient: client,
	}
	tomagoClient, err := NewTomagoClient(config)
	if err != nil {
		t.Fatalf("New tomago client fail: %v", err)
	}
	ccoinClient = tomagoClient.GetCCoinClient()
}

func TestIssueSucc(t *testing.T) {
	//init gock & Entityclient
	initCcoinClient(t)
	defer gock.Off()

	const (
		id     = "did:ara:001"
		coinID = "33333"
	)

	//request body & response body
	reqBody := &structs.IssueBody{}
	payload := &structs.TomagoResponse{
		Code:           0,
		Message:        "",
		Id:             id,
		CoinId:         coinID,
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/coins/issue").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

	//do create asset
	resp, err := ccoinClient.Issue(header, reqBody)
	if err != nil {
		t.Fatalf("issue fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != coinID {
		t.Fatalf("Coin Id should be %s, not %s", coinID, resp.CoinId)
	}
}

func TestTransferSucc(t *testing.T) {
	//init gock & Entityclient
	initCcoinClient(t)
	defer gock.Off()

	const (
		id     = "did:ara:001"
		coinID = "33333"
	)

	//request body & response body
	reqBody := &structs.TransferBody{}
	payload := &structs.TomagoResponse{
		Code:           0,
		Message:        "",
		Id:             id,
		CoinId:         coinID,
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/coins/transfer").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

	//do create asset
	resp, err := ccoinClient.Transfer(header, reqBody)
	if err != nil {
		t.Fatalf("transfer fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != coinID {
		t.Fatalf("Coin Id should be %s, not %s", coinID, resp.CoinId)
	}
}

func TestRollbackSucc(t *testing.T) {
	//init gock & Entityclient
	initCcoinClient(t)
	defer gock.Off()

	const (
		id     = "did:ara:001"
		coinID = "33333"
	)

	//request body & response body
	reqBody := &structs.RollbackBody{}
	payload := &structs.TomagoResponse{
		Code:           0,
		Message:        "",
		Id:             id,
		CoinId:         coinID,
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/coins/rollback").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

	//do create asset
	resp, err := ccoinClient.Rollback(header, reqBody)
	if err != nil {
		t.Fatalf("rollback fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != coinID {
		t.Fatalf("Coin Id should be %s, not %s", coinID, resp.CoinId)
	}
}

func TestInterestSucc(t *testing.T) {
	//init gock & Entityclient
	initCcoinClient(t)
	defer gock.Off()

	const (
		id     = "did:ara:001"
		coinID = "33333"
	)

	//request body & response body
	reqBody := &structs.InterestBody{}
	payload := &structs.TomagoResponse{
		Code:           0,
		Message:        "",
		Id:             id,
		CoinId:         coinID,
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/coins/interest").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

	//do create asset
	resp, err := ccoinClient.Interest(header, reqBody)
	if err != nil {
		t.Fatalf("interest fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != coinID {
		t.Fatalf("Coin Id should be %s, not %s", coinID, resp.CoinId)
	}
}

func TestWithdrawSucc(t *testing.T) {
	//init gock & Entityclient
	initCcoinClient(t)
	defer gock.Off()

	const (
		id     = "did:ara:001"
		coinID = "33333"
	)

	//request body & response body
	reqBody := &structs.WithdrawBody{}
	payload := &structs.TomagoResponse{
		Code:           0,
		Message:        "",
		Id:             id,
		CoinId:         coinID,
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/coins/withdraw").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

	//do create asset
	resp, err := ccoinClient.Withdraw(header, reqBody)
	if err != nil {
		t.Fatalf("withdraw fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != coinID {
		t.Fatalf("Coin Id should be %s, not %s", coinID, resp.CoinId)
	}
}
