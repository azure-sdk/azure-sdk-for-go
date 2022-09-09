//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualNetworkTapsClient contains the methods for the VirtualNetworkTaps group.
// Don't use this type directly, use NewVirtualNetworkTapsClient() instead.
type VirtualNetworkTapsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualNetworkTapsClient creates a new instance of VirtualNetworkTapsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualNetworkTapsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworkTapsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualNetworkTapsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Virtual Network Tap.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// tapName - The name of the virtual network tap.
// parameters - Parameters supplied to the create or update virtual network tap operation.
// options - VirtualNetworkTapsClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkTapsClient.BeginCreateOrUpdate
// method.
func (client *VirtualNetworkTapsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap, options *VirtualNetworkTapsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualNetworkTapsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, tapName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualNetworkTapsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworkTapsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a Virtual Network Tap.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *VirtualNetworkTapsClient) createOrUpdate(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap, options *VirtualNetworkTapsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, tapName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkTapsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, tapName string, parameters VirtualNetworkTap, options *VirtualNetworkTapsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if tapName == "" {
		return nil, errors.New("parameter tapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified virtual network tap.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// tapName - The name of the virtual network tap.
// options - VirtualNetworkTapsClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkTapsClient.BeginDelete
// method.
func (client *VirtualNetworkTapsClient) BeginDelete(ctx context.Context, resourceGroupName string, tapName string, options *VirtualNetworkTapsClientBeginDeleteOptions) (*runtime.Poller[VirtualNetworkTapsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, tapName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualNetworkTapsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworkTapsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified virtual network tap.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *VirtualNetworkTapsClient) deleteOperation(ctx context.Context, resourceGroupName string, tapName string, options *VirtualNetworkTapsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, tapName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworkTapsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, tapName string, options *VirtualNetworkTapsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if tapName == "" {
		return nil, errors.New("parameter tapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified virtual network tap.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// tapName - The name of virtual network tap.
// options - VirtualNetworkTapsClientGetOptions contains the optional parameters for the VirtualNetworkTapsClient.Get method.
func (client *VirtualNetworkTapsClient) Get(ctx context.Context, resourceGroupName string, tapName string, options *VirtualNetworkTapsClientGetOptions) (VirtualNetworkTapsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, tapName, options)
	if err != nil {
		return VirtualNetworkTapsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworkTapsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkTapsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkTapsClient) getCreateRequest(ctx context.Context, resourceGroupName string, tapName string, options *VirtualNetworkTapsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if tapName == "" {
		return nil, errors.New("parameter tapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkTapsClient) getHandleResponse(resp *http.Response) (VirtualNetworkTapsClientGetResponse, error) {
	result := VirtualNetworkTapsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkTap); err != nil {
		return VirtualNetworkTapsClientGetResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all the VirtualNetworkTaps in a subscription.
// Generated from API version 2022-01-01
// options - VirtualNetworkTapsClientListAllOptions contains the optional parameters for the VirtualNetworkTapsClient.ListAll
// method.
func (client *VirtualNetworkTapsClient) NewListAllPager(options *VirtualNetworkTapsClientListAllOptions) *runtime.Pager[VirtualNetworkTapsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworkTapsClientListAllResponse]{
		More: func(page VirtualNetworkTapsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworkTapsClientListAllResponse) (VirtualNetworkTapsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworkTapsClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualNetworkTapsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworkTapsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *VirtualNetworkTapsClient) listAllCreateRequest(ctx context.Context, options *VirtualNetworkTapsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworkTaps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *VirtualNetworkTapsClient) listAllHandleResponse(resp *http.Response) (VirtualNetworkTapsClientListAllResponse, error) {
	result := VirtualNetworkTapsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkTapListResult); err != nil {
		return VirtualNetworkTapsClientListAllResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all the VirtualNetworkTaps in a subscription.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// options - VirtualNetworkTapsClientListByResourceGroupOptions contains the optional parameters for the VirtualNetworkTapsClient.ListByResourceGroup
// method.
func (client *VirtualNetworkTapsClient) NewListByResourceGroupPager(resourceGroupName string, options *VirtualNetworkTapsClientListByResourceGroupOptions) *runtime.Pager[VirtualNetworkTapsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworkTapsClientListByResourceGroupResponse]{
		More: func(page VirtualNetworkTapsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworkTapsClientListByResourceGroupResponse) (VirtualNetworkTapsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworkTapsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualNetworkTapsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworkTapsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualNetworkTapsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualNetworkTapsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualNetworkTapsClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualNetworkTapsClientListByResourceGroupResponse, error) {
	result := VirtualNetworkTapsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkTapListResult); err != nil {
		return VirtualNetworkTapsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates an VirtualNetworkTap tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// tapName - The name of the tap.
// tapParameters - Parameters supplied to update VirtualNetworkTap tags.
// options - VirtualNetworkTapsClientUpdateTagsOptions contains the optional parameters for the VirtualNetworkTapsClient.UpdateTags
// method.
func (client *VirtualNetworkTapsClient) UpdateTags(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject, options *VirtualNetworkTapsClientUpdateTagsOptions) (VirtualNetworkTapsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, tapName, tapParameters, options)
	if err != nil {
		return VirtualNetworkTapsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworkTapsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkTapsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VirtualNetworkTapsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, tapName string, tapParameters TagsObject, options *VirtualNetworkTapsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if tapName == "" {
		return nil, errors.New("parameter tapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tapName}", url.PathEscape(tapName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, tapParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VirtualNetworkTapsClient) updateTagsHandleResponse(resp *http.Response) (VirtualNetworkTapsClientUpdateTagsResponse, error) {
	result := VirtualNetworkTapsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkTap); err != nil {
		return VirtualNetworkTapsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
