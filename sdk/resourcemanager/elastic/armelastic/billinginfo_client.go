//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armelastic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BillingInfoClient contains the methods for the BillingInfo group.
// Don't use this type directly, use NewBillingInfoClient() instead.
type BillingInfoClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBillingInfoClient creates a new instance of BillingInfoClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBillingInfoClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BillingInfoClient, error) {
	cl, err := arm.NewClient(moduleName+".BillingInfoClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BillingInfoClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get marketplace and organization info mapped to the given monitor.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - options - BillingInfoClientGetOptions contains the optional parameters for the BillingInfoClient.Get method.
func (client *BillingInfoClient) Get(ctx context.Context, resourceGroupName string, monitorName string, options *BillingInfoClientGetOptions) (BillingInfoClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return BillingInfoClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BillingInfoClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BillingInfoClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BillingInfoClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *BillingInfoClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/getBillingInfo"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BillingInfoClient) getHandleResponse(resp *http.Response) (BillingInfoClientGetResponse, error) {
	result := BillingInfoClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingInfoResponse); err != nil {
		return BillingInfoClientGetResponse{}, err
	}
	return result, nil
}
