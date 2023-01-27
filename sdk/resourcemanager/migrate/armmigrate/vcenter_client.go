//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmigrate

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

// VCenterClient contains the methods for the VCenter group.
// Don't use this type directly, use NewVCenterClient() instead.
type VCenterClient struct {
	host string
	pl   runtime.Pipeline
}

// NewVCenterClient creates a new instance of VCenterClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVCenterClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*VCenterClient, error) {
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
	client := &VCenterClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// DeleteVCenter - Method to delete vCenter in site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - vcenterName - VCenter ARM name.
//   - options - VCenterClientDeleteVCenterOptions contains the optional parameters for the VCenterClient.DeleteVCenter method.
func (client *VCenterClient) DeleteVCenter(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, options *VCenterClientDeleteVCenterOptions) (VCenterClientDeleteVCenterResponse, error) {
	req, err := client.deleteVCenterCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, vcenterName, options)
	if err != nil {
		return VCenterClientDeleteVCenterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VCenterClientDeleteVCenterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return VCenterClientDeleteVCenterResponse{}, runtime.NewResponseError(resp)
	}
	return VCenterClientDeleteVCenterResponse{}, nil
}

// deleteVCenterCreateRequest creates the DeleteVCenter request.
func (client *VCenterClient) deleteVCenterCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, options *VCenterClientDeleteVCenterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/vCenters/{vcenterName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// NewGetAllVCentersInSitePager - Method to get all vCenters in a site.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - options - VCenterClientGetAllVCentersInSiteOptions contains the optional parameters for the VCenterClient.NewGetAllVCentersInSitePager
//     method.
func (client *VCenterClient) NewGetAllVCentersInSitePager(subscriptionID string, resourceGroupName string, siteName string, options *VCenterClientGetAllVCentersInSiteOptions) *runtime.Pager[VCenterClientGetAllVCentersInSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[VCenterClientGetAllVCentersInSiteResponse]{
		More: func(page VCenterClientGetAllVCentersInSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VCenterClientGetAllVCentersInSiteResponse) (VCenterClientGetAllVCentersInSiteResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getAllVCentersInSiteCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VCenterClientGetAllVCentersInSiteResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VCenterClientGetAllVCentersInSiteResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VCenterClientGetAllVCentersInSiteResponse{}, runtime.NewResponseError(resp)
			}
			return client.getAllVCentersInSiteHandleResponse(resp)
		},
	})
}

// getAllVCentersInSiteCreateRequest creates the GetAllVCentersInSite request.
func (client *VCenterClient) getAllVCentersInSiteCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, options *VCenterClientGetAllVCentersInSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/vCenters"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAllVCentersInSiteHandleResponse handles the GetAllVCentersInSite response.
func (client *VCenterClient) getAllVCentersInSiteHandleResponse(resp *http.Response) (VCenterClientGetAllVCentersInSiteResponse, error) {
	result := VCenterClientGetAllVCentersInSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VCenterCollection); err != nil {
		return VCenterClientGetAllVCentersInSiteResponse{}, err
	}
	return result, nil
}

// GetVCenter - Method to get a vCenter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - vcenterName - VCenter ARM name.
//   - options - VCenterClientGetVCenterOptions contains the optional parameters for the VCenterClient.GetVCenter method.
func (client *VCenterClient) GetVCenter(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, options *VCenterClientGetVCenterOptions) (VCenterClientGetVCenterResponse, error) {
	req, err := client.getVCenterCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, vcenterName, options)
	if err != nil {
		return VCenterClientGetVCenterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VCenterClientGetVCenterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VCenterClientGetVCenterResponse{}, runtime.NewResponseError(resp)
	}
	return client.getVCenterHandleResponse(resp)
}

// getVCenterCreateRequest creates the GetVCenter request.
func (client *VCenterClient) getVCenterCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, options *VCenterClientGetVCenterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/vCenters/{vcenterName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getVCenterHandleResponse handles the GetVCenter response.
func (client *VCenterClient) getVCenterHandleResponse(resp *http.Response) (VCenterClientGetVCenterResponse, error) {
	result := VCenterClientGetVCenterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VCenter); err != nil {
		return VCenterClientGetVCenterResponse{}, err
	}
	return result, nil
}

// PutVCenter - Method to create or update a vCenter in site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - vcenterName - VCenter ARM name.
//   - body - Put vCenter body.
//   - options - VCenterClientPutVCenterOptions contains the optional parameters for the VCenterClient.PutVCenter method.
func (client *VCenterClient) PutVCenter(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, body VCenter, options *VCenterClientPutVCenterOptions) (VCenterClientPutVCenterResponse, error) {
	req, err := client.putVCenterCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, vcenterName, body, options)
	if err != nil {
		return VCenterClientPutVCenterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VCenterClientPutVCenterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return VCenterClientPutVCenterResponse{}, runtime.NewResponseError(resp)
	}
	return client.putVCenterHandleResponse(resp)
}

// putVCenterCreateRequest creates the PutVCenter request.
func (client *VCenterClient) putVCenterCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, vcenterName string, body VCenter, options *VCenterClientPutVCenterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/vCenters/{vcenterName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, body)
}

// putVCenterHandleResponse handles the PutVCenter response.
func (client *VCenterClient) putVCenterHandleResponse(resp *http.Response) (VCenterClientPutVCenterResponse, error) {
	result := VCenterClientPutVCenterResponse{}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	return result, nil
}
