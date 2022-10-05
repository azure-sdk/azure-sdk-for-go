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
	"strconv"
	"strings"
)

// ScopeConnectionsClient contains the methods for the ScopeConnections group.
// Don't use this type directly, use NewScopeConnectionsClient() instead.
type ScopeConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewScopeConnectionsClient creates a new instance of ScopeConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewScopeConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ScopeConnectionsClient, error) {
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
	client := &ScopeConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates scope connection from Network Manager
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// networkManagerName - The name of the network manager.
// scopeConnectionName - Name for the cross-tenant connection.
// parameters - Scope connection to be created/updated.
// options - ScopeConnectionsClientCreateOrUpdateOptions contains the optional parameters for the ScopeConnectionsClient.CreateOrUpdate
// method.
func (client *ScopeConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string, parameters ScopeConnection, options *ScopeConnectionsClientCreateOrUpdateOptions) (ScopeConnectionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkManagerName, scopeConnectionName, parameters, options)
	if err != nil {
		return ScopeConnectionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScopeConnectionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ScopeConnectionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ScopeConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string, parameters ScopeConnection, options *ScopeConnectionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/scopeConnections/{scopeConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if scopeConnectionName == "" {
		return nil, errors.New("parameter scopeConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeConnectionName}", url.PathEscape(scopeConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ScopeConnectionsClient) createOrUpdateHandleResponse(resp *http.Response) (ScopeConnectionsClientCreateOrUpdateResponse, error) {
	result := ScopeConnectionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScopeConnection); err != nil {
		return ScopeConnectionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the pending scope connection created by this network manager.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// networkManagerName - The name of the network manager.
// scopeConnectionName - Name for the cross-tenant connection.
// options - ScopeConnectionsClientDeleteOptions contains the optional parameters for the ScopeConnectionsClient.Delete method.
func (client *ScopeConnectionsClient) Delete(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string, options *ScopeConnectionsClientDeleteOptions) (ScopeConnectionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkManagerName, scopeConnectionName, options)
	if err != nil {
		return ScopeConnectionsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScopeConnectionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ScopeConnectionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ScopeConnectionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ScopeConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string, options *ScopeConnectionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/scopeConnections/{scopeConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if scopeConnectionName == "" {
		return nil, errors.New("parameter scopeConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeConnectionName}", url.PathEscape(scopeConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get specified scope connection created by this Network Manager.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// networkManagerName - The name of the network manager.
// scopeConnectionName - Name for the cross-tenant connection.
// options - ScopeConnectionsClientGetOptions contains the optional parameters for the ScopeConnectionsClient.Get method.
func (client *ScopeConnectionsClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string, options *ScopeConnectionsClientGetOptions) (ScopeConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkManagerName, scopeConnectionName, options)
	if err != nil {
		return ScopeConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScopeConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScopeConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ScopeConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string, options *ScopeConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/scopeConnections/{scopeConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if scopeConnectionName == "" {
		return nil, errors.New("parameter scopeConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeConnectionName}", url.PathEscape(scopeConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ScopeConnectionsClient) getHandleResponse(resp *http.Response) (ScopeConnectionsClientGetResponse, error) {
	result := ScopeConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScopeConnection); err != nil {
		return ScopeConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all scope connections created by this network manager.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// networkManagerName - The name of the network manager.
// options - ScopeConnectionsClientListOptions contains the optional parameters for the ScopeConnectionsClient.List method.
func (client *ScopeConnectionsClient) NewListPager(resourceGroupName string, networkManagerName string, options *ScopeConnectionsClientListOptions) *runtime.Pager[ScopeConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScopeConnectionsClientListResponse]{
		More: func(page ScopeConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScopeConnectionsClientListResponse) (ScopeConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkManagerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ScopeConnectionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ScopeConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ScopeConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ScopeConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, options *ScopeConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/scopeConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ScopeConnectionsClient) listHandleResponse(resp *http.Response) (ScopeConnectionsClientListResponse, error) {
	result := ScopeConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScopeConnectionListResult); err != nil {
		return ScopeConnectionsClientListResponse{}, err
	}
	return result, nil
}
