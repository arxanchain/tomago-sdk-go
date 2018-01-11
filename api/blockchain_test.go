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
	"github.com/arxanchain/sdk-go-common/structs"
	gock "gopkg.in/h2non/gock.v1"
)

var (
	chaincodeClient structs.IBlockchainClient
)

func initBlockchainClient(t *testing.T) {
	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	tomagoClient, err := NewTomagoClient(&api.Config{Address: "http://127.0.0.1:8003", HttpClient: client})
	if err != nil {
		t.Fatalf("New tomago client fail: %v", err)
	}
	chaincodeClient = tomagoClient.GetBlockchainClient()
}

func TestInvokeAssetSucc(t *testing.T) {
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
			Channel:     channel,
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
	header.Set("Channel-Id", "dacc")

	//do create blockchain
	resp, err := chaincodeClient.Invoke(header, reqBody)
	if err != nil {
		t.Fatalf("create chaincode fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
}

func TestInvokeAssetFail(t *testing.T) {
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
			Channel:     channel,
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
	header.Set("Channel-Id", "dacc")

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
			Channel:     channel,
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
	header.Set("Channel-Id", "dacc")

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
