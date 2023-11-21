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

// WebhookClient contains the methods for the Webhook group.
// Don't use this type directly, use NewWebhookClient() instead.
type WebhookClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWebhookClient creates a new instance of WebhookClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWebhookClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebhookClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WebhookClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create the webhook identified by webhook name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-10-31
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - webhookName - The webhook name.
//   - parameters - The create or update parameters for webhook.
//   - options - WebhookClientCreateOrUpdateOptions contains the optional parameters for the WebhookClient.CreateOrUpdate method.
func (client *WebhookClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string, parameters WebhookCreateOrUpdateParameters, options *WebhookClientCreateOrUpdateOptions) (WebhookClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "WebhookClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, webhookName, parameters, options)
	if err != nil {
		return WebhookClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebhookClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WebhookClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WebhookClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string, parameters WebhookCreateOrUpdateParameters, options *WebhookClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks/{webhookName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-10-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WebhookClient) createOrUpdateHandleResponse(resp *http.Response) (WebhookClientCreateOrUpdateResponse, error) {
	result := WebhookClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Webhook); err != nil {
		return WebhookClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the webhook by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-10-31
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - webhookName - The webhook name.
//   - options - WebhookClientDeleteOptions contains the optional parameters for the WebhookClient.Delete method.
func (client *WebhookClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string, options *WebhookClientDeleteOptions) (WebhookClientDeleteResponse, error) {
	var err error
	const operationName = "WebhookClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, webhookName, options)
	if err != nil {
		return WebhookClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebhookClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebhookClientDeleteResponse{}, err
	}
	return WebhookClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WebhookClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string, options *WebhookClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks/{webhookName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-10-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GenerateURI - Generates a Uri for use in creating a webhook.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-10-31
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - WebhookClientGenerateURIOptions contains the optional parameters for the WebhookClient.GenerateURI method.
func (client *WebhookClient) GenerateURI(ctx context.Context, resourceGroupName string, automationAccountName string, options *WebhookClientGenerateURIOptions) (WebhookClientGenerateURIResponse, error) {
	var err error
	const operationName = "WebhookClient.GenerateURI"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.generateURICreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return WebhookClientGenerateURIResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebhookClientGenerateURIResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebhookClientGenerateURIResponse{}, err
	}
	resp, err := client.generateURIHandleResponse(httpResp)
	return resp, err
}

// generateURICreateRequest creates the GenerateURI request.
func (client *WebhookClient) generateURICreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *WebhookClientGenerateURIOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks/generateUri"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-10-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generateURIHandleResponse handles the GenerateURI response.
func (client *WebhookClient) generateURIHandleResponse(resp *http.Response) (WebhookClientGenerateURIResponse, error) {
	result := WebhookClientGenerateURIResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return WebhookClientGenerateURIResponse{}, err
	}
	return result, nil
}

// Get - Retrieve the webhook identified by webhook name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-10-31
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - webhookName - The webhook name.
//   - options - WebhookClientGetOptions contains the optional parameters for the WebhookClient.Get method.
func (client *WebhookClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string, options *WebhookClientGetOptions) (WebhookClientGetResponse, error) {
	var err error
	const operationName = "WebhookClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, webhookName, options)
	if err != nil {
		return WebhookClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebhookClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebhookClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WebhookClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string, options *WebhookClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks/{webhookName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-10-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebhookClient) getHandleResponse(resp *http.Response) (WebhookClientGetResponse, error) {
	result := WebhookClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Webhook); err != nil {
		return WebhookClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of webhooks.
//
// Generated from API version 2015-10-31
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - WebhookClientListByAutomationAccountOptions contains the optional parameters for the WebhookClient.NewListByAutomationAccountPager
//     method.
func (client *WebhookClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *WebhookClientListByAutomationAccountOptions) *runtime.Pager[WebhookClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebhookClientListByAutomationAccountResponse]{
		More: func(page WebhookClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebhookClientListByAutomationAccountResponse) (WebhookClientListByAutomationAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WebhookClient.NewListByAutomationAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			}, nil)
			if err != nil {
				return WebhookClientListByAutomationAccountResponse{}, err
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *WebhookClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *WebhookClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2015-10-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *WebhookClient) listByAutomationAccountHandleResponse(resp *http.Response) (WebhookClientListByAutomationAccountResponse, error) {
	result := WebhookClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebhookListResult); err != nil {
		return WebhookClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update the webhook identified by webhook name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-10-31
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - webhookName - The webhook name.
//   - parameters - The update parameters for webhook.
//   - options - WebhookClientUpdateOptions contains the optional parameters for the WebhookClient.Update method.
func (client *WebhookClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string, parameters WebhookUpdateParameters, options *WebhookClientUpdateOptions) (WebhookClientUpdateResponse, error) {
	var err error
	const operationName = "WebhookClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, webhookName, parameters, options)
	if err != nil {
		return WebhookClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebhookClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebhookClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *WebhookClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, webhookName string, parameters WebhookUpdateParameters, options *WebhookClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks/{webhookName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if webhookName == "" {
		return nil, errors.New("parameter webhookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webhookName}", url.PathEscape(webhookName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-10-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *WebhookClient) updateHandleResponse(resp *http.Response) (WebhookClientUpdateResponse, error) {
	result := WebhookClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Webhook); err != nil {
		return WebhookClientUpdateResponse{}, err
	}
	return result, nil
}
