//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

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

// SessionHostConfigurationsClient contains the methods for the SessionHostConfigurations group.
// Don't use this type directly, use NewSessionHostConfigurationsClient() instead.
type SessionHostConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSessionHostConfigurationsClient creates a new instance of SessionHostConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSessionHostConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SessionHostConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SessionHostConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a SessionHostConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-08-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - sessionHostConfiguration - Object containing SessionHostConfiguration definitions.
//   - options - SessionHostConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the SessionHostConfigurationsClient.BeginCreateOrUpdate
//     method.
func (client *SessionHostConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostConfiguration SessionHostConfiguration, options *SessionHostConfigurationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SessionHostConfigurationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, hostPoolName, sessionHostConfiguration, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SessionHostConfigurationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SessionHostConfigurationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update a SessionHostConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-08-preview
func (client *SessionHostConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostConfiguration SessionHostConfiguration, options *SessionHostConfigurationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SessionHostConfigurationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostPoolName, sessionHostConfiguration, options)
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
func (client *SessionHostConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostConfiguration SessionHostConfiguration, options *SessionHostConfigurationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHostConfigurations/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-08-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sessionHostConfiguration); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get a SessionHostConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-08-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - SessionHostConfigurationsClientGetOptions contains the optional parameters for the SessionHostConfigurationsClient.Get
//     method.
func (client *SessionHostConfigurationsClient) Get(ctx context.Context, resourceGroupName string, hostPoolName string, options *SessionHostConfigurationsClientGetOptions) (SessionHostConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "SessionHostConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return SessionHostConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionHostConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SessionHostConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SessionHostConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *SessionHostConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHostConfigurations/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-08-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SessionHostConfigurationsClient) getHandleResponse(resp *http.Response) (SessionHostConfigurationsClientGetResponse, error) {
	result := SessionHostConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionHostConfiguration); err != nil {
		return SessionHostConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHostPoolPager - List sessionHostConfigurations.
//
// Generated from API version 2024-08-08-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - SessionHostConfigurationsClientListByHostPoolOptions contains the optional parameters for the SessionHostConfigurationsClient.NewListByHostPoolPager
//     method.
func (client *SessionHostConfigurationsClient) NewListByHostPoolPager(resourceGroupName string, hostPoolName string, options *SessionHostConfigurationsClientListByHostPoolOptions) *runtime.Pager[SessionHostConfigurationsClientListByHostPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[SessionHostConfigurationsClientListByHostPoolResponse]{
		More: func(page SessionHostConfigurationsClientListByHostPoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SessionHostConfigurationsClientListByHostPoolResponse) (SessionHostConfigurationsClientListByHostPoolResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SessionHostConfigurationsClient.NewListByHostPoolPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByHostPoolCreateRequest(ctx, resourceGroupName, hostPoolName, options)
			}, nil)
			if err != nil {
				return SessionHostConfigurationsClientListByHostPoolResponse{}, err
			}
			return client.listByHostPoolHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByHostPoolCreateRequest creates the ListByHostPool request.
func (client *SessionHostConfigurationsClient) listByHostPoolCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *SessionHostConfigurationsClientListByHostPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHostConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-08-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHostPoolHandleResponse handles the ListByHostPool response.
func (client *SessionHostConfigurationsClient) listByHostPoolHandleResponse(resp *http.Response) (SessionHostConfigurationsClientListByHostPoolResponse, error) {
	result := SessionHostConfigurationsClientListByHostPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionHostConfigurationList); err != nil {
		return SessionHostConfigurationsClientListByHostPoolResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a SessionHostConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-08-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - SessionHostConfigurationsClientBeginUpdateOptions contains the optional parameters for the SessionHostConfigurationsClient.BeginUpdate
//     method.
func (client *SessionHostConfigurationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, hostPoolName string, options *SessionHostConfigurationsClientBeginUpdateOptions) (*runtime.Poller[SessionHostConfigurationsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, hostPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SessionHostConfigurationsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SessionHostConfigurationsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a SessionHostConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-08-preview
func (client *SessionHostConfigurationsClient) update(ctx context.Context, resourceGroupName string, hostPoolName string, options *SessionHostConfigurationsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SessionHostConfigurationsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostPoolName, options)
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
func (client *SessionHostConfigurationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *SessionHostConfigurationsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHostConfigurations/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-08-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.SessionHostConfiguration != nil {
		if err := runtime.MarshalAsJSON(req, *options.SessionHostConfiguration); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}
