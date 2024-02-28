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

// SitesClient contains the methods for the Sites group.
// Don't use this type directly, use NewSitesClient() instead.
type SitesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSitesClient creates a new instance of SitesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSitesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SitesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SitesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a mobile network site. Must be created in the same location as its parent mobile
// network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - siteName - The name of the mobile network site.
//   - parameters - Parameters supplied to the create or update mobile network site operation.
//   - options - SitesClientBeginCreateOrUpdateOptions contains the optional parameters for the SitesClient.BeginCreateOrUpdate
//     method.
func (client *SitesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters Site, options *SitesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SitesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, mobileNetworkName, siteName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SitesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SitesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a mobile network site. Must be created in the same location as its parent mobile network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
func (client *SitesClient) createOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters Site, options *SitesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SitesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, mobileNetworkName, siteName, parameters, options)
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
func (client *SitesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters Site, options *SitesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/sites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified mobile network site. This will also delete any network functions that are a part of
// this site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - siteName - The name of the mobile network site.
//   - options - SitesClientBeginDeleteOptions contains the optional parameters for the SitesClient.BeginDelete method.
func (client *SitesClient) BeginDelete(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, options *SitesClientBeginDeleteOptions) (*runtime.Poller[SitesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, mobileNetworkName, siteName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SitesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SitesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified mobile network site. This will also delete any network functions that are a part of this
// site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
func (client *SitesClient) deleteOperation(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, options *SitesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SitesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, mobileNetworkName, siteName, options)
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
func (client *SitesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, options *SitesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/sites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDeletePacketCore - Deletes a packet core under the specified mobile network site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - siteName - The name of the mobile network site.
//   - parameters - Parameters supplied to delete a packet core under a site.
//   - options - SitesClientBeginDeletePacketCoreOptions contains the optional parameters for the SitesClient.BeginDeletePacketCore
//     method.
func (client *SitesClient) BeginDeletePacketCore(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters SiteDeletePacketCore, options *SitesClientBeginDeletePacketCoreOptions) (*runtime.Poller[SitesClientDeletePacketCoreResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deletePacketCore(ctx, resourceGroupName, mobileNetworkName, siteName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SitesClientDeletePacketCoreResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SitesClientDeletePacketCoreResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// DeletePacketCore - Deletes a packet core under the specified mobile network site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
func (client *SitesClient) deletePacketCore(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters SiteDeletePacketCore, options *SitesClientBeginDeletePacketCoreOptions) (*http.Response, error) {
	var err error
	const operationName = "SitesClient.BeginDeletePacketCore"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deletePacketCoreCreateRequest(ctx, resourceGroupName, mobileNetworkName, siteName, parameters, options)
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

// deletePacketCoreCreateRequest creates the DeletePacketCore request.
func (client *SitesClient) deletePacketCoreCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters SiteDeletePacketCore, options *SitesClientBeginDeletePacketCoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/sites/{siteName}/deletePacketCore"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets information about the specified mobile network site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - siteName - The name of the mobile network site.
//   - options - SitesClientGetOptions contains the optional parameters for the SitesClient.Get method.
func (client *SitesClient) Get(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, options *SitesClientGetOptions) (SitesClientGetResponse, error) {
	var err error
	const operationName = "SitesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, mobileNetworkName, siteName, options)
	if err != nil {
		return SitesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SitesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SitesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SitesClient) getCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, options *SitesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/sites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SitesClient) getHandleResponse(resp *http.Response) (SitesClientGetResponse, error) {
	result := SitesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Site); err != nil {
		return SitesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByMobileNetworkPager - Lists all sites in the mobile network.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - options - SitesClientListByMobileNetworkOptions contains the optional parameters for the SitesClient.NewListByMobileNetworkPager
//     method.
func (client *SitesClient) NewListByMobileNetworkPager(resourceGroupName string, mobileNetworkName string, options *SitesClientListByMobileNetworkOptions) *runtime.Pager[SitesClientListByMobileNetworkResponse] {
	return runtime.NewPager(runtime.PagingHandler[SitesClientListByMobileNetworkResponse]{
		More: func(page SitesClientListByMobileNetworkResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SitesClientListByMobileNetworkResponse) (SitesClientListByMobileNetworkResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SitesClient.NewListByMobileNetworkPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByMobileNetworkCreateRequest(ctx, resourceGroupName, mobileNetworkName, options)
			}, nil)
			if err != nil {
				return SitesClientListByMobileNetworkResponse{}, err
			}
			return client.listByMobileNetworkHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByMobileNetworkCreateRequest creates the ListByMobileNetwork request.
func (client *SitesClient) listByMobileNetworkCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, options *SitesClientListByMobileNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/sites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMobileNetworkHandleResponse handles the ListByMobileNetwork response.
func (client *SitesClient) listByMobileNetworkHandleResponse(resp *http.Response) (SitesClientListByMobileNetworkResponse, error) {
	result := SitesClientListByMobileNetworkResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SiteListResult); err != nil {
		return SitesClientListByMobileNetworkResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates site tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - mobileNetworkName - The name of the mobile network.
//   - siteName - The name of the mobile network site.
//   - parameters - Parameters supplied to update network site tags.
//   - options - SitesClientUpdateTagsOptions contains the optional parameters for the SitesClient.UpdateTags method.
func (client *SitesClient) UpdateTags(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters TagsObject, options *SitesClientUpdateTagsOptions) (SitesClientUpdateTagsResponse, error) {
	var err error
	const operationName = "SitesClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, mobileNetworkName, siteName, parameters, options)
	if err != nil {
		return SitesClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SitesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SitesClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SitesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, siteName string, parameters TagsObject, options *SitesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/sites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *SitesClient) updateTagsHandleResponse(resp *http.Response) (SitesClientUpdateTagsResponse, error) {
	result := SitesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Site); err != nil {
		return SitesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
