//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// ContentTemplateClient contains the methods for the ContentTemplate group.
// Don't use this type directly, use NewContentTemplateClient() instead.
type ContentTemplateClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContentTemplateClient creates a new instance of ContentTemplateClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContentTemplateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContentTemplateClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContentTemplateClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Delete - Delete an installed template.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - templateID - template Id
//   - options - ContentTemplateClientDeleteOptions contains the optional parameters for the ContentTemplateClient.Delete method.
func (client *ContentTemplateClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, templateID string, options *ContentTemplateClientDeleteOptions) (ContentTemplateClientDeleteResponse, error) {
	var err error
	const operationName = "ContentTemplateClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, templateID, options)
	if err != nil {
		return ContentTemplateClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentTemplateClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ContentTemplateClientDeleteResponse{}, err
	}
	return ContentTemplateClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContentTemplateClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, templateID string, options *ContentTemplateClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/contentTemplates/{templateId}"
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
	if templateID == "" {
		return nil, errors.New("parameter templateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateId}", url.PathEscape(templateID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a template byt its identifier. Expandable properties:
// * properties/mainTemplate
// * properties/dependantTemplates
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - templateID - template Id
//   - options - ContentTemplateClientGetOptions contains the optional parameters for the ContentTemplateClient.Get method.
func (client *ContentTemplateClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, templateID string, options *ContentTemplateClientGetOptions) (ContentTemplateClientGetResponse, error) {
	var err error
	const operationName = "ContentTemplateClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, templateID, options)
	if err != nil {
		return ContentTemplateClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentTemplateClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContentTemplateClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ContentTemplateClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, templateID string, options *ContentTemplateClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/contentTemplates/{templateId}"
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
	if templateID == "" {
		return nil, errors.New("parameter templateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateId}", url.PathEscape(templateID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContentTemplateClient) getHandleResponse(resp *http.Response) (ContentTemplateClientGetResponse, error) {
	result := ContentTemplateClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateModel); err != nil {
		return ContentTemplateClientGetResponse{}, err
	}
	return result, nil
}

// Install - Install a template.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - templateID - template Id
//   - templateInstallationProperties - Template installation properties
//   - options - ContentTemplateClientInstallOptions contains the optional parameters for the ContentTemplateClient.Install method.
func (client *ContentTemplateClient) Install(ctx context.Context, resourceGroupName string, workspaceName string, templateID string, templateInstallationProperties TemplateModel, options *ContentTemplateClientInstallOptions) (ContentTemplateClientInstallResponse, error) {
	var err error
	const operationName = "ContentTemplateClient.Install"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.installCreateRequest(ctx, resourceGroupName, workspaceName, templateID, templateInstallationProperties, options)
	if err != nil {
		return ContentTemplateClientInstallResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentTemplateClientInstallResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ContentTemplateClientInstallResponse{}, err
	}
	resp, err := client.installHandleResponse(httpResp)
	return resp, err
}

// installCreateRequest creates the Install request.
func (client *ContentTemplateClient) installCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, templateID string, templateInstallationProperties TemplateModel, options *ContentTemplateClientInstallOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/contentTemplates/{templateId}"
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
	if templateID == "" {
		return nil, errors.New("parameter templateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateId}", url.PathEscape(templateID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, templateInstallationProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// installHandleResponse handles the Install response.
func (client *ContentTemplateClient) installHandleResponse(resp *http.Response) (ContentTemplateClientInstallResponse, error) {
	result := ContentTemplateClientInstallResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateModel); err != nil {
		return ContentTemplateClientInstallResponse{}, err
	}
	return result, nil
}
