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

// VirtualHubIPConfigurationClient contains the methods for the VirtualHubIPConfiguration group.
// Don't use this type directly, use NewVirtualHubIPConfigurationClient() instead.
type VirtualHubIPConfigurationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualHubIPConfigurationClient creates a new instance of VirtualHubIPConfigurationClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualHubIPConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualHubIPConfigurationClient, error) {
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
	client := &VirtualHubIPConfigurationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a VirtualHubIpConfiguration resource if it doesn't exist else updates the existing VirtualHubIpConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// ipConfigName - The name of the ipconfig.
// parameters - Hub Ip Configuration parameters.
// options - VirtualHubIPConfigurationClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualHubIPConfigurationClient.BeginCreateOrUpdate
// method.
func (client *VirtualHubIPConfigurationClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, parameters HubIPConfiguration, options *VirtualHubIPConfigurationClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualHubIPConfigurationClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, ipConfigName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualHubIPConfigurationClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualHubIPConfigurationClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a VirtualHubIpConfiguration resource if it doesn't exist else updates the existing VirtualHubIpConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
func (client *VirtualHubIPConfigurationClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, parameters HubIPConfiguration, options *VirtualHubIPConfigurationClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, ipConfigName, parameters, options)
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
func (client *VirtualHubIPConfigurationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, parameters HubIPConfiguration, options *VirtualHubIPConfigurationClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations/{ipConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if ipConfigName == "" {
		return nil, errors.New("parameter ipConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigName}", url.PathEscape(ipConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a VirtualHubIpConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - The resource group name of the VirtualHubBgpConnection.
// virtualHubName - The name of the VirtualHub.
// ipConfigName - The name of the ipconfig.
// options - VirtualHubIPConfigurationClientBeginDeleteOptions contains the optional parameters for the VirtualHubIPConfigurationClient.BeginDelete
// method.
func (client *VirtualHubIPConfigurationClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *VirtualHubIPConfigurationClientBeginDeleteOptions) (*runtime.Poller[VirtualHubIPConfigurationClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, ipConfigName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualHubIPConfigurationClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualHubIPConfigurationClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a VirtualHubIpConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
func (client *VirtualHubIPConfigurationClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *VirtualHubIPConfigurationClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, ipConfigName, options)
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
func (client *VirtualHubIPConfigurationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *VirtualHubIPConfigurationClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations/{ipConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if ipConfigName == "" {
		return nil, errors.New("parameter ipConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigName}", url.PathEscape(ipConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a Virtual Hub Ip configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// ipConfigName - The name of the ipconfig.
// options - VirtualHubIPConfigurationClientGetOptions contains the optional parameters for the VirtualHubIPConfigurationClient.Get
// method.
func (client *VirtualHubIPConfigurationClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *VirtualHubIPConfigurationClientGetOptions) (VirtualHubIPConfigurationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, ipConfigName, options)
	if err != nil {
		return VirtualHubIPConfigurationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualHubIPConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualHubIPConfigurationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualHubIPConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *VirtualHubIPConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations/{ipConfigName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if ipConfigName == "" {
		return nil, errors.New("parameter ipConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigName}", url.PathEscape(ipConfigName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualHubIPConfigurationClient) getHandleResponse(resp *http.Response) (VirtualHubIPConfigurationClientGetResponse, error) {
	result := VirtualHubIPConfigurationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HubIPConfiguration); err != nil {
		return VirtualHubIPConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves the details of all VirtualHubIpConfigurations.
// Generated from API version 2022-09-01
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// options - VirtualHubIPConfigurationClientListOptions contains the optional parameters for the VirtualHubIPConfigurationClient.List
// method.
func (client *VirtualHubIPConfigurationClient) NewListPager(resourceGroupName string, virtualHubName string, options *VirtualHubIPConfigurationClientListOptions) *runtime.Pager[VirtualHubIPConfigurationClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualHubIPConfigurationClientListResponse]{
		More: func(page VirtualHubIPConfigurationClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualHubIPConfigurationClientListResponse) (VirtualHubIPConfigurationClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualHubIPConfigurationClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualHubIPConfigurationClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualHubIPConfigurationClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualHubIPConfigurationClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubIPConfigurationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/ipConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualHubIPConfigurationClient) listHandleResponse(resp *http.Response) (VirtualHubIPConfigurationClientListResponse, error) {
	result := VirtualHubIPConfigurationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVirtualHubIPConfigurationResults); err != nil {
		return VirtualHubIPConfigurationClientListResponse{}, err
	}
	return result, nil
}
