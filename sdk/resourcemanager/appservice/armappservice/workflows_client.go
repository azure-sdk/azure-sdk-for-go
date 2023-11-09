//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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

// WorkflowsClient contains the methods for the Workflows group.
// Don't use this type directly, use NewWorkflowsClient() instead.
type WorkflowsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkflowsClient creates a new instance of WorkflowsClient with the specified values.
//   - subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkflowsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkflowsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkflowsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// RegenerateAccessKey - Regenerates the callback URL access key for request triggers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Site name.
//   - workflowName - The workflow name.
//   - keyType - The access key type.
//   - options - WorkflowsClientRegenerateAccessKeyOptions contains the optional parameters for the WorkflowsClient.RegenerateAccessKey
//     method.
func (client *WorkflowsClient) RegenerateAccessKey(ctx context.Context, resourceGroupName string, name string, workflowName string, keyType RegenerateActionParameter, options *WorkflowsClientRegenerateAccessKeyOptions) (WorkflowsClientRegenerateAccessKeyResponse, error) {
	var err error
	const operationName = "WorkflowsClient.RegenerateAccessKey"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regenerateAccessKeyCreateRequest(ctx, resourceGroupName, name, workflowName, keyType, options)
	if err != nil {
		return WorkflowsClientRegenerateAccessKeyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowsClientRegenerateAccessKeyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowsClientRegenerateAccessKeyResponse{}, err
	}
	return WorkflowsClientRegenerateAccessKeyResponse{}, nil
}

// regenerateAccessKeyCreateRequest creates the RegenerateAccessKey request.
func (client *WorkflowsClient) regenerateAccessKeyCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, keyType RegenerateActionParameter, options *WorkflowsClientRegenerateAccessKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/regenerateAccessKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, keyType); err != nil {
		return nil, err
	}
	return req, nil
}

// Validate - Validates the workflow definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Site name.
//   - workflowName - The workflow name.
//   - validate - The workflow.
//   - options - WorkflowsClientValidateOptions contains the optional parameters for the WorkflowsClient.Validate method.
func (client *WorkflowsClient) Validate(ctx context.Context, resourceGroupName string, name string, workflowName string, validate Workflow, options *WorkflowsClientValidateOptions) (WorkflowsClientValidateResponse, error) {
	var err error
	const operationName = "WorkflowsClient.Validate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateCreateRequest(ctx, resourceGroupName, name, workflowName, validate, options)
	if err != nil {
		return WorkflowsClientValidateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowsClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowsClientValidateResponse{}, err
	}
	return WorkflowsClientValidateResponse{}, nil
}

// validateCreateRequest creates the Validate request.
func (client *WorkflowsClient) validateCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, validate Workflow, options *WorkflowsClientValidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/validate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, validate); err != nil {
		return nil, err
	}
	return req, nil
}
