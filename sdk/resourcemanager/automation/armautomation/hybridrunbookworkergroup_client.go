//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation

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

// HybridRunbookWorkerGroupClient contains the methods for the HybridRunbookWorkerGroup group.
// Don't use this type directly, use NewHybridRunbookWorkerGroupClient() instead.
type HybridRunbookWorkerGroupClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHybridRunbookWorkerGroupClient creates a new instance of HybridRunbookWorkerGroupClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHybridRunbookWorkerGroupClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HybridRunbookWorkerGroupClient, error) {
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
	client := &HybridRunbookWorkerGroupClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create a hybrid runbook worker group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-22
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// hybridRunbookWorkerGroupName - The hybrid runbook worker group name
// hybridRunbookWorkerGroupCreationParameters - The create or update parameters for hybrid runbook worker group.
// options - HybridRunbookWorkerGroupClientCreateOptions contains the optional parameters for the HybridRunbookWorkerGroupClient.Create
// method.
func (client *HybridRunbookWorkerGroupClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerGroupCreationParameters HybridRunbookWorkerGroupCreateOrUpdateParameters, options *HybridRunbookWorkerGroupClientCreateOptions) (HybridRunbookWorkerGroupClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerGroupCreationParameters, options)
	if err != nil {
		return HybridRunbookWorkerGroupClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridRunbookWorkerGroupClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridRunbookWorkerGroupClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *HybridRunbookWorkerGroupClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerGroupCreationParameters HybridRunbookWorkerGroupCreateOrUpdateParameters, options *HybridRunbookWorkerGroupClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, hybridRunbookWorkerGroupCreationParameters)
}

// createHandleResponse handles the Create response.
func (client *HybridRunbookWorkerGroupClient) createHandleResponse(resp *http.Response) (HybridRunbookWorkerGroupClientCreateResponse, error) {
	result := HybridRunbookWorkerGroupClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridRunbookWorkerGroup); err != nil {
		return HybridRunbookWorkerGroupClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a hybrid runbook worker group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-22
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// hybridRunbookWorkerGroupName - The hybrid runbook worker group name
// options - HybridRunbookWorkerGroupClientDeleteOptions contains the optional parameters for the HybridRunbookWorkerGroupClient.Delete
// method.
func (client *HybridRunbookWorkerGroupClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, options *HybridRunbookWorkerGroupClientDeleteOptions) (HybridRunbookWorkerGroupClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, options)
	if err != nil {
		return HybridRunbookWorkerGroupClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridRunbookWorkerGroupClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridRunbookWorkerGroupClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return HybridRunbookWorkerGroupClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HybridRunbookWorkerGroupClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, options *HybridRunbookWorkerGroupClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve a hybrid runbook worker group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-22
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// hybridRunbookWorkerGroupName - The hybrid runbook worker group name
// options - HybridRunbookWorkerGroupClientGetOptions contains the optional parameters for the HybridRunbookWorkerGroupClient.Get
// method.
func (client *HybridRunbookWorkerGroupClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, options *HybridRunbookWorkerGroupClientGetOptions) (HybridRunbookWorkerGroupClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, options)
	if err != nil {
		return HybridRunbookWorkerGroupClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridRunbookWorkerGroupClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridRunbookWorkerGroupClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HybridRunbookWorkerGroupClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, options *HybridRunbookWorkerGroupClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HybridRunbookWorkerGroupClient) getHandleResponse(resp *http.Response) (HybridRunbookWorkerGroupClientGetResponse, error) {
	result := HybridRunbookWorkerGroupClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridRunbookWorkerGroup); err != nil {
		return HybridRunbookWorkerGroupClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of hybrid runbook worker groups.
// Generated from API version 2022-02-22
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - HybridRunbookWorkerGroupClientListByAutomationAccountOptions contains the optional parameters for the HybridRunbookWorkerGroupClient.ListByAutomationAccount
// method.
func (client *HybridRunbookWorkerGroupClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *HybridRunbookWorkerGroupClientListByAutomationAccountOptions) *runtime.Pager[HybridRunbookWorkerGroupClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[HybridRunbookWorkerGroupClientListByAutomationAccountResponse]{
		More: func(page HybridRunbookWorkerGroupClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HybridRunbookWorkerGroupClientListByAutomationAccountResponse) (HybridRunbookWorkerGroupClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HybridRunbookWorkerGroupClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HybridRunbookWorkerGroupClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HybridRunbookWorkerGroupClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *HybridRunbookWorkerGroupClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *HybridRunbookWorkerGroupClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2022-02-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *HybridRunbookWorkerGroupClient) listByAutomationAccountHandleResponse(resp *http.Response) (HybridRunbookWorkerGroupClientListByAutomationAccountResponse, error) {
	result := HybridRunbookWorkerGroupClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridRunbookWorkerGroupsListResult); err != nil {
		return HybridRunbookWorkerGroupClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update a hybrid runbook worker group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-22
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// hybridRunbookWorkerGroupName - The hybrid runbook worker group name
// hybridRunbookWorkerGroupUpdationParameters - The hybrid runbook worker group
// options - HybridRunbookWorkerGroupClientUpdateOptions contains the optional parameters for the HybridRunbookWorkerGroupClient.Update
// method.
func (client *HybridRunbookWorkerGroupClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerGroupUpdationParameters HybridRunbookWorkerGroupCreateOrUpdateParameters, options *HybridRunbookWorkerGroupClientUpdateOptions) (HybridRunbookWorkerGroupClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerGroupUpdationParameters, options)
	if err != nil {
		return HybridRunbookWorkerGroupClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridRunbookWorkerGroupClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridRunbookWorkerGroupClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *HybridRunbookWorkerGroupClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerGroupUpdationParameters HybridRunbookWorkerGroupCreateOrUpdateParameters, options *HybridRunbookWorkerGroupClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, hybridRunbookWorkerGroupUpdationParameters)
}

// updateHandleResponse handles the Update response.
func (client *HybridRunbookWorkerGroupClient) updateHandleResponse(resp *http.Response) (HybridRunbookWorkerGroupClientUpdateResponse, error) {
	result := HybridRunbookWorkerGroupClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridRunbookWorkerGroup); err != nil {
		return HybridRunbookWorkerGroupClientUpdateResponse{}, err
	}
	return result, nil
}
