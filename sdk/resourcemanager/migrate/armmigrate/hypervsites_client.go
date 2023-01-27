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

// HyperVSitesClient contains the methods for the HyperVSites group.
// Don't use this type directly, use NewHyperVSitesClient() instead.
type HyperVSitesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewHyperVSitesClient creates a new instance of HyperVSitesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHyperVSitesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*HyperVSitesClient, error) {
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
	client := &HyperVSitesClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// DeleteSite - Method to delete a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - options - HyperVSitesClientDeleteSiteOptions contains the optional parameters for the HyperVSitesClient.DeleteSite method.
func (client *HyperVSitesClient) DeleteSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, options *HyperVSitesClientDeleteSiteOptions) (HyperVSitesClientDeleteSiteResponse, error) {
	req, err := client.deleteSiteCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, options)
	if err != nil {
		return HyperVSitesClientDeleteSiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HyperVSitesClientDeleteSiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return HyperVSitesClientDeleteSiteResponse{}, runtime.NewResponseError(resp)
	}
	return HyperVSitesClientDeleteSiteResponse{}, nil
}

// deleteSiteCreateRequest creates the DeleteSite request.
func (client *HyperVSitesClient) deleteSiteCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, options *HyperVSitesClientDeleteSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// GetSite - Method to get a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - options - HyperVSitesClientGetSiteOptions contains the optional parameters for the HyperVSitesClient.GetSite method.
func (client *HyperVSitesClient) GetSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, options *HyperVSitesClientGetSiteOptions) (HyperVSitesClientGetSiteResponse, error) {
	req, err := client.getSiteCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, options)
	if err != nil {
		return HyperVSitesClientGetSiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HyperVSitesClientGetSiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HyperVSitesClientGetSiteResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSiteHandleResponse(resp)
}

// getSiteCreateRequest creates the GetSite request.
func (client *HyperVSitesClient) getSiteCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, options *HyperVSitesClientGetSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}"
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSiteHandleResponse handles the GetSite response.
func (client *HyperVSitesClient) getSiteHandleResponse(resp *http.Response) (HyperVSitesClientGetSiteResponse, error) {
	result := HyperVSitesClientGetSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HyperVSite); err != nil {
		return HyperVSitesClientGetSiteResponse{}, err
	}
	return result, nil
}

// NewGetSiteHealthSummaryPager - Method to get site health summary.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - options - HyperVSitesClientGetSiteHealthSummaryOptions contains the optional parameters for the HyperVSitesClient.NewGetSiteHealthSummaryPager
//     method.
func (client *HyperVSitesClient) NewGetSiteHealthSummaryPager(subscriptionID string, resourceGroupName string, siteName string, options *HyperVSitesClientGetSiteHealthSummaryOptions) *runtime.Pager[HyperVSitesClientGetSiteHealthSummaryResponse] {
	return runtime.NewPager(runtime.PagingHandler[HyperVSitesClientGetSiteHealthSummaryResponse]{
		More: func(page HyperVSitesClientGetSiteHealthSummaryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HyperVSitesClientGetSiteHealthSummaryResponse) (HyperVSitesClientGetSiteHealthSummaryResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getSiteHealthSummaryCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HyperVSitesClientGetSiteHealthSummaryResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HyperVSitesClientGetSiteHealthSummaryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HyperVSitesClientGetSiteHealthSummaryResponse{}, runtime.NewResponseError(resp)
			}
			return client.getSiteHealthSummaryHandleResponse(resp)
		},
	})
}

// getSiteHealthSummaryCreateRequest creates the GetSiteHealthSummary request.
func (client *HyperVSitesClient) getSiteHealthSummaryCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, options *HyperVSitesClientGetSiteHealthSummaryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/healthSummary"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSiteHealthSummaryHandleResponse handles the GetSiteHealthSummary response.
func (client *HyperVSitesClient) getSiteHealthSummaryHandleResponse(resp *http.Response) (HyperVSitesClientGetSiteHealthSummaryResponse, error) {
	result := HyperVSitesClientGetSiteHealthSummaryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SiteHealthSummaryCollection); err != nil {
		return HyperVSitesClientGetSiteHealthSummaryResponse{}, err
	}
	return result, nil
}

// GetSiteUsage - Method to get site usage.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - options - HyperVSitesClientGetSiteUsageOptions contains the optional parameters for the HyperVSitesClient.GetSiteUsage
//     method.
func (client *HyperVSitesClient) GetSiteUsage(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, options *HyperVSitesClientGetSiteUsageOptions) (HyperVSitesClientGetSiteUsageResponse, error) {
	req, err := client.getSiteUsageCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, options)
	if err != nil {
		return HyperVSitesClientGetSiteUsageResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HyperVSitesClientGetSiteUsageResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HyperVSitesClientGetSiteUsageResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSiteUsageHandleResponse(resp)
}

// getSiteUsageCreateRequest creates the GetSiteUsage request.
func (client *HyperVSitesClient) getSiteUsageCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, options *HyperVSitesClientGetSiteUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/summary"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSiteUsageHandleResponse handles the GetSiteUsage response.
func (client *HyperVSitesClient) getSiteUsageHandleResponse(resp *http.Response) (HyperVSitesClientGetSiteUsageResponse, error) {
	result := HyperVSitesClientGetSiteUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HyperVSiteUsage); err != nil {
		return HyperVSitesClientGetSiteUsageResponse{}, err
	}
	return result, nil
}

// PatchSite - Method to patch an existing site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - body - Body with site details.
//   - options - HyperVSitesClientPatchSiteOptions contains the optional parameters for the HyperVSitesClient.PatchSite method.
func (client *HyperVSitesClient) PatchSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body HyperVSite, options *HyperVSitesClientPatchSiteOptions) (HyperVSitesClientPatchSiteResponse, error) {
	req, err := client.patchSiteCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, body, options)
	if err != nil {
		return HyperVSitesClientPatchSiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HyperVSitesClientPatchSiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return HyperVSitesClientPatchSiteResponse{}, runtime.NewResponseError(resp)
	}
	return client.patchSiteHandleResponse(resp)
}

// patchSiteCreateRequest creates the PatchSite request.
func (client *HyperVSitesClient) patchSiteCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body HyperVSite, options *HyperVSitesClientPatchSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// patchSiteHandleResponse handles the PatchSite response.
func (client *HyperVSitesClient) patchSiteHandleResponse(resp *http.Response) (HyperVSitesClientPatchSiteResponse, error) {
	result := HyperVSitesClientPatchSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HyperVSite); err != nil {
		return HyperVSitesClientPatchSiteResponse{}, err
	}
	return result, nil
}

// PutSite - Method to create or update a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - body - Body with site details.
//   - options - HyperVSitesClientPutSiteOptions contains the optional parameters for the HyperVSitesClient.PutSite method.
func (client *HyperVSitesClient) PutSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body HyperVSite, options *HyperVSitesClientPutSiteOptions) (HyperVSitesClientPutSiteResponse, error) {
	req, err := client.putSiteCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, body, options)
	if err != nil {
		return HyperVSitesClientPutSiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HyperVSitesClientPutSiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return HyperVSitesClientPutSiteResponse{}, runtime.NewResponseError(resp)
	}
	return client.putSiteHandleResponse(resp)
}

// putSiteCreateRequest creates the PutSite request.
func (client *HyperVSitesClient) putSiteCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body HyperVSite, options *HyperVSitesClientPutSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// putSiteHandleResponse handles the PutSite response.
func (client *HyperVSitesClient) putSiteHandleResponse(resp *http.Response) (HyperVSitesClientPutSiteResponse, error) {
	result := HyperVSitesClientPutSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HyperVSite); err != nil {
		return HyperVSitesClientPutSiteResponse{}, err
	}
	return result, nil
}

// RefreshSite - Method to refresh a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - subscriptionID - The ID of the target subscription.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - options - HyperVSitesClientRefreshSiteOptions contains the optional parameters for the HyperVSitesClient.RefreshSite method.
func (client *HyperVSitesClient) RefreshSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, options *HyperVSitesClientRefreshSiteOptions) (HyperVSitesClientRefreshSiteResponse, error) {
	req, err := client.refreshSiteCreateRequest(ctx, subscriptionID, resourceGroupName, siteName, options)
	if err != nil {
		return HyperVSitesClientRefreshSiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HyperVSitesClientRefreshSiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return HyperVSitesClientRefreshSiteResponse{}, runtime.NewResponseError(resp)
	}
	return client.refreshSiteHandleResponse(resp)
}

// refreshSiteCreateRequest creates the RefreshSite request.
func (client *HyperVSitesClient) refreshSiteCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, options *HyperVSitesClientRefreshSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/refresh"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// refreshSiteHandleResponse handles the RefreshSite response.
func (client *HyperVSitesClient) refreshSiteHandleResponse(resp *http.Response) (HyperVSitesClientRefreshSiteResponse, error) {
	result := HyperVSitesClientRefreshSiteResponse{}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	return result, nil
}
