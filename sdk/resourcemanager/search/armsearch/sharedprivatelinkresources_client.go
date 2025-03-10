// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

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

// SharedPrivateLinkResourcesClient contains the methods for the SharedPrivateLinkResources group.
// Don't use this type directly, use NewSharedPrivateLinkResourcesClient() instead.
type SharedPrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSharedPrivateLinkResourcesClient creates a new instance of SharedPrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription. You can obtain this value from the Azure Resource
//     Manager API, command line tools, or the portal.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSharedPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SharedPrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SharedPrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Initiates the creation or update of a shared private link resource managed by the search service
// in the given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-01
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the search service associated with the specified resource group.
//   - sharedPrivateLinkResourceName - The name of the shared private link resource managed by the search service within the specified
//     resource group.
//   - sharedPrivateLinkResource - The definition of the shared private link resource to create or update.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.BeginCreateOrUpdate
//     method.
func (client *SharedPrivateLinkResourcesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, sharedPrivateLinkResource SharedPrivateLinkResource, searchManagementRequestOptions *SearchManagementRequestOptions, options *SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SharedPrivateLinkResourcesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, searchServiceName, sharedPrivateLinkResourceName, sharedPrivateLinkResource, searchManagementRequestOptions, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SharedPrivateLinkResourcesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SharedPrivateLinkResourcesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Initiates the creation or update of a shared private link resource managed by the search service in the
// given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-01
func (client *SharedPrivateLinkResourcesClient) createOrUpdate(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, sharedPrivateLinkResource SharedPrivateLinkResource, searchManagementRequestOptions *SearchManagementRequestOptions, options *SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SharedPrivateLinkResourcesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, searchServiceName, sharedPrivateLinkResourceName, sharedPrivateLinkResource, searchManagementRequestOptions, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SharedPrivateLinkResourcesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, sharedPrivateLinkResource SharedPrivateLinkResource, searchManagementRequestOptions *SearchManagementRequestOptions, _ *SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	if err := runtime.MarshalAsJSON(req, sharedPrivateLinkResource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Initiates the deletion of the shared private link resource from the search service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-01
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the search service associated with the specified resource group.
//   - sharedPrivateLinkResourceName - The name of the shared private link resource managed by the search service within the specified
//     resource group.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - SharedPrivateLinkResourcesClientBeginDeleteOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.BeginDelete
//     method.
func (client *SharedPrivateLinkResourcesClient) BeginDelete(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*runtime.Poller[SharedPrivateLinkResourcesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, searchServiceName, sharedPrivateLinkResourceName, searchManagementRequestOptions, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SharedPrivateLinkResourcesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SharedPrivateLinkResourcesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Initiates the deletion of the shared private link resource from the search service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-01
func (client *SharedPrivateLinkResourcesClient) deleteOperation(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SharedPrivateLinkResourcesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, searchServiceName, sharedPrivateLinkResourceName, searchManagementRequestOptions, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SharedPrivateLinkResourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, searchManagementRequestOptions *SearchManagementRequestOptions, _ *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	return req, nil
}

// Get - Gets the details of the shared private link resource managed by the search service in the given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-01
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the search service associated with the specified resource group.
//   - sharedPrivateLinkResourceName - The name of the shared private link resource managed by the search service within the specified
//     resource group.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - SharedPrivateLinkResourcesClientGetOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.Get
//     method.
func (client *SharedPrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *SharedPrivateLinkResourcesClientGetOptions) (SharedPrivateLinkResourcesClientGetResponse, error) {
	var err error
	const operationName = "SharedPrivateLinkResourcesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, searchServiceName, sharedPrivateLinkResourceName, searchManagementRequestOptions, options)
	if err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SharedPrivateLinkResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, searchManagementRequestOptions *SearchManagementRequestOptions, _ *SharedPrivateLinkResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharedPrivateLinkResourcesClient) getHandleResponse(resp *http.Response) (SharedPrivateLinkResourcesClientGetResponse, error) {
	result := SharedPrivateLinkResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedPrivateLinkResource); err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServicePager - Gets a list of all shared private link resources managed by the given service.
//
// Generated from API version 2025-05-01
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the search service associated with the specified resource group.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - SharedPrivateLinkResourcesClientListByServiceOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.NewListByServicePager
//     method.
func (client *SharedPrivateLinkResourcesClient) NewListByServicePager(resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *SharedPrivateLinkResourcesClientListByServiceOptions) *runtime.Pager[SharedPrivateLinkResourcesClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[SharedPrivateLinkResourcesClientListByServiceResponse]{
		More: func(page SharedPrivateLinkResourcesClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SharedPrivateLinkResourcesClientListByServiceResponse) (SharedPrivateLinkResourcesClientListByServiceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SharedPrivateLinkResourcesClient.NewListByServicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServiceCreateRequest(ctx, resourceGroupName, searchServiceName, searchManagementRequestOptions, options)
			}, nil)
			if err != nil {
				return SharedPrivateLinkResourcesClientListByServiceResponse{}, err
			}
			return client.listByServiceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *SharedPrivateLinkResourcesClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, _ *SharedPrivateLinkResourcesClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/sharedPrivateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *SharedPrivateLinkResourcesClient) listByServiceHandleResponse(resp *http.Response) (SharedPrivateLinkResourcesClientListByServiceResponse, error) {
	result := SharedPrivateLinkResourcesClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedPrivateLinkResourceListResult); err != nil {
		return SharedPrivateLinkResourcesClientListByServiceResponse{}, err
	}
	return result, nil
}
