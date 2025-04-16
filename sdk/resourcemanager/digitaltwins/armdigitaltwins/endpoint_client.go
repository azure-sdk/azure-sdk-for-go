// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins

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

// EndpointClient contains the methods for the DigitalTwinsEndpoint group.
// Don't use this type directly, use NewEndpointClient() instead.
type EndpointClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEndpointClient creates a new instance of EndpointClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEndpointClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EndpointClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EndpointClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update DigitalTwinsInstance endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-31
//   - resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
//   - resourceName - The name of the DigitalTwinsInstance.
//   - endpointName - Name of Endpoint Resource.
//   - endpointDescription - The DigitalTwinsInstance endpoint metadata and security metadata.
//   - options - EndpointClientBeginCreateOrUpdateOptions contains the optional parameters for the EndpointClient.BeginCreateOrUpdate
//     method.
func (client *EndpointClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, endpointDescription EndpointResource, options *EndpointClientBeginCreateOrUpdateOptions) (*runtime.Poller[EndpointClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, endpointName, endpointDescription, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[EndpointClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[EndpointClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update DigitalTwinsInstance endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-31
func (client *EndpointClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, endpointDescription EndpointResource, options *EndpointClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "EndpointClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, endpointName, endpointDescription, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EndpointClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, endpointDescription EndpointResource, _ *EndpointClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, endpointDescription); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a DigitalTwinsInstance endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-31
//   - resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
//   - resourceName - The name of the DigitalTwinsInstance.
//   - endpointName - Name of Endpoint Resource.
//   - options - EndpointClientBeginDeleteOptions contains the optional parameters for the EndpointClient.BeginDelete method.
func (client *EndpointClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, options *EndpointClientBeginDeleteOptions) (*runtime.Poller[EndpointClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, endpointName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[EndpointClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[EndpointClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a DigitalTwinsInstance endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-31
func (client *EndpointClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, options *EndpointClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "EndpointClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, endpointName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EndpointClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, _ *EndpointClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get DigitalTwinsInstances Endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-31
//   - resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
//   - resourceName - The name of the DigitalTwinsInstance.
//   - endpointName - Name of Endpoint Resource.
//   - options - EndpointClientGetOptions contains the optional parameters for the EndpointClient.Get method.
func (client *EndpointClient) Get(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, options *EndpointClientGetOptions) (EndpointClientGetResponse, error) {
	var err error
	const operationName = "EndpointClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, endpointName, options)
	if err != nil {
		return EndpointClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EndpointClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EndpointClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EndpointClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, endpointName string, _ *EndpointClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EndpointClient) getHandleResponse(resp *http.Response) (EndpointClientGetResponse, error) {
	result := EndpointClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointResource); err != nil {
		return EndpointClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get DigitalTwinsInstance Endpoints.
//
// Generated from API version 2023-01-31
//   - resourceGroupName - The name of the resource group that contains the DigitalTwinsInstance.
//   - resourceName - The name of the DigitalTwinsInstance.
//   - options - EndpointClientListOptions contains the optional parameters for the EndpointClient.NewListPager method.
func (client *EndpointClient) NewListPager(resourceGroupName string, resourceName string, options *EndpointClientListOptions) *runtime.Pager[EndpointClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EndpointClientListResponse]{
		More: func(page EndpointClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EndpointClientListResponse) (EndpointClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EndpointClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return EndpointClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *EndpointClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, _ *EndpointClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EndpointClient) listHandleResponse(resp *http.Response) (EndpointClientListResponse, error) {
	result := EndpointClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointResourceListResult); err != nil {
		return EndpointClientListResponse{}, err
	}
	return result, nil
}
