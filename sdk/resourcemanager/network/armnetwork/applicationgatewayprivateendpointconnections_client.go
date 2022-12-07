//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ApplicationGatewayPrivateEndpointConnectionsClient contains the methods for the ApplicationGatewayPrivateEndpointConnections group.
// Don't use this type directly, use NewApplicationGatewayPrivateEndpointConnectionsClient() instead.
type ApplicationGatewayPrivateEndpointConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationGatewayPrivateEndpointConnectionsClient creates a new instance of ApplicationGatewayPrivateEndpointConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApplicationGatewayPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationGatewayPrivateEndpointConnectionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationGatewayPrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginDelete - Deletes the specified private endpoint connection on application gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// applicationGatewayName - The name of the application gateway.
// connectionName - The name of the application gateway private endpoint connection.
// options - ApplicationGatewayPrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the
// ApplicationGatewayPrivateEndpointConnectionsClient.BeginDelete method.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[ApplicationGatewayPrivateEndpointConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, applicationGatewayName, connectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ApplicationGatewayPrivateEndpointConnectionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationGatewayPrivateEndpointConnectionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified private endpoint connection on application gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationGatewayName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified private endpoint connection on application gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// applicationGatewayName - The name of the application gateway.
// connectionName - The name of the application gateway private endpoint connection.
// options - ApplicationGatewayPrivateEndpointConnectionsClientGetOptions contains the optional parameters for the ApplicationGatewayPrivateEndpointConnectionsClient.Get
// method.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsClientGetOptions) (ApplicationGatewayPrivateEndpointConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationGatewayName, connectionName, options)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationGatewayPrivateEndpointConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *ApplicationGatewayPrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (ApplicationGatewayPrivateEndpointConnectionsClientGetResponse, error) {
	result := ApplicationGatewayPrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGatewayPrivateEndpointConnection); err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all private endpoint connections on an application gateway.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// applicationGatewayName - The name of the application gateway.
// options - ApplicationGatewayPrivateEndpointConnectionsClientListOptions contains the optional parameters for the ApplicationGatewayPrivateEndpointConnectionsClient.List
// method.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) NewListPager(resourceGroupName string, applicationGatewayName string, options *ApplicationGatewayPrivateEndpointConnectionsClientListOptions) *runtime.Pager[ApplicationGatewayPrivateEndpointConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationGatewayPrivateEndpointConnectionsClientListResponse]{
		More: func(page ApplicationGatewayPrivateEndpointConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationGatewayPrivateEndpointConnectionsClientListResponse) (ApplicationGatewayPrivateEndpointConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, applicationGatewayName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationGatewayPrivateEndpointConnectionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationGatewayPrivateEndpointConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationGatewayPrivateEndpointConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, options *ApplicationGatewayPrivateEndpointConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) listHandleResponse(resp *http.Response) (ApplicationGatewayPrivateEndpointConnectionsClientListResponse, error) {
	result := ApplicationGatewayPrivateEndpointConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGatewayPrivateEndpointConnectionListResult); err != nil {
		return ApplicationGatewayPrivateEndpointConnectionsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the specified private endpoint connection on application gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// applicationGatewayName - The name of the application gateway.
// connectionName - The name of the application gateway private endpoint connection.
// parameters - Parameters supplied to update application gateway private endpoint connection operation.
// options - ApplicationGatewayPrivateEndpointConnectionsClientBeginUpdateOptions contains the optional parameters for the
// ApplicationGatewayPrivateEndpointConnectionsClient.BeginUpdate method.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection, options *ApplicationGatewayPrivateEndpointConnectionsClientBeginUpdateOptions) (*runtime.Poller[ApplicationGatewayPrivateEndpointConnectionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, applicationGatewayName, connectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ApplicationGatewayPrivateEndpointConnectionsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationGatewayPrivateEndpointConnectionsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates the specified private endpoint connection on application gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) update(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection, options *ApplicationGatewayPrivateEndpointConnectionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, applicationGatewayName, connectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ApplicationGatewayPrivateEndpointConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters ApplicationGatewayPrivateEndpointConnection, options *ApplicationGatewayPrivateEndpointConnectionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateEndpointConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
