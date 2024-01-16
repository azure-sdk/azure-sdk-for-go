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

// SessionHostManagementsOperationStatusClient contains the methods for the SessionHostManagementsOperationStatus group.
// Don't use this type directly, use NewSessionHostManagementsOperationStatusClient() instead.
type SessionHostManagementsOperationStatusClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSessionHostManagementsOperationStatusClient creates a new instance of SessionHostManagementsOperationStatusClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSessionHostManagementsOperationStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SessionHostManagementsOperationStatusClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SessionHostManagementsOperationStatusClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get Operation status for SessionHostManagement
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-16-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - operationID - The Guid of the operation.
//   - options - SessionHostManagementsOperationStatusClientGetOptions contains the optional parameters for the SessionHostManagementsOperationStatusClient.Get
//     method.
func (client *SessionHostManagementsOperationStatusClient) Get(ctx context.Context, resourceGroupName string, hostPoolName string, operationID string, options *SessionHostManagementsOperationStatusClientGetOptions) (SessionHostManagementsOperationStatusClientGetResponse, error) {
	var err error
	const operationName = "SessionHostManagementsOperationStatusClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostPoolName, operationID, options)
	if err != nil {
		return SessionHostManagementsOperationStatusClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionHostManagementsOperationStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SessionHostManagementsOperationStatusClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SessionHostManagementsOperationStatusClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, operationID string, options *SessionHostManagementsOperationStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHostManagements/default/operationStatuses/{operationId}"
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
	reqQP.Set("api-version", "2024-01-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SessionHostManagementsOperationStatusClient) getHandleResponse(resp *http.Response) (SessionHostManagementsOperationStatusClientGetResponse, error) {
	result := SessionHostManagementsOperationStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionHostManagementOperationStatus); err != nil {
		return SessionHostManagementsOperationStatusClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get Operation status for SessionHostManagement
//
// Generated from API version 2024-01-16-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - SessionHostManagementsOperationStatusClientListOptions contains the optional parameters for the SessionHostManagementsOperationStatusClient.NewListPager
//     method.
func (client *SessionHostManagementsOperationStatusClient) NewListPager(resourceGroupName string, hostPoolName string, options *SessionHostManagementsOperationStatusClientListOptions) *runtime.Pager[SessionHostManagementsOperationStatusClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SessionHostManagementsOperationStatusClientListResponse]{
		More: func(page SessionHostManagementsOperationStatusClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SessionHostManagementsOperationStatusClientListResponse) (SessionHostManagementsOperationStatusClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SessionHostManagementsOperationStatusClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, hostPoolName, options)
			}, nil)
			if err != nil {
				return SessionHostManagementsOperationStatusClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SessionHostManagementsOperationStatusClient) listCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *SessionHostManagementsOperationStatusClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHostManagements/default/operationStatuses"
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
	reqQP.Set("api-version", "2024-01-16-preview")
	if options != nil && options.IsLatest != nil {
		reqQP.Set("isLatest", strconv.FormatBool(*options.IsLatest))
	}
	if options != nil && options.Type != nil {
		reqQP.Set("type", *options.Type)
	}
	if options != nil && options.IsNonTerminal != nil {
		reqQP.Set("isNonTerminal", strconv.FormatBool(*options.IsNonTerminal))
	}
	if options != nil && options.CorrelationID != nil {
		reqQP.Set("correlationId", *options.CorrelationID)
	}
	if options != nil && options.Action != nil {
		reqQP.Set("action", *options.Action)
	}
	if options != nil && options.IsInitiatingOperation != nil {
		reqQP.Set("isInitiatingOperation", strconv.FormatBool(*options.IsInitiatingOperation))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SessionHostManagementsOperationStatusClient) listHandleResponse(resp *http.Response) (SessionHostManagementsOperationStatusClientListResponse, error) {
	result := SessionHostManagementsOperationStatusClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionHostManagementOperationStatusList); err != nil {
		return SessionHostManagementsOperationStatusClientListResponse{}, err
	}
	return result, nil
}
