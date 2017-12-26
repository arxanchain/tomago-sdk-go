package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/arxanchain/sdk-go-common/rest/api"
	"github.com/arxanchain/sdk-go-common/structs"
	gock "gopkg.in/h2non/gock.v1"
)

var (
	entityClient structs.IEntityClient
)

func initEntityClient(t *testing.T) {
	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	tomagoClient, err := NewTomagoClient(&api.Config{Address: "http://127.0.0.1:8003", HttpClient: client})
	if err != nil {
		t.Fatalf("New tomago client fail: %v", err)
	}
	entityClient = tomagoClient.GetEntityClient()
}

func TestCreateEntitySucc(t *testing.T) {
	//init gock & Entityclient
	initEntityClient(t)
	defer gock.Off()

	const (
		id     = "did:ara:001"
		coinID = "33333"
	)

	//request body & response body
	reqBody := &structs.EntityBody{
		Id:           id,
		EnrollmentId: "alice",
		CallbackUrl:  "http://172.16.199.6:8091/v2/test",
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
		Post("/v2/entities").
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")

	//do create asset
	resp, err := entityClient.CreateEntity(header, reqBody)
	if err != nil {
		t.Fatalf("create entity fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != coinID {
		t.Fatalf("Coin Id should be %s, not %s", coinID, resp.CoinId)
	}
}

func TestUpdateEntitySucc(t *testing.T) {
	//init gock & Entityclient
	initEntityClient(t)
	defer gock.Off()

	const (
		id       = "did:ara:001"
		coinID   = "33333"
		entityID = "8e114136-e6f8-4dc0-8605-ad4f8a1e0d35"
	)

	//request body & response body
	reqBody := &structs.EntityBody{
		Id:           id,
		EnrollmentId: "alice",
		CallbackUrl:  "http://172.16.199.6:8091/v2/test",
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
		Put("/v2/entities/" + entityID).
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")

	//do create asset
	resp, err := entityClient.UpdateEntity(header, entityID, reqBody)
	if err != nil {
		t.Fatalf("create entity fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
	if resp.CoinId != coinID {
		t.Fatalf("Coin Id should be %s, not %s", coinID, resp.CoinId)
	}
}

func TestQueryEntitySucc(t *testing.T) {
	//init gock & Entityclient
	initEntityClient(t)
	defer gock.Off()

	const (
		id       = "did:ara:001"
		coinID   = "33333"
		entityID = "8e114136-e6f8-4dc0-8605-ad4f8a1e0d35"
	)

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
		Get("/v2/entities/" + entityID).
		Reply(200).
		JSON(byPayload)

	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "dacc")

	//do create asset
	resp, err := entityClient.QueryEntity(header, entityID)
	if err != nil {
		t.Fatalf("create entity fail: %v", err)
	}
	if resp == nil {
		t.Fatalf("response should not be nil")
	}
}
