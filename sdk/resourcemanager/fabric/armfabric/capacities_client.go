//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfabric

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

// CapacitiesClient contains the methods for the FabricCapacities group.
// Don't use this type directly, use NewCapacitiesClient() instead.
type CapacitiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCapacitiesClient creates a new instance of CapacitiesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCapacitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CapacitiesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CapacitiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Implements local CheckNameAvailability operations
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - location - The name of the Azure region.
//   - body - The CheckAvailability request
//   - options - CapacitiesClientCheckNameAvailabilityOptions contains the optional parameters for the CapacitiesClient.CheckNameAvailability
//     method.
func (client *CapacitiesClient) CheckNameAvailability(ctx context.Context, location string, body CheckNameAvailabilityRequest, options *CapacitiesClientCheckNameAvailabilityOptions) (CapacitiesClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "CapacitiesClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, body, options)
	if err != nil {
		return CapacitiesClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapacitiesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CapacitiesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *CapacitiesClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, body CheckNameAvailabilityRequest, options *CapacitiesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Fabric/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *CapacitiesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (CapacitiesClientCheckNameAvailabilityResponse, error) {
	result := CapacitiesClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return CapacitiesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create a FabricCapacity
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - capacityName - The name of the Microsoft Fabric capacity. It must be a minimum of 3 characters, and a maximum of 63.
//   - resource - Resource create parameters.
//   - options - CapacitiesClientBeginCreateOrUpdateOptions contains the optional parameters for the CapacitiesClient.BeginCreateOrUpdate
//     method.
func (client *CapacitiesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, capacityName string, resource Capacity, options *CapacitiesClientBeginCreateOrUpdateOptions) (*runtime.Poller[CapacitiesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, capacityName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CapacitiesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CapacitiesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a FabricCapacity
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
func (client *CapacitiesClient) createOrUpdate(ctx context.Context, resourceGroupName string, capacityName string, resource Capacity, options *CapacitiesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CapacitiesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, capacityName, resource, options)
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
func (client *CapacitiesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, capacityName string, resource Capacity, options *CapacitiesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Fabric/capacities/{capacityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityName == "" {
		return nil, errors.New("parameter capacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityName}", url.PathEscape(capacityName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete a FabricCapacity
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - capacityName - The name of the Microsoft Fabric capacity. It must be a minimum of 3 characters, and a maximum of 63.
//   - options - CapacitiesClientDeleteOptions contains the optional parameters for the CapacitiesClient.Delete method.
func (client *CapacitiesClient) Delete(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientDeleteOptions) (CapacitiesClientDeleteResponse, error) {
	var err error
	const operationName = "CapacitiesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, capacityName, options)
	if err != nil {
		return CapacitiesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapacitiesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CapacitiesClientDeleteResponse{}, err
	}
	return CapacitiesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CapacitiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Fabric/capacities/{capacityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityName == "" {
		return nil, errors.New("parameter capacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityName}", url.PathEscape(capacityName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a FabricCapacity
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - capacityName - The name of the Microsoft Fabric capacity. It must be a minimum of 3 characters, and a maximum of 63.
//   - options - CapacitiesClientGetOptions contains the optional parameters for the CapacitiesClient.Get method.
func (client *CapacitiesClient) Get(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientGetOptions) (CapacitiesClientGetResponse, error) {
	var err error
	const operationName = "CapacitiesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, capacityName, options)
	if err != nil {
		return CapacitiesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapacitiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CapacitiesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CapacitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Fabric/capacities/{capacityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityName == "" {
		return nil, errors.New("parameter capacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityName}", url.PathEscape(capacityName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CapacitiesClient) getHandleResponse(resp *http.Response) (CapacitiesClientGetResponse, error) {
	result := CapacitiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Capacity); err != nil {
		return CapacitiesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List FabricCapacity resources by resource group
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - CapacitiesClientListByResourceGroupOptions contains the optional parameters for the CapacitiesClient.NewListByResourceGroupPager
//     method.
func (client *CapacitiesClient) NewListByResourceGroupPager(resourceGroupName string, options *CapacitiesClientListByResourceGroupOptions) *runtime.Pager[CapacitiesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapacitiesClientListByResourceGroupResponse]{
		More: func(page CapacitiesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapacitiesClientListByResourceGroupResponse) (CapacitiesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CapacitiesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return CapacitiesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CapacitiesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CapacitiesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Fabric/capacities"
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
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CapacitiesClient) listByResourceGroupHandleResponse(resp *http.Response) (CapacitiesClientListByResourceGroupResponse, error) {
	result := CapacitiesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityListResult); err != nil {
		return CapacitiesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List FabricCapacity resources by subscription ID
//
// Generated from API version 2023-11-01
//   - options - CapacitiesClientListBySubscriptionOptions contains the optional parameters for the CapacitiesClient.NewListBySubscriptionPager
//     method.
func (client *CapacitiesClient) NewListBySubscriptionPager(options *CapacitiesClientListBySubscriptionOptions) *runtime.Pager[CapacitiesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapacitiesClientListBySubscriptionResponse]{
		More: func(page CapacitiesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapacitiesClientListBySubscriptionResponse) (CapacitiesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CapacitiesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CapacitiesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CapacitiesClient) listBySubscriptionCreateRequest(ctx context.Context, options *CapacitiesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Fabric/capacities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CapacitiesClient) listBySubscriptionHandleResponse(resp *http.Response) (CapacitiesClientListBySubscriptionResponse, error) {
	result := CapacitiesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityListResult); err != nil {
		return CapacitiesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListSKUsPager - List eligible SKUs for Microsoft Fabric resource provider
//
// Generated from API version 2023-11-01
//   - options - CapacitiesClientListSKUsOptions contains the optional parameters for the CapacitiesClient.NewListSKUsPager method.
func (client *CapacitiesClient) NewListSKUsPager(options *CapacitiesClientListSKUsOptions) *runtime.Pager[CapacitiesClientListSKUsResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapacitiesClientListSKUsResponse]{
		More: func(page CapacitiesClientListSKUsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapacitiesClientListSKUsResponse) (CapacitiesClientListSKUsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CapacitiesClient.NewListSKUsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSKUsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CapacitiesClientListSKUsResponse{}, err
			}
			return client.listSKUsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *CapacitiesClient) listSKUsCreateRequest(ctx context.Context, options *CapacitiesClientListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Fabric/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *CapacitiesClient) listSKUsHandleResponse(resp *http.Response) (CapacitiesClientListSKUsResponse, error) {
	result := CapacitiesClientListSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RpSKUEnumerationForNewResourceResult); err != nil {
		return CapacitiesClientListSKUsResponse{}, err
	}
	return result, nil
}

// NewListSKUsForCapacityPager - List eligible SKUs for a Microsoft Fabric resource
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - capacityName - The name of the capacity.
//   - options - CapacitiesClientListSKUsForCapacityOptions contains the optional parameters for the CapacitiesClient.NewListSKUsForCapacityPager
//     method.
func (client *CapacitiesClient) NewListSKUsForCapacityPager(resourceGroupName string, capacityName string, options *CapacitiesClientListSKUsForCapacityOptions) *runtime.Pager[CapacitiesClientListSKUsForCapacityResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapacitiesClientListSKUsForCapacityResponse]{
		More: func(page CapacitiesClientListSKUsForCapacityResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapacitiesClientListSKUsForCapacityResponse) (CapacitiesClientListSKUsForCapacityResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CapacitiesClient.NewListSKUsForCapacityPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSKUsForCapacityCreateRequest(ctx, resourceGroupName, capacityName, options)
			}, nil)
			if err != nil {
				return CapacitiesClientListSKUsForCapacityResponse{}, err
			}
			return client.listSKUsForCapacityHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSKUsForCapacityCreateRequest creates the ListSKUsForCapacity request.
func (client *CapacitiesClient) listSKUsForCapacityCreateRequest(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientListSKUsForCapacityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Fabric/capacities/{capacityName}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityName == "" {
		return nil, errors.New("parameter capacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityName}", url.PathEscape(capacityName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSKUsForCapacityHandleResponse handles the ListSKUsForCapacity response.
func (client *CapacitiesClient) listSKUsForCapacityHandleResponse(resp *http.Response) (CapacitiesClientListSKUsForCapacityResponse, error) {
	result := CapacitiesClientListSKUsForCapacityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RpSKUEnumerationForExistingResourceResult); err != nil {
		return CapacitiesClientListSKUsForCapacityResponse{}, err
	}
	return result, nil
}

// BeginResume - Resume operation of the specified Fabric capacity instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - capacityName - The name of the Microsoft Fabric capacity. It must be a minimum of 3 characters, and a maximum of 63.
//   - options - CapacitiesClientBeginResumeOptions contains the optional parameters for the CapacitiesClient.BeginResume method.
func (client *CapacitiesClient) BeginResume(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientBeginResumeOptions) (*runtime.Poller[CapacitiesClientResumeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resume(ctx, resourceGroupName, capacityName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CapacitiesClientResumeResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CapacitiesClientResumeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Resume - Resume operation of the specified Fabric capacity instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
func (client *CapacitiesClient) resume(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientBeginResumeOptions) (*http.Response, error) {
	var err error
	const operationName = "CapacitiesClient.BeginResume"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.resumeCreateRequest(ctx, resourceGroupName, capacityName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// resumeCreateRequest creates the Resume request.
func (client *CapacitiesClient) resumeCreateRequest(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientBeginResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Fabric/capacities/{capacityName}/resume"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityName == "" {
		return nil, errors.New("parameter capacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityName}", url.PathEscape(capacityName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginSuspend - Suspend operation of the specified Fabric capacity instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - capacityName - The name of the Microsoft Fabric capacity. It must be a minimum of 3 characters, and a maximum of 63.
//   - options - CapacitiesClientBeginSuspendOptions contains the optional parameters for the CapacitiesClient.BeginSuspend method.
func (client *CapacitiesClient) BeginSuspend(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientBeginSuspendOptions) (*runtime.Poller[CapacitiesClientSuspendResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.suspend(ctx, resourceGroupName, capacityName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CapacitiesClientSuspendResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CapacitiesClientSuspendResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Suspend - Suspend operation of the specified Fabric capacity instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
func (client *CapacitiesClient) suspend(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientBeginSuspendOptions) (*http.Response, error) {
	var err error
	const operationName = "CapacitiesClient.BeginSuspend"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.suspendCreateRequest(ctx, resourceGroupName, capacityName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// suspendCreateRequest creates the Suspend request.
func (client *CapacitiesClient) suspendCreateRequest(ctx context.Context, resourceGroupName string, capacityName string, options *CapacitiesClientBeginSuspendOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Fabric/capacities/{capacityName}/suspend"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityName == "" {
		return nil, errors.New("parameter capacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityName}", url.PathEscape(capacityName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Update a FabricCapacity
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - capacityName - The name of the Microsoft Fabric capacity. It must be a minimum of 3 characters, and a maximum of 63.
//   - properties - The resource properties to be updated.
//   - options - CapacitiesClientUpdateOptions contains the optional parameters for the CapacitiesClient.Update method.
func (client *CapacitiesClient) Update(ctx context.Context, resourceGroupName string, capacityName string, properties CapacityUpdate, options *CapacitiesClientUpdateOptions) (CapacitiesClientUpdateResponse, error) {
	var err error
	const operationName = "CapacitiesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, capacityName, properties, options)
	if err != nil {
		return CapacitiesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapacitiesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CapacitiesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *CapacitiesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, capacityName string, properties CapacityUpdate, options *CapacitiesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Fabric/capacities/{capacityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityName == "" {
		return nil, errors.New("parameter capacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityName}", url.PathEscape(capacityName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *CapacitiesClient) updateHandleResponse(resp *http.Response) (CapacitiesClientUpdateResponse, error) {
	result := CapacitiesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Capacity); err != nil {
		return CapacitiesClientUpdateResponse{}, err
	}
	return result, nil
}
