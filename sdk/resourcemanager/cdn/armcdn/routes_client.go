// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// RoutesClient contains the methods for the Routes group.
// Don't use this type directly, use NewRoutesClient() instead.
type RoutesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRoutesClient creates a new instance of RoutesClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRoutesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RoutesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RoutesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a new route with the specified route name under the specified subscription, resource group, profile,
// and AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - routeName - Name of the routing rule.
//   - route - Route properties
//   - options - RoutesClientBeginCreateOptions contains the optional parameters for the RoutesClient.BeginCreate method.
func (client *RoutesClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, route Route, options *RoutesClientBeginCreateOptions) (*runtime.Poller[RoutesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, profileName, endpointName, routeName, route, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RoutesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RoutesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates a new route with the specified route name under the specified subscription, resource group, profile, and
// AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *RoutesClient) create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, route Route, options *RoutesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "RoutesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, endpointName, routeName, route, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *RoutesClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, route Route, _ *RoutesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, route); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an existing route with the specified route name under the specified subscription, resource group,
// profile, and AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - routeName - Name of the routing rule.
//   - options - RoutesClientBeginDeleteOptions contains the optional parameters for the RoutesClient.BeginDelete method.
func (client *RoutesClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, options *RoutesClientBeginDeleteOptions) (*runtime.Poller[RoutesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, endpointName, routeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RoutesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RoutesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes an existing route with the specified route name under the specified subscription, resource group, profile,
// and AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *RoutesClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, options *RoutesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RoutesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, endpointName, routeName, options)
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
func (client *RoutesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, _ *RoutesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing route with the specified route name under the specified subscription, resource group, profile, and
// AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - routeName - Name of the routing rule.
//   - options - RoutesClientGetOptions contains the optional parameters for the RoutesClient.Get method.
func (client *RoutesClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, options *RoutesClientGetOptions) (RoutesClientGetResponse, error) {
	var err error
	const operationName = "RoutesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, endpointName, routeName, options)
	if err != nil {
		return RoutesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoutesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoutesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RoutesClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, _ *RoutesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoutesClient) getHandleResponse(resp *http.Response) (RoutesClientGetResponse, error) {
	result := RoutesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Route); err != nil {
		return RoutesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByEndpointPager - Lists all of the existing origins within a profile.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - options - RoutesClientListByEndpointOptions contains the optional parameters for the RoutesClient.NewListByEndpointPager
//     method.
func (client *RoutesClient) NewListByEndpointPager(resourceGroupName string, profileName string, endpointName string, options *RoutesClientListByEndpointOptions) *runtime.Pager[RoutesClientListByEndpointResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoutesClientListByEndpointResponse]{
		More: func(page RoutesClientListByEndpointResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoutesClientListByEndpointResponse) (RoutesClientListByEndpointResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RoutesClient.NewListByEndpointPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByEndpointCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
			}, nil)
			if err != nil {
				return RoutesClientListByEndpointResponse{}, err
			}
			return client.listByEndpointHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByEndpointCreateRequest creates the ListByEndpoint request.
func (client *RoutesClient) listByEndpointCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, _ *RoutesClientListByEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEndpointHandleResponse handles the ListByEndpoint response.
func (client *RoutesClient) listByEndpointHandleResponse(resp *http.Response) (RoutesClientListByEndpointResponse, error) {
	result := RoutesClientListByEndpointResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteListResult); err != nil {
		return RoutesClientListByEndpointResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an existing route with the specified route name under the specified subscription, resource group,
// profile, and AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - routeName - Name of the routing rule.
//   - routeUpdateProperties - Route update properties
//   - options - RoutesClientBeginUpdateOptions contains the optional parameters for the RoutesClient.BeginUpdate method.
func (client *RoutesClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, routeUpdateProperties RouteUpdateParameters, options *RoutesClientBeginUpdateOptions) (*runtime.Poller[RoutesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, profileName, endpointName, routeName, routeUpdateProperties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RoutesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RoutesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates an existing route with the specified route name under the specified subscription, resource group, profile,
// and AzureFrontDoor endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *RoutesClient) update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, routeUpdateProperties RouteUpdateParameters, options *RoutesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "RoutesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, endpointName, routeName, routeUpdateProperties, options)
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
func (client *RoutesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, routeName string, routeUpdateProperties RouteUpdateParameters, _ *RoutesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/afdEndpoints/{endpointName}/routes/{routeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if routeName == "" {
		return nil, errors.New("parameter routeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeName}", url.PathEscape(routeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, routeUpdateProperties); err != nil {
		return nil, err
	}
	return req, nil
}
