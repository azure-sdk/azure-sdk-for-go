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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DscConfigurationClient contains the methods for the DscConfiguration group.
// Don't use this type directly, use NewDscConfigurationClient() instead.
type DscConfigurationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDscConfigurationClient creates a new instance of DscConfigurationClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDscConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DscConfigurationClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DscConfigurationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdateWithJSON - Create the configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - configurationName - The create or update parameters for configuration.
//   - parameters - The create or update parameters for configuration.
//   - options - DscConfigurationClientCreateOrUpdateWithJSONOptions contains the optional parameters for the DscConfigurationClient.CreateOrUpdateWithJSON
//     method.
func (client *DscConfigurationClient) CreateOrUpdateWithJSON(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters DscConfigurationCreateOrUpdateParameters, options *DscConfigurationClientCreateOrUpdateWithJSONOptions) (DscConfigurationClientCreateOrUpdateWithJSONResponse, error) {
	var err error
	const operationName = "DscConfigurationClient.CreateOrUpdateWithJSON"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateWithJSONCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, parameters, options)
	if err != nil {
		return DscConfigurationClientCreateOrUpdateWithJSONResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscConfigurationClientCreateOrUpdateWithJSONResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DscConfigurationClientCreateOrUpdateWithJSONResponse{}, err
	}
	resp, err := client.createOrUpdateWithJSONHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateWithJSONCreateRequest creates the CreateOrUpdateWithJSON request.
func (client *DscConfigurationClient) createOrUpdateWithJSONCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters DscConfigurationCreateOrUpdateParameters, options *DscConfigurationClientCreateOrUpdateWithJSONOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateWithJSONHandleResponse handles the CreateOrUpdateWithJSON response.
func (client *DscConfigurationClient) createOrUpdateWithJSONHandleResponse(resp *http.Response) (DscConfigurationClientCreateOrUpdateWithJSONResponse, error) {
	result := DscConfigurationClientCreateOrUpdateWithJSONResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationClientCreateOrUpdateWithJSONResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateWithText - Create the configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - configurationName - The create or update parameters for configuration.
//   - parameters - The create or update parameters for configuration.
//   - options - DscConfigurationClientCreateOrUpdateWithTextOptions contains the optional parameters for the DscConfigurationClient.CreateOrUpdateWithText
//     method.
func (client *DscConfigurationClient) CreateOrUpdateWithText(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters string, options *DscConfigurationClientCreateOrUpdateWithTextOptions) (DscConfigurationClientCreateOrUpdateWithTextResponse, error) {
	var err error
	const operationName = "DscConfigurationClient.CreateOrUpdateWithText"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateWithTextCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, parameters, options)
	if err != nil {
		return DscConfigurationClientCreateOrUpdateWithTextResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscConfigurationClientCreateOrUpdateWithTextResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DscConfigurationClientCreateOrUpdateWithTextResponse{}, err
	}
	resp, err := client.createOrUpdateWithTextHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateWithTextCreateRequest creates the CreateOrUpdateWithText request.
func (client *DscConfigurationClient) createOrUpdateWithTextCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters string, options *DscConfigurationClientCreateOrUpdateWithTextOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	body := streaming.NopCloser(strings.NewReader(parameters))
	if err := req.SetBody(body, "text/plain; charset=utf-8"); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateWithTextHandleResponse handles the CreateOrUpdateWithText response.
func (client *DscConfigurationClient) createOrUpdateWithTextHandleResponse(resp *http.Response) (DscConfigurationClientCreateOrUpdateWithTextResponse, error) {
	result := DscConfigurationClientCreateOrUpdateWithTextResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationClientCreateOrUpdateWithTextResponse{}, err
	}
	return result, nil
}

// Delete - Delete the dsc configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - configurationName - The configuration name.
//   - options - DscConfigurationClientDeleteOptions contains the optional parameters for the DscConfigurationClient.Delete method.
func (client *DscConfigurationClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientDeleteOptions) (DscConfigurationClientDeleteResponse, error) {
	var err error
	const operationName = "DscConfigurationClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscConfigurationClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DscConfigurationClientDeleteResponse{}, err
	}
	return DscConfigurationClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DscConfigurationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - configurationName - The configuration name.
//   - options - DscConfigurationClientGetOptions contains the optional parameters for the DscConfigurationClient.Get method.
func (client *DscConfigurationClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientGetOptions) (DscConfigurationClientGetResponse, error) {
	var err error
	const operationName = "DscConfigurationClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DscConfigurationClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DscConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DscConfigurationClient) getHandleResponse(resp *http.Response) (DscConfigurationClientGetResponse, error) {
	result := DscConfigurationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// GetContent - Retrieve the configuration script identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - configurationName - The configuration name.
//   - options - DscConfigurationClientGetContentOptions contains the optional parameters for the DscConfigurationClient.GetContent
//     method.
func (client *DscConfigurationClient) GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientGetContentOptions) (DscConfigurationClientGetContentResponse, error) {
	var err error
	const operationName = "DscConfigurationClient.GetContent"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getContentCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationClientGetContentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscConfigurationClientGetContentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DscConfigurationClientGetContentResponse{}, err
	}
	return DscConfigurationClientGetContentResponse{}, nil
}

// getContentCreateRequest creates the GetContent request.
func (client *DscConfigurationClient) getContentCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientGetContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}/content"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"text/powershell"}
	return req, nil
}

// NewListByAutomationAccountPager - Retrieve a list of configurations.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - DscConfigurationClientListByAutomationAccountOptions contains the optional parameters for the DscConfigurationClient.NewListByAutomationAccountPager
//     method.
func (client *DscConfigurationClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *DscConfigurationClientListByAutomationAccountOptions) *runtime.Pager[DscConfigurationClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[DscConfigurationClientListByAutomationAccountResponse]{
		More: func(page DscConfigurationClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DscConfigurationClientListByAutomationAccountResponse) (DscConfigurationClientListByAutomationAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DscConfigurationClient.NewListByAutomationAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			}, nil)
			if err != nil {
				return DscConfigurationClientListByAutomationAccountResponse{}, err
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *DscConfigurationClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *DscConfigurationClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations"
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
	if options != nil && options.Inlinecount != nil {
		reqQP.Set("$inlinecount", *options.Inlinecount)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *DscConfigurationClient) listByAutomationAccountHandleResponse(resp *http.Response) (DscConfigurationClientListByAutomationAccountResponse, error) {
	result := DscConfigurationClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfigurationListResult); err != nil {
		return DscConfigurationClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// UpdateWithJSON - Create the configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - configurationName - The create or update parameters for configuration.
//   - options - DscConfigurationClientUpdateWithJSONOptions contains the optional parameters for the DscConfigurationClient.UpdateWithJSON
//     method.
func (client *DscConfigurationClient) UpdateWithJSON(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientUpdateWithJSONOptions) (DscConfigurationClientUpdateWithJSONResponse, error) {
	var err error
	const operationName = "DscConfigurationClient.UpdateWithJSON"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateWithJSONCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationClientUpdateWithJSONResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscConfigurationClientUpdateWithJSONResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DscConfigurationClientUpdateWithJSONResponse{}, err
	}
	resp, err := client.updateWithJSONHandleResponse(httpResp)
	return resp, err
}

// updateWithJSONCreateRequest creates the UpdateWithJSON request.
func (client *DscConfigurationClient) updateWithJSONCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientUpdateWithJSONOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateWithJSONHandleResponse handles the UpdateWithJSON response.
func (client *DscConfigurationClient) updateWithJSONHandleResponse(resp *http.Response) (DscConfigurationClientUpdateWithJSONResponse, error) {
	result := DscConfigurationClientUpdateWithJSONResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationClientUpdateWithJSONResponse{}, err
	}
	return result, nil
}

// UpdateWithText - Create the configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - configurationName - The create or update parameters for configuration.
//   - options - DscConfigurationClientUpdateWithTextOptions contains the optional parameters for the DscConfigurationClient.UpdateWithText
//     method.
func (client *DscConfigurationClient) UpdateWithText(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientUpdateWithTextOptions) (DscConfigurationClientUpdateWithTextResponse, error) {
	var err error
	const operationName = "DscConfigurationClient.UpdateWithText"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateWithTextCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationClientUpdateWithTextResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscConfigurationClientUpdateWithTextResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DscConfigurationClientUpdateWithTextResponse{}, err
	}
	resp, err := client.updateWithTextHandleResponse(httpResp)
	return resp, err
}

// updateWithTextCreateRequest creates the UpdateWithText request.
func (client *DscConfigurationClient) updateWithTextCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientUpdateWithTextOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		body := streaming.NopCloser(strings.NewReader(*options.Parameters))
		if err := req.SetBody(body, "text/plain; charset=utf-8"); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateWithTextHandleResponse handles the UpdateWithText response.
func (client *DscConfigurationClient) updateWithTextHandleResponse(resp *http.Response) (DscConfigurationClientUpdateWithTextResponse, error) {
	result := DscConfigurationClientUpdateWithTextResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationClientUpdateWithTextResponse{}, err
	}
	return result, nil
}
