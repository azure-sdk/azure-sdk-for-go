// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armservicenetworking

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

// TrafficControllerInterfaceClient contains the methods for the TrafficControllerInterface group.
// Don't use this type directly, use NewTrafficControllerInterfaceClient() instead.
type TrafficControllerInterfaceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTrafficControllerInterfaceClient creates a new instance of TrafficControllerInterfaceClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTrafficControllerInterfaceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TrafficControllerInterfaceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TrafficControllerInterfaceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a TrafficController
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - resource - Resource create parameters.
//   - options - TrafficControllerInterfaceClientBeginCreateOrUpdateOptions contains the optional parameters for the TrafficControllerInterfaceClient.BeginCreateOrUpdate
//     method.
func (client *TrafficControllerInterfaceClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, trafficControllerName string, resource TrafficController, options *TrafficControllerInterfaceClientBeginCreateOrUpdateOptions) (*runtime.Poller[TrafficControllerInterfaceClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, trafficControllerName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TrafficControllerInterfaceClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TrafficControllerInterfaceClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a TrafficController
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
func (client *TrafficControllerInterfaceClient) createOrUpdate(ctx context.Context, resourceGroupName string, trafficControllerName string, resource TrafficController, options *TrafficControllerInterfaceClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "TrafficControllerInterfaceClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, trafficControllerName, resource, options)
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
func (client *TrafficControllerInterfaceClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, resource TrafficController, _ *TrafficControllerInterfaceClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a TrafficController
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - options - TrafficControllerInterfaceClientBeginDeleteOptions contains the optional parameters for the TrafficControllerInterfaceClient.BeginDelete
//     method.
func (client *TrafficControllerInterfaceClient) BeginDelete(ctx context.Context, resourceGroupName string, trafficControllerName string, options *TrafficControllerInterfaceClientBeginDeleteOptions) (*runtime.Poller[TrafficControllerInterfaceClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, trafficControllerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TrafficControllerInterfaceClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TrafficControllerInterfaceClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a TrafficController
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
func (client *TrafficControllerInterfaceClient) deleteOperation(ctx context.Context, resourceGroupName string, trafficControllerName string, options *TrafficControllerInterfaceClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "TrafficControllerInterfaceClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, trafficControllerName, options)
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
func (client *TrafficControllerInterfaceClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, _ *TrafficControllerInterfaceClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a TrafficController
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - options - TrafficControllerInterfaceClientGetOptions contains the optional parameters for the TrafficControllerInterfaceClient.Get
//     method.
func (client *TrafficControllerInterfaceClient) Get(ctx context.Context, resourceGroupName string, trafficControllerName string, options *TrafficControllerInterfaceClientGetOptions) (TrafficControllerInterfaceClientGetResponse, error) {
	var err error
	const operationName = "TrafficControllerInterfaceClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, trafficControllerName, options)
	if err != nil {
		return TrafficControllerInterfaceClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TrafficControllerInterfaceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TrafficControllerInterfaceClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TrafficControllerInterfaceClient) getCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, _ *TrafficControllerInterfaceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TrafficControllerInterfaceClient) getHandleResponse(resp *http.Response) (TrafficControllerInterfaceClientGetResponse, error) {
	result := TrafficControllerInterfaceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrafficController); err != nil {
		return TrafficControllerInterfaceClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List TrafficController resources by resource group
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - TrafficControllerInterfaceClientListByResourceGroupOptions contains the optional parameters for the TrafficControllerInterfaceClient.NewListByResourceGroupPager
//     method.
func (client *TrafficControllerInterfaceClient) NewListByResourceGroupPager(resourceGroupName string, options *TrafficControllerInterfaceClientListByResourceGroupOptions) *runtime.Pager[TrafficControllerInterfaceClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[TrafficControllerInterfaceClientListByResourceGroupResponse]{
		More: func(page TrafficControllerInterfaceClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TrafficControllerInterfaceClientListByResourceGroupResponse) (TrafficControllerInterfaceClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TrafficControllerInterfaceClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return TrafficControllerInterfaceClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *TrafficControllerInterfaceClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *TrafficControllerInterfaceClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers"
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
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *TrafficControllerInterfaceClient) listByResourceGroupHandleResponse(resp *http.Response) (TrafficControllerInterfaceClientListByResourceGroupResponse, error) {
	result := TrafficControllerInterfaceClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrafficControllerListResult); err != nil {
		return TrafficControllerInterfaceClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List TrafficController resources by subscription ID
//
// Generated from API version 2025-03-01-preview
//   - options - TrafficControllerInterfaceClientListBySubscriptionOptions contains the optional parameters for the TrafficControllerInterfaceClient.NewListBySubscriptionPager
//     method.
func (client *TrafficControllerInterfaceClient) NewListBySubscriptionPager(options *TrafficControllerInterfaceClientListBySubscriptionOptions) *runtime.Pager[TrafficControllerInterfaceClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[TrafficControllerInterfaceClientListBySubscriptionResponse]{
		More: func(page TrafficControllerInterfaceClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TrafficControllerInterfaceClientListBySubscriptionResponse) (TrafficControllerInterfaceClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TrafficControllerInterfaceClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TrafficControllerInterfaceClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *TrafficControllerInterfaceClient) listBySubscriptionCreateRequest(ctx context.Context, _ *TrafficControllerInterfaceClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceNetworking/trafficControllers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *TrafficControllerInterfaceClient) listBySubscriptionHandleResponse(resp *http.Response) (TrafficControllerInterfaceClientListBySubscriptionResponse, error) {
	result := TrafficControllerInterfaceClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrafficControllerListResult); err != nil {
		return TrafficControllerInterfaceClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a TrafficController
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - properties - The resource properties to be updated.
//   - options - TrafficControllerInterfaceClientUpdateOptions contains the optional parameters for the TrafficControllerInterfaceClient.Update
//     method.
func (client *TrafficControllerInterfaceClient) Update(ctx context.Context, resourceGroupName string, trafficControllerName string, properties TrafficControllerUpdate, options *TrafficControllerInterfaceClientUpdateOptions) (TrafficControllerInterfaceClientUpdateResponse, error) {
	var err error
	const operationName = "TrafficControllerInterfaceClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, trafficControllerName, properties, options)
	if err != nil {
		return TrafficControllerInterfaceClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TrafficControllerInterfaceClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TrafficControllerInterfaceClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *TrafficControllerInterfaceClient) updateCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, properties TrafficControllerUpdate, _ *TrafficControllerInterfaceClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TrafficControllerInterfaceClient) updateHandleResponse(resp *http.Response) (TrafficControllerInterfaceClientUpdateResponse, error) {
	result := TrafficControllerInterfaceClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrafficController); err != nil {
		return TrafficControllerInterfaceClientUpdateResponse{}, err
	}
	return result, nil
}
