//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

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

// ManagementPoliciesClient contains the methods for the ManagementPolicies group.
// Don't use this type directly, use NewManagementPoliciesClient() instead.
type ManagementPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagementPoliciesClient creates a new instance of ManagementPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagementPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagementPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Sets the managementpolicy to the specified storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - managementPolicyName - The name of the Storage Account Management Policy. It should always be 'default'
//   - properties - The ManagementPolicy set to a storage account.
//   - options - ManagementPoliciesClientCreateOrUpdateOptions contains the optional parameters for the ManagementPoliciesClient.CreateOrUpdate
//     method.
func (client *ManagementPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, properties ManagementPolicy, options *ManagementPoliciesClientCreateOrUpdateOptions) (ManagementPoliciesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ManagementPoliciesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, managementPolicyName, properties, options)
	if err != nil {
		return ManagementPoliciesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagementPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagementPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, properties ManagementPolicy, options *ManagementPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if managementPolicyName == "" {
		return nil, errors.New("parameter managementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementPolicyName}", url.PathEscape(string(managementPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagementPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ManagementPoliciesClientCreateOrUpdateResponse, error) {
	result := ManagementPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementPolicy); err != nil {
		return ManagementPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the managementpolicy associated with the specified storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - managementPolicyName - The name of the Storage Account Management Policy. It should always be 'default'
//   - options - ManagementPoliciesClientDeleteOptions contains the optional parameters for the ManagementPoliciesClient.Delete
//     method.
func (client *ManagementPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, options *ManagementPoliciesClientDeleteOptions) (ManagementPoliciesClientDeleteResponse, error) {
	var err error
	const operationName = "ManagementPoliciesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, managementPolicyName, options)
	if err != nil {
		return ManagementPoliciesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ManagementPoliciesClientDeleteResponse{}, err
	}
	return ManagementPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagementPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, options *ManagementPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if managementPolicyName == "" {
		return nil, errors.New("parameter managementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementPolicyName}", url.PathEscape(string(managementPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the managementpolicy associated with the specified storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - managementPolicyName - The name of the Storage Account Management Policy. It should always be 'default'
//   - options - ManagementPoliciesClientGetOptions contains the optional parameters for the ManagementPoliciesClient.Get method.
func (client *ManagementPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, options *ManagementPoliciesClientGetOptions) (ManagementPoliciesClientGetResponse, error) {
	var err error
	const operationName = "ManagementPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, managementPolicyName, options)
	if err != nil {
		return ManagementPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagementPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagementPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, options *ManagementPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if managementPolicyName == "" {
		return nil, errors.New("parameter managementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementPolicyName}", url.PathEscape(string(managementPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagementPoliciesClient) getHandleResponse(resp *http.Response) (ManagementPoliciesClientGetResponse, error) {
	result := ManagementPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementPolicy); err != nil {
		return ManagementPoliciesClientGetResponse{}, err
	}
	return result, nil
}
