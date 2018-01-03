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
	"net/http"

	cerr "github.com/arxanchain/sdk-go-common/errors"
	"github.com/arxanchain/sdk-go-common/rest"
	restapi "github.com/arxanchain/sdk-go-common/rest/api"
	"github.com/arxanchain/sdk-go-common/structs"
)

// CCoinClient colored coin client struct
type CCoinClient struct {
	c *restapi.Client
}

// Issue is used to issue colored coins
// Response:
//   result:
// 	   Code: error code
//     Message: error message
//     TransactionIds: blockchain transaction id list, this api returns one transaction id
//   err: create asset succ, return nil; others return non-nil.
//
func (t *CCoinClient) Issue(header http.Header, body *structs.IssueBody) (result *structs.TomagoResponse, err error) {
	r := t.c.NewRequest("POST", "/v2/coins/issue")
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

// Transfer is used to transfer colored coins
// Response:
//   result:
// 	   Code: error code
//     Message: error message
//     TransactionIds: blockchain transaction id list, this api returns one transaction id
//   err: create asset succ, return nil; others return non-nil.
//
func (t *CCoinClient) Transfer(header http.Header, body *structs.TransferBody) (result *structs.TomagoResponse, err error) {
	r := t.c.NewRequest("POST", "/v2/coins/transfer")
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

// Rollback is used to rollback transaction finished before
// Response:
//   result:
// 	   Code: error code
//     Message: error message
//     TransactionIds: blockchain transaction id list
//   err: create asset succ, return nil; others return non-nil.
//
func (t *CCoinClient) Rollback(header http.Header, body *structs.RollbackBody) (result *structs.TomagoResponse, err error) {
	r := t.c.NewRequest("POST", "/v2/coins/rollback")
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

// Interest is used to charge interest
// Response:
//   result:
// 	   Code: error code
//     Message: error message
//     TransactionIds: blockchain transaction id list
//   err: create asset succ, return nil; others return non-nil.
//
func (t *CCoinClient) Interest(header http.Header, body *structs.InterestBody) (result *structs.TomagoResponse, err error) {
	r := t.c.NewRequest("POST", "/v2/coins/interest")
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

// Withdraw is used to withdraw colored coin to RMB
// Response:
//   result:
// 	   Code: error code
//     Message: error message
//     TransactionIds: blockchain transaction id list, this api returns one transaction id
//   err: create asset succ, return nil; others return non-nil.
//
func (t *CCoinClient) Withdraw(header http.Header, body *structs.WithdrawBody) (result *structs.TomagoResponse, err error) {
	r := t.c.NewRequest("POST", "/v2/coins/withdraw")
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
