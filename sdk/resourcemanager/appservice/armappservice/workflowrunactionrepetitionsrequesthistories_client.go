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
	"strings"
)

// WorkflowRunActionRepetitionsRequestHistoriesClient contains the methods for the WorkflowRunActionRepetitionsRequestHistories group.
// Don't use this type directly, use NewWorkflowRunActionRepetitionsRequestHistoriesClient() instead.
type WorkflowRunActionRepetitionsRequestHistoriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowRunActionRepetitionsRequestHistoriesClient creates a new instance of WorkflowRunActionRepetitionsRequestHistoriesClient with the specified values.
//   - subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkflowRunActionRepetitionsRequestHistoriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkflowRunActionRepetitionsRequestHistoriesClient, error) {
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
	client := &WorkflowRunActionRepetitionsRequestHistoriesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a workflow run repetition request history.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Site name.
//   - workflowName - The workflow name.
//   - runName - The workflow run name.
//   - actionName - The workflow action name.
//   - repetitionName - The workflow repetition.
//   - requestHistoryName - The request history name.
//   - options - WorkflowRunActionRepetitionsRequestHistoriesClientGetOptions contains the optional parameters for the WorkflowRunActionRepetitionsRequestHistoriesClient.Get
//     method.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) Get(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string, requestHistoryName string, options *WorkflowRunActionRepetitionsRequestHistoriesClientGetOptions) (WorkflowRunActionRepetitionsRequestHistoriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, workflowName, runName, actionName, repetitionName, requestHistoryName, options)
	if err != nil {
		return WorkflowRunActionRepetitionsRequestHistoriesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunActionRepetitionsRequestHistoriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunActionRepetitionsRequestHistoriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string, requestHistoryName string, options *WorkflowRunActionRepetitionsRequestHistoriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/requestHistories/{requestHistoryName}"
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
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	if repetitionName == "" {
		return nil, errors.New("parameter repetitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{repetitionName}", url.PathEscape(repetitionName))
	if requestHistoryName == "" {
		return nil, errors.New("parameter requestHistoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{requestHistoryName}", url.PathEscape(requestHistoryName))
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
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) getHandleResponse(resp *http.Response) (WorkflowRunActionRepetitionsRequestHistoriesClientGetResponse, error) {
	result := WorkflowRunActionRepetitionsRequestHistoriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RequestHistory); err != nil {
		return WorkflowRunActionRepetitionsRequestHistoriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List a workflow run repetition request history.
//
// Generated from API version 2022-03-01
//   - resourceGroupName - Name of the resource group to which the resource belongs.
//   - name - Site name.
//   - workflowName - The workflow name.
//   - runName - The workflow run name.
//   - actionName - The workflow action name.
//   - repetitionName - The workflow repetition.
//   - options - WorkflowRunActionRepetitionsRequestHistoriesClientListOptions contains the optional parameters for the WorkflowRunActionRepetitionsRequestHistoriesClient.NewListPager
//     method.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) NewListPager(resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string, options *WorkflowRunActionRepetitionsRequestHistoriesClientListOptions) *runtime.Pager[WorkflowRunActionRepetitionsRequestHistoriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkflowRunActionRepetitionsRequestHistoriesClientListResponse]{
		More: func(page WorkflowRunActionRepetitionsRequestHistoriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkflowRunActionRepetitionsRequestHistoriesClientListResponse) (WorkflowRunActionRepetitionsRequestHistoriesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, name, workflowName, runName, actionName, repetitionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkflowRunActionRepetitionsRequestHistoriesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkflowRunActionRepetitionsRequestHistoriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkflowRunActionRepetitionsRequestHistoriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string, repetitionName string, options *WorkflowRunActionRepetitionsRequestHistoriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/requestHistories"
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
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	if repetitionName == "" {
		return nil, errors.New("parameter repetitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{repetitionName}", url.PathEscape(repetitionName))
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

// listHandleResponse handles the List response.
func (client *WorkflowRunActionRepetitionsRequestHistoriesClient) listHandleResponse(resp *http.Response) (WorkflowRunActionRepetitionsRequestHistoriesClientListResponse, error) {
	result := WorkflowRunActionRepetitionsRequestHistoriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RequestHistoryListResult); err != nil {
		return WorkflowRunActionRepetitionsRequestHistoriesClientListResponse{}, err
	}
	return result, nil
}
