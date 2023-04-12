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

// IntegrationRuntimeAuthKeysClient contains the methods for the IntegrationRuntimeAuthKeys group.
// Don't use this type directly, use NewIntegrationRuntimeAuthKeysClient() instead.
type IntegrationRuntimeAuthKeysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationRuntimeAuthKeysClient creates a new instance of IntegrationRuntimeAuthKeysClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationRuntimeAuthKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationRuntimeAuthKeysClient, error) {
	cl, err := arm.NewClient(moduleName+".IntegrationRuntimeAuthKeysClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationRuntimeAuthKeysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - List authentication keys in an integration runtime
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - integrationRuntimeName - Integration runtime name
//   - options - IntegrationRuntimeAuthKeysClientListOptions contains the optional parameters for the IntegrationRuntimeAuthKeysClient.List
//     method.
func (client *IntegrationRuntimeAuthKeysClient) List(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeAuthKeysClientListOptions) (IntegrationRuntimeAuthKeysClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeAuthKeysClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimeAuthKeysClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeAuthKeysClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *IntegrationRuntimeAuthKeysClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeAuthKeysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/listAuthKeys"
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
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationRuntimeAuthKeysClient) listHandleResponse(resp *http.Response) (IntegrationRuntimeAuthKeysClientListResponse, error) {
	result := IntegrationRuntimeAuthKeysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeAuthKeys); err != nil {
		return IntegrationRuntimeAuthKeysClientListResponse{}, err
	}
	return result, nil
}

// Regenerate - Regenerate the authentication key for an integration runtime
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - integrationRuntimeName - Integration runtime name
//   - regenerateKeyParameters - The parameters for regenerating integration runtime authentication key.
//   - options - IntegrationRuntimeAuthKeysClientRegenerateOptions contains the optional parameters for the IntegrationRuntimeAuthKeysClient.Regenerate
//     method.
func (client *IntegrationRuntimeAuthKeysClient) Regenerate(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, regenerateKeyParameters IntegrationRuntimeRegenerateKeyParameters, options *IntegrationRuntimeAuthKeysClientRegenerateOptions) (IntegrationRuntimeAuthKeysClientRegenerateResponse, error) {
	req, err := client.regenerateCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, regenerateKeyParameters, options)
	if err != nil {
		return IntegrationRuntimeAuthKeysClientRegenerateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimeAuthKeysClientRegenerateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeAuthKeysClientRegenerateResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateHandleResponse(resp)
}

// regenerateCreateRequest creates the Regenerate request.
func (client *IntegrationRuntimeAuthKeysClient) regenerateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, regenerateKeyParameters IntegrationRuntimeRegenerateKeyParameters, options *IntegrationRuntimeAuthKeysClientRegenerateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/regenerateAuthKey"
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
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, regenerateKeyParameters)
}

// regenerateHandleResponse handles the Regenerate response.
func (client *IntegrationRuntimeAuthKeysClient) regenerateHandleResponse(resp *http.Response) (IntegrationRuntimeAuthKeysClientRegenerateResponse, error) {
	result := IntegrationRuntimeAuthKeysClientRegenerateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeAuthKeys); err != nil {
		return IntegrationRuntimeAuthKeysClientRegenerateResponse{}, err
	}
	return result, nil
}
