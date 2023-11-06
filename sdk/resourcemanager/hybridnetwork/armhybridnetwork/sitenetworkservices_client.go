//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// SiteNetworkServicesClient contains the methods for the SiteNetworkServices group.
// Don't use this type directly, use NewSiteNetworkServicesClient() instead.
type SiteNetworkServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSiteNetworkServicesClient creates a new instance of SiteNetworkServicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSiteNetworkServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SiteNetworkServicesClient, error) {
	cl, err := arm.NewClient(moduleName+".SiteNetworkServicesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SiteNetworkServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a network site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteNetworkServiceName - The name of the site network service.
//   - parameters - Parameters supplied to the create or update site network service operation.
//   - options - SiteNetworkServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the SiteNetworkServicesClient.BeginCreateOrUpdate
//     method.
func (client *SiteNetworkServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, siteNetworkServiceName string, parameters SiteNetworkService, options *SiteNetworkServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SiteNetworkServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, siteNetworkServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SiteNetworkServicesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SiteNetworkServicesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a network site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *SiteNetworkServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, siteNetworkServiceName string, parameters SiteNetworkService, options *SiteNetworkServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, siteNetworkServiceName, parameters, options)
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
func (client *SiteNetworkServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, siteNetworkServiceName string, parameters SiteNetworkService, options *SiteNetworkServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/siteNetworkServices/{siteNetworkServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteNetworkServiceName == "" {
		return nil, errors.New("parameter siteNetworkServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteNetworkServiceName}", url.PathEscape(siteNetworkServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified site network service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteNetworkServiceName - The name of the site network service.
//   - options - SiteNetworkServicesClientBeginDeleteOptions contains the optional parameters for the SiteNetworkServicesClient.BeginDelete
//     method.
func (client *SiteNetworkServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, siteNetworkServiceName string, options *SiteNetworkServicesClientBeginDeleteOptions) (*runtime.Poller[SiteNetworkServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, siteNetworkServiceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SiteNetworkServicesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SiteNetworkServicesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified site network service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *SiteNetworkServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, siteNetworkServiceName string, options *SiteNetworkServicesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, siteNetworkServiceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SiteNetworkServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, siteNetworkServiceName string, options *SiteNetworkServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/siteNetworkServices/{siteNetworkServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteNetworkServiceName == "" {
		return nil, errors.New("parameter siteNetworkServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteNetworkServiceName}", url.PathEscape(siteNetworkServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified site network service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteNetworkServiceName - The name of the site network service.
//   - options - SiteNetworkServicesClientGetOptions contains the optional parameters for the SiteNetworkServicesClient.Get method.
func (client *SiteNetworkServicesClient) Get(ctx context.Context, resourceGroupName string, siteNetworkServiceName string, options *SiteNetworkServicesClientGetOptions) (SiteNetworkServicesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteNetworkServiceName, options)
	if err != nil {
		return SiteNetworkServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SiteNetworkServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SiteNetworkServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SiteNetworkServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteNetworkServiceName string, options *SiteNetworkServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/siteNetworkServices/{siteNetworkServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteNetworkServiceName == "" {
		return nil, errors.New("parameter siteNetworkServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteNetworkServiceName}", url.PathEscape(siteNetworkServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SiteNetworkServicesClient) getHandleResponse(resp *http.Response) (SiteNetworkServicesClientGetResponse, error) {
	result := SiteNetworkServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SiteNetworkService); err != nil {
		return SiteNetworkServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all site network services.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - SiteNetworkServicesClientListByResourceGroupOptions contains the optional parameters for the SiteNetworkServicesClient.NewListByResourceGroupPager
//     method.
func (client *SiteNetworkServicesClient) NewListByResourceGroupPager(resourceGroupName string, options *SiteNetworkServicesClientListByResourceGroupOptions) *runtime.Pager[SiteNetworkServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SiteNetworkServicesClientListByResourceGroupResponse]{
		More: func(page SiteNetworkServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SiteNetworkServicesClientListByResourceGroupResponse) (SiteNetworkServicesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SiteNetworkServicesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SiteNetworkServicesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SiteNetworkServicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SiteNetworkServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SiteNetworkServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/siteNetworkServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SiteNetworkServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (SiteNetworkServicesClientListByResourceGroupResponse, error) {
	result := SiteNetworkServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SiteNetworkServiceListResult); err != nil {
		return SiteNetworkServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all sites in the network service in a subscription.
//
// Generated from API version 2023-09-01
//   - options - SiteNetworkServicesClientListBySubscriptionOptions contains the optional parameters for the SiteNetworkServicesClient.NewListBySubscriptionPager
//     method.
func (client *SiteNetworkServicesClient) NewListBySubscriptionPager(options *SiteNetworkServicesClientListBySubscriptionOptions) *runtime.Pager[SiteNetworkServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SiteNetworkServicesClientListBySubscriptionResponse]{
		More: func(page SiteNetworkServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SiteNetworkServicesClientListBySubscriptionResponse) (SiteNetworkServicesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SiteNetworkServicesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SiteNetworkServicesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SiteNetworkServicesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SiteNetworkServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *SiteNetworkServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/siteNetworkServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SiteNetworkServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (SiteNetworkServicesClientListBySubscriptionResponse, error) {
	result := SiteNetworkServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SiteNetworkServiceListResult); err != nil {
		return SiteNetworkServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates a site update tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteNetworkServiceName - The name of the site network service.
//   - parameters - Parameters supplied to update network site tags.
//   - options - SiteNetworkServicesClientUpdateTagsOptions contains the optional parameters for the SiteNetworkServicesClient.UpdateTags
//     method.
func (client *SiteNetworkServicesClient) UpdateTags(ctx context.Context, resourceGroupName string, siteNetworkServiceName string, parameters TagsObject, options *SiteNetworkServicesClientUpdateTagsOptions) (SiteNetworkServicesClientUpdateTagsResponse, error) {
	var err error
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, siteNetworkServiceName, parameters, options)
	if err != nil {
		return SiteNetworkServicesClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SiteNetworkServicesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SiteNetworkServicesClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SiteNetworkServicesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, siteNetworkServiceName string, parameters TagsObject, options *SiteNetworkServicesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/siteNetworkServices/{siteNetworkServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteNetworkServiceName == "" {
		return nil, errors.New("parameter siteNetworkServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteNetworkServiceName}", url.PathEscape(siteNetworkServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *SiteNetworkServicesClient) updateTagsHandleResponse(resp *http.Response) (SiteNetworkServicesClientUpdateTagsResponse, error) {
	result := SiteNetworkServicesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SiteNetworkService); err != nil {
		return SiteNetworkServicesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
