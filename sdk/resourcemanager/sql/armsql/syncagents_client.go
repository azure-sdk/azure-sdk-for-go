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

// SyncAgentsClient contains the methods for the SyncAgents group.
// Don't use this type directly, use NewSyncAgentsClient() instead.
type SyncAgentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSyncAgentsClient creates a new instance of SyncAgentsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSyncAgentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SyncAgentsClient, error) {
	cl, err := arm.NewClient(moduleName+".SyncAgentsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SyncAgentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a sync agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server on which the sync agent is hosted.
//   - syncAgentName - The name of the sync agent.
//   - parameters - The requested sync agent resource state.
//   - options - SyncAgentsClientBeginCreateOrUpdateOptions contains the optional parameters for the SyncAgentsClient.BeginCreateOrUpdate
//     method.
func (client *SyncAgentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, parameters SyncAgent, options *SyncAgentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SyncAgentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, syncAgentName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SyncAgentsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SyncAgentsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a sync agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
func (client *SyncAgentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, parameters SyncAgent, options *SyncAgentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, syncAgentName, parameters, options)
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
func (client *SyncAgentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, parameters SyncAgent, options *SyncAgentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if syncAgentName == "" {
		return nil, errors.New("parameter syncAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncAgentName}", url.PathEscape(syncAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a sync agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server on which the sync agent is hosted.
//   - syncAgentName - The name of the sync agent.
//   - options - SyncAgentsClientBeginDeleteOptions contains the optional parameters for the SyncAgentsClient.BeginDelete method.
func (client *SyncAgentsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsClientBeginDeleteOptions) (*runtime.Poller[SyncAgentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, syncAgentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SyncAgentsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SyncAgentsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a sync agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
func (client *SyncAgentsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, syncAgentName, options)
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
func (client *SyncAgentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if syncAgentName == "" {
		return nil, errors.New("parameter syncAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncAgentName}", url.PathEscape(syncAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// GenerateKey - Generates a sync agent key.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server on which the sync agent is hosted.
//   - syncAgentName - The name of the sync agent.
//   - options - SyncAgentsClientGenerateKeyOptions contains the optional parameters for the SyncAgentsClient.GenerateKey method.
func (client *SyncAgentsClient) GenerateKey(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsClientGenerateKeyOptions) (SyncAgentsClientGenerateKeyResponse, error) {
	var err error
	req, err := client.generateKeyCreateRequest(ctx, resourceGroupName, serverName, syncAgentName, options)
	if err != nil {
		return SyncAgentsClientGenerateKeyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SyncAgentsClientGenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SyncAgentsClientGenerateKeyResponse{}, err
	}
	resp, err := client.generateKeyHandleResponse(httpResp)
	return resp, err
}

// generateKeyCreateRequest creates the GenerateKey request.
func (client *SyncAgentsClient) generateKeyCreateRequest(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsClientGenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}/generateKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if syncAgentName == "" {
		return nil, errors.New("parameter syncAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncAgentName}", url.PathEscape(syncAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generateKeyHandleResponse handles the GenerateKey response.
func (client *SyncAgentsClient) generateKeyHandleResponse(resp *http.Response) (SyncAgentsClientGenerateKeyResponse, error) {
	result := SyncAgentsClientGenerateKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncAgentKeyProperties); err != nil {
		return SyncAgentsClientGenerateKeyResponse{}, err
	}
	return result, nil
}

// Get - Gets a sync agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server on which the sync agent is hosted.
//   - syncAgentName - The name of the sync agent.
//   - options - SyncAgentsClientGetOptions contains the optional parameters for the SyncAgentsClient.Get method.
func (client *SyncAgentsClient) Get(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsClientGetOptions) (SyncAgentsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, syncAgentName, options)
	if err != nil {
		return SyncAgentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SyncAgentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SyncAgentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SyncAgentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if syncAgentName == "" {
		return nil, errors.New("parameter syncAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncAgentName}", url.PathEscape(syncAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SyncAgentsClient) getHandleResponse(resp *http.Response) (SyncAgentsClientGetResponse, error) {
	result := SyncAgentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncAgent); err != nil {
		return SyncAgentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Lists sync agents in a server.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server on which the sync agent is hosted.
//   - options - SyncAgentsClientListByServerOptions contains the optional parameters for the SyncAgentsClient.NewListByServerPager
//     method.
func (client *SyncAgentsClient) NewListByServerPager(resourceGroupName string, serverName string, options *SyncAgentsClientListByServerOptions) *runtime.Pager[SyncAgentsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[SyncAgentsClientListByServerResponse]{
		More: func(page SyncAgentsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SyncAgentsClientListByServerResponse) (SyncAgentsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SyncAgentsClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SyncAgentsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SyncAgentsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *SyncAgentsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *SyncAgentsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents"
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
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *SyncAgentsClient) listByServerHandleResponse(resp *http.Response) (SyncAgentsClientListByServerResponse, error) {
	result := SyncAgentsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncAgentListResult); err != nil {
		return SyncAgentsClientListByServerResponse{}, err
	}
	return result, nil
}

// NewListLinkedDatabasesPager - Lists databases linked to a sync agent.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server on which the sync agent is hosted.
//   - syncAgentName - The name of the sync agent.
//   - options - SyncAgentsClientListLinkedDatabasesOptions contains the optional parameters for the SyncAgentsClient.NewListLinkedDatabasesPager
//     method.
func (client *SyncAgentsClient) NewListLinkedDatabasesPager(resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsClientListLinkedDatabasesOptions) *runtime.Pager[SyncAgentsClientListLinkedDatabasesResponse] {
	return runtime.NewPager(runtime.PagingHandler[SyncAgentsClientListLinkedDatabasesResponse]{
		More: func(page SyncAgentsClientListLinkedDatabasesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SyncAgentsClientListLinkedDatabasesResponse) (SyncAgentsClientListLinkedDatabasesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listLinkedDatabasesCreateRequest(ctx, resourceGroupName, serverName, syncAgentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SyncAgentsClientListLinkedDatabasesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SyncAgentsClientListLinkedDatabasesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SyncAgentsClientListLinkedDatabasesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listLinkedDatabasesHandleResponse(resp)
		},
	})
}

// listLinkedDatabasesCreateRequest creates the ListLinkedDatabases request.
func (client *SyncAgentsClient) listLinkedDatabasesCreateRequest(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, options *SyncAgentsClientListLinkedDatabasesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}/linkedDatabases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if syncAgentName == "" {
		return nil, errors.New("parameter syncAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncAgentName}", url.PathEscape(syncAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listLinkedDatabasesHandleResponse handles the ListLinkedDatabases response.
func (client *SyncAgentsClient) listLinkedDatabasesHandleResponse(resp *http.Response) (SyncAgentsClientListLinkedDatabasesResponse, error) {
	result := SyncAgentsClientListLinkedDatabasesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncAgentLinkedDatabaseListResult); err != nil {
		return SyncAgentsClientListLinkedDatabasesResponse{}, err
	}
	return result, nil
}
