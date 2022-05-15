//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// AlertRulesClient contains the methods for the AlertRules group.
// Don't use this type directly, use NewAlertRulesClient() instead.
type AlertRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAlertRulesClient creates a new instance of AlertRulesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAlertRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertRulesClient, error) {
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
	client := &AlertRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// ruleID - Alert rule ID
// alertRule - The alert rule
// options - AlertRulesClientCreateOrUpdateOptions contains the optional parameters for the AlertRulesClient.CreateOrUpdate
// method.
func (client *AlertRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, alertRule AlertRuleClassification, options *AlertRulesClientCreateOrUpdateOptions) (AlertRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, ruleID, alertRule, options)
	if err != nil {
		return AlertRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AlertRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AlertRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, alertRule AlertRuleClassification, options *AlertRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, alertRule)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AlertRulesClient) createOrUpdateHandleResponse(resp *http.Response) (AlertRulesClientCreateOrUpdateResponse, error) {
	result := AlertRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return AlertRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// ruleID - Alert rule ID
// options - AlertRulesClientDeleteOptions contains the optional parameters for the AlertRulesClient.Delete method.
func (client *AlertRulesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, options *AlertRulesClientDeleteOptions) (AlertRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, ruleID, options)
	if err != nil {
		return AlertRulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AlertRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AlertRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AlertRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, options *AlertRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// ruleID - Alert rule ID
// options - AlertRulesClientGetOptions contains the optional parameters for the AlertRulesClient.Get method.
func (client *AlertRulesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, options *AlertRulesClientGetOptions) (AlertRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, ruleID, options)
	if err != nil {
		return AlertRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AlertRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, options *AlertRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertRulesClient) getHandleResponse(resp *http.Response) (AlertRulesClientGetResponse, error) {
	result := AlertRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return AlertRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all alert rules.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - AlertRulesClientListOptions contains the optional parameters for the AlertRulesClient.List method.
func (client *AlertRulesClient) NewListPager(resourceGroupName string, workspaceName string, options *AlertRulesClientListOptions) *runtime.Pager[AlertRulesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AlertRulesClientListResponse]{
		More: func(page AlertRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertRulesClientListResponse) (AlertRulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AlertRulesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AlertRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AlertRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *AlertRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AlertRulesClient) listHandleResponse(resp *http.Response) (AlertRulesClientListResponse, error) {
	result := AlertRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRulesList); err != nil {
		return AlertRulesClientListResponse{}, err
	}
	return result, nil
}
