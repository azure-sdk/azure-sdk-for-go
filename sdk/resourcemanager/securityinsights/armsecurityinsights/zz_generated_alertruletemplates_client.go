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

// AlertRuleTemplatesClient contains the methods for the AlertRuleTemplates group.
// Don't use this type directly, use NewAlertRuleTemplatesClient() instead.
type AlertRuleTemplatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAlertRuleTemplatesClient creates a new instance of AlertRuleTemplatesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAlertRuleTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertRuleTemplatesClient, error) {
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
	client := &AlertRuleTemplatesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the alert rule template.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// alertRuleTemplateID - Alert rule template ID
// options - AlertRuleTemplatesClientGetOptions contains the optional parameters for the AlertRuleTemplatesClient.Get method.
func (client *AlertRuleTemplatesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, alertRuleTemplateID string, options *AlertRuleTemplatesClientGetOptions) (AlertRuleTemplatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, alertRuleTemplateID, options)
	if err != nil {
		return AlertRuleTemplatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertRuleTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertRuleTemplatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AlertRuleTemplatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, alertRuleTemplateID string, options *AlertRuleTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRuleTemplates/{alertRuleTemplateId}"
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
	if alertRuleTemplateID == "" {
		return nil, errors.New("parameter alertRuleTemplateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertRuleTemplateId}", url.PathEscape(alertRuleTemplateID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertRuleTemplatesClient) getHandleResponse(resp *http.Response) (AlertRuleTemplatesClientGetResponse, error) {
	result := AlertRuleTemplatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return AlertRuleTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all alert rule templates.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - AlertRuleTemplatesClientListOptions contains the optional parameters for the AlertRuleTemplatesClient.List method.
func (client *AlertRuleTemplatesClient) NewListPager(resourceGroupName string, workspaceName string, options *AlertRuleTemplatesClientListOptions) *runtime.Pager[AlertRuleTemplatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AlertRuleTemplatesClientListResponse]{
		More: func(page AlertRuleTemplatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertRuleTemplatesClientListResponse) (AlertRuleTemplatesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AlertRuleTemplatesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AlertRuleTemplatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertRuleTemplatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AlertRuleTemplatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *AlertRuleTemplatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRuleTemplates"
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
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AlertRuleTemplatesClient) listHandleResponse(resp *http.Response) (AlertRuleTemplatesClientListResponse, error) {
	result := AlertRuleTemplatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleTemplatesList); err != nil {
		return AlertRuleTemplatesClientListResponse{}, err
	}
	return result, nil
}
