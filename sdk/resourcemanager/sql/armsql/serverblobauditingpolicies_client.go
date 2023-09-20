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

// ServerBlobAuditingPoliciesClient contains the methods for the ServerBlobAuditingPolicies group.
// Don't use this type directly, use NewServerBlobAuditingPoliciesClient() instead.
type ServerBlobAuditingPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerBlobAuditingPoliciesClient creates a new instance of ServerBlobAuditingPoliciesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerBlobAuditingPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerBlobAuditingPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".ServerBlobAuditingPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerBlobAuditingPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a server's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - parameters - Properties of blob auditing policy
//   - options - ServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerBlobAuditingPoliciesClient.BeginCreateOrUpdate
//     method.
func (client *ServerBlobAuditingPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters ServerBlobAuditingPolicy, options *ServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerBlobAuditingPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ServerBlobAuditingPoliciesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ServerBlobAuditingPoliciesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a server's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *ServerBlobAuditingPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters ServerBlobAuditingPolicy, options *ServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, parameters, options)
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
func (client *ServerBlobAuditingPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, parameters ServerBlobAuditingPolicy, options *ServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingSettings/{blobAuditingPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape("default"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets a server's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ServerBlobAuditingPoliciesClientGetOptions contains the optional parameters for the ServerBlobAuditingPoliciesClient.Get
//     method.
func (client *ServerBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, options *ServerBlobAuditingPoliciesClientGetOptions) (ServerBlobAuditingPoliciesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerBlobAuditingPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerBlobAuditingPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerBlobAuditingPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerBlobAuditingPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerBlobAuditingPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingSettings/{blobAuditingPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape("default"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerBlobAuditingPoliciesClient) getHandleResponse(resp *http.Response) (ServerBlobAuditingPoliciesClientGetResponse, error) {
	result := ServerBlobAuditingPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerBlobAuditingPolicy); err != nil {
		return ServerBlobAuditingPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Lists auditing settings of a server.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ServerBlobAuditingPoliciesClientListByServerOptions contains the optional parameters for the ServerBlobAuditingPoliciesClient.NewListByServerPager
//     method.
func (client *ServerBlobAuditingPoliciesClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerBlobAuditingPoliciesClientListByServerOptions) *runtime.Pager[ServerBlobAuditingPoliciesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerBlobAuditingPoliciesClientListByServerResponse]{
		More: func(page ServerBlobAuditingPoliciesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerBlobAuditingPoliciesClientListByServerResponse) (ServerBlobAuditingPoliciesClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServerBlobAuditingPoliciesClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServerBlobAuditingPoliciesClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerBlobAuditingPoliciesClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerBlobAuditingPoliciesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerBlobAuditingPoliciesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingSettings"
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
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerBlobAuditingPoliciesClient) listByServerHandleResponse(resp *http.Response) (ServerBlobAuditingPoliciesClientListByServerResponse, error) {
	result := ServerBlobAuditingPoliciesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerBlobAuditingPolicyListResult); err != nil {
		return ServerBlobAuditingPoliciesClientListByServerResponse{}, err
	}
	return result, nil
}
