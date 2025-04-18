// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// FirewallPolicyIdpsSignaturesFilterValuesClient contains the methods for the FirewallPolicyIdpsSignaturesFilterValues group.
// Don't use this type directly, use NewFirewallPolicyIdpsSignaturesFilterValuesClient() instead.
type FirewallPolicyIdpsSignaturesFilterValuesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFirewallPolicyIdpsSignaturesFilterValuesClient creates a new instance of FirewallPolicyIdpsSignaturesFilterValuesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFirewallPolicyIdpsSignaturesFilterValuesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FirewallPolicyIdpsSignaturesFilterValuesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FirewallPolicyIdpsSignaturesFilterValuesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - Retrieves the current filter values for the signatures overrides
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group.
//   - firewallPolicyName - The name of the Firewall Policy.
//   - options - FirewallPolicyIdpsSignaturesFilterValuesClientListOptions contains the optional parameters for the FirewallPolicyIdpsSignaturesFilterValuesClient.List
//     method.
func (client *FirewallPolicyIdpsSignaturesFilterValuesClient) List(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters SignatureOverridesFilterValuesQuery, options *FirewallPolicyIdpsSignaturesFilterValuesClientListOptions) (FirewallPolicyIdpsSignaturesFilterValuesClientListResponse, error) {
	var err error
	const operationName = "FirewallPolicyIdpsSignaturesFilterValuesClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, firewallPolicyName, parameters, options)
	if err != nil {
		return FirewallPolicyIdpsSignaturesFilterValuesClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallPolicyIdpsSignaturesFilterValuesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FirewallPolicyIdpsSignaturesFilterValuesClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *FirewallPolicyIdpsSignaturesFilterValuesClient) listCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters SignatureOverridesFilterValuesQuery, _ *FirewallPolicyIdpsSignaturesFilterValuesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/listIdpsFilterOptions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FirewallPolicyIdpsSignaturesFilterValuesClient) listHandleResponse(resp *http.Response) (FirewallPolicyIdpsSignaturesFilterValuesClientListResponse, error) {
	result := FirewallPolicyIdpsSignaturesFilterValuesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SignatureOverridesFilterValuesResponse); err != nil {
		return FirewallPolicyIdpsSignaturesFilterValuesClientListResponse{}, err
	}
	return result, nil
}
