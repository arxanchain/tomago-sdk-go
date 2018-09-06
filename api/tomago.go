/*
Copyright ArxanFintech Technology Ltd. 2017-2018 All Rights Reserved.

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
	"fmt"

	restapi "github.com/arxanchain/sdk-go-common/rest/api"
	structs "github.com/arxanchain/sdk-go-common/structs/tomago"
)

// TomagoClient tomago client struct
type TomagoClient struct {
	c                *restapi.Client
	blockchainClient *BlockchainClient
}

// NewTomagoClient returns a handle to the agent endpoints
func NewTomagoClient(config *restapi.Config) (*TomagoClient, error) {
	if config == nil {
		return nil, fmt.Errorf("config must be set")
	}
	if config.RouteTag == "" {
		config.RouteTag = "tomago"
	}

	c, err := restapi.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &TomagoClient{c: c}, nil
}

// GetBlockchainClient Get blockchain client
func (t *TomagoClient) GetBlockchainClient() structs.IBlockchainClient {
	if t.blockchainClient == nil {
		t.blockchainClient = &BlockchainClient{c: t.c}
	}
	return t.blockchainClient
}
