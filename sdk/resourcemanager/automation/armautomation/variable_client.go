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

// VariableClient contains the methods for the Variable group.
// Don't use this type directly, use NewVariableClient() instead.
type VariableClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVariableClient creates a new instance of VariableClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVariableClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VariableClient, error) {
	cl, err := arm.NewClient(moduleName+".VariableClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VariableClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a variable.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - variableName - The variable name.
//   - parameters - The parameters supplied to the create or update variable operation.
//   - options - VariableClientCreateOrUpdateOptions contains the optional parameters for the VariableClient.CreateOrUpdate method.
func (client *VariableClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string, parameters VariableCreateOrUpdateParameters, options *VariableClientCreateOrUpdateOptions) (VariableClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, variableName, parameters, options)
	if err != nil {
		return VariableClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VariableClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return VariableClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VariableClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string, parameters VariableCreateOrUpdateParameters, options *VariableClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/variables/{variableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if variableName == "" {
		return nil, errors.New("parameter variableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{variableName}", url.PathEscape(variableName))
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
func (client *VariableClient) createOrUpdateHandleResponse(resp *http.Response) (VariableClientCreateOrUpdateResponse, error) {
	result := VariableClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Variable); err != nil {
		return VariableClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the variable.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - variableName - The name of variable.
//   - options - VariableClientDeleteOptions contains the optional parameters for the VariableClient.Delete method.
func (client *VariableClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string, options *VariableClientDeleteOptions) (VariableClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, variableName, options)
	if err != nil {
		return VariableClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VariableClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VariableClientDeleteResponse{}, err
	}
	return VariableClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VariableClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string, options *VariableClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/variables/{variableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if variableName == "" {
		return nil, errors.New("parameter variableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{variableName}", url.PathEscape(variableName))
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

// Get - Retrieve the variable identified by variable name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - variableName - The name of variable.
//   - options - VariableClientGetOptions contains the optional parameters for the VariableClient.Get method.
func (client *VariableClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string, options *VariableClientGetOptions) (VariableClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, variableName, options)
	if err != nil {
		return VariableClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VariableClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VariableClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VariableClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string, options *VariableClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/variables/{variableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if variableName == "" {
		return nil, errors.New("parameter variableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{variableName}", url.PathEscape(variableName))
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
func (client *VariableClient) getHandleResponse(resp *http.Response) (VariableClientGetResponse, error) {
	result := VariableClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Variable); err != nil {
		return VariableClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of variables.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - VariableClientListByAutomationAccountOptions contains the optional parameters for the VariableClient.NewListByAutomationAccountPager
//     method.
func (client *VariableClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *VariableClientListByAutomationAccountOptions) *runtime.Pager[VariableClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[VariableClientListByAutomationAccountResponse]{
		More: func(page VariableClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VariableClientListByAutomationAccountResponse) (VariableClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VariableClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VariableClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VariableClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *VariableClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *VariableClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/variables"
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
func (client *VariableClient) listByAutomationAccountHandleResponse(resp *http.Response) (VariableClientListByAutomationAccountResponse, error) {
	result := VariableClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VariableListResult); err != nil {
		return VariableClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update a variable.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - variableName - The variable name.
//   - parameters - The parameters supplied to the update variable operation.
//   - options - VariableClientUpdateOptions contains the optional parameters for the VariableClient.Update method.
func (client *VariableClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string, parameters VariableUpdateParameters, options *VariableClientUpdateOptions) (VariableClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, variableName, parameters, options)
	if err != nil {
		return VariableClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VariableClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VariableClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *VariableClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, variableName string, parameters VariableUpdateParameters, options *VariableClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/variables/{variableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if variableName == "" {
		return nil, errors.New("parameter variableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{variableName}", url.PathEscape(variableName))
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
func (client *VariableClient) updateHandleResponse(resp *http.Response) (VariableClientUpdateResponse, error) {
	result := VariableClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Variable); err != nil {
		return VariableClientUpdateResponse{}, err
	}
	return result, nil
}
