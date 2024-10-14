//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceregistry

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

// AssetEndpointProfilesClient contains the methods for the AssetEndpointProfiles group.
// Don't use this type directly, use NewAssetEndpointProfilesClient() instead.
type AssetEndpointProfilesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssetEndpointProfilesClient creates a new instance of AssetEndpointProfilesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssetEndpointProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssetEndpointProfilesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssetEndpointProfilesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrReplace - Create a AssetEndpointProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - assetEndpointProfileName - Asset Endpoint Profile name parameter.
//   - resource - Resource create parameters.
//   - options - AssetEndpointProfilesClientBeginCreateOrReplaceOptions contains the optional parameters for the AssetEndpointProfilesClient.BeginCreateOrReplace
//     method.
func (client *AssetEndpointProfilesClient) BeginCreateOrReplace(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, resource AssetEndpointProfile, options *AssetEndpointProfilesClientBeginCreateOrReplaceOptions) (*runtime.Poller[AssetEndpointProfilesClientCreateOrReplaceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrReplace(ctx, resourceGroupName, assetEndpointProfileName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AssetEndpointProfilesClientCreateOrReplaceResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AssetEndpointProfilesClientCreateOrReplaceResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrReplace - Create a AssetEndpointProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *AssetEndpointProfilesClient) createOrReplace(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, resource AssetEndpointProfile, options *AssetEndpointProfilesClientBeginCreateOrReplaceOptions) (*http.Response, error) {
	var err error
	const operationName = "AssetEndpointProfilesClient.BeginCreateOrReplace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrReplaceCreateRequest(ctx, resourceGroupName, assetEndpointProfileName, resource, options)
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

// createOrReplaceCreateRequest creates the CreateOrReplace request.
func (client *AssetEndpointProfilesClient) createOrReplaceCreateRequest(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, resource AssetEndpointProfile, options *AssetEndpointProfilesClientBeginCreateOrReplaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/{assetEndpointProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if assetEndpointProfileName == "" {
		return nil, errors.New("parameter assetEndpointProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetEndpointProfileName}", url.PathEscape(assetEndpointProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a AssetEndpointProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - assetEndpointProfileName - Asset Endpoint Profile name parameter.
//   - options - AssetEndpointProfilesClientBeginDeleteOptions contains the optional parameters for the AssetEndpointProfilesClient.BeginDelete
//     method.
func (client *AssetEndpointProfilesClient) BeginDelete(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, options *AssetEndpointProfilesClientBeginDeleteOptions) (*runtime.Poller[AssetEndpointProfilesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, assetEndpointProfileName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AssetEndpointProfilesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AssetEndpointProfilesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a AssetEndpointProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *AssetEndpointProfilesClient) deleteOperation(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, options *AssetEndpointProfilesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AssetEndpointProfilesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, assetEndpointProfileName, options)
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
func (client *AssetEndpointProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, options *AssetEndpointProfilesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/{assetEndpointProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if assetEndpointProfileName == "" {
		return nil, errors.New("parameter assetEndpointProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetEndpointProfileName}", url.PathEscape(assetEndpointProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a AssetEndpointProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - assetEndpointProfileName - Asset Endpoint Profile name parameter.
//   - options - AssetEndpointProfilesClientGetOptions contains the optional parameters for the AssetEndpointProfilesClient.Get
//     method.
func (client *AssetEndpointProfilesClient) Get(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, options *AssetEndpointProfilesClientGetOptions) (AssetEndpointProfilesClientGetResponse, error) {
	var err error
	const operationName = "AssetEndpointProfilesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, assetEndpointProfileName, options)
	if err != nil {
		return AssetEndpointProfilesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssetEndpointProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssetEndpointProfilesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AssetEndpointProfilesClient) getCreateRequest(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, options *AssetEndpointProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/{assetEndpointProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if assetEndpointProfileName == "" {
		return nil, errors.New("parameter assetEndpointProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetEndpointProfileName}", url.PathEscape(assetEndpointProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssetEndpointProfilesClient) getHandleResponse(resp *http.Response) (AssetEndpointProfilesClientGetResponse, error) {
	result := AssetEndpointProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetEndpointProfile); err != nil {
		return AssetEndpointProfilesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List AssetEndpointProfile resources by resource group
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - AssetEndpointProfilesClientListByResourceGroupOptions contains the optional parameters for the AssetEndpointProfilesClient.NewListByResourceGroupPager
//     method.
func (client *AssetEndpointProfilesClient) NewListByResourceGroupPager(resourceGroupName string, options *AssetEndpointProfilesClientListByResourceGroupOptions) *runtime.Pager[AssetEndpointProfilesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssetEndpointProfilesClientListByResourceGroupResponse]{
		More: func(page AssetEndpointProfilesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssetEndpointProfilesClientListByResourceGroupResponse) (AssetEndpointProfilesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssetEndpointProfilesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AssetEndpointProfilesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AssetEndpointProfilesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AssetEndpointProfilesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceRegistry/assetEndpointProfiles"
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
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AssetEndpointProfilesClient) listByResourceGroupHandleResponse(resp *http.Response) (AssetEndpointProfilesClientListByResourceGroupResponse, error) {
	result := AssetEndpointProfilesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetEndpointProfileListResult); err != nil {
		return AssetEndpointProfilesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List AssetEndpointProfile resources by subscription ID
//
// Generated from API version 2024-11-01
//   - options - AssetEndpointProfilesClientListBySubscriptionOptions contains the optional parameters for the AssetEndpointProfilesClient.NewListBySubscriptionPager
//     method.
func (client *AssetEndpointProfilesClient) NewListBySubscriptionPager(options *AssetEndpointProfilesClientListBySubscriptionOptions) *runtime.Pager[AssetEndpointProfilesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssetEndpointProfilesClientListBySubscriptionResponse]{
		More: func(page AssetEndpointProfilesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssetEndpointProfilesClientListBySubscriptionResponse) (AssetEndpointProfilesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssetEndpointProfilesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AssetEndpointProfilesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AssetEndpointProfilesClient) listBySubscriptionCreateRequest(ctx context.Context, options *AssetEndpointProfilesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DeviceRegistry/assetEndpointProfiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AssetEndpointProfilesClient) listBySubscriptionHandleResponse(resp *http.Response) (AssetEndpointProfilesClientListBySubscriptionResponse, error) {
	result := AssetEndpointProfilesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetEndpointProfileListResult); err != nil {
		return AssetEndpointProfilesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a AssetEndpointProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - assetEndpointProfileName - Asset Endpoint Profile name parameter.
//   - properties - The resource properties to be updated.
//   - options - AssetEndpointProfilesClientBeginUpdateOptions contains the optional parameters for the AssetEndpointProfilesClient.BeginUpdate
//     method.
func (client *AssetEndpointProfilesClient) BeginUpdate(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, properties AssetEndpointProfileUpdate, options *AssetEndpointProfilesClientBeginUpdateOptions) (*runtime.Poller[AssetEndpointProfilesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, assetEndpointProfileName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AssetEndpointProfilesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AssetEndpointProfilesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a AssetEndpointProfile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *AssetEndpointProfilesClient) update(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, properties AssetEndpointProfileUpdate, options *AssetEndpointProfilesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AssetEndpointProfilesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, assetEndpointProfileName, properties, options)
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
func (client *AssetEndpointProfilesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, assetEndpointProfileName string, properties AssetEndpointProfileUpdate, options *AssetEndpointProfilesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/{assetEndpointProfileName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if assetEndpointProfileName == "" {
		return nil, errors.New("parameter assetEndpointProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetEndpointProfileName}", url.PathEscape(assetEndpointProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
