//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// ContainerAppsSessionPoolsClient contains the methods for the ContainerAppsSessionPools group.
// Don't use this type directly, use NewContainerAppsSessionPoolsClient() instead.
type ContainerAppsSessionPoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainerAppsSessionPoolsClient creates a new instance of ContainerAppsSessionPoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainerAppsSessionPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsSessionPoolsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerAppsSessionPoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a Session Pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sessionPoolName - Name of the Session Pool.
//   - sessionPoolEnvelope - Properties used to create a session pool
//   - options - ContainerAppsSessionPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the ContainerAppsSessionPoolsClient.BeginCreateOrUpdate
//     method.
func (client *ContainerAppsSessionPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, sessionPoolName string, sessionPoolEnvelope SessionPool, options *ContainerAppsSessionPoolsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ContainerAppsSessionPoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, sessionPoolName, sessionPoolEnvelope, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerAppsSessionPoolsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerAppsSessionPoolsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update a Session Pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *ContainerAppsSessionPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, sessionPoolName string, sessionPoolEnvelope SessionPool, options *ContainerAppsSessionPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerAppsSessionPoolsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, sessionPoolName, sessionPoolEnvelope, options)
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
func (client *ContainerAppsSessionPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, sessionPoolName string, sessionPoolEnvelope SessionPool, options *ContainerAppsSessionPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/sessionPools/{sessionPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sessionPoolName == "" {
		return nil, errors.New("parameter sessionPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sessionPoolName}", url.PathEscape(sessionPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sessionPoolEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a Session Pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sessionPoolName - Name of the Session Pool.
//   - options - ContainerAppsSessionPoolsClientBeginDeleteOptions contains the optional parameters for the ContainerAppsSessionPoolsClient.BeginDelete
//     method.
func (client *ContainerAppsSessionPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, sessionPoolName string, options *ContainerAppsSessionPoolsClientBeginDeleteOptions) (*runtime.Poller[ContainerAppsSessionPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sessionPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerAppsSessionPoolsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerAppsSessionPoolsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Session Pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *ContainerAppsSessionPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, sessionPoolName string, options *ContainerAppsSessionPoolsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerAppsSessionPoolsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sessionPoolName, options)
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
func (client *ContainerAppsSessionPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sessionPoolName string, options *ContainerAppsSessionPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/sessionPools/{sessionPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sessionPoolName == "" {
		return nil, errors.New("parameter sessionPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sessionPoolName}", url.PathEscape(sessionPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the properties of a Session Pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sessionPoolName - Name of the Session Pool.
//   - options - ContainerAppsSessionPoolsClientGetOptions contains the optional parameters for the ContainerAppsSessionPoolsClient.Get
//     method.
func (client *ContainerAppsSessionPoolsClient) Get(ctx context.Context, resourceGroupName string, sessionPoolName string, options *ContainerAppsSessionPoolsClientGetOptions) (ContainerAppsSessionPoolsClientGetResponse, error) {
	var err error
	const operationName = "ContainerAppsSessionPoolsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, sessionPoolName, options)
	if err != nil {
		return ContainerAppsSessionPoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsSessionPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsSessionPoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ContainerAppsSessionPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, sessionPoolName string, options *ContainerAppsSessionPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/sessionPools/{sessionPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sessionPoolName == "" {
		return nil, errors.New("parameter sessionPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sessionPoolName}", url.PathEscape(sessionPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainerAppsSessionPoolsClient) getHandleResponse(resp *http.Response) (ContainerAppsSessionPoolsClientGetResponse, error) {
	result := ContainerAppsSessionPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionPool); err != nil {
		return ContainerAppsSessionPoolsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get the Session Pools in a given subscription.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ContainerAppsSessionPoolsClientListByResourceGroupOptions contains the optional parameters for the ContainerAppsSessionPoolsClient.NewListByResourceGroupPager
//     method.
func (client *ContainerAppsSessionPoolsClient) NewListByResourceGroupPager(resourceGroupName string, options *ContainerAppsSessionPoolsClientListByResourceGroupOptions) *runtime.Pager[ContainerAppsSessionPoolsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsSessionPoolsClientListByResourceGroupResponse]{
		More: func(page ContainerAppsSessionPoolsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsSessionPoolsClientListByResourceGroupResponse) (ContainerAppsSessionPoolsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContainerAppsSessionPoolsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ContainerAppsSessionPoolsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ContainerAppsSessionPoolsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ContainerAppsSessionPoolsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/sessionPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ContainerAppsSessionPoolsClient) listByResourceGroupHandleResponse(resp *http.Response) (ContainerAppsSessionPoolsClientListByResourceGroupResponse, error) {
	result := ContainerAppsSessionPoolsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionPoolCollection); err != nil {
		return ContainerAppsSessionPoolsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get the Session Pools in a given subscription.
//
// Generated from API version 2024-02-02-preview
//   - options - ContainerAppsSessionPoolsClientListBySubscriptionOptions contains the optional parameters for the ContainerAppsSessionPoolsClient.NewListBySubscriptionPager
//     method.
func (client *ContainerAppsSessionPoolsClient) NewListBySubscriptionPager(options *ContainerAppsSessionPoolsClientListBySubscriptionOptions) *runtime.Pager[ContainerAppsSessionPoolsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsSessionPoolsClientListBySubscriptionResponse]{
		More: func(page ContainerAppsSessionPoolsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsSessionPoolsClientListBySubscriptionResponse) (ContainerAppsSessionPoolsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContainerAppsSessionPoolsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ContainerAppsSessionPoolsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ContainerAppsSessionPoolsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ContainerAppsSessionPoolsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.App/sessionPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ContainerAppsSessionPoolsClient) listBySubscriptionHandleResponse(resp *http.Response) (ContainerAppsSessionPoolsClientListBySubscriptionResponse, error) {
	result := ContainerAppsSessionPoolsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionPoolCollection); err != nil {
		return ContainerAppsSessionPoolsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patches a Session Pool using JSON Merge Patch
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sessionPoolName - Name of the Session Pool.
//   - sessionPoolEnvelope - Properties used to create a session pool
//   - options - ContainerAppsSessionPoolsClientBeginUpdateOptions contains the optional parameters for the ContainerAppsSessionPoolsClient.BeginUpdate
//     method.
func (client *ContainerAppsSessionPoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, sessionPoolName string, sessionPoolEnvelope SessionPoolPatchProperties, options *ContainerAppsSessionPoolsClientBeginUpdateOptions) (*runtime.Poller[ContainerAppsSessionPoolsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, sessionPoolName, sessionPoolEnvelope, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerAppsSessionPoolsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerAppsSessionPoolsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patches a Session Pool using JSON Merge Patch
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *ContainerAppsSessionPoolsClient) update(ctx context.Context, resourceGroupName string, sessionPoolName string, sessionPoolEnvelope SessionPoolPatchProperties, options *ContainerAppsSessionPoolsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerAppsSessionPoolsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sessionPoolName, sessionPoolEnvelope, options)
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

// updateCreateRequest creates the Update request.
func (client *ContainerAppsSessionPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sessionPoolName string, sessionPoolEnvelope SessionPoolPatchProperties, options *ContainerAppsSessionPoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/sessionPools/{sessionPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sessionPoolName == "" {
		return nil, errors.New("parameter sessionPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sessionPoolName}", url.PathEscape(sessionPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sessionPoolEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}
