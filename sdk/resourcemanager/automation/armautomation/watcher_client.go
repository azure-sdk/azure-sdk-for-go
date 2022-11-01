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

// WatcherClient contains the methods for the Watcher group.
// Don't use this type directly, use NewWatcherClient() instead.
type WatcherClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWatcherClient creates a new instance of WatcherClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWatcherClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WatcherClient, error) {
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
	client := &WatcherClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create the watcher identified by watcher name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// watcherName - The watcher name.
// parameters - The create or update parameters for watcher.
// options - WatcherClientCreateOrUpdateOptions contains the optional parameters for the WatcherClient.CreateOrUpdate method.
func (client *WatcherClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, parameters Watcher, options *WatcherClientCreateOrUpdateOptions) (WatcherClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, watcherName, parameters, options)
	if err != nil {
		return WatcherClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WatcherClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return WatcherClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WatcherClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, parameters Watcher, options *WatcherClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WatcherClient) createOrUpdateHandleResponse(resp *http.Response) (WatcherClientCreateOrUpdateResponse, error) {
	result := WatcherClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Watcher); err != nil {
		return WatcherClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the watcher by name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// watcherName - The watcher name.
// options - WatcherClientDeleteOptions contains the optional parameters for the WatcherClient.Delete method.
func (client *WatcherClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, options *WatcherClientDeleteOptions) (WatcherClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, watcherName, options)
	if err != nil {
		return WatcherClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WatcherClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WatcherClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WatcherClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WatcherClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, options *WatcherClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the watcher identified by watcher name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// watcherName - The watcher name.
// options - WatcherClientGetOptions contains the optional parameters for the WatcherClient.Get method.
func (client *WatcherClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, options *WatcherClientGetOptions) (WatcherClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, watcherName, options)
	if err != nil {
		return WatcherClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WatcherClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WatcherClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WatcherClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, options *WatcherClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WatcherClient) getHandleResponse(resp *http.Response) (WatcherClientGetResponse, error) {
	result := WatcherClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Watcher); err != nil {
		return WatcherClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of watchers.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - WatcherClientListByAutomationAccountOptions contains the optional parameters for the WatcherClient.ListByAutomationAccount
// method.
func (client *WatcherClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *WatcherClientListByAutomationAccountOptions) *runtime.Pager[WatcherClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[WatcherClientListByAutomationAccountResponse]{
		More: func(page WatcherClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WatcherClientListByAutomationAccountResponse) (WatcherClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WatcherClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WatcherClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WatcherClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *WatcherClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *WatcherClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers"
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
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *WatcherClient) listByAutomationAccountHandleResponse(resp *http.Response) (WatcherClientListByAutomationAccountResponse, error) {
	result := WatcherClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WatcherListResult); err != nil {
		return WatcherClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Start - Resume the watcher identified by watcher name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// watcherName - The watcher name.
// options - WatcherClientStartOptions contains the optional parameters for the WatcherClient.Start method.
func (client *WatcherClient) Start(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, options *WatcherClientStartOptions) (WatcherClientStartResponse, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, automationAccountName, watcherName, options)
	if err != nil {
		return WatcherClientStartResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WatcherClientStartResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WatcherClientStartResponse{}, runtime.NewResponseError(resp)
	}
	return WatcherClientStartResponse{}, nil
}

// startCreateRequest creates the Start request.
func (client *WatcherClient) startCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, options *WatcherClientStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}/start"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Stop - Resume the watcher identified by watcher name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// watcherName - The watcher name.
// options - WatcherClientStopOptions contains the optional parameters for the WatcherClient.Stop method.
func (client *WatcherClient) Stop(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, options *WatcherClientStopOptions) (WatcherClientStopResponse, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, automationAccountName, watcherName, options)
	if err != nil {
		return WatcherClientStopResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WatcherClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WatcherClientStopResponse{}, runtime.NewResponseError(resp)
	}
	return WatcherClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *WatcherClient) stopCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, options *WatcherClientStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}/stop"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Update the watcher identified by watcher name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// watcherName - The watcher name.
// parameters - The update parameters for watcher.
// options - WatcherClientUpdateOptions contains the optional parameters for the WatcherClient.Update method.
func (client *WatcherClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, parameters WatcherUpdateParameters, options *WatcherClientUpdateOptions) (WatcherClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, watcherName, parameters, options)
	if err != nil {
		return WatcherClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WatcherClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WatcherClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *WatcherClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, watcherName string, parameters WatcherUpdateParameters, options *WatcherClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/watchers/{watcherName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if watcherName == "" {
		return nil, errors.New("parameter watcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{watcherName}", url.PathEscape(watcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *WatcherClient) updateHandleResponse(resp *http.Response) (WatcherClientUpdateResponse, error) {
	result := WatcherClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Watcher); err != nil {
		return WatcherClientUpdateResponse{}, err
	}
	return result, nil
}
