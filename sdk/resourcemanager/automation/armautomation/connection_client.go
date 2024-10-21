//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// ConnectionClient contains the methods for the Connection group.
// Don't use this type directly, use NewConnectionClient() instead.
type ConnectionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectionClient creates a new instance of ConnectionClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectionClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - connectionName - The parameters supplied to the create or update connection operation.
//   - parameters - The parameters supplied to the create or update connection operation.
//   - options - ConnectionClientCreateOrUpdateOptions contains the optional parameters for the ConnectionClient.CreateOrUpdate
//     method.
func (client *ConnectionClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, parameters ConnectionCreateOrUpdateParameters, options *ConnectionClientCreateOrUpdateOptions) (ConnectionClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ConnectionClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, connectionName, parameters, options)
	if err != nil {
		return ConnectionClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectionClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ConnectionClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, parameters ConnectionCreateOrUpdateParameters, options *ConnectionClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectionClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectionClientCreateOrUpdateResponse, error) {
	result := ConnectionClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Connection); err != nil {
		return ConnectionClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - connectionName - The name of connection.
//   - options - ConnectionClientDeleteOptions contains the optional parameters for the ConnectionClient.Delete method.
func (client *ConnectionClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, options *ConnectionClientDeleteOptions) (ConnectionClientDeleteResponse, error) {
	var err error
	const operationName = "ConnectionClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, connectionName, options)
	if err != nil {
		return ConnectionClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectionClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConnectionClientDeleteResponse{}, err
	}
	return ConnectionClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, options *ConnectionClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the connection identified by connection name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - connectionName - The name of connection.
//   - options - ConnectionClientGetOptions contains the optional parameters for the ConnectionClient.Get method.
func (client *ConnectionClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, options *ConnectionClientGetOptions) (ConnectionClientGetResponse, error) {
	var err error
	const operationName = "ConnectionClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, connectionName, options)
	if err != nil {
		return ConnectionClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectionClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConnectionClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, options *ConnectionClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectionClient) getHandleResponse(resp *http.Response) (ConnectionClientGetResponse, error) {
	result := ConnectionClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Connection); err != nil {
		return ConnectionClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of connections.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - ConnectionClientListByAutomationAccountOptions contains the optional parameters for the ConnectionClient.NewListByAutomationAccountPager
//     method.
func (client *ConnectionClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *ConnectionClientListByAutomationAccountOptions) *runtime.Pager[ConnectionClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectionClientListByAutomationAccountResponse]{
		More: func(page ConnectionClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectionClientListByAutomationAccountResponse) (ConnectionClientListByAutomationAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConnectionClient.NewListByAutomationAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			}, nil)
			if err != nil {
				return ConnectionClientListByAutomationAccountResponse{}, err
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *ConnectionClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *ConnectionClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *ConnectionClient) listByAutomationAccountHandleResponse(resp *http.Response) (ConnectionClientListByAutomationAccountResponse, error) {
	result := ConnectionClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionListResult); err != nil {
		return ConnectionClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update a connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - connectionName - The parameters supplied to the update a connection operation.
//   - parameters - The parameters supplied to the update a connection operation.
//   - options - ConnectionClientUpdateOptions contains the optional parameters for the ConnectionClient.Update method.
func (client *ConnectionClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, parameters ConnectionUpdateParameters, options *ConnectionClientUpdateOptions) (ConnectionClientUpdateResponse, error) {
	var err error
	const operationName = "ConnectionClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, connectionName, parameters, options)
	if err != nil {
		return ConnectionClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectionClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectionClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ConnectionClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, parameters ConnectionUpdateParameters, options *ConnectionClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ConnectionClient) updateHandleResponse(resp *http.Response) (ConnectionClientUpdateResponse, error) {
	result := ConnectionClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Connection); err != nil {
		return ConnectionClientUpdateResponse{}, err
	}
	return result, nil
}
