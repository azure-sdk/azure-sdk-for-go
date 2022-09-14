//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

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

// DedicatedHostGroupsClient contains the methods for the DedicatedHostGroups group.
// Don't use this type directly, use NewDedicatedHostGroupsClient() instead.
type DedicatedHostGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDedicatedHostGroupsClient creates a new instance of DedicatedHostGroupsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDedicatedHostGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DedicatedHostGroupsClient, error) {
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
	client := &DedicatedHostGroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a dedicated host group. For details of Dedicated Host and Dedicated Host Groups please
// see Dedicated Host Documentation [https://go.microsoft.com/fwlink/?linkid=2082596]
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - The name of the resource group.
// hostGroupName - The name of the dedicated host group.
// parameters - Parameters supplied to the Create Dedicated Host Group.
// options - DedicatedHostGroupsClientCreateOrUpdateOptions contains the optional parameters for the DedicatedHostGroupsClient.CreateOrUpdate
// method.
func (client *DedicatedHostGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroup, options *DedicatedHostGroupsClientCreateOrUpdateOptions) (DedicatedHostGroupsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostGroupName, parameters, options)
	if err != nil {
		return DedicatedHostGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DedicatedHostGroupsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DedicatedHostGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroup, options *DedicatedHostGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DedicatedHostGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (DedicatedHostGroupsClientCreateOrUpdateResponse, error) {
	result := DedicatedHostGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroup); err != nil {
		return DedicatedHostGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a dedicated host group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - The name of the resource group.
// hostGroupName - The name of the dedicated host group.
// options - DedicatedHostGroupsClientDeleteOptions contains the optional parameters for the DedicatedHostGroupsClient.Delete
// method.
func (client *DedicatedHostGroupsClient) Delete(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsClientDeleteOptions) (DedicatedHostGroupsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostGroupName, options)
	if err != nil {
		return DedicatedHostGroupsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DedicatedHostGroupsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DedicatedHostGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DedicatedHostGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about a dedicated host group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - The name of the resource group.
// hostGroupName - The name of the dedicated host group.
// options - DedicatedHostGroupsClientGetOptions contains the optional parameters for the DedicatedHostGroupsClient.Get method.
func (client *DedicatedHostGroupsClient) Get(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsClientGetOptions) (DedicatedHostGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostGroupName, options)
	if err != nil {
		return DedicatedHostGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedHostGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DedicatedHostGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedHostGroupsClient) getHandleResponse(resp *http.Response) (DedicatedHostGroupsClientGetResponse, error) {
	result := DedicatedHostGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroup); err != nil {
		return DedicatedHostGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all of the dedicated host groups in the specified resource group. Use the nextLink
// property in the response to get the next page of dedicated host groups.
// Generated from API version 2022-03-01
// resourceGroupName - The name of the resource group.
// options - DedicatedHostGroupsClientListByResourceGroupOptions contains the optional parameters for the DedicatedHostGroupsClient.ListByResourceGroup
// method.
func (client *DedicatedHostGroupsClient) NewListByResourceGroupPager(resourceGroupName string, options *DedicatedHostGroupsClientListByResourceGroupOptions) *runtime.Pager[DedicatedHostGroupsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DedicatedHostGroupsClientListByResourceGroupResponse]{
		More: func(page DedicatedHostGroupsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHostGroupsClientListByResourceGroupResponse) (DedicatedHostGroupsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DedicatedHostGroupsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DedicatedHostGroupsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DedicatedHostGroupsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DedicatedHostGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DedicatedHostGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups"
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
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DedicatedHostGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (DedicatedHostGroupsClientListByResourceGroupResponse, error) {
	result := DedicatedHostGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroupListResult); err != nil {
		return DedicatedHostGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all of the dedicated host groups in the subscription. Use the nextLink property in the
// response to get the next page of dedicated host groups.
// Generated from API version 2022-03-01
// options - DedicatedHostGroupsClientListBySubscriptionOptions contains the optional parameters for the DedicatedHostGroupsClient.ListBySubscription
// method.
func (client *DedicatedHostGroupsClient) NewListBySubscriptionPager(options *DedicatedHostGroupsClientListBySubscriptionOptions) *runtime.Pager[DedicatedHostGroupsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DedicatedHostGroupsClientListBySubscriptionResponse]{
		More: func(page DedicatedHostGroupsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHostGroupsClientListBySubscriptionResponse) (DedicatedHostGroupsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DedicatedHostGroupsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DedicatedHostGroupsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DedicatedHostGroupsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DedicatedHostGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DedicatedHostGroupsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/hostGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DedicatedHostGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (DedicatedHostGroupsClientListBySubscriptionResponse, error) {
	result := DedicatedHostGroupsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroupListResult); err != nil {
		return DedicatedHostGroupsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an dedicated host group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - The name of the resource group.
// hostGroupName - The name of the dedicated host group.
// parameters - Parameters supplied to the Update Dedicated Host Group operation.
// options - DedicatedHostGroupsClientUpdateOptions contains the optional parameters for the DedicatedHostGroupsClient.Update
// method.
func (client *DedicatedHostGroupsClient) Update(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroupUpdate, options *DedicatedHostGroupsClientUpdateOptions) (DedicatedHostGroupsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostGroupName, parameters, options)
	if err != nil {
		return DedicatedHostGroupsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedHostGroupsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DedicatedHostGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, parameters DedicatedHostGroupUpdate, options *DedicatedHostGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *DedicatedHostGroupsClient) updateHandleResponse(resp *http.Response) (DedicatedHostGroupsClientUpdateResponse, error) {
	result := DedicatedHostGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostGroup); err != nil {
		return DedicatedHostGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
