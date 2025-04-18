// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqldiscovery

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

// MySQLSitesClient contains the methods for the MySQLSites group.
// Don't use this type directly, use NewMySQLSitesClient() instead.
type MySQLSitesClient struct {
	internal       *arm.Client
	subscriptionID string
	siteName       string
}

// NewMySQLSitesClient creates a new instance of MySQLSitesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - siteName - The name of Site
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMySQLSitesClient(subscriptionID string, siteName string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MySQLSitesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MySQLSitesClient{
		subscriptionID: subscriptionID,
		siteName:       siteName,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Puts the MySQLSites resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - body - The machines to add to the assessment
//   - options - MySQLSitesClientBeginCreateOrUpdateOptions contains the optional parameters for the MySQLSitesClient.BeginCreateOrUpdate
//     method.
func (client *MySQLSitesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, body MySQLSite, options *MySQLSitesClientBeginCreateOrUpdateOptions) (*runtime.Poller[MySQLSitesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MySQLSitesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MySQLSitesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Puts the MySQLSites resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-30-preview
func (client *MySQLSitesClient) createOrUpdate(ctx context.Context, resourceGroupName string, body MySQLSite, options *MySQLSitesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MySQLSitesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, body, options)
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
func (client *MySQLSitesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, body MySQLSite, _ *MySQLSitesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MySQLDiscovery/MySQLSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Deletes the MySQLSites resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - MySQLSitesClientDeleteOptions contains the optional parameters for the MySQLSitesClient.Delete method.
func (client *MySQLSitesClient) Delete(ctx context.Context, resourceGroupName string, options *MySQLSitesClientDeleteOptions) (MySQLSitesClientDeleteResponse, error) {
	var err error
	const operationName = "MySQLSitesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return MySQLSitesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MySQLSitesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MySQLSitesClientDeleteResponse{}, err
	}
	return MySQLSitesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MySQLSitesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, _ *MySQLSitesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MySQLDiscovery/MySQLSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the MySQLSites resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - MySQLSitesClientGetOptions contains the optional parameters for the MySQLSitesClient.Get method.
func (client *MySQLSitesClient) Get(ctx context.Context, resourceGroupName string, options *MySQLSitesClientGetOptions) (MySQLSitesClientGetResponse, error) {
	var err error
	const operationName = "MySQLSitesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return MySQLSitesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MySQLSitesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MySQLSitesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MySQLSitesClient) getCreateRequest(ctx context.Context, resourceGroupName string, _ *MySQLSitesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MySQLDiscovery/MySQLSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MySQLSitesClient) getHandleResponse(resp *http.Response) (MySQLSitesClientGetResponse, error) {
	result := MySQLSitesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MySQLSite); err != nil {
		return MySQLSitesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lits the MySQLSites resource in a resource group.
//
// Generated from API version 2024-09-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - MySQLSitesClientListByResourceGroupOptions contains the optional parameters for the MySQLSitesClient.NewListByResourceGroupPager
//     method.
func (client *MySQLSitesClient) NewListByResourceGroupPager(resourceGroupName string, options *MySQLSitesClientListByResourceGroupOptions) *runtime.Pager[MySQLSitesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[MySQLSitesClientListByResourceGroupResponse]{
		More: func(page MySQLSitesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MySQLSitesClientListByResourceGroupResponse) (MySQLSitesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MySQLSitesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return MySQLSitesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *MySQLSitesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *MySQLSitesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MySQLDiscovery/MySQLSites"
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
	reqQP.Set("api-version", "2024-09-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *MySQLSitesClient) listByResourceGroupHandleResponse(resp *http.Response) (MySQLSitesClientListByResourceGroupResponse, error) {
	result := MySQLSitesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MySQLSiteList); err != nil {
		return MySQLSitesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists the MySQLSites resource in a subscription.
//
// Generated from API version 2024-09-30-preview
//   - options - MySQLSitesClientListBySubscriptionOptions contains the optional parameters for the MySQLSitesClient.NewListBySubscriptionPager
//     method.
func (client *MySQLSitesClient) NewListBySubscriptionPager(options *MySQLSitesClientListBySubscriptionOptions) *runtime.Pager[MySQLSitesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[MySQLSitesClientListBySubscriptionResponse]{
		More: func(page MySQLSitesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MySQLSitesClientListBySubscriptionResponse) (MySQLSitesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MySQLSitesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return MySQLSitesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *MySQLSitesClient) listBySubscriptionCreateRequest(ctx context.Context, _ *MySQLSitesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MySQLDiscovery/MySQLSites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *MySQLSitesClient) listBySubscriptionHandleResponse(resp *http.Response) (MySQLSitesClientListBySubscriptionResponse, error) {
	result := MySQLSitesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MySQLSiteList); err != nil {
		return MySQLSitesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginTriggerRefresh - Trigger Refresh Refresh action
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - MySQLSitesClientBeginTriggerRefreshOptions contains the optional parameters for the MySQLSitesClient.BeginTriggerRefresh
//     method.
func (client *MySQLSitesClient) BeginTriggerRefresh(ctx context.Context, resourceGroupName string, options *MySQLSitesClientBeginTriggerRefreshOptions) (*runtime.Poller[MySQLSitesClientTriggerRefreshResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.triggerRefresh(ctx, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MySQLSitesClientTriggerRefreshResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MySQLSitesClientTriggerRefreshResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// TriggerRefresh - Trigger Refresh Refresh action
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-30-preview
func (client *MySQLSitesClient) triggerRefresh(ctx context.Context, resourceGroupName string, options *MySQLSitesClientBeginTriggerRefreshOptions) (*http.Response, error) {
	var err error
	const operationName = "MySQLSitesClient.BeginTriggerRefresh"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.triggerRefreshCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// triggerRefreshCreateRequest creates the TriggerRefresh request.
func (client *MySQLSitesClient) triggerRefreshCreateRequest(ctx context.Context, resourceGroupName string, _ *MySQLSitesClientBeginTriggerRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MySQLDiscovery/MySQLSites/{siteName}/refresh"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Patch the MySQLSites resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - body - Clouds patch payload.
//   - options - MySQLSitesClientBeginUpdateOptions contains the optional parameters for the MySQLSitesClient.BeginUpdate method.
func (client *MySQLSitesClient) BeginUpdate(ctx context.Context, resourceGroupName string, body MySQLSitesResourcePatch, options *MySQLSitesClientBeginUpdateOptions) (*runtime.Poller[MySQLSitesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MySQLSitesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MySQLSitesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch the MySQLSites resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-30-preview
func (client *MySQLSitesClient) update(ctx context.Context, resourceGroupName string, body MySQLSitesResourcePatch, options *MySQLSitesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MySQLSitesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, body, options)
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

// updateCreateRequest creates the Update request.
func (client *MySQLSitesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, body MySQLSitesResourcePatch, _ *MySQLSitesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MySQLDiscovery/MySQLSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
