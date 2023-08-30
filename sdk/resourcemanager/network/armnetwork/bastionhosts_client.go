//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// BastionHostsClient contains the methods for the BastionHosts group.
// Don't use this type directly, use NewBastionHostsClient() instead.
type BastionHostsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBastionHostsClient creates a new instance of BastionHostsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBastionHostsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BastionHostsClient, error) {
	cl, err := arm.NewClient(moduleName+".BastionHostsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BastionHostsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the specified Bastion Host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - bastionHostName - The name of the Bastion Host.
//   - parameters - Parameters supplied to the create or update Bastion Host operation.
//   - options - BastionHostsClientBeginCreateOrUpdateOptions contains the optional parameters for the BastionHostsClient.BeginCreateOrUpdate
//     method.
func (client *BastionHostsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, bastionHostName string, parameters BastionHost, options *BastionHostsClientBeginCreateOrUpdateOptions) (*runtime.Poller[BastionHostsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, bastionHostName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BastionHostsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[BastionHostsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates the specified Bastion Host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *BastionHostsClient) createOrUpdate(ctx context.Context, resourceGroupName string, bastionHostName string, parameters BastionHost, options *BastionHostsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, bastionHostName, parameters, options)
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
func (client *BastionHostsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, parameters BastionHost, options *BastionHostsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// BeginDelete - Deletes the specified Bastion Host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - bastionHostName - The name of the Bastion Host.
//   - options - BastionHostsClientBeginDeleteOptions contains the optional parameters for the BastionHostsClient.BeginDelete
//     method.
func (client *BastionHostsClient) BeginDelete(ctx context.Context, resourceGroupName string, bastionHostName string, options *BastionHostsClientBeginDeleteOptions) (*runtime.Poller[BastionHostsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, bastionHostName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BastionHostsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[BastionHostsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified Bastion Host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *BastionHostsClient) deleteOperation(ctx context.Context, resourceGroupName string, bastionHostName string, options *BastionHostsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, bastionHostName, options)
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
func (client *BastionHostsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, options *BastionHostsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// Get - Gets the specified Bastion Host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - bastionHostName - The name of the Bastion Host.
//   - options - BastionHostsClientGetOptions contains the optional parameters for the BastionHostsClient.Get method.
func (client *BastionHostsClient) Get(ctx context.Context, resourceGroupName string, bastionHostName string, options *BastionHostsClientGetOptions) (BastionHostsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, bastionHostName, options)
	if err != nil {
		return BastionHostsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BastionHostsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BastionHostsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BastionHostsClient) getCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, options *BastionHostsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *BastionHostsClient) getHandleResponse(resp *http.Response) (BastionHostsClientGetResponse, error) {
	result := BastionHostsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BastionHost); err != nil {
		return BastionHostsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all Bastion Hosts in a subscription.
//
// Generated from API version 2023-06-01
//   - options - BastionHostsClientListOptions contains the optional parameters for the BastionHostsClient.NewListPager method.
func (client *BastionHostsClient) NewListPager(options *BastionHostsClientListOptions) *runtime.Pager[BastionHostsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BastionHostsClientListResponse]{
		More: func(page BastionHostsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BastionHostsClientListResponse) (BastionHostsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BastionHostsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BastionHostsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BastionHostsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BastionHostsClient) listCreateRequest(ctx context.Context, options *BastionHostsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/bastionHosts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listHandleResponse handles the List response.
func (client *BastionHostsClient) listHandleResponse(resp *http.Response) (BastionHostsClientListResponse, error) {
	result := BastionHostsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BastionHostListResult); err != nil {
		return BastionHostsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all Bastion Hosts in a resource group.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - options - BastionHostsClientListByResourceGroupOptions contains the optional parameters for the BastionHostsClient.NewListByResourceGroupPager
//     method.
func (client *BastionHostsClient) NewListByResourceGroupPager(resourceGroupName string, options *BastionHostsClientListByResourceGroupOptions) *runtime.Pager[BastionHostsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[BastionHostsClientListByResourceGroupResponse]{
		More: func(page BastionHostsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BastionHostsClientListByResourceGroupResponse) (BastionHostsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BastionHostsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BastionHostsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BastionHostsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *BastionHostsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *BastionHostsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *BastionHostsClient) listByResourceGroupHandleResponse(resp *http.Response) (BastionHostsClientListByResourceGroupResponse, error) {
	result := BastionHostsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BastionHostListResult); err != nil {
		return BastionHostsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdateTags - Updates Tags for BastionHost resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - bastionHostName - The name of the Bastion Host.
//   - parameters - Parameters supplied to update BastionHost tags.
//   - options - BastionHostsClientBeginUpdateTagsOptions contains the optional parameters for the BastionHostsClient.BeginUpdateTags
//     method.
func (client *BastionHostsClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, bastionHostName string, parameters TagsObject, options *BastionHostsClientBeginUpdateTagsOptions) (*runtime.Poller[BastionHostsClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, bastionHostName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BastionHostsClientUpdateTagsResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[BastionHostsClientUpdateTagsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateTags - Updates Tags for BastionHost resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *BastionHostsClient) updateTags(ctx context.Context, resourceGroupName string, bastionHostName string, parameters TagsObject, options *BastionHostsClientBeginUpdateTagsOptions) (*http.Response, error) {
	var err error
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, bastionHostName, parameters, options)
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

// updateTagsCreateRequest creates the UpdateTags request.
func (client *BastionHostsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, parameters TagsObject, options *BastionHostsClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if bastionHostName == "" {
		return nil, errors.New("parameter bastionHostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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
