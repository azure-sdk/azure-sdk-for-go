//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// ServerConfigurationOptionsClient contains the methods for the ServerConfigurationOptions group.
// Don't use this type directly, use NewServerConfigurationOptionsClient() instead.
type ServerConfigurationOptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerConfigurationOptionsClient creates a new instance of ServerConfigurationOptionsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerConfigurationOptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerConfigurationOptionsClient, error) {
	cl, err := arm.NewClient(moduleName+".ServerConfigurationOptionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerConfigurationOptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Updates managed instance server configuration option.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - serverConfigurationOptionName - The name of the server configuration option.
//   - parameters - Server configuration option parameters.
//   - options - ServerConfigurationOptionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerConfigurationOptionsClient.BeginCreateOrUpdate
//     method.
func (client *ServerConfigurationOptionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, serverConfigurationOptionName ServerConfigurationOptionName, parameters ServerConfigurationOption, options *ServerConfigurationOptionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerConfigurationOptionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, serverConfigurationOptionName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServerConfigurationOptionsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServerConfigurationOptionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Updates managed instance server configuration option.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
func (client *ServerConfigurationOptionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, serverConfigurationOptionName ServerConfigurationOptionName, parameters ServerConfigurationOption, options *ServerConfigurationOptionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, serverConfigurationOptionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerConfigurationOptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, serverConfigurationOptionName ServerConfigurationOptionName, parameters ServerConfigurationOption, options *ServerConfigurationOptionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/serverConfigurationOptions/{serverConfigurationOptionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if serverConfigurationOptionName == "" {
		return nil, errors.New("parameter serverConfigurationOptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverConfigurationOptionName}", url.PathEscape(string(serverConfigurationOptionName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Gets managed instance server configuration option.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - serverConfigurationOptionName - The name of the server configuration option.
//   - options - ServerConfigurationOptionsClientGetOptions contains the optional parameters for the ServerConfigurationOptionsClient.Get
//     method.
func (client *ServerConfigurationOptionsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, serverConfigurationOptionName ServerConfigurationOptionName, options *ServerConfigurationOptionsClientGetOptions) (ServerConfigurationOptionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, serverConfigurationOptionName, options)
	if err != nil {
		return ServerConfigurationOptionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerConfigurationOptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerConfigurationOptionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerConfigurationOptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, serverConfigurationOptionName ServerConfigurationOptionName, options *ServerConfigurationOptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/serverConfigurationOptions/{serverConfigurationOptionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if serverConfigurationOptionName == "" {
		return nil, errors.New("parameter serverConfigurationOptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverConfigurationOptionName}", url.PathEscape(string(serverConfigurationOptionName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerConfigurationOptionsClient) getHandleResponse(resp *http.Response) (ServerConfigurationOptionsClientGetResponse, error) {
	result := ServerConfigurationOptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConfigurationOption); err != nil {
		return ServerConfigurationOptionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagedInstancePager - Gets a list of managed instance server configuration options.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ServerConfigurationOptionsClientListByManagedInstanceOptions contains the optional parameters for the ServerConfigurationOptionsClient.NewListByManagedInstancePager
//     method.
func (client *ServerConfigurationOptionsClient) NewListByManagedInstancePager(resourceGroupName string, managedInstanceName string, options *ServerConfigurationOptionsClientListByManagedInstanceOptions) *runtime.Pager[ServerConfigurationOptionsClientListByManagedInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerConfigurationOptionsClientListByManagedInstanceResponse]{
		More: func(page ServerConfigurationOptionsClientListByManagedInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerConfigurationOptionsClientListByManagedInstanceResponse) (ServerConfigurationOptionsClientListByManagedInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByManagedInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServerConfigurationOptionsClientListByManagedInstanceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServerConfigurationOptionsClientListByManagedInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerConfigurationOptionsClientListByManagedInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagedInstanceHandleResponse(resp)
		},
	})
}

// listByManagedInstanceCreateRequest creates the ListByManagedInstance request.
func (client *ServerConfigurationOptionsClient) listByManagedInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ServerConfigurationOptionsClientListByManagedInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/serverConfigurationOptions"
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
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedInstanceHandleResponse handles the ListByManagedInstance response.
func (client *ServerConfigurationOptionsClient) listByManagedInstanceHandleResponse(resp *http.Response) (ServerConfigurationOptionsClientListByManagedInstanceResponse, error) {
	result := ServerConfigurationOptionsClientListByManagedInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConfigurationOptionListResult); err != nil {
		return ServerConfigurationOptionsClientListByManagedInstanceResponse{}, err
	}
	return result, nil
}
