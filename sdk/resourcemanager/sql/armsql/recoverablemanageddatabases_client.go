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

// RecoverableManagedDatabasesClient contains the methods for the RecoverableManagedDatabases group.
// Don't use this type directly, use NewRecoverableManagedDatabasesClient() instead.
type RecoverableManagedDatabasesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRecoverableManagedDatabasesClient creates a new instance of RecoverableManagedDatabasesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRecoverableManagedDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RecoverableManagedDatabasesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RecoverableManagedDatabasesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a recoverable managed database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - RecoverableManagedDatabasesClientGetOptions contains the optional parameters for the RecoverableManagedDatabasesClient.Get
//     method.
func (client *RecoverableManagedDatabasesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, recoverableDatabaseName string, options *RecoverableManagedDatabasesClientGetOptions) (RecoverableManagedDatabasesClientGetResponse, error) {
	var err error
	const operationName = "RecoverableManagedDatabasesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, recoverableDatabaseName, options)
	if err != nil {
		return RecoverableManagedDatabasesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecoverableManagedDatabasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecoverableManagedDatabasesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RecoverableManagedDatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, recoverableDatabaseName string, _ *RecoverableManagedDatabasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/recoverableDatabases/{recoverableDatabaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if recoverableDatabaseName == "" {
		return nil, errors.New("parameter recoverableDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recoverableDatabaseName}", url.PathEscape(recoverableDatabaseName))
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
func (client *RecoverableManagedDatabasesClient) getHandleResponse(resp *http.Response) (RecoverableManagedDatabasesClientGetResponse, error) {
	result := RecoverableManagedDatabasesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableManagedDatabase); err != nil {
		return RecoverableManagedDatabasesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInstancePager - Gets a list of recoverable managed databases.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - RecoverableManagedDatabasesClientListByInstanceOptions contains the optional parameters for the RecoverableManagedDatabasesClient.NewListByInstancePager
//     method.
func (client *RecoverableManagedDatabasesClient) NewListByInstancePager(resourceGroupName string, managedInstanceName string, options *RecoverableManagedDatabasesClientListByInstanceOptions) *runtime.Pager[RecoverableManagedDatabasesClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[RecoverableManagedDatabasesClientListByInstanceResponse]{
		More: func(page RecoverableManagedDatabasesClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecoverableManagedDatabasesClientListByInstanceResponse) (RecoverableManagedDatabasesClientListByInstanceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RecoverableManagedDatabasesClient.NewListByInstancePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			}, nil)
			if err != nil {
				return RecoverableManagedDatabasesClientListByInstanceResponse{}, err
			}
			return client.listByInstanceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *RecoverableManagedDatabasesClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, _ *RecoverableManagedDatabasesClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/recoverableDatabases"
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

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *RecoverableManagedDatabasesClient) listByInstanceHandleResponse(resp *http.Response) (RecoverableManagedDatabasesClientListByInstanceResponse, error) {
	result := RecoverableManagedDatabasesClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableManagedDatabaseListResult); err != nil {
		return RecoverableManagedDatabasesClientListByInstanceResponse{}, err
	}
	return result, nil
}
