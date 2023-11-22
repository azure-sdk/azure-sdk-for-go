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

// ServerAzureADAdministratorsClient contains the methods for the ServerAzureADAdministrators group.
// Don't use this type directly, use NewServerAzureADAdministratorsClient() instead.
type ServerAzureADAdministratorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerAzureADAdministratorsClient creates a new instance of ServerAzureADAdministratorsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerAzureADAdministratorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerAzureADAdministratorsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerAzureADAdministratorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an existing Azure Active Directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - administratorName - The name of server active directory administrator.
//   - parameters - The requested Azure Active Directory administrator Resource state.
//   - options - ServerAzureADAdministratorsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerAzureADAdministratorsClient.BeginCreateOrUpdate
//     method.
func (client *ServerAzureADAdministratorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, parameters ServerAzureADAdministrator, options *ServerAzureADAdministratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerAzureADAdministratorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, administratorName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerAzureADAdministratorsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerAzureADAdministratorsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates an existing Azure Active Directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *ServerAzureADAdministratorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, parameters ServerAzureADAdministrator, options *ServerAzureADAdministratorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerAzureADAdministratorsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, administratorName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerAzureADAdministratorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, parameters ServerAzureADAdministrator, options *ServerAzureADAdministratorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
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

// BeginDelete - Deletes the Azure Active Directory administrator with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - administratorName - The name of server active directory administrator.
//   - options - ServerAzureADAdministratorsClientBeginDeleteOptions contains the optional parameters for the ServerAzureADAdministratorsClient.BeginDelete
//     method.
func (client *ServerAzureADAdministratorsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *ServerAzureADAdministratorsClientBeginDeleteOptions) (*runtime.Poller[ServerAzureADAdministratorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, administratorName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerAzureADAdministratorsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerAzureADAdministratorsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the Azure Active Directory administrator with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *ServerAzureADAdministratorsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *ServerAzureADAdministratorsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerAzureADAdministratorsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, administratorName, options)
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
func (client *ServerAzureADAdministratorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *ServerAzureADAdministratorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a Azure Active Directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - administratorName - The name of server active directory administrator.
//   - options - ServerAzureADAdministratorsClientGetOptions contains the optional parameters for the ServerAzureADAdministratorsClient.Get
//     method.
func (client *ServerAzureADAdministratorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *ServerAzureADAdministratorsClientGetOptions) (ServerAzureADAdministratorsClientGetResponse, error) {
	var err error
	const operationName = "ServerAzureADAdministratorsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, administratorName, options)
	if err != nil {
		return ServerAzureADAdministratorsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerAzureADAdministratorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerAzureADAdministratorsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerAzureADAdministratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, administratorName AdministratorName, options *ServerAzureADAdministratorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if administratorName == "" {
		return nil, errors.New("parameter administratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{administratorName}", url.PathEscape(string(administratorName)))
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

// getHandleResponse handles the Get response.
func (client *ServerAzureADAdministratorsClient) getHandleResponse(resp *http.Response) (ServerAzureADAdministratorsClientGetResponse, error) {
	result := ServerAzureADAdministratorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAzureADAdministrator); err != nil {
		return ServerAzureADAdministratorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of Azure Active Directory administrators in a server.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ServerAzureADAdministratorsClientListByServerOptions contains the optional parameters for the ServerAzureADAdministratorsClient.NewListByServerPager
//     method.
func (client *ServerAzureADAdministratorsClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerAzureADAdministratorsClientListByServerOptions) *runtime.Pager[ServerAzureADAdministratorsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerAzureADAdministratorsClientListByServerResponse]{
		More: func(page ServerAzureADAdministratorsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerAzureADAdministratorsClientListByServerResponse) (ServerAzureADAdministratorsClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerAzureADAdministratorsClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return ServerAzureADAdministratorsClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerAzureADAdministratorsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAzureADAdministratorsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
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

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerAzureADAdministratorsClient) listByServerHandleResponse(resp *http.Response) (ServerAzureADAdministratorsClientListByServerResponse, error) {
	result := ServerAzureADAdministratorsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdministratorListResult); err != nil {
		return ServerAzureADAdministratorsClientListByServerResponse{}, err
	}
	return result, nil
}
