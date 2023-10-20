//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// DomainWhoisClient contains the methods for the DomainWhois group.
// Don't use this type directly, use NewDomainWhoisClient() instead.
type DomainWhoisClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDomainWhoisClient creates a new instance of DomainWhoisClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDomainWhoisClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DomainWhoisClient, error) {
	cl, err := arm.NewClient(moduleName+".DomainWhoisClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DomainWhoisClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get whois information for a single domain name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - domain - Domain name to be enriched
//   - options - DomainWhoisClientGetOptions contains the optional parameters for the DomainWhoisClient.Get method.
func (client *DomainWhoisClient) Get(ctx context.Context, resourceGroupName string, domain string, options *DomainWhoisClientGetOptions) (DomainWhoisClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, domain, options)
	if err != nil {
		return DomainWhoisClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DomainWhoisClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DomainWhoisClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DomainWhoisClient) getCreateRequest(ctx context.Context, resourceGroupName string, domain string, options *DomainWhoisClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityInsights/enrichment/domain/whois/"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	reqQP.Set("domain", domain)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DomainWhoisClient) getHandleResponse(resp *http.Response) (DomainWhoisClientGetResponse, error) {
	result := DomainWhoisClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnrichmentDomainWhois); err != nil {
		return DomainWhoisClientGetResponse{}, err
	}
	return result, nil
}
