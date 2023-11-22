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

// DatabaseExtensionsClient contains the methods for the DatabaseExtensions group.
// Don't use this type directly, use NewDatabaseExtensionsClient() instead.
type DatabaseExtensionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseExtensionsClient creates a new instance of DatabaseExtensionsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseExtensionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseExtensionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Perform a database extension operation, like database import, database export, or polybase import
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - parameters - The database import request parameters.
//   - options - DatabaseExtensionsClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabaseExtensionsClient.BeginCreateOrUpdate
//     method.
func (client *DatabaseExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, extensionName string, parameters DatabaseExtensions, options *DatabaseExtensionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DatabaseExtensionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, databaseName, extensionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabaseExtensionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabaseExtensionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Perform a database extension operation, like database import, database export, or polybase import
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *DatabaseExtensionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, extensionName string, parameters DatabaseExtensions, options *DatabaseExtensionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabaseExtensionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, extensionName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatabaseExtensionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, extensionName string, parameters DatabaseExtensions, options *DatabaseExtensionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/extensions/{extensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets a database extension. This will return resource not found as it is not supported.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - DatabaseExtensionsClientGetOptions contains the optional parameters for the DatabaseExtensionsClient.Get method.
func (client *DatabaseExtensionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, extensionName string, options *DatabaseExtensionsClientGetOptions) (DatabaseExtensionsClientGetResponse, error) {
	var err error
	const operationName = "DatabaseExtensionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, extensionName, options)
	if err != nil {
		return DatabaseExtensionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseExtensionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseExtensionsClientGetResponse{}, err
	}
	return DatabaseExtensionsClientGetResponse{}, nil
}

// getCreateRequest creates the Get request.
func (client *DatabaseExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, extensionName string, options *DatabaseExtensionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/extensions/{extensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// NewListByDatabasePager - List database extension. This will return an empty list as it is not supported.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - DatabaseExtensionsClientListByDatabaseOptions contains the optional parameters for the DatabaseExtensionsClient.NewListByDatabasePager
//     method.
func (client *DatabaseExtensionsClient) NewListByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *DatabaseExtensionsClientListByDatabaseOptions) *runtime.Pager[DatabaseExtensionsClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabaseExtensionsClientListByDatabaseResponse]{
		More: func(page DatabaseExtensionsClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatabaseExtensionsClientListByDatabaseResponse) (DatabaseExtensionsClientListByDatabaseResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DatabaseExtensionsClient.NewListByDatabasePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			}, nil)
			if err != nil {
				return DatabaseExtensionsClientListByDatabaseResponse{}, err
			}
			return client.listByDatabaseHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *DatabaseExtensionsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DatabaseExtensionsClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/extensions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *DatabaseExtensionsClient) listByDatabaseHandleResponse(resp *http.Response) (DatabaseExtensionsClientListByDatabaseResponse, error) {
	result := DatabaseExtensionsClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportExportExtensionsOperationListResult); err != nil {
		return DatabaseExtensionsClientListByDatabaseResponse{}, err
	}
	return result, nil
}
