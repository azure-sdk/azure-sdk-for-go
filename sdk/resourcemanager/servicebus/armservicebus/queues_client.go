//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// QueuesClient contains the methods for the Queues group.
// Don't use this type directly, use NewQueuesClient() instead.
type QueuesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewQueuesClient creates a new instance of QueuesClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQueuesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QueuesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &QueuesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Service Bus queue. This operation is idempotent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - queueName - The queue name.
//   - parameters - Parameters supplied to create or update a queue resource.
//   - options - QueuesClientCreateOrUpdateOptions contains the optional parameters for the QueuesClient.CreateOrUpdate method.
func (client *QueuesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, parameters SBQueue, options *QueuesClientCreateOrUpdateOptions) (QueuesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "QueuesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, queueName, parameters, options)
	if err != nil {
		return QueuesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueuesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueuesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *QueuesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, parameters SBQueue, options *QueuesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *QueuesClient) createOrUpdateHandleResponse(resp *http.Response) (QueuesClientCreateOrUpdateResponse, error) {
	result := QueuesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBQueue); err != nil {
		return QueuesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateAuthorizationRule - Creates an authorization rule for a queue.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - queueName - The queue name.
//   - authorizationRuleName - The authorization rule name.
//   - parameters - The shared access authorization rule.
//   - options - QueuesClientCreateOrUpdateAuthorizationRuleOptions contains the optional parameters for the QueuesClient.CreateOrUpdateAuthorizationRule
//     method.
func (client *QueuesClient) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, parameters SBAuthorizationRule, options *QueuesClientCreateOrUpdateAuthorizationRuleOptions) (QueuesClientCreateOrUpdateAuthorizationRuleResponse, error) {
	var err error
	const operationName = "QueuesClient.CreateOrUpdateAuthorizationRule"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, queueName, authorizationRuleName, parameters, options)
	if err != nil {
		return QueuesClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueuesClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueuesClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	resp, err := client.createOrUpdateAuthorizationRuleHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateAuthorizationRuleCreateRequest creates the CreateOrUpdateAuthorizationRule request.
func (client *QueuesClient) createOrUpdateAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, parameters SBAuthorizationRule, options *QueuesClientCreateOrUpdateAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateAuthorizationRuleHandleResponse handles the CreateOrUpdateAuthorizationRule response.
func (client *QueuesClient) createOrUpdateAuthorizationRuleHandleResponse(resp *http.Response) (QueuesClientCreateOrUpdateAuthorizationRuleResponse, error) {
	result := QueuesClientCreateOrUpdateAuthorizationRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRule); err != nil {
		return QueuesClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a queue from the specified namespace in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - queueName - The queue name.
//   - options - QueuesClientDeleteOptions contains the optional parameters for the QueuesClient.Delete method.
func (client *QueuesClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, options *QueuesClientDeleteOptions) (QueuesClientDeleteResponse, error) {
	var err error
	const operationName = "QueuesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, queueName, options)
	if err != nil {
		return QueuesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueuesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueuesClientDeleteResponse{}, err
	}
	return QueuesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *QueuesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, options *QueuesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteAuthorizationRule - Deletes a queue authorization rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - queueName - The queue name.
//   - authorizationRuleName - The authorization rule name.
//   - options - QueuesClientDeleteAuthorizationRuleOptions contains the optional parameters for the QueuesClient.DeleteAuthorizationRule
//     method.
func (client *QueuesClient) DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesClientDeleteAuthorizationRuleOptions) (QueuesClientDeleteAuthorizationRuleResponse, error) {
	var err error
	const operationName = "QueuesClient.DeleteAuthorizationRule"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, queueName, authorizationRuleName, options)
	if err != nil {
		return QueuesClientDeleteAuthorizationRuleResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueuesClientDeleteAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueuesClientDeleteAuthorizationRuleResponse{}, err
	}
	return QueuesClientDeleteAuthorizationRuleResponse{}, nil
}

// deleteAuthorizationRuleCreateRequest creates the DeleteAuthorizationRule request.
func (client *QueuesClient) deleteAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesClientDeleteAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns a description for the specified queue.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - queueName - The queue name.
//   - options - QueuesClientGetOptions contains the optional parameters for the QueuesClient.Get method.
func (client *QueuesClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, options *QueuesClientGetOptions) (QueuesClientGetResponse, error) {
	var err error
	const operationName = "QueuesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, queueName, options)
	if err != nil {
		return QueuesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueuesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueuesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *QueuesClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, options *QueuesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QueuesClient) getHandleResponse(resp *http.Response) (QueuesClientGetResponse, error) {
	result := QueuesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBQueue); err != nil {
		return QueuesClientGetResponse{}, err
	}
	return result, nil
}

// GetAuthorizationRule - Gets an authorization rule for a queue by rule name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - queueName - The queue name.
//   - authorizationRuleName - The authorization rule name.
//   - options - QueuesClientGetAuthorizationRuleOptions contains the optional parameters for the QueuesClient.GetAuthorizationRule
//     method.
func (client *QueuesClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesClientGetAuthorizationRuleOptions) (QueuesClientGetAuthorizationRuleResponse, error) {
	var err error
	const operationName = "QueuesClient.GetAuthorizationRule"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, queueName, authorizationRuleName, options)
	if err != nil {
		return QueuesClientGetAuthorizationRuleResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueuesClientGetAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueuesClientGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.getAuthorizationRuleHandleResponse(httpResp)
	return resp, err
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *QueuesClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesClientGetAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *QueuesClient) getAuthorizationRuleHandleResponse(resp *http.Response) (QueuesClientGetAuthorizationRuleResponse, error) {
	result := QueuesClientGetAuthorizationRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRule); err != nil {
		return QueuesClientGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// NewListAuthorizationRulesPager - Gets all authorization rules for a queue.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - queueName - The queue name.
//   - options - QueuesClientListAuthorizationRulesOptions contains the optional parameters for the QueuesClient.NewListAuthorizationRulesPager
//     method.
func (client *QueuesClient) NewListAuthorizationRulesPager(resourceGroupName string, namespaceName string, queueName string, options *QueuesClientListAuthorizationRulesOptions) *runtime.Pager[QueuesClientListAuthorizationRulesResponse] {
	return runtime.NewPager(runtime.PagingHandler[QueuesClientListAuthorizationRulesResponse]{
		More: func(page QueuesClientListAuthorizationRulesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QueuesClientListAuthorizationRulesResponse) (QueuesClientListAuthorizationRulesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "QueuesClient.NewListAuthorizationRulesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, queueName, options)
			}, nil)
			if err != nil {
				return QueuesClientListAuthorizationRulesResponse{}, err
			}
			return client.listAuthorizationRulesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *QueuesClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, options *QueuesClientListAuthorizationRulesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *QueuesClient) listAuthorizationRulesHandleResponse(resp *http.Response) (QueuesClientListAuthorizationRulesResponse, error) {
	result := QueuesClientListAuthorizationRulesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBAuthorizationRuleListResult); err != nil {
		return QueuesClientListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// NewListByNamespacePager - Gets the queues within a namespace.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - options - QueuesClientListByNamespaceOptions contains the optional parameters for the QueuesClient.NewListByNamespacePager
//     method.
func (client *QueuesClient) NewListByNamespacePager(resourceGroupName string, namespaceName string, options *QueuesClientListByNamespaceOptions) *runtime.Pager[QueuesClientListByNamespaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[QueuesClientListByNamespaceResponse]{
		More: func(page QueuesClientListByNamespaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QueuesClientListByNamespaceResponse) (QueuesClientListByNamespaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "QueuesClient.NewListByNamespacePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
			}, nil)
			if err != nil {
				return QueuesClientListByNamespaceResponse{}, err
			}
			return client.listByNamespaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *QueuesClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *QueuesClientListByNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNamespaceHandleResponse handles the ListByNamespace response.
func (client *QueuesClient) listByNamespaceHandleResponse(resp *http.Response) (QueuesClientListByNamespaceResponse, error) {
	result := QueuesClientListByNamespaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SBQueueListResult); err != nil {
		return QueuesClientListByNamespaceResponse{}, err
	}
	return result, nil
}

// ListKeys - Primary and secondary connection strings to the queue.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - queueName - The queue name.
//   - authorizationRuleName - The authorization rule name.
//   - options - QueuesClientListKeysOptions contains the optional parameters for the QueuesClient.ListKeys method.
func (client *QueuesClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesClientListKeysOptions) (QueuesClientListKeysResponse, error) {
	var err error
	const operationName = "QueuesClient.ListKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, queueName, authorizationRuleName, options)
	if err != nil {
		return QueuesClientListKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueuesClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueuesClientListKeysResponse{}, err
	}
	resp, err := client.listKeysHandleResponse(httpResp)
	return resp, err
}

// listKeysCreateRequest creates the ListKeys request.
func (client *QueuesClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, options *QueuesClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}/ListKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *QueuesClient) listKeysHandleResponse(resp *http.Response) (QueuesClientListKeysResponse, error) {
	result := QueuesClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return QueuesClientListKeysResponse{}, err
	}
	return result, nil
}

// RegenerateKeys - Regenerates the primary or secondary connection strings to the queue.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - queueName - The queue name.
//   - authorizationRuleName - The authorization rule name.
//   - parameters - Parameters supplied to regenerate the authorization rule.
//   - options - QueuesClientRegenerateKeysOptions contains the optional parameters for the QueuesClient.RegenerateKeys method.
func (client *QueuesClient) RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *QueuesClientRegenerateKeysOptions) (QueuesClientRegenerateKeysResponse, error) {
	var err error
	const operationName = "QueuesClient.RegenerateKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, namespaceName, queueName, authorizationRuleName, parameters, options)
	if err != nil {
		return QueuesClientRegenerateKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueuesClientRegenerateKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueuesClientRegenerateKeysResponse{}, err
	}
	resp, err := client.regenerateKeysHandleResponse(httpResp)
	return resp, err
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *QueuesClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *QueuesClientRegenerateKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}/regenerateKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if queueName == "" {
		return nil, errors.New("parameter queueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queueName}", url.PathEscape(queueName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *QueuesClient) regenerateKeysHandleResponse(resp *http.Response) (QueuesClientRegenerateKeysResponse, error) {
	result := QueuesClientRegenerateKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return QueuesClientRegenerateKeysResponse{}, err
	}
	return result, nil
}
