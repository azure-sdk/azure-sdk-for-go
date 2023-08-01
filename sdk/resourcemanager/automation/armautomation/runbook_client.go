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

// RunbookClient contains the methods for the Runbook group.
// Don't use this type directly, use NewRunbookClient() instead.
type RunbookClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRunbookClient creates a new instance of RunbookClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRunbookClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RunbookClient, error) {
	cl, err := arm.NewClient(moduleName+".RunbookClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RunbookClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create the runbook identified by runbook name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - runbookName - The runbook name.
//   - parameters - The create or update parameters for runbook. Provide either content link for a published runbook or draft,
//     not both.
//   - options - RunbookClientCreateOrUpdateOptions contains the optional parameters for the RunbookClient.CreateOrUpdate method.
func (client *RunbookClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters RunbookCreateOrUpdateParameters, options *RunbookClientCreateOrUpdateOptions) (RunbookClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, parameters, options)
	if err != nil {
		return RunbookClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RunbookClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return RunbookClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RunbookClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters RunbookCreateOrUpdateParameters, options *RunbookClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RunbookClient) createOrUpdateHandleResponse(resp *http.Response) (RunbookClientCreateOrUpdateResponse, error) {
	result := RunbookClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Runbook); err != nil {
		return RunbookClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the runbook by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - runbookName - The runbook name.
//   - options - RunbookClientDeleteOptions contains the optional parameters for the RunbookClient.Delete method.
func (client *RunbookClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookClientDeleteOptions) (RunbookClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return RunbookClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RunbookClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RunbookClientDeleteResponse{}, err
	}
	return RunbookClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RunbookClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the runbook identified by runbook name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - runbookName - The runbook name.
//   - options - RunbookClientGetOptions contains the optional parameters for the RunbookClient.Get method.
func (client *RunbookClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookClientGetOptions) (RunbookClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return RunbookClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RunbookClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RunbookClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RunbookClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RunbookClient) getHandleResponse(resp *http.Response) (RunbookClientGetResponse, error) {
	result := RunbookClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Runbook); err != nil {
		return RunbookClientGetResponse{}, err
	}
	return result, nil
}

// GetContent - Retrieve the content of runbook identified by runbook name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - runbookName - The runbook name.
//   - options - RunbookClientGetContentOptions contains the optional parameters for the RunbookClient.GetContent method.
func (client *RunbookClient) GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookClientGetContentOptions) (RunbookClientGetContentResponse, error) {
	var err error
	req, err := client.getContentCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return RunbookClientGetContentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RunbookClientGetContentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RunbookClientGetContentResponse{}, err
	}
	return RunbookClientGetContentResponse{}, nil
}

// getContentCreateRequest creates the GetContent request.
func (client *RunbookClient) getContentCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookClientGetContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/content"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"text/powershell"}
	return req, nil
}

// NewListByAutomationAccountPager - Retrieve a list of runbooks.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - RunbookClientListByAutomationAccountOptions contains the optional parameters for the RunbookClient.NewListByAutomationAccountPager
//     method.
func (client *RunbookClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *RunbookClientListByAutomationAccountOptions) *runtime.Pager[RunbookClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[RunbookClientListByAutomationAccountResponse]{
		More: func(page RunbookClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RunbookClientListByAutomationAccountResponse) (RunbookClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RunbookClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RunbookClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RunbookClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *RunbookClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *RunbookClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *RunbookClient) listByAutomationAccountHandleResponse(resp *http.Response) (RunbookClientListByAutomationAccountResponse, error) {
	result := RunbookClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunbookListResult); err != nil {
		return RunbookClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// BeginPublish - Publish runbook draft.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - runbookName - The parameters supplied to the publish runbook operation.
//   - options - RunbookClientBeginPublishOptions contains the optional parameters for the RunbookClient.BeginPublish method.
func (client *RunbookClient) BeginPublish(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookClientBeginPublishOptions) (*runtime.Poller[RunbookClientPublishResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.publish(ctx, resourceGroupName, automationAccountName, runbookName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[RunbookClientPublishResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RunbookClientPublishResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Publish - Publish runbook draft.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
func (client *RunbookClient) publish(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookClientBeginPublishOptions) (*http.Response, error) {
	var err error
	req, err := client.publishCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// publishCreateRequest creates the Publish request.
func (client *RunbookClient) publishCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookClientBeginPublishOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/publish"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Update the runbook identified by runbook name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - runbookName - The runbook name.
//   - parameters - The update parameters for runbook.
//   - options - RunbookClientUpdateOptions contains the optional parameters for the RunbookClient.Update method.
func (client *RunbookClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters RunbookUpdateParameters, options *RunbookClientUpdateOptions) (RunbookClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, parameters, options)
	if err != nil {
		return RunbookClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RunbookClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RunbookClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *RunbookClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters RunbookUpdateParameters, options *RunbookClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *RunbookClient) updateHandleResponse(resp *http.Response) (RunbookClientUpdateResponse, error) {
	result := RunbookClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Runbook); err != nil {
		return RunbookClientUpdateResponse{}, err
	}
	return result, nil
}
