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
	assetClient structs.IAssetClient
)

func initAssetClient(t *testing.T) {
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
	assetClient = tomagoClient.GetAssetClient()
}

func TestCreateAssetSucc(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()

	const (
		id     = "did:ara:001"
		coinID = "33333"
	)

	//request body & response body
	reqBody := &structs.AssetBody{
		Id:         id,
		Name:       "票据001",
		Hash:       "sajfskjdfsdjfsdjfsj12923932kdfjds",
		ParentId:   "9867900-shdfjk",
		Owner:      "dd37fb3b-79d8-405e-8292-916de58d8663",
		ExpireTime: 10,
		Metadata: `{
			"name": "Army",
			"type": 1,
			"order_id": "678997"
		}`,
	}
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
		Post("/v2/assets").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

	//do create asset
	resp, err := assetClient.CreateAsset(header, reqBody)
	if err != nil {
		t.Fatalf("create asset fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != coinID {
		t.Fatalf("Coin Id should be %s, not %s", coinID, resp.CoinId)
	}
}

func TestCreateAssetFail(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()

	const (
		id     = "did:ara:001"
		coinID = "33333"
	)

	//request body & response body
	reqBody := &structs.AssetBody{
		Id:         id,
		Name:       "票据001",
		Hash:       "sajfskjdfsdjfsdjfsj12923932kdfjds",
		ParentId:   "9867900-shdfjk",
		Owner:      "dd37fb3b-79d8-405e-8292-916de58d8663",
		ExpireTime: 10,
		Metadata: `{
			"name": "Army",
			"type": 1,
			"order_id": "678997"
		}`,
	}

	//mock http request
	gock.New("http://127.0.0.1:8003").
		Post("/v2/assets").
		Reply(401)

	//set http header
	header := http.Header{}
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

	//do create asset
	resp, err := assetClient.CreateAsset(header, reqBody)
	if err == nil {
		t.Fatalf("create asset fail err should not be nil")
	}
	if resp != nil {
		t.Fatalf("create asset fail response should be nil")
	}
}

func TestCreateAssetErrCode(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()

	const (
		id      = "did:ara:001"
		coinID  = "33333"
		errCode = 5000
		errMsg  = "Register Entity Fail"
	)

	//request body & response body
	reqBody := &structs.AssetBody{
		Id:         id,
		Name:       "票据001",
		Hash:       "sajfskjdfsdjfsdjfsj12923932kdfjds",
		ParentId:   "9867900-shdfjk",
		Owner:      "dd37fb3b-79d8-405e-8292-916de58d8663",
		ExpireTime: 10,
		Metadata: `{
			"name": "Army",
			"type": 1,
			"order_id": "678997"
		}`,
	}
	payload := &structs.TomagoResponse{
		Code:           errCode,
		Message:        errMsg,
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
		Post("/v2/assets").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")

	//do create asset
	resp, err := assetClient.CreateAsset(header, reqBody)
	if err == nil {
		t.Fatalf("create asset fail err should not be nil")
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
func TestUpdateAssetSucc(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()
	const (
		id        = "did:ara:001"
		mychannel = "dcc"
	)
	payload := &structs.TomagoResponse{
		Code:           0,
		Message:        "",
		Id:             id,
		CoinId:         "2222",
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	gock.New("http://127.0.0.1:8003").
		Put("/v2/assets/zzz").
		Reply(200).
		JSON(byPayload)
	//set http header
	header := http.Header{}
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")
	reqBody := &structs.AssetBody{
		Id:         id,
		Name:       "票据001",
		Hash:       "sajfskjdfsdjfsdjfsj12923932kdfjds",
		ParentId:   "9867900-shdfjk",
		Owner:      "dd37fb3b-79d8-405e-8292-916de58d8663",
		ExpireTime: 10,
		Metadata: `{
			"name": "Army",
			"type": 1,
			"order_id": "678997"
		}`,
	}
	resp, err := assetClient.UpdateAsset(header, "zzz", reqBody)
	if err != nil {
		t.Fatalf("update asset fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != "2222" {
		t.Fatalf("Coin Id should be %s, not %s", "2222", resp.CoinId)
	}
}

func TestUpdateAssetFail(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()
	const (
		id        = "did:ara:001"
		mychannel = "dcc"
	)

	gock.New("http://127.0.0.1:8003").
		Put("/v2/assets/zzz").
		Reply(401)
	//set http header
	header := http.Header{}
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")
	reqBody := &structs.AssetBody{
		Id:         id,
		Name:       "票据001",
		Hash:       "sajfskjdfsdjfsdjfsj12923932kdfjds",
		ParentId:   "9867900-shdfjk",
		Owner:      "dd37fb3b-79d8-405e-8292-916de58d8663",
		ExpireTime: 10,
		Metadata: `{
			"name": "Army",
			"type": 1,
			"order_id": "678997"
		}`,
	}
	resp, err := assetClient.UpdateAsset(header, "zzz", reqBody)
	if err == nil {
		t.Fatalf("update asset fail err should be nil")
	}
	if resp != nil {
		t.Fatalf("response should be nil")
	}
}

func TestUpdateAssetErrCode(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()
	const (
		id        = "did:ara:001"
		mychannel = "dcc"
		errCode   = 5001
		errMsg    = "Update Entity Fail"
	)
	payload := &structs.TomagoResponse{
		Code:           errCode,
		Message:        errMsg,
		Id:             id,
		CoinId:         "2222",
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	gock.New("http://127.0.0.1:8003").
		Put("/v2/assets/zzz").
		Reply(200).
		JSON(byPayload)
	//set http header
	header := http.Header{}
	header.Set("Callback-Url", "http://172.16.199.6:8091/v2/test")
	reqBody := &structs.AssetBody{
		Id:         id,
		Name:       "票据001",
		Hash:       "sajfskjdfsdjfsdjfsj12923932kdfjds",
		ParentId:   "9867900-shdfjk",
		Owner:      "dd37fb3b-79d8-405e-8292-916de58d8663",
		ExpireTime: 10,
		Metadata: `{
			"name": "Army",
			"type": 1,
			"order_id": "678997"
		}`,
	}
	resp, err := assetClient.UpdateAsset(header, "zzz", reqBody)
	if err == nil {
		t.Fatalf("update asset fail err should not be nil")
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

func TestQueryAssetSucc(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()
	const (
		id        = "did:ara:001"
		mychannel = "dcc"
	)
	payload := &structs.TomagoResponse{
		Code:           0,
		Message:        "",
		Id:             id,
		CoinId:         "2222",
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}
	gock.New("http://127.0.0.1:8003").
		Get("/v2/assets/zzz").
		Reply(200).
		JSON(byPayload)
	//set http header
	header := http.Header{}
	resp, err := assetClient.QueryAsset(header, "zzz")
	if err != nil {
		t.Fatalf("create asset fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
}

func TestQueryAssetFail(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()
	const (
		id        = "did:ara:001"
		mychannel = "dcc"
	)

	gock.New("http://127.0.0.1:8003").
		Get("/v2/assets/zzz").
		Reply(401)
	//set http header
	header := http.Header{}
	resp, err := assetClient.QueryAsset(header, "zzz")
	if err == nil {
		t.Fatalf("query asset fail err should not be nil")
	}
	if resp != nil {
		t.Fatalf("response should be nil")
	}
}

func TestQueryAssetErrCode(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()
	const (
		id        = "did:ara:001"
		mychannel = "dcc"
		errCode   = 5002
		errMsg    = "Query Entity Fail"
	)
	payload := &structs.TomagoResponse{
		Code:           errCode,
		Message:        errMsg,
		Id:             id,
		CoinId:         "2222",
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}
	gock.New("http://127.0.0.1:8003").
		Get("/v2/assets/zzz").
		Reply(200).
		JSON(byPayload)
	//set http header
	header := http.Header{}
	_, err = assetClient.QueryAsset(header, "zzz")
	if err == nil {
		t.Fatalf("query asset fail err should not be nil")
	}
	/*
		if resp == nil {
			t.Fatalf("response should not be nil")
		} */
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
}

func TestTransferAssetSucc(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()
	const (
		id        = "did:ara:001"
		mychannel = "dcc"
	)
	payload := &structs.TomagoResponse{
		Code:           0,
		Message:        "",
		Id:             id,
		CoinId:         "2222",
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	gock.New("http://127.0.0.1:8003").
		Post("/v2/assets/transfer").
		Reply(200).
		JSON(byPayload)
	//set http header
	header := http.Header{}
	header.Set("Callback-Url", "zzz")
	reqBody := &structs.TransferAssetBody{
		From:   "ss",
		To:     "xx",
		Assets: []string{""},
		Fees:   nil,
	}
	resp, err := assetClient.TransferAsset(header, reqBody)
	if err != nil {
		t.Fatalf("create asset fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != "2222" {
		t.Fatalf("Coin Id should be %s, not %s", "2222", resp.CoinId)
	}
}
func TestTransferAssetFail(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()
	const (
		id        = "did:ara:001"
		mychannel = "dcc"
	)
	payload := &structs.TomagoResponse{
		Code:           0,
		Message:        "",
		Id:             id,
		CoinId:         "2222",
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	gock.New("http://127.0.0.1:8003").
		Post("/v2/assets/transfer").
		Reply(200).
		JSON(byPayload)
	//set http header
	header := http.Header{}
	header.Set("Callback-Url", "zzz")
	reqBody := &structs.TransferAssetBody{
		From:   "ss",
		To:     "xx",
		Assets: []string{""},
		Fees:   nil,
	}
	resp, err := assetClient.TransferAsset(header, reqBody)
	if err != nil {
		t.Fatalf("create asset fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != "2222" {
		t.Fatalf("Coin Id should be %s, not %s", "2222", resp.CoinId)
	}
}
func TestTransferAssetErrCode(t *testing.T) {
	//init gock & assetclient
	initAssetClient(t)
	defer gock.Off()
	const (
		id        = "did:ara:001"
		mychannel = "dcc"
		errCode   = 5012
		errMsg    = "Transfer Asset Fail"
	)
	payload := &structs.TomagoResponse{
		Code:           errCode,
		Message:        errMsg,
		Id:             id,
		CoinId:         "2222",
		TransactionIds: []string{""},
	}
	byPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	gock.New("http://127.0.0.1:8003").
		Post("/v2/assets/transfer").
		Reply(200).
		JSON(byPayload)
	//set http header
	header := http.Header{}
	header.Set("Callback-Url", "zzz")
	reqBody := &structs.TransferAssetBody{
		From:   "ss",
		To:     "xx",
		Assets: []string{""},
		Fees:   nil,
	}
	resp, err := assetClient.TransferAsset(header, reqBody)
	if err == nil {
		t.Fatalf("transfer asset fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
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
}
