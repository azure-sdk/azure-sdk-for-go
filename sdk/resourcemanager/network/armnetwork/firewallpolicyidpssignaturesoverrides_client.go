//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// FirewallPolicyIdpsSignaturesOverridesClient contains the methods for the FirewallPolicyIdpsSignaturesOverrides group.
// Don't use this type directly, use NewFirewallPolicyIdpsSignaturesOverridesClient() instead.
type FirewallPolicyIdpsSignaturesOverridesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFirewallPolicyIdpsSignaturesOverridesClient creates a new instance of FirewallPolicyIdpsSignaturesOverridesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFirewallPolicyIdpsSignaturesOverridesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FirewallPolicyIdpsSignaturesOverridesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &FirewallPolicyIdpsSignaturesOverridesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Returns all signatures overrides for a specific policy.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// options - FirewallPolicyIdpsSignaturesOverridesClientGetOptions contains the optional parameters for the FirewallPolicyIdpsSignaturesOverridesClient.Get
// method.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPolicyIdpsSignaturesOverridesClientGetOptions) (FirewallPolicyIdpsSignaturesOverridesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
	if err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FirewallPolicyIdpsSignaturesOverridesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) getCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPolicyIdpsSignaturesOverridesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/signatureOverrides/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) getHandleResponse(resp *http.Response) (FirewallPolicyIdpsSignaturesOverridesClientGetResponse, error) {
	result := FirewallPolicyIdpsSignaturesOverridesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SignaturesOverrides); err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns all signatures overrides objects for a specific policy as a list containing a single value.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// options - FirewallPolicyIdpsSignaturesOverridesClientListOptions contains the optional parameters for the FirewallPolicyIdpsSignaturesOverridesClient.List
// method.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) List(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPolicyIdpsSignaturesOverridesClientListOptions) (FirewallPolicyIdpsSignaturesOverridesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
	if err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FirewallPolicyIdpsSignaturesOverridesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) listCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPolicyIdpsSignaturesOverridesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/signatureOverrides"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) listHandleResponse(resp *http.Response) (FirewallPolicyIdpsSignaturesOverridesClientListResponse, error) {
	result := FirewallPolicyIdpsSignaturesOverridesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SignaturesOverridesList); err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientListResponse{}, err
	}
	return result, nil
}

// Patch - Will update the status of policy's signature overrides for IDPS
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// parameters - Will contain all properties of the object to put
// options - FirewallPolicyIdpsSignaturesOverridesClientPatchOptions contains the optional parameters for the FirewallPolicyIdpsSignaturesOverridesClient.Patch
// method.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) Patch(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters SignaturesOverrides, options *FirewallPolicyIdpsSignaturesOverridesClientPatchOptions) (FirewallPolicyIdpsSignaturesOverridesClientPatchResponse, error) {
	req, err := client.patchCreateRequest(ctx, resourceGroupName, firewallPolicyName, parameters, options)
	if err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientPatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientPatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FirewallPolicyIdpsSignaturesOverridesClientPatchResponse{}, runtime.NewResponseError(resp)
	}
	return client.patchHandleResponse(resp)
}

// patchCreateRequest creates the Patch request.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) patchCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters SignaturesOverrides, options *FirewallPolicyIdpsSignaturesOverridesClientPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/signatureOverrides/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// patchHandleResponse handles the Patch response.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) patchHandleResponse(resp *http.Response) (FirewallPolicyIdpsSignaturesOverridesClientPatchResponse, error) {
	result := FirewallPolicyIdpsSignaturesOverridesClientPatchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SignaturesOverrides); err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientPatchResponse{}, err
	}
	return result, nil
}

// Put - Will override/create a new signature overrides for the policy's IDPS
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// parameters - Will contain all properties of the object to put
// options - FirewallPolicyIdpsSignaturesOverridesClientPutOptions contains the optional parameters for the FirewallPolicyIdpsSignaturesOverridesClient.Put
// method.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) Put(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters SignaturesOverrides, options *FirewallPolicyIdpsSignaturesOverridesClientPutOptions) (FirewallPolicyIdpsSignaturesOverridesClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, resourceGroupName, firewallPolicyName, parameters, options)
	if err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientPutResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FirewallPolicyIdpsSignaturesOverridesClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) putCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters SignaturesOverrides, options *FirewallPolicyIdpsSignaturesOverridesClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/signatureOverrides/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// putHandleResponse handles the Put response.
func (client *FirewallPolicyIdpsSignaturesOverridesClient) putHandleResponse(resp *http.Response) (FirewallPolicyIdpsSignaturesOverridesClientPutResponse, error) {
	result := FirewallPolicyIdpsSignaturesOverridesClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SignaturesOverrides); err != nil {
		return FirewallPolicyIdpsSignaturesOverridesClientPutResponse{}, err
	}
	return result, nil
}
