//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmpcnetworkfunction

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

// NetworkFunctionsClient contains the methods for the NetworkFunctions group.
// Don't use this type directly, use NewNetworkFunctionsClient() instead.
type NetworkFunctionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkFunctionsClient creates a new instance of NetworkFunctionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkFunctionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkFunctionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkFunctionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a NetworkFunctionResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFunctionName - The name of the network function
//   - resource - Resource create parameters.
//   - options - NetworkFunctionsClientBeginCreateOrUpdateOptions contains the optional parameters for the NetworkFunctionsClient.BeginCreateOrUpdate
//     method.
func (client *NetworkFunctionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkFunctionName string, resource NetworkFunctionResource, options *NetworkFunctionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NetworkFunctionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkFunctionName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkFunctionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkFunctionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a NetworkFunctionResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
func (client *NetworkFunctionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkFunctionName string, resource NetworkFunctionResource, options *NetworkFunctionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkFunctionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkFunctionName, resource, options)
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
func (client *NetworkFunctionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkFunctionName string, resource NetworkFunctionResource, options *NetworkFunctionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/networkFunctions/{networkFunctionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFunctionName == "" {
		return nil, errors.New("parameter networkFunctionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionName}", url.PathEscape(networkFunctionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete a NetworkFunctionResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFunctionName - The name of the network function
//   - options - NetworkFunctionsClientDeleteOptions contains the optional parameters for the NetworkFunctionsClient.Delete method.
func (client *NetworkFunctionsClient) Delete(ctx context.Context, resourceGroupName string, networkFunctionName string, options *NetworkFunctionsClientDeleteOptions) (NetworkFunctionsClientDeleteResponse, error) {
	var err error
	const operationName = "NetworkFunctionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkFunctionName, options)
	if err != nil {
		return NetworkFunctionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkFunctionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NetworkFunctionsClientDeleteResponse{}, err
	}
	return NetworkFunctionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkFunctionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkFunctionName string, options *NetworkFunctionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/networkFunctions/{networkFunctionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFunctionName == "" {
		return nil, errors.New("parameter networkFunctionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionName}", url.PathEscape(networkFunctionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a NetworkFunctionResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFunctionName - The name of the network function
//   - options - NetworkFunctionsClientGetOptions contains the optional parameters for the NetworkFunctionsClient.Get method.
func (client *NetworkFunctionsClient) Get(ctx context.Context, resourceGroupName string, networkFunctionName string, options *NetworkFunctionsClientGetOptions) (NetworkFunctionsClientGetResponse, error) {
	var err error
	const operationName = "NetworkFunctionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkFunctionName, options)
	if err != nil {
		return NetworkFunctionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkFunctionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkFunctionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkFunctionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkFunctionName string, options *NetworkFunctionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/networkFunctions/{networkFunctionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFunctionName == "" {
		return nil, errors.New("parameter networkFunctionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionName}", url.PathEscape(networkFunctionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkFunctionsClient) getHandleResponse(resp *http.Response) (NetworkFunctionsClientGetResponse, error) {
	result := NetworkFunctionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionResource); err != nil {
		return NetworkFunctionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List NetworkFunctionResource resources by resource group
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - NetworkFunctionsClientListByResourceGroupOptions contains the optional parameters for the NetworkFunctionsClient.NewListByResourceGroupPager
//     method.
func (client *NetworkFunctionsClient) NewListByResourceGroupPager(resourceGroupName string, options *NetworkFunctionsClientListByResourceGroupOptions) *runtime.Pager[NetworkFunctionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkFunctionsClientListByResourceGroupResponse]{
		More: func(page NetworkFunctionsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkFunctionsClientListByResourceGroupResponse) (NetworkFunctionsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkFunctionsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return NetworkFunctionsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkFunctionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkFunctionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/networkFunctions"
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
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *NetworkFunctionsClient) listByResourceGroupHandleResponse(resp *http.Response) (NetworkFunctionsClientListByResourceGroupResponse, error) {
	result := NetworkFunctionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionResourceListResult); err != nil {
		return NetworkFunctionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List NetworkFunctionResource resources by subscription ID
//
// Generated from API version 2023-05-15-preview
//   - options - NetworkFunctionsClientListBySubscriptionOptions contains the optional parameters for the NetworkFunctionsClient.NewListBySubscriptionPager
//     method.
func (client *NetworkFunctionsClient) NewListBySubscriptionPager(options *NetworkFunctionsClientListBySubscriptionOptions) *runtime.Pager[NetworkFunctionsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkFunctionsClientListBySubscriptionResponse]{
		More: func(page NetworkFunctionsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkFunctionsClientListBySubscriptionResponse) (NetworkFunctionsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkFunctionsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return NetworkFunctionsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *NetworkFunctionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *NetworkFunctionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MobilePacketCore/networkFunctions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *NetworkFunctionsClient) listBySubscriptionHandleResponse(resp *http.Response) (NetworkFunctionsClientListBySubscriptionResponse, error) {
	result := NetworkFunctionsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionResourceListResult); err != nil {
		return NetworkFunctionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Update a NetworkFunctionResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFunctionName - The name of the network function
//   - properties - The resource properties to be updated.
//   - options - NetworkFunctionsClientUpdateTagsOptions contains the optional parameters for the NetworkFunctionsClient.UpdateTags
//     method.
func (client *NetworkFunctionsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkFunctionName string, properties NetworkFunctionResourceTagsUpdate, options *NetworkFunctionsClientUpdateTagsOptions) (NetworkFunctionsClientUpdateTagsResponse, error) {
	var err error
	const operationName = "NetworkFunctionsClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkFunctionName, properties, options)
	if err != nil {
		return NetworkFunctionsClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkFunctionsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkFunctionsClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *NetworkFunctionsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkFunctionName string, properties NetworkFunctionResourceTagsUpdate, options *NetworkFunctionsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/networkFunctions/{networkFunctionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFunctionName == "" {
		return nil, errors.New("parameter networkFunctionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionName}", url.PathEscape(networkFunctionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *NetworkFunctionsClient) updateTagsHandleResponse(resp *http.Response) (NetworkFunctionsClientUpdateTagsResponse, error) {
	result := NetworkFunctionsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionResource); err != nil {
		return NetworkFunctionsClientUpdateTagsResponse{}, err
	}
	return result, nil
}