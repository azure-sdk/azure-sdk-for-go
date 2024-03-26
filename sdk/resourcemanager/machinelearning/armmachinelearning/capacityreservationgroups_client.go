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

// CapacityReservationGroupsClient contains the methods for the CapacityReservationGroups group.
// Don't use this type directly, use NewCapacityReservationGroupsClient() instead.
type CapacityReservationGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCapacityReservationGroupsClient creates a new instance of CapacityReservationGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCapacityReservationGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CapacityReservationGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CapacityReservationGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update CapacityReservationGroup
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - groupID - Group ID
//   - body - Capacity Reservation Group payload to create
//   - options - CapacityReservationGroupsClientCreateOrUpdateOptions contains the optional parameters for the CapacityReservationGroupsClient.CreateOrUpdate
//     method.
func (client *CapacityReservationGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, groupID string, body CapacityReservationGroup, options *CapacityReservationGroupsClientCreateOrUpdateOptions) (CapacityReservationGroupsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "CapacityReservationGroupsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, groupID, body, options)
	if err != nil {
		return CapacityReservationGroupsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapacityReservationGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return CapacityReservationGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CapacityReservationGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, groupID string, body CapacityReservationGroup, options *CapacityReservationGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/capacityReserverationGroups/{groupId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CapacityReservationGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (CapacityReservationGroupsClientCreateOrUpdateResponse, error) {
	result := CapacityReservationGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroup); err != nil {
		return CapacityReservationGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete CapacityReservationGroup
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - groupID - Group ID
//   - options - CapacityReservationGroupsClientDeleteOptions contains the optional parameters for the CapacityReservationGroupsClient.Delete
//     method.
func (client *CapacityReservationGroupsClient) Delete(ctx context.Context, resourceGroupName string, groupID string, options *CapacityReservationGroupsClientDeleteOptions) (CapacityReservationGroupsClientDeleteResponse, error) {
	var err error
	const operationName = "CapacityReservationGroupsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, groupID, options)
	if err != nil {
		return CapacityReservationGroupsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapacityReservationGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CapacityReservationGroupsClientDeleteResponse{}, err
	}
	return CapacityReservationGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CapacityReservationGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, groupID string, options *CapacityReservationGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/capacityReserverationGroups/{groupId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get CapacityReservationGroup
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - groupID - Group ID
//   - options - CapacityReservationGroupsClientGetOptions contains the optional parameters for the CapacityReservationGroupsClient.Get
//     method.
func (client *CapacityReservationGroupsClient) Get(ctx context.Context, resourceGroupName string, groupID string, options *CapacityReservationGroupsClientGetOptions) (CapacityReservationGroupsClientGetResponse, error) {
	var err error
	const operationName = "CapacityReservationGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, groupID, options)
	if err != nil {
		return CapacityReservationGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapacityReservationGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CapacityReservationGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CapacityReservationGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, groupID string, options *CapacityReservationGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/capacityReserverationGroups/{groupId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CapacityReservationGroupsClient) getHandleResponse(resp *http.Response) (CapacityReservationGroupsClientGetResponse, error) {
	result := CapacityReservationGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroup); err != nil {
		return CapacityReservationGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists CapacityReservationGroups
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - CapacityReservationGroupsClientListOptions contains the optional parameters for the CapacityReservationGroupsClient.NewListPager
//     method.
func (client *CapacityReservationGroupsClient) NewListPager(resourceGroupName string, options *CapacityReservationGroupsClientListOptions) *runtime.Pager[CapacityReservationGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapacityReservationGroupsClientListResponse]{
		More: func(page CapacityReservationGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapacityReservationGroupsClientListResponse) (CapacityReservationGroupsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CapacityReservationGroupsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return CapacityReservationGroupsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CapacityReservationGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *CapacityReservationGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/capacityReserverationGroups"
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
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CapacityReservationGroupsClient) listHandleResponse(resp *http.Response) (CapacityReservationGroupsClientListResponse, error) {
	result := CapacityReservationGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroupTrackedResourceArmPaginatedResult); err != nil {
		return CapacityReservationGroupsClientListResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List CapacityReservationGroups by subscription
//
// Generated from API version 2024-04-01-preview
//   - options - CapacityReservationGroupsClientListBySubscriptionOptions contains the optional parameters for the CapacityReservationGroupsClient.NewListBySubscriptionPager
//     method.
func (client *CapacityReservationGroupsClient) NewListBySubscriptionPager(options *CapacityReservationGroupsClientListBySubscriptionOptions) *runtime.Pager[CapacityReservationGroupsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapacityReservationGroupsClientListBySubscriptionResponse]{
		More: func(page CapacityReservationGroupsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapacityReservationGroupsClientListBySubscriptionResponse) (CapacityReservationGroupsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CapacityReservationGroupsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CapacityReservationGroupsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CapacityReservationGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CapacityReservationGroupsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/capacityReserverationGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CapacityReservationGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (CapacityReservationGroupsClientListBySubscriptionResponse, error) {
	result := CapacityReservationGroupsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroupTrackedResourceArmPaginatedResult); err != nil {
		return CapacityReservationGroupsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update CapacityReservationGroup
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - groupID - Group ID
//   - body - Capacity Reservation Group payload to update
//   - options - CapacityReservationGroupsClientUpdateOptions contains the optional parameters for the CapacityReservationGroupsClient.Update
//     method.
func (client *CapacityReservationGroupsClient) Update(ctx context.Context, resourceGroupName string, groupID string, body PartialMinimalTrackedResourceWithSKUAndIdentity, options *CapacityReservationGroupsClientUpdateOptions) (CapacityReservationGroupsClientUpdateResponse, error) {
	var err error
	const operationName = "CapacityReservationGroupsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, groupID, body, options)
	if err != nil {
		return CapacityReservationGroupsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapacityReservationGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CapacityReservationGroupsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *CapacityReservationGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, groupID string, body PartialMinimalTrackedResourceWithSKUAndIdentity, options *CapacityReservationGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/capacityReserverationGroups/{groupId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *CapacityReservationGroupsClient) updateHandleResponse(resp *http.Response) (CapacityReservationGroupsClientUpdateResponse, error) {
	result := CapacityReservationGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroup); err != nil {
		return CapacityReservationGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
