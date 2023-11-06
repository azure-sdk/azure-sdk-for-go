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

// Python3PackageClient contains the methods for the Python3Package group.
// Don't use this type directly, use NewPython3PackageClient() instead.
type Python3PackageClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPython3PackageClient creates a new instance of Python3PackageClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPython3PackageClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Python3PackageClient, error) {
	cl, err := arm.NewClient(moduleName+".Python3PackageClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Python3PackageClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or Update the python 3 package identified by package name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - packageName - The name of python package.
//   - parameters - The create or update parameters for python package.
//   - options - Python3PackageClientCreateOrUpdateOptions contains the optional parameters for the Python3PackageClient.CreateOrUpdate
//     method.
func (client *Python3PackageClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageCreateParameters, options *Python3PackageClientCreateOrUpdateOptions) (Python3PackageClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, packageName, parameters, options)
	if err != nil {
		return Python3PackageClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Python3PackageClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return Python3PackageClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *Python3PackageClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageCreateParameters, options *Python3PackageClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python3Packages/{packageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *Python3PackageClient) createOrUpdateHandleResponse(resp *http.Response) (Python3PackageClientCreateOrUpdateResponse, error) {
	result := Python3PackageClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Module); err != nil {
		return Python3PackageClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the python 3 package by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - packageName - The python package name.
//   - options - Python3PackageClientDeleteOptions contains the optional parameters for the Python3PackageClient.Delete method.
func (client *Python3PackageClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, options *Python3PackageClientDeleteOptions) (Python3PackageClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, packageName, options)
	if err != nil {
		return Python3PackageClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Python3PackageClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return Python3PackageClientDeleteResponse{}, err
	}
	return Python3PackageClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *Python3PackageClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, options *Python3PackageClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python3Packages/{packageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the python 3 package identified by package name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - packageName - The python package name.
//   - options - Python3PackageClientGetOptions contains the optional parameters for the Python3PackageClient.Get method.
func (client *Python3PackageClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, options *Python3PackageClientGetOptions) (Python3PackageClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, packageName, options)
	if err != nil {
		return Python3PackageClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Python3PackageClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return Python3PackageClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *Python3PackageClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, options *Python3PackageClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python3Packages/{packageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *Python3PackageClient) getHandleResponse(resp *http.Response) (Python3PackageClientGetResponse, error) {
	result := Python3PackageClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Module); err != nil {
		return Python3PackageClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of python 3 packages.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - Python3PackageClientListByAutomationAccountOptions contains the optional parameters for the Python3PackageClient.NewListByAutomationAccountPager
//     method.
func (client *Python3PackageClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *Python3PackageClientListByAutomationAccountOptions) *runtime.Pager[Python3PackageClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[Python3PackageClientListByAutomationAccountResponse]{
		More: func(page Python3PackageClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *Python3PackageClientListByAutomationAccountResponse) (Python3PackageClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return Python3PackageClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return Python3PackageClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return Python3PackageClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *Python3PackageClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *Python3PackageClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python3Packages"
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
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *Python3PackageClient) listByAutomationAccountHandleResponse(resp *http.Response) (Python3PackageClientListByAutomationAccountResponse, error) {
	result := Python3PackageClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModuleListResult); err != nil {
		return Python3PackageClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update the python 3 package identified by package name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - packageName - The name of python package.
//   - parameters - The update parameters for python package.
//   - options - Python3PackageClientUpdateOptions contains the optional parameters for the Python3PackageClient.Update method.
func (client *Python3PackageClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageUpdateParameters, options *Python3PackageClientUpdateOptions) (Python3PackageClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, packageName, parameters, options)
	if err != nil {
		return Python3PackageClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Python3PackageClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return Python3PackageClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *Python3PackageClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageUpdateParameters, options *Python3PackageClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python3Packages/{packageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *Python3PackageClient) updateHandleResponse(resp *http.Response) (Python3PackageClientUpdateResponse, error) {
	result := Python3PackageClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Module); err != nil {
		return Python3PackageClientUpdateResponse{}, err
	}
	return result, nil
}
