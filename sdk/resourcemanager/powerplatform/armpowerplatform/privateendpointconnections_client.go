//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerplatform

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

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".PrivateEndpointConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Approve or reject a private endpoint connection with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - enterprisePolicyName - EnterprisePolicy for the Microsoft Azure subscription.
//   - privateEndpointConnectionName - The name of the private endpoint connection.
//   - parameters - Parameters supplied to create or update a private endpoint connection.
//   - options - PrivateEndpointConnectionsClientCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.CreateOrUpdate
//     method.
func (client *PrivateEndpointConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, enterprisePolicyName string, privateEndpointConnectionName string, parameters PrivateEndpointConnection, options *PrivateEndpointConnectionsClientCreateOrUpdateOptions) (PrivateEndpointConnectionsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, enterprisePolicyName, privateEndpointConnectionName, parameters, options)
	if err != nil {
		return PrivateEndpointConnectionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateEndpointConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, enterprisePolicyName string, privateEndpointConnectionName string, parameters PrivateEndpointConnection, options *PrivateEndpointConnectionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerPlatform/enterprisePolicies/{enterprisePolicyName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if enterprisePolicyName == "" {
		return nil, errors.New("parameter enterprisePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{enterprisePolicyName}", url.PathEscape(enterprisePolicyName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PrivateEndpointConnectionsClient) createOrUpdateHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientCreateOrUpdateResponse, error) {
	result := PrivateEndpointConnectionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a private endpoint connection with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - enterprisePolicyName - EnterprisePolicy for the Microsoft Azure subscription.
//   - privateEndpointConnectionName - The name of the private endpoint connection.
//   - options - PrivateEndpointConnectionsClientDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Delete
//     method.
func (client *PrivateEndpointConnectionsClient) Delete(ctx context.Context, resourceGroupName string, enterprisePolicyName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientDeleteOptions) (PrivateEndpointConnectionsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, enterprisePolicyName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientDeleteResponse{}, err
	}
	return PrivateEndpointConnectionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, enterprisePolicyName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerPlatform/enterprisePolicies/{enterprisePolicyName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if enterprisePolicyName == "" {
		return nil, errors.New("parameter enterprisePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{enterprisePolicyName}", url.PathEscape(enterprisePolicyName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - enterprisePolicyName - EnterprisePolicy for the Microsoft Azure subscription.
//   - privateEndpointConnectionName - The name of the private endpoint connection.
//   - options - PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
//     method.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, enterprisePolicyName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, enterprisePolicyName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, enterprisePolicyName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerPlatform/enterprisePolicies/{enterprisePolicyName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if enterprisePolicyName == "" {
		return nil, errors.New("parameter enterprisePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{enterprisePolicyName}", url.PathEscape(enterprisePolicyName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetResponse, error) {
	result := PrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByEnterprisePolicyPager - List all private endpoint connections on an EnterprisePolicy.
//
// Generated from API version 2020-10-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - enterprisePolicyName - EnterprisePolicy for the Microsoft Azure subscription.
//   - options - PrivateEndpointConnectionsClientListByEnterprisePolicyOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByEnterprisePolicyPager
//     method.
func (client *PrivateEndpointConnectionsClient) NewListByEnterprisePolicyPager(resourceGroupName string, enterprisePolicyName string, options *PrivateEndpointConnectionsClientListByEnterprisePolicyOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByEnterprisePolicyResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateEndpointConnectionsClientListByEnterprisePolicyResponse]{
		More: func(page PrivateEndpointConnectionsClientListByEnterprisePolicyResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionsClientListByEnterprisePolicyResponse) (PrivateEndpointConnectionsClientListByEnterprisePolicyResponse, error) {
			req, err := client.listByEnterprisePolicyCreateRequest(ctx, resourceGroupName, enterprisePolicyName, options)
			if err != nil {
				return PrivateEndpointConnectionsClientListByEnterprisePolicyResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateEndpointConnectionsClientListByEnterprisePolicyResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateEndpointConnectionsClientListByEnterprisePolicyResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByEnterprisePolicyHandleResponse(resp)
		},
	})
}

// listByEnterprisePolicyCreateRequest creates the ListByEnterprisePolicy request.
func (client *PrivateEndpointConnectionsClient) listByEnterprisePolicyCreateRequest(ctx context.Context, resourceGroupName string, enterprisePolicyName string, options *PrivateEndpointConnectionsClientListByEnterprisePolicyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerPlatform/enterprisePolicies/{enterprisePolicyName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if enterprisePolicyName == "" {
		return nil, errors.New("parameter enterprisePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{enterprisePolicyName}", url.PathEscape(enterprisePolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEnterprisePolicyHandleResponse handles the ListByEnterprisePolicy response.
func (client *PrivateEndpointConnectionsClient) listByEnterprisePolicyHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListByEnterprisePolicyResponse, error) {
	result := PrivateEndpointConnectionsClientListByEnterprisePolicyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsClientListByEnterprisePolicyResponse{}, err
	}
	return result, nil
}
