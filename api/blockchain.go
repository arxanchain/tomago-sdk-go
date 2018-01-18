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
	"net/http"

	cerr "github.com/arxanchain/sdk-go-common/errors"
	"github.com/arxanchain/sdk-go-common/rest"
	restapi "github.com/arxanchain/sdk-go-common/rest/api"
	"github.com/arxanchain/sdk-go-common/structs"
)

// BlockchainClient asset client struct
type BlockchainClient struct {
	c *restapi.Client
}

// Invoke is used for starting a transaction to blockchain platform
// Response:
//   result:
// 	  Code: error code
//    Message: return blockchain trade id; else return failure cause
//    err: start transaction succ, return nil; others return non-nil.
//
func (t *BlockchainClient) Invoke(header http.Header, body *structs.PayloadWithTags) (result *structs.ChaincodeResponse, err error) {
	r := t.c.NewRequest("POST", "/v2/blockchain/invoke")
	r.SetHeaders(header)
	err = r.SetBody(body)
	if err != nil {
		return
	}

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

// Query is used for inquery the status that user transactions recorded in Blockchain
// Response:
//   result:
// 	   Code: error code
//     Message: error message
//   err: create asset succ, return nil; others return non-nil.
//
func (t *BlockchainClient) Query(header http.Header, body *structs.PayloadWithTags) (result *structs.ChaincodeResponse, err error) {
	r := t.c.NewRequest("POST", "/v2/blockchain/query")
	r.SetHeaders(header)
	err = r.SetBody(body)
	if err != nil {
		return
	}
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

// QueryTxn is used for inquery status of transactions in Blockchain, such as the blockchain that transactions are in, transaction time, transaction information
// Response:
//   result:
// 	   Code: error code
//     Message: error message
//   err: create asset succ, return nil; others return non-nil.
//
func (t *BlockchainClient) QueryTxn(header http.Header, txnid string) (result *structs.TransactionResponse, err error) {
	// Build request
	r := t.c.NewRequest("GET", "/v2/blockchain/transaction/"+txnid)
	r.SetHeaders(header)

	// Do request
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
