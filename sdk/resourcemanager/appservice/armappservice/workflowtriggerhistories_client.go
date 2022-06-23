//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice

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

// WorkflowTriggerHistoriesClient contains the methods for the WorkflowTriggerHistories group.
// Don't use this type directly, use NewWorkflowTriggerHistoriesClient() instead.
type WorkflowTriggerHistoriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowTriggerHistoriesClient creates a new instance of WorkflowTriggerHistoriesClient with the specified values.
// subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkflowTriggerHistoriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkflowTriggerHistoriesClient, error) {
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
	client := &WorkflowTriggerHistoriesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a workflow trigger history.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Site name.
// workflowName - The workflow name.
// triggerName - The workflow trigger name.
// historyName - The workflow trigger history name. Corresponds to the run name for triggers that resulted in a run.
// options - WorkflowTriggerHistoriesClientGetOptions contains the optional parameters for the WorkflowTriggerHistoriesClient.Get
// method.
func (client *WorkflowTriggerHistoriesClient) Get(ctx context.Context, resourceGroupName string, name string, workflowName string, triggerName string, historyName string, options *WorkflowTriggerHistoriesClientGetOptions) (WorkflowTriggerHistoriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, workflowName, triggerName, historyName, options)
	if err != nil {
		return WorkflowTriggerHistoriesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowTriggerHistoriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowTriggerHistoriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowTriggerHistoriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, triggerName string, historyName string, options *WorkflowTriggerHistoriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/triggers/{triggerName}/histories/{historyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if historyName == "" {
		return nil, errors.New("parameter historyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyName}", url.PathEscape(historyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowTriggerHistoriesClient) getHandleResponse(resp *http.Response) (WorkflowTriggerHistoriesClientGetResponse, error) {
	result := WorkflowTriggerHistoriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerHistory); err != nil {
		return WorkflowTriggerHistoriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of workflow trigger histories.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Site name.
// workflowName - The workflow name.
// triggerName - The workflow trigger name.
// options - WorkflowTriggerHistoriesClientListOptions contains the optional parameters for the WorkflowTriggerHistoriesClient.List
// method.
func (client *WorkflowTriggerHistoriesClient) NewListPager(resourceGroupName string, name string, workflowName string, triggerName string, options *WorkflowTriggerHistoriesClientListOptions) *runtime.Pager[WorkflowTriggerHistoriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkflowTriggerHistoriesClientListResponse]{
		More: func(page WorkflowTriggerHistoriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkflowTriggerHistoriesClientListResponse) (WorkflowTriggerHistoriesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, name, workflowName, triggerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkflowTriggerHistoriesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkflowTriggerHistoriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkflowTriggerHistoriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkflowTriggerHistoriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, triggerName string, options *WorkflowTriggerHistoriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/triggers/{triggerName}/histories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowTriggerHistoriesClient) listHandleResponse(resp *http.Response) (WorkflowTriggerHistoriesClientListResponse, error) {
	result := WorkflowTriggerHistoriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerHistoryListResult); err != nil {
		return WorkflowTriggerHistoriesClientListResponse{}, err
	}
	return result, nil
}

// BeginResubmit - Resubmits a workflow run based on the trigger history.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Site name.
// workflowName - The workflow name.
// triggerName - The workflow trigger name.
// historyName - The workflow trigger history name. Corresponds to the run name for triggers that resulted in a run.
// options - WorkflowTriggerHistoriesClientBeginResubmitOptions contains the optional parameters for the WorkflowTriggerHistoriesClient.BeginResubmit
// method.
func (client *WorkflowTriggerHistoriesClient) BeginResubmit(ctx context.Context, resourceGroupName string, name string, workflowName string, triggerName string, historyName string, options *WorkflowTriggerHistoriesClientBeginResubmitOptions) (*runtime.Poller[WorkflowTriggerHistoriesClientResubmitResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resubmit(ctx, resourceGroupName, name, workflowName, triggerName, historyName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[WorkflowTriggerHistoriesClientResubmitResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[WorkflowTriggerHistoriesClientResubmitResponse](options.ResumeToken, client.pl, nil)
	}
}

// Resubmit - Resubmits a workflow run based on the trigger history.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
func (client *WorkflowTriggerHistoriesClient) resubmit(ctx context.Context, resourceGroupName string, name string, workflowName string, triggerName string, historyName string, options *WorkflowTriggerHistoriesClientBeginResubmitOptions) (*http.Response, error) {
	req, err := client.resubmitCreateRequest(ctx, resourceGroupName, name, workflowName, triggerName, historyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// resubmitCreateRequest creates the Resubmit request.
func (client *WorkflowTriggerHistoriesClient) resubmitCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, triggerName string, historyName string, options *WorkflowTriggerHistoriesClientBeginResubmitOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/triggers/{triggerName}/histories/{historyName}/resubmit"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if historyName == "" {
		return nil, errors.New("parameter historyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyName}", url.PathEscape(historyName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
