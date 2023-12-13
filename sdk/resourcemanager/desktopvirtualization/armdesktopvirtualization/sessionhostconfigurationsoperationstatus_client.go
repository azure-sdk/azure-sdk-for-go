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
	"strconv"
	"strings"
)

// SessionHostConfigurationsOperationStatusClient contains the methods for the SessionHostConfigurationsOperationStatus group.
// Don't use this type directly, use NewSessionHostConfigurationsOperationStatusClient() instead.
type SessionHostConfigurationsOperationStatusClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSessionHostConfigurationsOperationStatusClient creates a new instance of SessionHostConfigurationsOperationStatusClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSessionHostConfigurationsOperationStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SessionHostConfigurationsOperationStatusClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SessionHostConfigurationsOperationStatusClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get Operation status for SessionHostManagement
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - operationID - The Guid of the operation.
//   - options - SessionHostConfigurationsOperationStatusClientGetOptions contains the optional parameters for the SessionHostConfigurationsOperationStatusClient.Get
//     method.
func (client *SessionHostConfigurationsOperationStatusClient) Get(ctx context.Context, resourceGroupName string, hostPoolName string, operationID string, options *SessionHostConfigurationsOperationStatusClientGetOptions) (SessionHostConfigurationsOperationStatusClientGetResponse, error) {
	var err error
	const operationName = "SessionHostConfigurationsOperationStatusClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostPoolName, operationID, options)
	if err != nil {
		return SessionHostConfigurationsOperationStatusClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionHostConfigurationsOperationStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SessionHostConfigurationsOperationStatusClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SessionHostConfigurationsOperationStatusClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, operationID string, options *SessionHostConfigurationsOperationStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHostConfigurations/default/operationStatuses/{operationId}"
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
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SessionHostConfigurationsOperationStatusClient) getHandleResponse(resp *http.Response) (SessionHostConfigurationsOperationStatusClientGetResponse, error) {
	result := SessionHostConfigurationsOperationStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionHostConfigurationOperationStatus); err != nil {
		return SessionHostConfigurationsOperationStatusClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get Operation status for SessionHostConfiguration
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - SessionHostConfigurationsOperationStatusClientListOptions contains the optional parameters for the SessionHostConfigurationsOperationStatusClient.NewListPager
//     method.
func (client *SessionHostConfigurationsOperationStatusClient) NewListPager(resourceGroupName string, hostPoolName string, options *SessionHostConfigurationsOperationStatusClientListOptions) *runtime.Pager[SessionHostConfigurationsOperationStatusClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SessionHostConfigurationsOperationStatusClientListResponse]{
		More: func(page SessionHostConfigurationsOperationStatusClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SessionHostConfigurationsOperationStatusClientListResponse) (SessionHostConfigurationsOperationStatusClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SessionHostConfigurationsOperationStatusClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, hostPoolName, options)
			}, nil)
			if err != nil {
				return SessionHostConfigurationsOperationStatusClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SessionHostConfigurationsOperationStatusClient) listCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *SessionHostConfigurationsOperationStatusClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHostConfigurations/default/operationStatuses"
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
	reqQP.Set("api-version", "2023-11-01-preview")
	if options != nil && options.IsLatest != nil {
		reqQP.Set("isLatest", strconv.FormatBool(*options.IsLatest))
	}
	if options != nil && options.IsNonTerminal != nil {
		reqQP.Set("isNonTerminal", strconv.FormatBool(*options.IsNonTerminal))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SessionHostConfigurationsOperationStatusClient) listHandleResponse(resp *http.Response) (SessionHostConfigurationsOperationStatusClientListResponse, error) {
	result := SessionHostConfigurationsOperationStatusClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionHostConfigurationOperationStatusList); err != nil {
		return SessionHostConfigurationsOperationStatusClientListResponse{}, err
	}
	return result, nil
}
