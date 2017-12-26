/*
Copyright ArxanFintech Technology Ltd. 2017 All Rights Reserved.

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
	"io/ioutil"
	"net/http"

	cerr "github.com/arxanchain/sdk-go-common/errors"
	"github.com/arxanchain/sdk-go-common/rest"
	restapi "github.com/arxanchain/sdk-go-common/rest/api"
	"github.com/arxanchain/sdk-go-common/structs"
)

// EntityClient entity client struct
type EntityClient struct {
	c *restapi.Client
}

// CreateEntity is used to create an entity
// Response:
//   result:
// 	   Code: error code
//     Message: error message
//     Id: entity id returned by server
//     TransactionIds: blockchain transaction id list, this api returns one transaction id
//   err: create entity succ, return nil; others return non-nil.
//
func (t *EntityClient) CreateEntity(header http.Header, body *structs.EntityBody) (result *structs.TomagoResponse, err error) {
	r := t.c.NewRequest("POST", "/v2/entities")
	r.SetHeaders(header)
	r.SetBody(body)

	_, resp, err := restapi.RequireOK(t.c.DoRequest(r))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = restapi.DecodeBody(resp, &result); err != nil {
		return
	}

	if result.Code != 0 {
		err = rest.CodedError(cerr.ErrCodeType(result.Code), result.Message)
		return
	}

	return
}

// UpdateEntity is used to update an entity
// Response:
//   result:
// 	   Code: error code
//     Message: error message
//     TransactionIds: blockchain transaction id list, this api returns one transaction id
//   err: create entity succ, return nil; others return non-nil.
//
func (t *EntityClient) UpdateEntity(header http.Header, id string, body *structs.EntityBody) (result *structs.TomagoResponse, err error) {
	r := t.c.NewRequest("PUT", "/v2/entities/"+id)
	r.SetHeaders(header)
	r.SetBody(body)

	_, resp, err := restapi.RequireOK(t.c.DoRequest(r))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = restapi.DecodeBody(resp, &result); err != nil {
		return
	}

	if result.Code != 0 {
		err = rest.CodedError(cerr.ErrCodeType(result.Code), result.Message)
		return
	}

	return
}

// QueryEntity is used to query the entity metadata
// Response:
//   payload: return entity payload if succ
//   err: create entity succ, return nil; others return non-nil.
//
func (t *EntityClient) QueryEntity(header http.Header, id string) (payload *structs.EntityPayload, err error) {
	// Build request
	r := t.c.NewRequest("GET", "/v2/entities/"+id)
	r.SetHeaders(header)

	// Do request
	_, resp, err := restapi.RequireOK(t.c.DoRequest(r))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Read response result
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Check status code
	var result *structs.TomagoResponse
	if err = json.Unmarshal(respData, &result); err != nil {
		return
	}

	if result.Code != 0 {
		err = rest.CodedError(cerr.ErrCodeType(result.Code), result.Message)
		return
	}

	// Parse entity payload
	if err = json.Unmarshal(respData, &payload); err != nil {
		return
	}

	return
}
