//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

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

// DataNetworksClient contains the methods for the DataNetworks group.
// Don't use this type directly, use NewDataNetworksClient() instead.
type DataNetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataNetworksClient creates a new instance of DataNetworksClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataNetworksClient, error) {
	cl, err := arm.NewClient(moduleName+".DataNetworksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataNetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a data network. Must be created in the same location as its parent mobile network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - dataNetworkName - The name of the data network.
//   - parameters - Parameters supplied to the create or update data network operation.
//   - options - DataNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the DataNetworksClient.BeginCreateOrUpdate
//     method.
func (client *DataNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, parameters DataNetwork, options *DataNetworksClientBeginCreateOrUpdateOptions) (*runtime.Poller[DataNetworksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DataNetworksClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DataNetworksClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a data network. Must be created in the same location as its parent mobile network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *DataNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, parameters DataNetwork, options *DataNetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, parameters, options)
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
func (client *DataNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, parameters DataNetwork, options *DataNetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/dataNetworks/{dataNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if dataNetworkName == "" {
		return nil, errors.New("parameter dataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataNetworkName}", url.PathEscape(dataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified data network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - dataNetworkName - The name of the data network.
//   - options - DataNetworksClientBeginDeleteOptions contains the optional parameters for the DataNetworksClient.BeginDelete
//     method.
func (client *DataNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, options *DataNetworksClientBeginDeleteOptions) (*runtime.Poller[DataNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DataNetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DataNetworksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified data network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *DataNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, options *DataNetworksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, options)
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
func (client *DataNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, options *DataNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/dataNetworks/{dataNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if dataNetworkName == "" {
		return nil, errors.New("parameter dataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataNetworkName}", url.PathEscape(dataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified data network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - dataNetworkName - The name of the data network.
//   - options - DataNetworksClientGetOptions contains the optional parameters for the DataNetworksClient.Get method.
func (client *DataNetworksClient) Get(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, options *DataNetworksClientGetOptions) (DataNetworksClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, options)
	if err != nil {
		return DataNetworksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataNetworksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DataNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, options *DataNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/dataNetworks/{dataNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if dataNetworkName == "" {
		return nil, errors.New("parameter dataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataNetworkName}", url.PathEscape(dataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataNetworksClient) getHandleResponse(resp *http.Response) (DataNetworksClientGetResponse, error) {
	result := DataNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataNetwork); err != nil {
		return DataNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByMobileNetworkPager - Lists all data networks in the mobile network.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - options - DataNetworksClientListByMobileNetworkOptions contains the optional parameters for the DataNetworksClient.NewListByMobileNetworkPager
//     method.
func (client *DataNetworksClient) NewListByMobileNetworkPager(resourceGroupName string, mobileNetworkName string, options *DataNetworksClientListByMobileNetworkOptions) *runtime.Pager[DataNetworksClientListByMobileNetworkResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataNetworksClientListByMobileNetworkResponse]{
		More: func(page DataNetworksClientListByMobileNetworkResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataNetworksClientListByMobileNetworkResponse) (DataNetworksClientListByMobileNetworkResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByMobileNetworkCreateRequest(ctx, resourceGroupName, mobileNetworkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataNetworksClientListByMobileNetworkResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DataNetworksClientListByMobileNetworkResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataNetworksClientListByMobileNetworkResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByMobileNetworkHandleResponse(resp)
		},
	})
}

// listByMobileNetworkCreateRequest creates the ListByMobileNetwork request.
func (client *DataNetworksClient) listByMobileNetworkCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, options *DataNetworksClientListByMobileNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/dataNetworks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMobileNetworkHandleResponse handles the ListByMobileNetwork response.
func (client *DataNetworksClient) listByMobileNetworkHandleResponse(resp *http.Response) (DataNetworksClientListByMobileNetworkResponse, error) {
	result := DataNetworksClientListByMobileNetworkResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataNetworkListResult); err != nil {
		return DataNetworksClientListByMobileNetworkResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates data network tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - dataNetworkName - The name of the data network.
//   - parameters - Parameters supplied to update data network tags.
//   - options - DataNetworksClientUpdateTagsOptions contains the optional parameters for the DataNetworksClient.UpdateTags method.
func (client *DataNetworksClient) UpdateTags(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, parameters TagsObject, options *DataNetworksClientUpdateTagsOptions) (DataNetworksClientUpdateTagsResponse, error) {
	var err error
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, mobileNetworkName, dataNetworkName, parameters, options)
	if err != nil {
		return DataNetworksClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataNetworksClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataNetworksClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *DataNetworksClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, dataNetworkName string, parameters TagsObject, options *DataNetworksClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/dataNetworks/{dataNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if dataNetworkName == "" {
		return nil, errors.New("parameter dataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataNetworkName}", url.PathEscape(dataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *DataNetworksClient) updateTagsHandleResponse(resp *http.Response) (DataNetworksClientUpdateTagsResponse, error) {
	result := DataNetworksClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataNetwork); err != nil {
		return DataNetworksClientUpdateTagsResponse{}, err
	}
	return result, nil
}
