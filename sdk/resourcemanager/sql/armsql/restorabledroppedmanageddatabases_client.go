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

// RestorableDroppedManagedDatabasesClient contains the methods for the RestorableDroppedManagedDatabases group.
// Don't use this type directly, use NewRestorableDroppedManagedDatabasesClient() instead.
type RestorableDroppedManagedDatabasesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRestorableDroppedManagedDatabasesClient creates a new instance of RestorableDroppedManagedDatabasesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestorableDroppedManagedDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableDroppedManagedDatabasesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RestorableDroppedManagedDatabasesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a restorable dropped managed database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - RestorableDroppedManagedDatabasesClientGetOptions contains the optional parameters for the RestorableDroppedManagedDatabasesClient.Get
//     method.
func (client *RestorableDroppedManagedDatabasesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, options *RestorableDroppedManagedDatabasesClientGetOptions) (RestorableDroppedManagedDatabasesClientGetResponse, error) {
	var err error
	const operationName = "RestorableDroppedManagedDatabasesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID, options)
	if err != nil {
		return RestorableDroppedManagedDatabasesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RestorableDroppedManagedDatabasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RestorableDroppedManagedDatabasesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RestorableDroppedManagedDatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, options *RestorableDroppedManagedDatabasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if restorableDroppedDatabaseID == "" {
		return nil, errors.New("parameter restorableDroppedDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorableDroppedDatabaseId}", url.PathEscape(restorableDroppedDatabaseID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RestorableDroppedManagedDatabasesClient) getHandleResponse(resp *http.Response) (RestorableDroppedManagedDatabasesClientGetResponse, error) {
	result := RestorableDroppedManagedDatabasesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDroppedManagedDatabase); err != nil {
		return RestorableDroppedManagedDatabasesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInstancePager - Gets a list of restorable dropped managed databases.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - RestorableDroppedManagedDatabasesClientListByInstanceOptions contains the optional parameters for the RestorableDroppedManagedDatabasesClient.NewListByInstancePager
//     method.
func (client *RestorableDroppedManagedDatabasesClient) NewListByInstancePager(resourceGroupName string, managedInstanceName string, options *RestorableDroppedManagedDatabasesClientListByInstanceOptions) *runtime.Pager[RestorableDroppedManagedDatabasesClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableDroppedManagedDatabasesClientListByInstanceResponse]{
		More: func(page RestorableDroppedManagedDatabasesClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RestorableDroppedManagedDatabasesClientListByInstanceResponse) (RestorableDroppedManagedDatabasesClientListByInstanceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RestorableDroppedManagedDatabasesClient.NewListByInstancePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			}, nil)
			if err != nil {
				return RestorableDroppedManagedDatabasesClientListByInstanceResponse{}, err
			}
			return client.listByInstanceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *RestorableDroppedManagedDatabasesClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *RestorableDroppedManagedDatabasesClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases"
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
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *RestorableDroppedManagedDatabasesClient) listByInstanceHandleResponse(resp *http.Response) (RestorableDroppedManagedDatabasesClientListByInstanceResponse, error) {
	result := RestorableDroppedManagedDatabasesClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDroppedManagedDatabaseListResult); err != nil {
		return RestorableDroppedManagedDatabasesClientListByInstanceResponse{}, err
	}
	return result, nil
}
