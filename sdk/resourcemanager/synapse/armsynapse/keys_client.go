//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse

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

// KeysClient contains the methods for the Keys group.
// Don't use this type directly, use NewKeysClient() instead.
type KeysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewKeysClient creates a new instance of KeysClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KeysClient, error) {
	cl, err := arm.NewClient(moduleName+".KeysClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &KeysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a workspace key
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - keyName - The name of the workspace key
//   - keyProperties - Key put request properties
//   - options - KeysClientCreateOrUpdateOptions contains the optional parameters for the KeysClient.CreateOrUpdate method.
func (client *KeysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, keyName string, keyProperties Key, options *KeysClientCreateOrUpdateOptions) (KeysClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, keyName, keyProperties, options)
	if err != nil {
		return KeysClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KeysClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KeysClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *KeysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, keyName string, keyProperties Key, options *KeysClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/keys/{keyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, keyProperties)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *KeysClient) createOrUpdateHandleResponse(resp *http.Response) (KeysClientCreateOrUpdateResponse, error) {
	result := KeysClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Key); err != nil {
		return KeysClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a workspace key
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - keyName - The name of the workspace key
//   - options - KeysClientDeleteOptions contains the optional parameters for the KeysClient.Delete method.
func (client *KeysClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, keyName string, options *KeysClientDeleteOptions) (KeysClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, keyName, options)
	if err != nil {
		return KeysClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KeysClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return KeysClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *KeysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, keyName string, options *KeysClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/keys/{keyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *KeysClient) deleteHandleResponse(resp *http.Response) (KeysClientDeleteResponse, error) {
	result := KeysClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Key); err != nil {
		return KeysClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Gets a workspace key
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - keyName - The name of the workspace key
//   - options - KeysClientGetOptions contains the optional parameters for the KeysClient.Get method.
func (client *KeysClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, keyName string, options *KeysClientGetOptions) (KeysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, keyName, options)
	if err != nil {
		return KeysClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KeysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KeysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *KeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, keyName string, options *KeysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/keys/{keyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
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
func (client *KeysClient) getHandleResponse(resp *http.Response) (KeysClientGetResponse, error) {
	result := KeysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Key); err != nil {
		return KeysClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Returns a list of keys in a workspace
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - KeysClientListByWorkspaceOptions contains the optional parameters for the KeysClient.NewListByWorkspacePager
//     method.
func (client *KeysClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *KeysClientListByWorkspaceOptions) *runtime.Pager[KeysClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[KeysClientListByWorkspaceResponse]{
		More: func(page KeysClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *KeysClientListByWorkspaceResponse) (KeysClientListByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return KeysClientListByWorkspaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return KeysClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return KeysClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *KeysClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *KeysClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/keys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
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

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *KeysClient) listByWorkspaceHandleResponse(resp *http.Response) (KeysClientListByWorkspaceResponse, error) {
	result := KeysClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KeyInfoListResult); err != nil {
		return KeysClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
