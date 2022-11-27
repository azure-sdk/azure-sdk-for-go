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
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagementClient contains the methods for the SynapseManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementClient, error) {
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
	client := &ManagementClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// IntegrationRuntimeStartOperationStatus - Get an integration runtime start operation status
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// integrationRuntimeName - Integration runtime name
// integrationRuntimeAction - Integration runtime operation id parameter name
// integrationRuntimeOperationID - Integration runtime operation id parameter name
// options - ManagementClientIntegrationRuntimeStartOperationStatusOptions contains the optional parameters for the ManagementClient.IntegrationRuntimeStartOperationStatus
// method.
func (client *ManagementClient) IntegrationRuntimeStartOperationStatus(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, integrationRuntimeAction string, integrationRuntimeOperationID string, options *ManagementClientIntegrationRuntimeStartOperationStatusOptions) (ManagementClientIntegrationRuntimeStartOperationStatusResponse, error) {
	req, err := client.integrationRuntimeStartOperationStatusCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, integrationRuntimeAction, integrationRuntimeOperationID, options)
	if err != nil {
		return ManagementClientIntegrationRuntimeStartOperationStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementClientIntegrationRuntimeStartOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientIntegrationRuntimeStartOperationStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.integrationRuntimeStartOperationStatusHandleResponse(resp)
}

// integrationRuntimeStartOperationStatusCreateRequest creates the IntegrationRuntimeStartOperationStatus request.
func (client *ManagementClient) integrationRuntimeStartOperationStatusCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, integrationRuntimeAction string, integrationRuntimeOperationID string, options *ManagementClientIntegrationRuntimeStartOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/{integrationRuntimeAction}/operationstatuses/{integrationRuntimeOperationId}"
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
	if integrationRuntimeAction == "" {
		return nil, errors.New("parameter integrationRuntimeAction cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeAction}", url.PathEscape(integrationRuntimeAction))
	if integrationRuntimeOperationID == "" {
		return nil, errors.New("parameter integrationRuntimeOperationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeOperationId}", url.PathEscape(integrationRuntimeOperationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// integrationRuntimeStartOperationStatusHandleResponse handles the IntegrationRuntimeStartOperationStatus response.
func (client *ManagementClient) integrationRuntimeStartOperationStatusHandleResponse(resp *http.Response) (ManagementClientIntegrationRuntimeStartOperationStatusResponse, error) {
	result := ManagementClientIntegrationRuntimeStartOperationStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeOperationStatus); err != nil {
		return ManagementClientIntegrationRuntimeStartOperationStatusResponse{}, err
	}
	return result, nil
}
