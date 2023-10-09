//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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

// BastionHostsClient contains the methods for the BastionHosts group.
// Don't use this type directly, use NewBastionHostsClient() instead.
type BastionHostsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBastionHostsClient creates a new instance of BastionHostsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
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

// BeginCreateOrUpdate - Create or replace an existing bastionHost. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - labName - The name of the lab.
//   - virtualNetworkName - The name of the virtual network.
//   - name - The name of the bastionhost.
//   - bastionHost - Profile of a Bastion Host
//   - options - BastionHostsClientBeginCreateOrUpdateOptions contains the optional parameters for the BastionHostsClient.BeginCreateOrUpdate
//     method.
func (client *BastionHostsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, name string, bastionHost BastionHost, options *BastionHostsClientBeginCreateOrUpdateOptions) (*runtime.Poller[BastionHostsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, labName, virtualNetworkName, name, bastionHost, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BastionHostsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[BastionHostsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or replace an existing bastionHost. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
func (client *BastionHostsClient) createOrUpdate(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, name string, bastionHost BastionHost, options *BastionHostsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, virtualNetworkName, name, bastionHost, options)
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
func (client *BastionHostsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, name string, bastionHost BastionHost, options *BastionHostsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{virtualNetworkName}/bastionhosts/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, bastionHost); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete bastionhost. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - labName - The name of the lab.
//   - virtualNetworkName - The name of the virtual network.
//   - name - The name of the bastionhost.
//   - options - BastionHostsClientBeginDeleteOptions contains the optional parameters for the BastionHostsClient.BeginDelete
//     method.
func (client *BastionHostsClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, name string, options *BastionHostsClientBeginDeleteOptions) (*runtime.Poller[BastionHostsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, labName, virtualNetworkName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BastionHostsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[BastionHostsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete bastionhost. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
func (client *BastionHostsClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, name string, options *BastionHostsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, virtualNetworkName, name, options)
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
func (client *BastionHostsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, name string, options *BastionHostsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{virtualNetworkName}/bastionhosts/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get bastionhost.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - labName - The name of the lab.
//   - virtualNetworkName - The name of the virtual network.
//   - name - The name of the bastionhost.
//   - options - BastionHostsClientGetOptions contains the optional parameters for the BastionHostsClient.Get method.
func (client *BastionHostsClient) Get(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, name string, options *BastionHostsClientGetOptions) (BastionHostsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, virtualNetworkName, name, options)
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
func (client *BastionHostsClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, name string, options *BastionHostsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{virtualNetworkName}/bastionhosts/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
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

// NewListPager - List bastionhosts in a given virtual network.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - labName - The name of the lab.
//   - virtualNetworkName - The name of the virtual network.
//   - options - BastionHostsClientListOptions contains the optional parameters for the BastionHostsClient.NewListPager method.
func (client *BastionHostsClient) NewListPager(resourceGroupName string, labName string, virtualNetworkName string, options *BastionHostsClientListOptions) *runtime.Pager[BastionHostsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BastionHostsClientListResponse]{
		More: func(page BastionHostsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BastionHostsClientListResponse) (BastionHostsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, labName, virtualNetworkName, options)
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
func (client *BastionHostsClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, options *BastionHostsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{virtualNetworkName}/bastionhosts"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BastionHostsClient) listHandleResponse(resp *http.Response) (BastionHostsClientListResponse, error) {
	result := BastionHostsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BastionHostList); err != nil {
		return BastionHostsClientListResponse{}, err
	}
	return result, nil
}

// Update - Allows modifying tags of bastionhosts. All other properties will be ignored.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - labName - The name of the lab.
//   - virtualNetworkName - The name of the virtual network.
//   - name - The name of the bastionhost.
//   - bastionHost - Allows modifying tags of bastionhosts. All other properties will be ignored.
//   - options - BastionHostsClientUpdateOptions contains the optional parameters for the BastionHostsClient.Update method.
func (client *BastionHostsClient) Update(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, name string, bastionHost BastionHostFragment, options *BastionHostsClientUpdateOptions) (BastionHostsClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, virtualNetworkName, name, bastionHost, options)
	if err != nil {
		return BastionHostsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BastionHostsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BastionHostsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *BastionHostsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, virtualNetworkName string, name string, bastionHost BastionHostFragment, options *BastionHostsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{virtualNetworkName}/bastionhosts/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, bastionHost); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *BastionHostsClient) updateHandleResponse(resp *http.Response) (BastionHostsClientUpdateResponse, error) {
	result := BastionHostsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BastionHost); err != nil {
		return BastionHostsClientUpdateResponse{}, err
	}
	return result, nil
}
