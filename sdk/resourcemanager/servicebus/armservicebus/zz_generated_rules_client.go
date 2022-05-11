//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

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
	"strconv"
	"strings"
)

// RulesClient contains the methods for the Rules group.
// Don't use this type directly, use NewRulesClient() instead.
type RulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRulesClient creates a new instance of RulesClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RulesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new rule and updates an existing rule
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// topicName - The topic name.
// subscriptionName - The subscription name.
// ruleName - The rule name.
// parameters - Parameters supplied to create a rule.
// options - RulesClientCreateOrUpdateOptions contains the optional parameters for the RulesClient.CreateOrUpdate method.
func (client *RulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string, parameters Rule, options *RulesClientCreateOrUpdateOptions) (RulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, topicName, subscriptionName, ruleName, parameters, options)
	if err != nil {
		return RulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string, parameters Rule, options *RulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if subscriptionName == "" {
		return nil, errors.New("parameter subscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionName}", url.PathEscape(subscriptionName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RulesClient) createOrUpdateHandleResponse(resp *http.Response) (RulesClientCreateOrUpdateResponse, error) {
	result := RulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Rule); err != nil {
		return RulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// topicName - The topic name.
// subscriptionName - The subscription name.
// ruleName - The rule name.
// options - RulesClientDeleteOptions contains the optional parameters for the RulesClient.Delete method.
func (client *RulesClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string, options *RulesClientDeleteOptions) (RulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, topicName, subscriptionName, ruleName, options)
	if err != nil {
		return RulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return RulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string, options *RulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if subscriptionName == "" {
		return nil, errors.New("parameter subscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionName}", url.PathEscape(subscriptionName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves the description for the specified rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// topicName - The topic name.
// subscriptionName - The subscription name.
// ruleName - The rule name.
// options - RulesClientGetOptions contains the optional parameters for the RulesClient.Get method.
func (client *RulesClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string, options *RulesClientGetOptions) (RulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, topicName, subscriptionName, ruleName, options)
	if err != nil {
		return RulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string, options *RulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if subscriptionName == "" {
		return nil, errors.New("parameter subscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionName}", url.PathEscape(subscriptionName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RulesClient) getHandleResponse(resp *http.Response) (RulesClientGetResponse, error) {
	result := RulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Rule); err != nil {
		return RulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionsPager - List all the rules within given topic-subscription
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// namespaceName - The namespace name
// topicName - The topic name.
// subscriptionName - The subscription name.
// options - RulesClientListBySubscriptionsOptions contains the optional parameters for the RulesClient.ListBySubscriptions
// method.
func (client *RulesClient) NewListBySubscriptionsPager(resourceGroupName string, namespaceName string, topicName string, subscriptionName string, options *RulesClientListBySubscriptionsOptions) *runtime.Pager[RulesClientListBySubscriptionsResponse] {
	return runtime.NewPager(runtime.PageProcessor[RulesClientListBySubscriptionsResponse]{
		More: func(page RulesClientListBySubscriptionsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RulesClientListBySubscriptionsResponse) (RulesClientListBySubscriptionsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionsCreateRequest(ctx, resourceGroupName, namespaceName, topicName, subscriptionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RulesClientListBySubscriptionsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RulesClientListBySubscriptionsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RulesClientListBySubscriptionsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionsHandleResponse(resp)
		},
	})
}

// listBySubscriptionsCreateRequest creates the ListBySubscriptions request.
func (client *RulesClient) listBySubscriptionsCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, options *RulesClientListBySubscriptionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if subscriptionName == "" {
		return nil, errors.New("parameter subscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionName}", url.PathEscape(subscriptionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionsHandleResponse handles the ListBySubscriptions response.
func (client *RulesClient) listBySubscriptionsHandleResponse(resp *http.Response) (RulesClientListBySubscriptionsResponse, error) {
	result := RulesClientListBySubscriptionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleListResult); err != nil {
		return RulesClientListBySubscriptionsResponse{}, err
	}
	return result, nil
}
