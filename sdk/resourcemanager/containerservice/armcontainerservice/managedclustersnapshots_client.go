//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

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

// ManagedClusterSnapshotsClient contains the methods for the ManagedClusterSnapshots group.
// Don't use this type directly, use NewManagedClusterSnapshotsClient() instead.
type ManagedClusterSnapshotsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedClusterSnapshotsClient creates a new instance of ManagedClusterSnapshotsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedClusterSnapshotsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedClusterSnapshotsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedClusterSnapshotsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a managed cluster snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - parameters - The managed cluster snapshot to create or update.
//   - options - ManagedClusterSnapshotsClientCreateOrUpdateOptions contains the optional parameters for the ManagedClusterSnapshotsClient.CreateOrUpdate
//     method.
func (client *ManagedClusterSnapshotsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters ManagedClusterSnapshot, options *ManagedClusterSnapshotsClientCreateOrUpdateOptions) (ManagedClusterSnapshotsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ManagedClusterSnapshotsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return ManagedClusterSnapshotsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedClusterSnapshotsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ManagedClusterSnapshotsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedClusterSnapshotsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters ManagedClusterSnapshot, options *ManagedClusterSnapshotsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclustersnapshots/{resourceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagedClusterSnapshotsClient) createOrUpdateHandleResponse(resp *http.Response) (ManagedClusterSnapshotsClientCreateOrUpdateResponse, error) {
	result := ManagedClusterSnapshotsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedClusterSnapshot); err != nil {
		return ManagedClusterSnapshotsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a managed cluster snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - options - ManagedClusterSnapshotsClientDeleteOptions contains the optional parameters for the ManagedClusterSnapshotsClient.Delete
//     method.
func (client *ManagedClusterSnapshotsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, options *ManagedClusterSnapshotsClientDeleteOptions) (ManagedClusterSnapshotsClientDeleteResponse, error) {
	var err error
	const operationName = "ManagedClusterSnapshotsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ManagedClusterSnapshotsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedClusterSnapshotsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ManagedClusterSnapshotsClientDeleteResponse{}, err
	}
	return ManagedClusterSnapshotsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagedClusterSnapshotsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ManagedClusterSnapshotsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclustersnapshots/{resourceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a managed cluster snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - options - ManagedClusterSnapshotsClientGetOptions contains the optional parameters for the ManagedClusterSnapshotsClient.Get
//     method.
func (client *ManagedClusterSnapshotsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *ManagedClusterSnapshotsClientGetOptions) (ManagedClusterSnapshotsClientGetResponse, error) {
	var err error
	const operationName = "ManagedClusterSnapshotsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ManagedClusterSnapshotsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedClusterSnapshotsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedClusterSnapshotsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedClusterSnapshotsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ManagedClusterSnapshotsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclustersnapshots/{resourceName}"
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
	reqQP.Set("api-version", "2023-11-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedClusterSnapshotsClient) getHandleResponse(resp *http.Response) (ManagedClusterSnapshotsClientGetResponse, error) {
	result := ManagedClusterSnapshotsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedClusterSnapshot); err != nil {
		return ManagedClusterSnapshotsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of managed cluster snapshots in the specified subscription.
//
// Generated from API version 2023-11-02-preview
//   - options - ManagedClusterSnapshotsClientListOptions contains the optional parameters for the ManagedClusterSnapshotsClient.NewListPager
//     method.
func (client *ManagedClusterSnapshotsClient) NewListPager(options *ManagedClusterSnapshotsClientListOptions) *runtime.Pager[ManagedClusterSnapshotsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedClusterSnapshotsClientListResponse]{
		More: func(page ManagedClusterSnapshotsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedClusterSnapshotsClientListResponse) (ManagedClusterSnapshotsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedClusterSnapshotsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ManagedClusterSnapshotsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ManagedClusterSnapshotsClient) listCreateRequest(ctx context.Context, options *ManagedClusterSnapshotsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/managedclustersnapshots"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedClusterSnapshotsClient) listHandleResponse(resp *http.Response) (ManagedClusterSnapshotsClientListResponse, error) {
	result := ManagedClusterSnapshotsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedClusterSnapshotListResult); err != nil {
		return ManagedClusterSnapshotsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists managed cluster snapshots in the specified subscription and resource group.
//
// Generated from API version 2023-11-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ManagedClusterSnapshotsClientListByResourceGroupOptions contains the optional parameters for the ManagedClusterSnapshotsClient.NewListByResourceGroupPager
//     method.
func (client *ManagedClusterSnapshotsClient) NewListByResourceGroupPager(resourceGroupName string, options *ManagedClusterSnapshotsClientListByResourceGroupOptions) *runtime.Pager[ManagedClusterSnapshotsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedClusterSnapshotsClientListByResourceGroupResponse]{
		More: func(page ManagedClusterSnapshotsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedClusterSnapshotsClientListByResourceGroupResponse) (ManagedClusterSnapshotsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedClusterSnapshotsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ManagedClusterSnapshotsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ManagedClusterSnapshotsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ManagedClusterSnapshotsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclustersnapshots"
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
	reqQP.Set("api-version", "2023-11-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ManagedClusterSnapshotsClient) listByResourceGroupHandleResponse(resp *http.Response) (ManagedClusterSnapshotsClientListByResourceGroupResponse, error) {
	result := ManagedClusterSnapshotsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedClusterSnapshotListResult); err != nil {
		return ManagedClusterSnapshotsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates tags on a managed cluster snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - parameters - Parameters supplied to the Update managed cluster snapshot Tags operation.
//   - options - ManagedClusterSnapshotsClientUpdateTagsOptions contains the optional parameters for the ManagedClusterSnapshotsClient.UpdateTags
//     method.
func (client *ManagedClusterSnapshotsClient) UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, parameters TagsObject, options *ManagedClusterSnapshotsClientUpdateTagsOptions) (ManagedClusterSnapshotsClientUpdateTagsResponse, error) {
	var err error
	const operationName = "ManagedClusterSnapshotsClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return ManagedClusterSnapshotsClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedClusterSnapshotsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedClusterSnapshotsClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ManagedClusterSnapshotsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters TagsObject, options *ManagedClusterSnapshotsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclustersnapshots/{resourceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ManagedClusterSnapshotsClient) updateTagsHandleResponse(resp *http.Response) (ManagedClusterSnapshotsClientUpdateTagsResponse, error) {
	result := ManagedClusterSnapshotsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedClusterSnapshot); err != nil {
		return ManagedClusterSnapshotsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
