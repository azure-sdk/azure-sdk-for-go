//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armagrifood

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DataManagerForAgricultureResourcesClient contains the methods for the DataManagerForAgricultureResources group.
// Don't use this type directly, use NewDataManagerForAgricultureResourcesClient() instead.
type DataManagerForAgricultureResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataManagerForAgricultureResourcesClient creates a new instance of DataManagerForAgricultureResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataManagerForAgricultureResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataManagerForAgricultureResourcesClient, error) {
	cl, err := arm.NewClient(moduleName+".DataManagerForAgricultureResourcesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataManagerForAgricultureResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update Data Manager For Agriculture resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - dataManagerForAgricultureResourceName - DataManagerForAgriculture resource name.
//   - request - Data Manager For Agriculture resource create or update request object.
//   - options - DataManagerForAgricultureResourcesClientCreateOrUpdateOptions contains the optional parameters for the DataManagerForAgricultureResourcesClient.CreateOrUpdate
//     method.
func (client *DataManagerForAgricultureResourcesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, request DataManagerForAgriculture, options *DataManagerForAgricultureResourcesClientCreateOrUpdateOptions) (DataManagerForAgricultureResourcesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, dataManagerForAgricultureResourceName, request, options)
	if err != nil {
		return DataManagerForAgricultureResourcesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataManagerForAgricultureResourcesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DataManagerForAgricultureResourcesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataManagerForAgricultureResourcesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, request DataManagerForAgriculture, options *DataManagerForAgricultureResourcesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{dataManagerForAgricultureResourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerForAgricultureResourceName == "" {
		return nil, errors.New("parameter dataManagerForAgricultureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerForAgricultureResourceName}", url.PathEscape(dataManagerForAgricultureResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataManagerForAgricultureResourcesClient) createOrUpdateHandleResponse(resp *http.Response) (DataManagerForAgricultureResourcesClientCreateOrUpdateResponse, error) {
	result := DataManagerForAgricultureResourcesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataManagerForAgriculture); err != nil {
		return DataManagerForAgricultureResourcesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Data Manager For Agriculture resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - dataManagerForAgricultureResourceName - DataManagerForAgriculture resource name.
//   - options - DataManagerForAgricultureResourcesClientDeleteOptions contains the optional parameters for the DataManagerForAgricultureResourcesClient.Delete
//     method.
func (client *DataManagerForAgricultureResourcesClient) Delete(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, options *DataManagerForAgricultureResourcesClientDeleteOptions) (DataManagerForAgricultureResourcesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dataManagerForAgricultureResourceName, options)
	if err != nil {
		return DataManagerForAgricultureResourcesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataManagerForAgricultureResourcesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DataManagerForAgricultureResourcesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DataManagerForAgricultureResourcesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataManagerForAgricultureResourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, options *DataManagerForAgricultureResourcesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{dataManagerForAgricultureResourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerForAgricultureResourceName == "" {
		return nil, errors.New("parameter dataManagerForAgricultureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerForAgricultureResourceName}", url.PathEscape(dataManagerForAgricultureResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get DataManagerForAgriculture resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - dataManagerForAgricultureResourceName - DataManagerForAgriculture resource name.
//   - options - DataManagerForAgricultureResourcesClientGetOptions contains the optional parameters for the DataManagerForAgricultureResourcesClient.Get
//     method.
func (client *DataManagerForAgricultureResourcesClient) Get(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, options *DataManagerForAgricultureResourcesClientGetOptions) (DataManagerForAgricultureResourcesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dataManagerForAgricultureResourceName, options)
	if err != nil {
		return DataManagerForAgricultureResourcesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataManagerForAgricultureResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataManagerForAgricultureResourcesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataManagerForAgricultureResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, options *DataManagerForAgricultureResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{dataManagerForAgricultureResourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerForAgricultureResourceName == "" {
		return nil, errors.New("parameter dataManagerForAgricultureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerForAgricultureResourceName}", url.PathEscape(dataManagerForAgricultureResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataManagerForAgricultureResourcesClient) getHandleResponse(resp *http.Response) (DataManagerForAgricultureResourcesClientGetResponse, error) {
	result := DataManagerForAgricultureResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataManagerForAgriculture); err != nil {
		return DataManagerForAgricultureResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists the DataManagerForAgriculture instances for a resource group.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DataManagerForAgricultureResourcesClientListByResourceGroupOptions contains the optional parameters for the DataManagerForAgricultureResourcesClient.NewListByResourceGroupPager
//     method.
func (client *DataManagerForAgricultureResourcesClient) NewListByResourceGroupPager(resourceGroupName string, options *DataManagerForAgricultureResourcesClientListByResourceGroupOptions) *runtime.Pager[DataManagerForAgricultureResourcesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataManagerForAgricultureResourcesClientListByResourceGroupResponse]{
		More: func(page DataManagerForAgricultureResourcesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataManagerForAgricultureResourcesClientListByResourceGroupResponse) (DataManagerForAgricultureResourcesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataManagerForAgricultureResourcesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DataManagerForAgricultureResourcesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataManagerForAgricultureResourcesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DataManagerForAgricultureResourcesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DataManagerForAgricultureResourcesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MaxPageSize != nil {
		reqQP.Set("$maxPageSize", strconv.FormatInt(int64(*options.MaxPageSize), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DataManagerForAgricultureResourcesClient) listByResourceGroupHandleResponse(resp *http.Response) (DataManagerForAgricultureResourcesClientListByResourceGroupResponse, error) {
	result := DataManagerForAgricultureResourcesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataManagerForAgricultureListResponse); err != nil {
		return DataManagerForAgricultureResourcesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists the DataManagerForAgriculture instances for a subscription.
//
// Generated from API version 2023-06-01-preview
//   - options - DataManagerForAgricultureResourcesClientListBySubscriptionOptions contains the optional parameters for the DataManagerForAgricultureResourcesClient.NewListBySubscriptionPager
//     method.
func (client *DataManagerForAgricultureResourcesClient) NewListBySubscriptionPager(options *DataManagerForAgricultureResourcesClientListBySubscriptionOptions) *runtime.Pager[DataManagerForAgricultureResourcesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataManagerForAgricultureResourcesClientListBySubscriptionResponse]{
		More: func(page DataManagerForAgricultureResourcesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataManagerForAgricultureResourcesClientListBySubscriptionResponse) (DataManagerForAgricultureResourcesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataManagerForAgricultureResourcesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DataManagerForAgricultureResourcesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataManagerForAgricultureResourcesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DataManagerForAgricultureResourcesClient) listBySubscriptionCreateRequest(ctx context.Context, options *DataManagerForAgricultureResourcesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AgFoodPlatform/farmBeats"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MaxPageSize != nil {
		reqQP.Set("$maxPageSize", strconv.FormatInt(int64(*options.MaxPageSize), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DataManagerForAgricultureResourcesClient) listBySubscriptionHandleResponse(resp *http.Response) (DataManagerForAgricultureResourcesClientListBySubscriptionResponse, error) {
	result := DataManagerForAgricultureResourcesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataManagerForAgricultureListResponse); err != nil {
		return DataManagerForAgricultureResourcesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a Data Manager For Agriculture resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - dataManagerForAgricultureResourceName - DataManagerForAgriculture resource name.
//   - request - Request object.
//   - options - DataManagerForAgricultureResourcesClientBeginUpdateOptions contains the optional parameters for the DataManagerForAgricultureResourcesClient.BeginUpdate
//     method.
func (client *DataManagerForAgricultureResourcesClient) BeginUpdate(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, request DataManagerForAgricultureUpdateRequestModel, options *DataManagerForAgricultureResourcesClientBeginUpdateOptions) (*runtime.Poller[DataManagerForAgricultureResourcesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, dataManagerForAgricultureResourceName, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DataManagerForAgricultureResourcesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DataManagerForAgricultureResourcesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update a Data Manager For Agriculture resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *DataManagerForAgricultureResourcesClient) update(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, request DataManagerForAgricultureUpdateRequestModel, options *DataManagerForAgricultureResourcesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dataManagerForAgricultureResourceName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DataManagerForAgricultureResourcesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, request DataManagerForAgricultureUpdateRequestModel, options *DataManagerForAgricultureResourcesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{dataManagerForAgricultureResourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerForAgricultureResourceName == "" {
		return nil, errors.New("parameter dataManagerForAgricultureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerForAgricultureResourceName}", url.PathEscape(dataManagerForAgricultureResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}
