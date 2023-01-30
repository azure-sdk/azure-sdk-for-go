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

// MasterSitesClient contains the methods for the MasterSites group.
// Don't use this type directly, use NewMasterSitesClient() instead.
type MasterSitesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMasterSitesClient creates a new instance of MasterSitesClient with the specified values.
//   - subscriptionID - Azure Subscription Id in which project was created.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMasterSitesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MasterSitesClient, error) {
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
	client := &MasterSitesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// DeleteSite - Method to delete a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-07-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - options - MasterSitesClientDeleteSiteOptions contains the optional parameters for the MasterSitesClient.DeleteSite method.
func (client *MasterSitesClient) DeleteSite(ctx context.Context, resourceGroupName string, siteName string, options *MasterSitesClientDeleteSiteOptions) (MasterSitesClientDeleteSiteResponse, error) {
	req, err := client.deleteSiteCreateRequest(ctx, resourceGroupName, siteName, options)
	if err != nil {
		return MasterSitesClientDeleteSiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MasterSitesClientDeleteSiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return MasterSitesClientDeleteSiteResponse{}, runtime.NewResponseError(resp)
	}
	return MasterSitesClientDeleteSiteResponse{}, nil
}

// deleteSiteCreateRequest creates the DeleteSite request.
func (client *MasterSitesClient) deleteSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *MasterSitesClientDeleteSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/MasterSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
	reqQP.Set("api-version", "2020-07-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetSite - Method to get a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-07-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - The site name.
//   - options - MasterSitesClientGetSiteOptions contains the optional parameters for the MasterSitesClient.GetSite method.
func (client *MasterSitesClient) GetSite(ctx context.Context, resourceGroupName string, siteName string, options *MasterSitesClientGetSiteOptions) (MasterSitesClientGetSiteResponse, error) {
	req, err := client.getSiteCreateRequest(ctx, resourceGroupName, siteName, options)
	if err != nil {
		return MasterSitesClientGetSiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MasterSitesClientGetSiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MasterSitesClientGetSiteResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSiteHandleResponse(resp)
}

// getSiteCreateRequest creates the GetSite request.
func (client *MasterSitesClient) getSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *MasterSitesClientGetSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/MasterSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
	reqQP.Set("api-version", "2020-07-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSiteHandleResponse handles the GetSite response.
func (client *MasterSitesClient) getSiteHandleResponse(resp *http.Response) (MasterSitesClientGetSiteResponse, error) {
	result := MasterSitesClientGetSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MasterSite); err != nil {
		return MasterSitesClientGetSiteResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all the sites in the resource group.
//
// Generated from API version 2020-07-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - MasterSitesClientListOptions contains the optional parameters for the MasterSitesClient.NewListPager method.
func (client *MasterSitesClient) NewListPager(resourceGroupName string, options *MasterSitesClientListOptions) *runtime.Pager[MasterSitesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MasterSitesClientListResponse]{
		More: func(page MasterSitesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MasterSitesClientListResponse) (MasterSitesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MasterSitesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MasterSitesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MasterSitesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MasterSitesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *MasterSitesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OffAzure/MasterSites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MasterSitesClient) listHandleResponse(resp *http.Response) (MasterSitesClientListResponse, error) {
	result := MasterSitesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MasterSiteList); err != nil {
		return MasterSitesClientListResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get all the sites in the subscription.
//
// Generated from API version 2020-07-07
//   - options - MasterSitesClientListBySubscriptionOptions contains the optional parameters for the MasterSitesClient.NewListBySubscriptionPager
//     method.
func (client *MasterSitesClient) NewListBySubscriptionPager(options *MasterSitesClientListBySubscriptionOptions) *runtime.Pager[MasterSitesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[MasterSitesClientListBySubscriptionResponse]{
		More: func(page MasterSitesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MasterSitesClientListBySubscriptionResponse) (MasterSitesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MasterSitesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MasterSitesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MasterSitesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *MasterSitesClient) listBySubscriptionCreateRequest(ctx context.Context, options *MasterSitesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OffAzure/MasterSites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *MasterSitesClient) listBySubscriptionHandleResponse(resp *http.Response) (MasterSitesClientListBySubscriptionResponse, error) {
	result := MasterSitesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MasterSiteList); err != nil {
		return MasterSitesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// PatchSite - Method to patch an existing site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-07-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - body - Body with site details.
//   - options - MasterSitesClientPatchSiteOptions contains the optional parameters for the MasterSitesClient.PatchSite method.
func (client *MasterSitesClient) PatchSite(ctx context.Context, resourceGroupName string, siteName string, body MasterSite, options *MasterSitesClientPatchSiteOptions) (MasterSitesClientPatchSiteResponse, error) {
	req, err := client.patchSiteCreateRequest(ctx, resourceGroupName, siteName, body, options)
	if err != nil {
		return MasterSitesClientPatchSiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MasterSitesClientPatchSiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return MasterSitesClientPatchSiteResponse{}, runtime.NewResponseError(resp)
	}
	return client.patchSiteHandleResponse(resp)
}

// patchSiteCreateRequest creates the PatchSite request.
func (client *MasterSitesClient) patchSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, body MasterSite, options *MasterSitesClientPatchSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/MasterSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
	reqQP.Set("api-version", "2020-07-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// patchSiteHandleResponse handles the PatchSite response.
func (client *MasterSitesClient) patchSiteHandleResponse(resp *http.Response) (MasterSitesClientPatchSiteResponse, error) {
	result := MasterSitesClientPatchSiteResponse{}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.MasterSite); err != nil {
		return MasterSitesClientPatchSiteResponse{}, err
	}
	return result, nil
}

// PutSite - Method to create or update a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-07-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - body - Body with site details.
//   - options - MasterSitesClientPutSiteOptions contains the optional parameters for the MasterSitesClient.PutSite method.
func (client *MasterSitesClient) PutSite(ctx context.Context, resourceGroupName string, siteName string, body MasterSite, options *MasterSitesClientPutSiteOptions) (MasterSitesClientPutSiteResponse, error) {
	req, err := client.putSiteCreateRequest(ctx, resourceGroupName, siteName, body, options)
	if err != nil {
		return MasterSitesClientPutSiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MasterSitesClientPutSiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return MasterSitesClientPutSiteResponse{}, runtime.NewResponseError(resp)
	}
	return client.putSiteHandleResponse(resp)
}

// putSiteCreateRequest creates the PutSite request.
func (client *MasterSitesClient) putSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, body MasterSite, options *MasterSitesClientPutSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/MasterSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
	reqQP.Set("api-version", "2020-07-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// putSiteHandleResponse handles the PutSite response.
func (client *MasterSitesClient) putSiteHandleResponse(resp *http.Response) (MasterSitesClientPutSiteResponse, error) {
	result := MasterSitesClientPutSiteResponse{}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	return result, nil
}
