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

// OrganizationsClient contains the methods for the Organizations group.
// Don't use this type directly, use NewOrganizationsClient() instead.
type OrganizationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOrganizationsClient creates a new instance of OrganizationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOrganizationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OrganizationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OrganizationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetAPIKey - Fetch User API Key from internal database, if it was generated and stored while creating the Elasticsearch
// Organization.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - options - OrganizationsClientGetAPIKeyOptions contains the optional parameters for the OrganizationsClient.GetAPIKey method.
func (client *OrganizationsClient) GetAPIKey(ctx context.Context, options *OrganizationsClientGetAPIKeyOptions) (OrganizationsClientGetAPIKeyResponse, error) {
	var err error
	const operationName = "OrganizationsClient.GetAPIKey"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAPIKeyCreateRequest(ctx, options)
	if err != nil {
		return OrganizationsClientGetAPIKeyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OrganizationsClientGetAPIKeyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OrganizationsClientGetAPIKeyResponse{}, err
	}
	resp, err := client.getAPIKeyHandleResponse(httpResp)
	return resp, err
}

// getAPIKeyCreateRequest creates the GetAPIKey request.
func (client *OrganizationsClient) getAPIKeyCreateRequest(ctx context.Context, options *OrganizationsClientGetAPIKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Elastic/getOrganizationApiKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// getAPIKeyHandleResponse handles the GetAPIKey response.
func (client *OrganizationsClient) getAPIKeyHandleResponse(resp *http.Response) (OrganizationsClientGetAPIKeyResponse, error) {
	result := OrganizationsClientGetAPIKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserAPIKeyResponse); err != nil {
		return OrganizationsClientGetAPIKeyResponse{}, err
	}
	return result, nil
}

// GetElasticToAzureSubscriptionMapping - Get Elastic Organization To Azure Subscription Mapping details for the logged-in
// user.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - options - OrganizationsClientGetElasticToAzureSubscriptionMappingOptions contains the optional parameters for the OrganizationsClient.GetElasticToAzureSubscriptionMapping
//     method.
func (client *OrganizationsClient) GetElasticToAzureSubscriptionMapping(ctx context.Context, options *OrganizationsClientGetElasticToAzureSubscriptionMappingOptions) (OrganizationsClientGetElasticToAzureSubscriptionMappingResponse, error) {
	var err error
	const operationName = "OrganizationsClient.GetElasticToAzureSubscriptionMapping"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getElasticToAzureSubscriptionMappingCreateRequest(ctx, options)
	if err != nil {
		return OrganizationsClientGetElasticToAzureSubscriptionMappingResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OrganizationsClientGetElasticToAzureSubscriptionMappingResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OrganizationsClientGetElasticToAzureSubscriptionMappingResponse{}, err
	}
	resp, err := client.getElasticToAzureSubscriptionMappingHandleResponse(httpResp)
	return resp, err
}

// getElasticToAzureSubscriptionMappingCreateRequest creates the GetElasticToAzureSubscriptionMapping request.
func (client *OrganizationsClient) getElasticToAzureSubscriptionMappingCreateRequest(ctx context.Context, options *OrganizationsClientGetElasticToAzureSubscriptionMappingOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Elastic/getElasticOrganizationToAzureSubscriptionMapping"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getElasticToAzureSubscriptionMappingHandleResponse handles the GetElasticToAzureSubscriptionMapping response.
func (client *OrganizationsClient) getElasticToAzureSubscriptionMappingHandleResponse(resp *http.Response) (OrganizationsClientGetElasticToAzureSubscriptionMappingResponse, error) {
	result := OrganizationsClientGetElasticToAzureSubscriptionMappingResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationToAzureSubscriptionMappingResponse); err != nil {
		return OrganizationsClientGetElasticToAzureSubscriptionMappingResponse{}, err
	}
	return result, nil
}
