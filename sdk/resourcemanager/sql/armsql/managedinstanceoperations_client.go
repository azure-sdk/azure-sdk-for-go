//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ManagedInstanceOperationsClient contains the methods for the ManagedInstanceOperations group.
// Don't use this type directly, use NewManagedInstanceOperationsClient() instead.
type ManagedInstanceOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedInstanceOperationsClient creates a new instance of ManagedInstanceOperationsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedInstanceOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedInstanceOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedInstanceOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Cancel - Cancels the asynchronous operation on the managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstanceOperationsClientCancelOptions contains the optional parameters for the ManagedInstanceOperationsClient.Cancel
//     method.
func (client *ManagedInstanceOperationsClient) Cancel(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsClientCancelOptions) (ManagedInstanceOperationsClientCancelResponse, error) {
	var err error
	const operationName = "ManagedInstanceOperationsClient.Cancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, managedInstanceName, operationID, options)
	if err != nil {
		return ManagedInstanceOperationsClientCancelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedInstanceOperationsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedInstanceOperationsClientCancelResponse{}, err
	}
	return ManagedInstanceOperationsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *ManagedInstanceOperationsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/operations/{operationId}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a management operation on a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstanceOperationsClientGetOptions contains the optional parameters for the ManagedInstanceOperationsClient.Get
//     method.
func (client *ManagedInstanceOperationsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsClientGetOptions) (ManagedInstanceOperationsClientGetResponse, error) {
	var err error
	const operationName = "ManagedInstanceOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, operationID, options)
	if err != nil {
		return ManagedInstanceOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedInstanceOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedInstanceOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/operations/{operationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedInstanceOperationsClient) getHandleResponse(resp *http.Response) (ManagedInstanceOperationsClientGetResponse, error) {
	result := ManagedInstanceOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceOperation); err != nil {
		return ManagedInstanceOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagedInstancePager - Gets a list of operations performed on the managed instance.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstanceOperationsClientListByManagedInstanceOptions contains the optional parameters for the ManagedInstanceOperationsClient.NewListByManagedInstancePager
//     method.
func (client *ManagedInstanceOperationsClient) NewListByManagedInstancePager(resourceGroupName string, managedInstanceName string, options *ManagedInstanceOperationsClientListByManagedInstanceOptions) *runtime.Pager[ManagedInstanceOperationsClientListByManagedInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstanceOperationsClientListByManagedInstanceResponse]{
		More: func(page ManagedInstanceOperationsClientListByManagedInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstanceOperationsClientListByManagedInstanceResponse) (ManagedInstanceOperationsClientListByManagedInstanceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedInstanceOperationsClient.NewListByManagedInstancePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByManagedInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			}, nil)
			if err != nil {
				return ManagedInstanceOperationsClientListByManagedInstanceResponse{}, err
			}
			return client.listByManagedInstanceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagedInstanceCreateRequest creates the ListByManagedInstance request.
func (client *ManagedInstanceOperationsClient) listByManagedInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstanceOperationsClientListByManagedInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/operations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedInstanceHandleResponse handles the ListByManagedInstance response.
func (client *ManagedInstanceOperationsClient) listByManagedInstanceHandleResponse(resp *http.Response) (ManagedInstanceOperationsClientListByManagedInstanceResponse, error) {
	result := ManagedInstanceOperationsClientListByManagedInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceOperationListResult); err != nil {
		return ManagedInstanceOperationsClientListByManagedInstanceResponse{}, err
	}
	return result, nil
}
