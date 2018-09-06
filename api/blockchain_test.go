/*
Copyright ArxanFintech Technology Ltd. 2018 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/arxanchain/sdk-go-common/rest"
	"github.com/arxanchain/sdk-go-common/rest/api"
	structs "github.com/arxanchain/sdk-go-common/structs/tomago"
	gock "gopkg.in/h2non/gock.v1"
)

var (
	chaincodeClient structs.IBlockchainClient
)

func initBlockchainClient(t *testing.T) {
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
	chaincodeClient = tomagoClient.GetBlockchainClient()
}

func TestInvokeSucc(t *testing.T) {
	//init gock & assetclient
	initBlockchainClient(t)
	defer gock.Off()

	const (
		channel     = "channel001"
		chaincodeID = "33333"
	)

	//request body & response body
	reqBody := &structs.PayloadWithTags{
		Payload: &structs.ChaincodeRequest{
			ChaincodeID: chaincodeID,
			Args:        []string{"invoke", "a", "b", "1"},
		},
	}
	ret := &structs.ChaincodeResponse{
		Result: "72b9c26c52e36e1824ca82901e0973de5081e7e8278a321c3c3f4bb719edf934",
	}
	byPayload, err := json.Marshal(ret)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/blockchain/invoke").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", channel)

	//do create blockchain
	resp, err := chaincodeClient.Invoke(header, reqBody)
	if err != nil {
		t.Fatalf("create chaincode fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
}

func TestInvokeFail(t *testing.T) {
	//init gock & assetclient
	initBlockchainClient(t)
	defer gock.Off()

	const (
		channel     = "channel001"
		chaincodeID = "mycc"
	)

	//request body & response body
	reqBody := &structs.PayloadWithTags{
		Payload: &structs.ChaincodeRequest{
			ChaincodeID: chaincodeID,
			Args:        []string{"invoke", "a", "b", "1"},
		},
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/blockchain/invoke").
		Reply(401)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", channel)

	//do create blockchain
	resp, err := chaincodeClient.Invoke(header, reqBody)
	if err == nil {
		t.Fatalf("create chaincode fail: %v or should not be nil", err)
	}
	if resp != nil {
		t.Fatalf("create chaincode fail response should not be nil")
	}
}

func TestInvokeErrCode(t *testing.T) {
	//init gock & assetclient
	initBlockchainClient(t)
	defer gock.Off()

	const (
		channel     = "channel001"
		chaincodeID = "mycc"
		errCode     = 5000
		errMsg      = "Invoke Transaction failed"
	)

	//request body & response body
	reqBody := &structs.PayloadWithTags{
		Payload: &structs.ChaincodeRequest{
			ChaincodeID: chaincodeID,
			Args:        []string{"invoke", "a", "b", "1"},
		},
	}
	payload := &structs.ChaincodeResponse{
		Code:    errCode,
		Message: errMsg,
		Result:  "",
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/blockchain/invoke").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", channel)

	//do create blockchain
	resp, err := chaincodeClient.Invoke(header, reqBody)
	if err == nil {
		t.Fatalf("Start invoke fail err should not be nil")
	}
	errWitherrCode, ok := err.(rest.HTTPCodedError)
	if !ok {
		t.Fatalf("err should be HTTPCodedError")
	}

	if errWitherrCode.Code() != errCode {
		t.Fatalf("err Code should be %d, not %d", errCode, errWitherrCode.Code())
	}

	if errWitherrCode.Error() != errMsg {
		t.Fatalf("errMsg should be %s, not %s", errMsg, errWitherrCode.Error())
	}

	if resp.Code != errCode {
		t.Fatalf("errCode should be %d, not %d", errCode, resp.Code)
	}
}

func TestQuerySucc(t *testing.T) {
	//init gock & assetclient
	initBlockchainClient(t)
	defer gock.Off()

	const (
		channel     = "channel001"
		chaincodeID = "mycc"
	)

	//request body & response body
	reqBody := &structs.PayloadWithTags{
		Payload: &structs.ChaincodeRequest{
			ChaincodeID: chaincodeID,
			Args:        []string{"query", "a"},
		},
	}
	ret := &structs.ChaincodeResponse{
		Result: "99",
	}
	byPayload, err := json.Marshal(ret)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/blockchain/query").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", channel)

	//do create blockchain
	resp, err := chaincodeClient.Query(header, reqBody)
	if err != nil {
		t.Fatalf("create chaincode fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
}

func TestQueryFail(t *testing.T) {
	//init gock & assetclient
	initBlockchainClient(t)
	defer gock.Off()

	const (
		channel     = "channel001"
		chaincodeID = "mycc"
	)

	//request body & response body
	reqBody := &structs.PayloadWithTags{
		Payload: &structs.ChaincodeRequest{
			ChaincodeID: chaincodeID,
			Args:        []string{"query", "a"},
		},
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/blockchain/query").
		Reply(401)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", channel)

	//do create blockchain
	resp, err := chaincodeClient.Query(header, reqBody)
	if err == nil {
		t.Fatalf("query chaincode fail: %v", err)
	}
	if resp != nil {
		t.Fatalf("response should not be nil")
	}
}

func TestQueryErrCode(t *testing.T) {
	//init gock & assetclient
	initBlockchainClient(t)
	defer gock.Off()

	const (
		channel     = "channel001"
		chaincodeID = "mycc"
		errCode     = 5000
		errMsg      = "Query chaincode failed"
	)

	//request body & response body
	reqBody := &structs.PayloadWithTags{
		Payload: &structs.ChaincodeRequest{
			ChaincodeID: chaincodeID,
			Args:        []string{"query", "a"},
		},
	}
	payload := &structs.ChaincodeResponse{
		Code:    errCode,
		Message: errMsg,
		Result:  "",
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/blockchain/query").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", channel)

	//do create blockchain
	resp, err := chaincodeClient.Query(header, reqBody)
	if err == nil {
		t.Fatalf("Start query chaincode fail err should not be nil")
	}
	errWitherrCode, ok := err.(rest.HTTPCodedError)
	if !ok {
		t.Fatalf("err should be HTTPCodedError")
	}

	if errWitherrCode.Code() != errCode {
		t.Fatalf("err Code should be %d, not %d", errCode, errWitherrCode.Code())
	}

	if errWitherrCode.Error() != errMsg {
		t.Fatalf("errMsg should be %s, not %s", errMsg, errWitherrCode.Error())
	}

	if resp.Code != errCode {
		t.Fatalf("errCode should be %d, not %d", errCode, resp.Code)
	}
}

func TestQueryTxnSucc(t *testing.T) {
	//init gock & assetclient
	initBlockchainClient(t)
	defer gock.Off()

	const (
		txid = "mycc"
	)

	//request body & response body
	ret := &structs.TransactionResponse{
		ChannelID:     "mychannel",
		ChaincodeID:   "mycc",
		TransactionID: "991d9f7658cb6515af4467c74842593158cf99b09c744f6d6137f751436707f9",
		Timestamp:     structs.Timestamp{Seconds: 1502867427, Nanos: 239380560},
		CreatorID:     []byte("CgdPcmcxTVNQEq4GLS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNMRENDQWRLZ0F3SUJBZ0lSQUtaSGhlQ1pQRStHTUxSVjJXWEJyMTB3Q2dZSUtvWkl6ajBFQXdJd2NERUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhHVEFYQmdOVkJBTVRFRzl5Clp6RXVaWGhoYlhCc1pTNWpiMjB3SGhjTk1UY3dOREl5TVRJd01qVTJXaGNOTWpjd05ESXdNVEl3TWpVMldqQmIKTVFzd0NRWURWUVFHRXdKVlV6RVRNQkVHQTFVRUNCTUtRMkZzYVdadmNtNXBZVEVXTUJRR0ExVUVCeE1OVTJGdQpJRVp5WVc1amFYTmpiekVmTUIwR0ExVUVBd3dXVlhObGNqRkFiM0puTVM1bGVHRnRjR3hsTG1OdmJUQlpNQk1HCkJ5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUFCRlVLdU5DbGl3VjlFNHRtU2JXV2QzdHYvNFpFNms0Q0dJaVkKYUtOSmpIWUk2WVZqbFRNRWwyTnJzU1djT01aMWF5cys5eEoyRXdqc1F2RGFpWkJuSlBlallqQmdNQTRHQTFVZApEd0VCL3dRRUF3SUZvREFUQmdOVkhTVUVEREFLQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQ3NHCkExVWRJd1FrTUNLQUlLSXRyelZyS3F0WGt1cFQ0MTltL003eDEvR3FLem9ya3R2NytXcEVqcUpxTUFvR0NDcUcKU000OUJBTUNBMGdBTUVVQ0lRRDNoc0hTMURTOU94N3RxNDZwN3gwUVdQOXljKytNN1hBN1BSZjhMN3dYL1FJZwpVMExkSVhKcmh4QVhYMjl0Qy9xRzJRR1BBNFQ1UVRDS1paY1ZOYUFUL0xRPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg"),
		PayloadSize:   1881,
		IsInvalID:     false,
		Payload:       "",
	}
	byPayload, err := json.Marshal(ret)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Get("/v2/blockchain/transaction/" + txid).
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")

	//do create blockchain
	resp, err := chaincodeClient.QueryTxn(header, txid)
	if err != nil {
		t.Fatalf("query transaction fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
}

func TestQueryTxnFail(t *testing.T) {
	//init gock & assetclient
	initBlockchainClient(t)
	defer gock.Off()

	const (
		txid = "mycc"
	)

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Get("/v2/blockchain/" + txid).
		Reply(401)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")

	//do create blockchain
	resp, err := chaincodeClient.QueryTxn(header, txid)
	if err == nil {
		t.Fatalf("query chaincode transaction fail: %v", err)
	}
	if resp != nil {
		t.Fatalf("response should not be nil")
	}
}

func TestQueryTxnErrCode(t *testing.T) {
	//init gock & assetclient
	initBlockchainClient(t)
	defer gock.Off()

	const (
		txid    = "mycc"
		errCode = 5000
		errMsg  = "Query Transaction failed"
	)

	//request body & response body
	payload := &structs.ChaincodeResponse{
		Result:  "",
		Code:    errCode,
		Message: errMsg,
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Get("/v2/blockchain/transaction/" + txid).
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")

	//do create blockchain
	resp, err := chaincodeClient.QueryTxn(header, txid)
	if err == nil {
		t.Fatalf("Start query transaction fail err should not be nil")
	}
	errWitherrCode, ok := err.(rest.HTTPCodedError)
	if !ok {
		t.Fatalf("err should be HTTPCodedError")
	}

	if errWitherrCode.Code() != errCode {
		t.Fatalf("err Code should be %d, not %d", errCode, errWitherrCode.Code())
	}

	if errWitherrCode.Error() != errMsg {
		t.Fatalf("errMsg should be %s, not %s", errMsg, errWitherrCode.Error())
	}

	if resp.Code != errCode {
		t.Fatalf("errCode should be %d, not %d", errCode, resp.Code)
	}
}
