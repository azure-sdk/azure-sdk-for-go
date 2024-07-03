//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ServersControllerClient contains the methods for the ServersController group.
// Don't use this type directly, use NewServersControllerClient() instead.
type ServersControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServersControllerClient creates a new instance of ServersControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServersControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServersControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServersControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// DeleteMachine - Delete a Server
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - machineName - A server machine name
//   - options - ServersControllerClientDeleteMachineOptions contains the optional parameters for the ServersControllerClient.DeleteMachine
//     method.
func (client *ServersControllerClient) DeleteMachine(ctx context.Context, resourceGroupName string, siteName string, machineName string, options *ServersControllerClientDeleteMachineOptions) (ServersControllerClientDeleteMachineResponse, error) {
	var err error
	const operationName = "ServersControllerClient.DeleteMachine"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteMachineCreateRequest(ctx, resourceGroupName, siteName, machineName, options)
	if err != nil {
		return ServersControllerClientDeleteMachineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServersControllerClientDeleteMachineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServersControllerClientDeleteMachineResponse{}, err
	}
	return ServersControllerClientDeleteMachineResponse{}, nil
}

// deleteMachineCreateRequest creates the DeleteMachine request.
func (client *ServersControllerClient) deleteMachineCreateRequest(ctx context.Context, resourceGroupName string, siteName string, machineName string, options *ServersControllerClientDeleteMachineOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/serverSites/{siteName}/machines/{machineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetMachine - Get a Server
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - machineName - A server machine name
//   - options - ServersControllerClientGetMachineOptions contains the optional parameters for the ServersControllerClient.GetMachine
//     method.
func (client *ServersControllerClient) GetMachine(ctx context.Context, resourceGroupName string, siteName string, machineName string, options *ServersControllerClientGetMachineOptions) (ServersControllerClientGetMachineResponse, error) {
	var err error
	const operationName = "ServersControllerClient.GetMachine"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getMachineCreateRequest(ctx, resourceGroupName, siteName, machineName, options)
	if err != nil {
		return ServersControllerClientGetMachineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServersControllerClientGetMachineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServersControllerClientGetMachineResponse{}, err
	}
	resp, err := client.getMachineHandleResponse(httpResp)
	return resp, err
}

// getMachineCreateRequest creates the GetMachine request.
func (client *ServersControllerClient) getMachineCreateRequest(ctx context.Context, resourceGroupName string, siteName string, machineName string, options *ServersControllerClientGetMachineOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/serverSites/{siteName}/machines/{machineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getMachineHandleResponse handles the GetMachine response.
func (client *ServersControllerClient) getMachineHandleResponse(resp *http.Response) (ServersControllerClientGetMachineResponse, error) {
	result := ServersControllerClientGetMachineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Server); err != nil {
		return ServersControllerClientGetMachineResponse{}, err
	}
	return result, nil
}

// NewListByServerSiteResourcePager - Get all machines in a site.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - ServersControllerClientListByServerSiteResourceOptions contains the optional parameters for the ServersControllerClient.NewListByServerSiteResourcePager
//     method.
func (client *ServersControllerClient) NewListByServerSiteResourcePager(resourceGroupName string, siteName string, options *ServersControllerClientListByServerSiteResourceOptions) *runtime.Pager[ServersControllerClientListByServerSiteResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServersControllerClientListByServerSiteResourceResponse]{
		More: func(page ServersControllerClientListByServerSiteResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServersControllerClientListByServerSiteResourceResponse) (ServersControllerClientListByServerSiteResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServersControllerClient.NewListByServerSiteResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerSiteResourceCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return ServersControllerClientListByServerSiteResourceResponse{}, err
			}
			return client.listByServerSiteResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerSiteResourceCreateRequest creates the ListByServerSiteResource request.
func (client *ServersControllerClient) listByServerSiteResourceCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *ServersControllerClientListByServerSiteResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/serverSites/{siteName}/machines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", *options.Top)
	}
	if options != nil && options.TotalRecordCount != nil {
		reqQP.Set("totalRecordCount", strconv.FormatInt(int64(*options.TotalRecordCount), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerSiteResourceHandleResponse handles the ListByServerSiteResource response.
func (client *ServersControllerClient) listByServerSiteResourceHandleResponse(resp *http.Response) (ServersControllerClientListByServerSiteResourceResponse, error) {
	result := ServersControllerClientListByServerSiteResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerListResult); err != nil {
		return ServersControllerClientListByServerSiteResourceResponse{}, err
	}
	return result, nil
}

// UpdateMachine - Update a Server machine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - machineName - A server machine name
//   - body - The resource properties to be updated.
//   - options - ServersControllerClientUpdateMachineOptions contains the optional parameters for the ServersControllerClient.UpdateMachine
//     method.
func (client *ServersControllerClient) UpdateMachine(ctx context.Context, resourceGroupName string, siteName string, machineName string, body ServerUpdate, options *ServersControllerClientUpdateMachineOptions) (ServersControllerClientUpdateMachineResponse, error) {
	var err error
	const operationName = "ServersControllerClient.UpdateMachine"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateMachineCreateRequest(ctx, resourceGroupName, siteName, machineName, body, options)
	if err != nil {
		return ServersControllerClientUpdateMachineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServersControllerClientUpdateMachineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServersControllerClientUpdateMachineResponse{}, err
	}
	resp, err := client.updateMachineHandleResponse(httpResp)
	return resp, err
}

// updateMachineCreateRequest creates the UpdateMachine request.
func (client *ServersControllerClient) updateMachineCreateRequest(ctx context.Context, resourceGroupName string, siteName string, machineName string, body ServerUpdate, options *ServersControllerClientUpdateMachineOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/serverSites/{siteName}/machines/{machineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// updateMachineHandleResponse handles the UpdateMachine response.
func (client *ServersControllerClient) updateMachineHandleResponse(resp *http.Response) (ServersControllerClientUpdateMachineResponse, error) {
	result := ServersControllerClientUpdateMachineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Server); err != nil {
		return ServersControllerClientUpdateMachineResponse{}, err
	}
	return result, nil
}
