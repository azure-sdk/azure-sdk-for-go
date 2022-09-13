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

// VirtualNetworkPeeringsClient contains the methods for the VirtualNetworkPeerings group.
// Don't use this type directly, use NewVirtualNetworkPeeringsClient() instead.
type VirtualNetworkPeeringsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualNetworkPeeringsClient creates a new instance of VirtualNetworkPeeringsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualNetworkPeeringsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworkPeeringsClient, error) {
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
	client := &VirtualNetworkPeeringsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a peering in the specified virtual network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// virtualNetworkName - The name of the virtual network.
// virtualNetworkPeeringName - The name of the peering.
// virtualNetworkPeeringParameters - Parameters supplied to the create or update virtual network peering operation.
// options - VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkPeeringsClient.BeginCreateOrUpdate
// method.
func (client *VirtualNetworkPeeringsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters VirtualNetworkPeering, options *VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualNetworkPeeringsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, virtualNetworkPeeringParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualNetworkPeeringsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworkPeeringsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a peering in the specified virtual network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *VirtualNetworkPeeringsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters VirtualNetworkPeering, options *VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, virtualNetworkPeeringParameters, options)
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
func (client *VirtualNetworkPeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters VirtualNetworkPeering, options *VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if virtualNetworkPeeringName == "" {
		return nil, errors.New("parameter virtualNetworkPeeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkPeeringName}", url.PathEscape(virtualNetworkPeeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SyncRemoteAddressSpace != nil {
		reqQP.Set("syncRemoteAddressSpace", string(*options.SyncRemoteAddressSpace))
	}
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, virtualNetworkPeeringParameters)
}

// BeginDelete - Deletes the specified virtual network peering.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// virtualNetworkName - The name of the virtual network.
// virtualNetworkPeeringName - The name of the virtual network peering.
// options - VirtualNetworkPeeringsClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkPeeringsClient.BeginDelete
// method.
func (client *VirtualNetworkPeeringsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *VirtualNetworkPeeringsClientBeginDeleteOptions) (*runtime.Poller[VirtualNetworkPeeringsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualNetworkPeeringsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworkPeeringsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified virtual network peering.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *VirtualNetworkPeeringsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *VirtualNetworkPeeringsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, options)
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
func (client *VirtualNetworkPeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *VirtualNetworkPeeringsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if virtualNetworkPeeringName == "" {
		return nil, errors.New("parameter virtualNetworkPeeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkPeeringName}", url.PathEscape(virtualNetworkPeeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified virtual network peering.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// virtualNetworkName - The name of the virtual network.
// virtualNetworkPeeringName - The name of the virtual network peering.
// options - VirtualNetworkPeeringsClientGetOptions contains the optional parameters for the VirtualNetworkPeeringsClient.Get
// method.
func (client *VirtualNetworkPeeringsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *VirtualNetworkPeeringsClientGetOptions) (VirtualNetworkPeeringsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName, options)
	if err != nil {
		return VirtualNetworkPeeringsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworkPeeringsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkPeeringsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkPeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *VirtualNetworkPeeringsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if virtualNetworkPeeringName == "" {
		return nil, errors.New("parameter virtualNetworkPeeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkPeeringName}", url.PathEscape(virtualNetworkPeeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkPeeringsClient) getHandleResponse(resp *http.Response) (VirtualNetworkPeeringsClientGetResponse, error) {
	result := VirtualNetworkPeeringsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkPeering); err != nil {
		return VirtualNetworkPeeringsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all virtual network peerings in a virtual network.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// virtualNetworkName - The name of the virtual network.
// options - VirtualNetworkPeeringsClientListOptions contains the optional parameters for the VirtualNetworkPeeringsClient.List
// method.
func (client *VirtualNetworkPeeringsClient) NewListPager(resourceGroupName string, virtualNetworkName string, options *VirtualNetworkPeeringsClientListOptions) *runtime.Pager[VirtualNetworkPeeringsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworkPeeringsClientListResponse]{
		More: func(page VirtualNetworkPeeringsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworkPeeringsClientListResponse) (VirtualNetworkPeeringsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworkPeeringsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualNetworkPeeringsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworkPeeringsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualNetworkPeeringsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworkPeeringsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualNetworkPeeringsClient) listHandleResponse(resp *http.Response) (VirtualNetworkPeeringsClientListResponse, error) {
	result := VirtualNetworkPeeringsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkPeeringListResult); err != nil {
		return VirtualNetworkPeeringsClientListResponse{}, err
	}
	return result, nil
}
