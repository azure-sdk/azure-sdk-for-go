//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

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

// ServerThreatProtectionSettingsClient contains the methods for the ServerThreatProtectionSettings group.
// Don't use this type directly, use NewServerThreatProtectionSettingsClient() instead.
type ServerThreatProtectionSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerThreatProtectionSettingsClient creates a new instance of ServerThreatProtectionSettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerThreatProtectionSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerThreatProtectionSettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerThreatProtectionSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a server's Advanced Threat Protection settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - threatProtectionName - The name of the Threat Protection state.
//   - parameters - The Advanced Threat Protection state for the flexible server.
//   - options - ServerThreatProtectionSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerThreatProtectionSettingsClient.BeginCreateOrUpdate
//     method.
func (client *ServerThreatProtectionSettingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, threatProtectionName ThreatProtectionName, parameters ServerThreatProtectionSettingsModel, options *ServerThreatProtectionSettingsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerThreatProtectionSettingsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, threatProtectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerThreatProtectionSettingsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerThreatProtectionSettingsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a server's Advanced Threat Protection settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
func (client *ServerThreatProtectionSettingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, threatProtectionName ThreatProtectionName, parameters ServerThreatProtectionSettingsModel, options *ServerThreatProtectionSettingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerThreatProtectionSettingsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, threatProtectionName, parameters, options)
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
func (client *ServerThreatProtectionSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, threatProtectionName ThreatProtectionName, parameters ServerThreatProtectionSettingsModel, options *ServerThreatProtectionSettingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/advancedThreatProtectionSettings/{threatProtectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if threatProtectionName == "" {
		return nil, errors.New("parameter threatProtectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{threatProtectionName}", url.PathEscape(string(threatProtectionName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get a server's Advanced Threat Protection settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - threatProtectionName - The name of the Threat Protection state.
//   - options - ServerThreatProtectionSettingsClientGetOptions contains the optional parameters for the ServerThreatProtectionSettingsClient.Get
//     method.
func (client *ServerThreatProtectionSettingsClient) Get(ctx context.Context, resourceGroupName string, serverName string, threatProtectionName ThreatProtectionName, options *ServerThreatProtectionSettingsClientGetOptions) (ServerThreatProtectionSettingsClientGetResponse, error) {
	var err error
	const operationName = "ServerThreatProtectionSettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, threatProtectionName, options)
	if err != nil {
		return ServerThreatProtectionSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerThreatProtectionSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerThreatProtectionSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerThreatProtectionSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, threatProtectionName ThreatProtectionName, options *ServerThreatProtectionSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/advancedThreatProtectionSettings/{threatProtectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if threatProtectionName == "" {
		return nil, errors.New("parameter threatProtectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{threatProtectionName}", url.PathEscape(string(threatProtectionName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerThreatProtectionSettingsClient) getHandleResponse(resp *http.Response) (ServerThreatProtectionSettingsClientGetResponse, error) {
	result := ServerThreatProtectionSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerThreatProtectionSettingsModel); err != nil {
		return ServerThreatProtectionSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Get a list of server's Threat Protection state.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - ServerThreatProtectionSettingsClientListByServerOptions contains the optional parameters for the ServerThreatProtectionSettingsClient.NewListByServerPager
//     method.
func (client *ServerThreatProtectionSettingsClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerThreatProtectionSettingsClientListByServerOptions) *runtime.Pager[ServerThreatProtectionSettingsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerThreatProtectionSettingsClientListByServerResponse]{
		More: func(page ServerThreatProtectionSettingsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerThreatProtectionSettingsClientListByServerResponse) (ServerThreatProtectionSettingsClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerThreatProtectionSettingsClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return ServerThreatProtectionSettingsClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerThreatProtectionSettingsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerThreatProtectionSettingsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/advancedThreatProtectionSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerThreatProtectionSettingsClient) listByServerHandleResponse(resp *http.Response) (ServerThreatProtectionSettingsClientListByServerResponse, error) {
	result := ServerThreatProtectionSettingsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerThreatProtectionListResult); err != nil {
		return ServerThreatProtectionSettingsClientListByServerResponse{}, err
	}
	return result, nil
}
