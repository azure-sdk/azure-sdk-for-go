//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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

// PrefixesClient contains the methods for the Prefixes group.
// Don't use this type directly, use NewPrefixesClient() instead.
type PrefixesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrefixesClient creates a new instance of PrefixesClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrefixesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrefixesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrefixesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new prefix with the specified name under the given subscription, resource group and peering
// service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The name of the resource group.
//   - peeringServiceName - The name of the peering service.
//   - prefixName - The name of the prefix.
//   - peeringServicePrefix - The properties needed to create a prefix.
//   - options - PrefixesClientCreateOrUpdateOptions contains the optional parameters for the PrefixesClient.CreateOrUpdate method.
func (client *PrefixesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, peeringServicePrefix ServicePrefix, options *PrefixesClientCreateOrUpdateOptions) (PrefixesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "PrefixesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, peeringServiceName, prefixName, peeringServicePrefix, options)
	if err != nil {
		return PrefixesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrefixesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return PrefixesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrefixesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, peeringServicePrefix ServicePrefix, options *PrefixesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if prefixName == "" {
		return nil, errors.New("parameter prefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{prefixName}", url.PathEscape(prefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, peeringServicePrefix); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PrefixesClient) createOrUpdateHandleResponse(resp *http.Response) (PrefixesClientCreateOrUpdateResponse, error) {
	result := PrefixesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServicePrefix); err != nil {
		return PrefixesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing prefix with the specified name under the given subscription, resource group and peering service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The name of the resource group.
//   - peeringServiceName - The name of the peering service.
//   - prefixName - The name of the prefix.
//   - options - PrefixesClientDeleteOptions contains the optional parameters for the PrefixesClient.Delete method.
func (client *PrefixesClient) Delete(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, options *PrefixesClientDeleteOptions) (PrefixesClientDeleteResponse, error) {
	var err error
	const operationName = "PrefixesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, peeringServiceName, prefixName, options)
	if err != nil {
		return PrefixesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrefixesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PrefixesClientDeleteResponse{}, err
	}
	return PrefixesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrefixesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, options *PrefixesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if prefixName == "" {
		return nil, errors.New("parameter prefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{prefixName}", url.PathEscape(prefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing prefix with the specified name under the given subscription, resource group and peering service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The name of the resource group.
//   - peeringServiceName - The name of the peering service.
//   - prefixName - The name of the prefix.
//   - options - PrefixesClientGetOptions contains the optional parameters for the PrefixesClient.Get method.
func (client *PrefixesClient) Get(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, options *PrefixesClientGetOptions) (PrefixesClientGetResponse, error) {
	var err error
	const operationName = "PrefixesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, peeringServiceName, prefixName, options)
	if err != nil {
		return PrefixesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrefixesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrefixesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrefixesClient) getCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, options *PrefixesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if prefixName == "" {
		return nil, errors.New("parameter prefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{prefixName}", url.PathEscape(prefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrefixesClient) getHandleResponse(resp *http.Response) (PrefixesClientGetResponse, error) {
	result := PrefixesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServicePrefix); err != nil {
		return PrefixesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPeeringServicePager - Lists all prefixes under the given subscription, resource group and peering service.
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The name of the resource group.
//   - peeringServiceName - The name of the peering service.
//   - options - PrefixesClientListByPeeringServiceOptions contains the optional parameters for the PrefixesClient.NewListByPeeringServicePager
//     method.
func (client *PrefixesClient) NewListByPeeringServicePager(resourceGroupName string, peeringServiceName string, options *PrefixesClientListByPeeringServiceOptions) *runtime.Pager[PrefixesClientListByPeeringServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrefixesClientListByPeeringServiceResponse]{
		More: func(page PrefixesClientListByPeeringServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrefixesClientListByPeeringServiceResponse) (PrefixesClientListByPeeringServiceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrefixesClient.NewListByPeeringServicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByPeeringServiceCreateRequest(ctx, resourceGroupName, peeringServiceName, options)
			}, nil)
			if err != nil {
				return PrefixesClientListByPeeringServiceResponse{}, err
			}
			return client.listByPeeringServiceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByPeeringServiceCreateRequest creates the ListByPeeringService request.
func (client *PrefixesClient) listByPeeringServiceCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, options *PrefixesClientListByPeeringServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPeeringServiceHandleResponse handles the ListByPeeringService response.
func (client *PrefixesClient) listByPeeringServiceHandleResponse(resp *http.Response) (PrefixesClientListByPeeringServiceResponse, error) {
	result := PrefixesClientListByPeeringServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServicePrefixListResult); err != nil {
		return PrefixesClientListByPeeringServiceResponse{}, err
	}
	return result, nil
}
