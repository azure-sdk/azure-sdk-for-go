//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

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

// MarketplaceSubscriptionsClient contains the methods for the MarketplaceSubscriptions group.
// Don't use this type directly, use NewMarketplaceSubscriptionsClient() instead.
type MarketplaceSubscriptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMarketplaceSubscriptionsClient creates a new instance of MarketplaceSubscriptionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMarketplaceSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MarketplaceSubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MarketplaceSubscriptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update Marketplace Subscription (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Marketplace Subscription name.
//   - body - Marketplace Subscription entity to apply during operation.
//   - options - MarketplaceSubscriptionsClientBeginCreateOrUpdateOptions contains the optional parameters for the MarketplaceSubscriptionsClient.BeginCreateOrUpdate
//     method.
func (client *MarketplaceSubscriptionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, name string, body MarketplaceSubscription, options *MarketplaceSubscriptionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[MarketplaceSubscriptionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, name, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MarketplaceSubscriptionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MarketplaceSubscriptionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update Marketplace Subscription (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
func (client *MarketplaceSubscriptionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, name string, body MarketplaceSubscription, options *MarketplaceSubscriptionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MarketplaceSubscriptionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, name, body, options)
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
func (client *MarketplaceSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, body MarketplaceSubscription, options *MarketplaceSubscriptionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/marketplaceSubscriptions/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete Marketplace Subscription (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Marketplace Subscription name.
//   - options - MarketplaceSubscriptionsClientBeginDeleteOptions contains the optional parameters for the MarketplaceSubscriptionsClient.BeginDelete
//     method.
func (client *MarketplaceSubscriptionsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *MarketplaceSubscriptionsClientBeginDeleteOptions) (*runtime.Poller[MarketplaceSubscriptionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MarketplaceSubscriptionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MarketplaceSubscriptionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete Marketplace Subscription (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
func (client *MarketplaceSubscriptionsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *MarketplaceSubscriptionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "MarketplaceSubscriptionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MarketplaceSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *MarketplaceSubscriptionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/marketplaceSubscriptions/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name.
//   - options - MarketplaceSubscriptionsClientGetOptions contains the optional parameters for the MarketplaceSubscriptionsClient.Get
//     method.
func (client *MarketplaceSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *MarketplaceSubscriptionsClientGetOptions) (MarketplaceSubscriptionsClientGetResponse, error) {
	var err error
	const operationName = "MarketplaceSubscriptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return MarketplaceSubscriptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MarketplaceSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MarketplaceSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MarketplaceSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *MarketplaceSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/marketplaceSubscriptions/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MarketplaceSubscriptionsClient) getHandleResponse(resp *http.Response) (MarketplaceSubscriptionsClientGetResponse, error) {
	result := MarketplaceSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplaceSubscription); err != nil {
		return MarketplaceSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List containers.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - options - MarketplaceSubscriptionsClientListOptions contains the optional parameters for the MarketplaceSubscriptionsClient.NewListPager
//     method.
func (client *MarketplaceSubscriptionsClient) NewListPager(resourceGroupName string, workspaceName string, options *MarketplaceSubscriptionsClientListOptions) *runtime.Pager[MarketplaceSubscriptionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MarketplaceSubscriptionsClientListResponse]{
		More: func(page MarketplaceSubscriptionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MarketplaceSubscriptionsClientListResponse) (MarketplaceSubscriptionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MarketplaceSubscriptionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return MarketplaceSubscriptionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MarketplaceSubscriptionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *MarketplaceSubscriptionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/marketplaceSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MarketplaceSubscriptionsClient) listHandleResponse(resp *http.Response) (MarketplaceSubscriptionsClientListResponse, error) {
	result := MarketplaceSubscriptionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplaceSubscriptionResourceArmPaginatedResult); err != nil {
		return MarketplaceSubscriptionsClientListResponse{}, err
	}
	return result, nil
}
