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

// WorkspaceManagedIdentitySQLControlSettingsClient contains the methods for the WorkspaceManagedIdentitySQLControlSettings group.
// Don't use this type directly, use NewWorkspaceManagedIdentitySQLControlSettingsClient() instead.
type WorkspaceManagedIdentitySQLControlSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceManagedIdentitySQLControlSettingsClient creates a new instance of WorkspaceManagedIdentitySQLControlSettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagedIdentitySQLControlSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagedIdentitySQLControlSettingsClient, error) {
	cl, err := arm.NewClient(moduleName+".WorkspaceManagedIdentitySQLControlSettingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagedIdentitySQLControlSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update Managed Identity Sql Control Settings
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - managedIdentitySQLControlSettings - Managed Identity Sql Control Settings
//   - options - WorkspaceManagedIdentitySQLControlSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for
//     the WorkspaceManagedIdentitySQLControlSettingsClient.BeginCreateOrUpdate method.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, managedIdentitySQLControlSettings ManagedIdentitySQLControlSettingsModel, options *WorkspaceManagedIdentitySQLControlSettingsClientBeginCreateOrUpdateOptions) (*runtime.Poller[WorkspaceManagedIdentitySQLControlSettingsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, managedIdentitySQLControlSettings, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WorkspaceManagedIdentitySQLControlSettingsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[WorkspaceManagedIdentitySQLControlSettingsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update Managed Identity Sql Control Settings
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, managedIdentitySQLControlSettings ManagedIdentitySQLControlSettingsModel, options *WorkspaceManagedIdentitySQLControlSettingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, managedIdentitySQLControlSettings, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, managedIdentitySQLControlSettings ManagedIdentitySQLControlSettingsModel, options *WorkspaceManagedIdentitySQLControlSettingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/managedIdentitySqlControlSettings/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, managedIdentitySQLControlSettings)
}

// Get - Get Managed Identity Sql Control Settings
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - WorkspaceManagedIdentitySQLControlSettingsClientGetOptions contains the optional parameters for the WorkspaceManagedIdentitySQLControlSettingsClient.Get
//     method.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedIdentitySQLControlSettingsClientGetOptions) (WorkspaceManagedIdentitySQLControlSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return WorkspaceManagedIdentitySQLControlSettingsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagedIdentitySQLControlSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedIdentitySQLControlSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedIdentitySQLControlSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/managedIdentitySqlControlSettings/default"
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
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagedIdentitySQLControlSettingsClient) getHandleResponse(resp *http.Response) (WorkspaceManagedIdentitySQLControlSettingsClientGetResponse, error) {
	result := WorkspaceManagedIdentitySQLControlSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedIdentitySQLControlSettingsModel); err != nil {
		return WorkspaceManagedIdentitySQLControlSettingsClientGetResponse{}, err
	}
	return result, nil
}
