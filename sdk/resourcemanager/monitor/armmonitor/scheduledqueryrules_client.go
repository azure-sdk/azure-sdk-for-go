//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// ScheduledQueryRulesClient contains the methods for the ScheduledQueryRules group.
// Don't use this type directly, use NewScheduledQueryRulesClient() instead.
type ScheduledQueryRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewScheduledQueryRulesClient creates a new instance of ScheduledQueryRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScheduledQueryRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ScheduledQueryRulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScheduledQueryRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a scheduled query rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleName - The name of the rule.
//   - parameters - The parameters of the rule to create or update.
//   - options - ScheduledQueryRulesClientCreateOrUpdateOptions contains the optional parameters for the ScheduledQueryRulesClient.CreateOrUpdate
//     method.
func (client *ScheduledQueryRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters ScheduledQueryRuleResource, options *ScheduledQueryRulesClientCreateOrUpdateOptions) (ScheduledQueryRulesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ScheduledQueryRulesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return ScheduledQueryRulesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledQueryRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledQueryRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ScheduledQueryRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters ScheduledQueryRuleResource, options *ScheduledQueryRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/scheduledQueryRules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ScheduledQueryRulesClient) createOrUpdateHandleResponse(resp *http.Response) (ScheduledQueryRulesClientCreateOrUpdateResponse, error) {
	result := ScheduledQueryRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduledQueryRuleResource); err != nil {
		return ScheduledQueryRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a scheduled query rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleName - The name of the rule.
//   - options - ScheduledQueryRulesClientDeleteOptions contains the optional parameters for the ScheduledQueryRulesClient.Delete
//     method.
func (client *ScheduledQueryRulesClient) Delete(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesClientDeleteOptions) (ScheduledQueryRulesClientDeleteResponse, error) {
	var err error
	const operationName = "ScheduledQueryRulesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return ScheduledQueryRulesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledQueryRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledQueryRulesClientDeleteResponse{}, err
	}
	return ScheduledQueryRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ScheduledQueryRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/scheduledQueryRules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve an scheduled query rule definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleName - The name of the rule.
//   - options - ScheduledQueryRulesClientGetOptions contains the optional parameters for the ScheduledQueryRulesClient.Get method.
func (client *ScheduledQueryRulesClient) Get(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesClientGetOptions) (ScheduledQueryRulesClientGetResponse, error) {
	var err error
	const operationName = "ScheduledQueryRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return ScheduledQueryRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledQueryRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledQueryRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ScheduledQueryRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *ScheduledQueryRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/scheduledQueryRules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ScheduledQueryRulesClient) getHandleResponse(resp *http.Response) (ScheduledQueryRulesClientGetResponse, error) {
	result := ScheduledQueryRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduledQueryRuleResource); err != nil {
		return ScheduledQueryRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Retrieve scheduled query rule definitions in a resource group.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ScheduledQueryRulesClientListByResourceGroupOptions contains the optional parameters for the ScheduledQueryRulesClient.NewListByResourceGroupPager
//     method.
func (client *ScheduledQueryRulesClient) NewListByResourceGroupPager(resourceGroupName string, options *ScheduledQueryRulesClientListByResourceGroupOptions) *runtime.Pager[ScheduledQueryRulesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScheduledQueryRulesClientListByResourceGroupResponse]{
		More: func(page ScheduledQueryRulesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScheduledQueryRulesClientListByResourceGroupResponse) (ScheduledQueryRulesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScheduledQueryRulesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ScheduledQueryRulesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ScheduledQueryRulesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ScheduledQueryRulesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/scheduledQueryRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ScheduledQueryRulesClient) listByResourceGroupHandleResponse(resp *http.Response) (ScheduledQueryRulesClientListByResourceGroupResponse, error) {
	result := ScheduledQueryRulesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduledQueryRuleResourceCollection); err != nil {
		return ScheduledQueryRulesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Retrieve a scheduled query rule definitions in a subscription.
//
// Generated from API version 2023-12-01
//   - options - ScheduledQueryRulesClientListBySubscriptionOptions contains the optional parameters for the ScheduledQueryRulesClient.NewListBySubscriptionPager
//     method.
func (client *ScheduledQueryRulesClient) NewListBySubscriptionPager(options *ScheduledQueryRulesClientListBySubscriptionOptions) *runtime.Pager[ScheduledQueryRulesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScheduledQueryRulesClientListBySubscriptionResponse]{
		More: func(page ScheduledQueryRulesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScheduledQueryRulesClientListBySubscriptionResponse) (ScheduledQueryRulesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScheduledQueryRulesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ScheduledQueryRulesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ScheduledQueryRulesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ScheduledQueryRulesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/scheduledQueryRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ScheduledQueryRulesClient) listBySubscriptionHandleResponse(resp *http.Response) (ScheduledQueryRulesClientListBySubscriptionResponse, error) {
	result := ScheduledQueryRulesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduledQueryRuleResourceCollection); err != nil {
		return ScheduledQueryRulesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a scheduled query rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleName - The name of the rule.
//   - parameters - The parameters of the rule to update.
//   - options - ScheduledQueryRulesClientUpdateOptions contains the optional parameters for the ScheduledQueryRulesClient.Update
//     method.
func (client *ScheduledQueryRulesClient) Update(ctx context.Context, resourceGroupName string, ruleName string, parameters ScheduledQueryRuleResourcePatch, options *ScheduledQueryRulesClientUpdateOptions) (ScheduledQueryRulesClientUpdateResponse, error) {
	var err error
	const operationName = "ScheduledQueryRulesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return ScheduledQueryRulesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScheduledQueryRulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScheduledQueryRulesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ScheduledQueryRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters ScheduledQueryRuleResourcePatch, options *ScheduledQueryRulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/scheduledQueryRules/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ScheduledQueryRulesClient) updateHandleResponse(resp *http.Response) (ScheduledQueryRulesClientUpdateResponse, error) {
	result := ScheduledQueryRulesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduledQueryRuleResource); err != nil {
		return ScheduledQueryRulesClientUpdateResponse{}, err
	}
	return result, nil
}
