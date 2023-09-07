//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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
	"time"
)

// JobTargetExecutionsClient contains the methods for the JobTargetExecutions group.
// Don't use this type directly, use NewJobTargetExecutionsClient() instead.
type JobTargetExecutionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewJobTargetExecutionsClient creates a new instance of JobTargetExecutionsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewJobTargetExecutionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobTargetExecutionsClient, error) {
	cl, err := arm.NewClient(moduleName+".JobTargetExecutionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &JobTargetExecutionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a target execution.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - jobName - The name of the job to get.
//   - jobExecutionID - The unique id of the job execution
//   - stepName - The name of the step.
//   - targetID - The target id.
//   - options - JobTargetExecutionsClientGetOptions contains the optional parameters for the JobTargetExecutionsClient.Get method.
func (client *JobTargetExecutionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, stepName string, targetID string, options *JobTargetExecutionsClientGetOptions) (JobTargetExecutionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, stepName, targetID, options)
	if err != nil {
		return JobTargetExecutionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobTargetExecutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JobTargetExecutionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *JobTargetExecutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, stepName string, targetID string, options *JobTargetExecutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}/steps/{stepName}/targets/{targetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobExecutionId}", url.PathEscape(jobExecutionID))
	if stepName == "" {
		return nil, errors.New("parameter stepName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stepName}", url.PathEscape(stepName))
	urlPath = strings.ReplaceAll(urlPath, "{targetId}", url.PathEscape(targetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobTargetExecutionsClient) getHandleResponse(resp *http.Response) (JobTargetExecutionsClientGetResponse, error) {
	result := JobTargetExecutionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecution); err != nil {
		return JobTargetExecutionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByJobExecutionPager - Lists target executions for all steps of a job execution.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - jobName - The name of the job to get.
//   - jobExecutionID - The id of the job execution
//   - options - JobTargetExecutionsClientListByJobExecutionOptions contains the optional parameters for the JobTargetExecutionsClient.NewListByJobExecutionPager
//     method.
func (client *JobTargetExecutionsClient) NewListByJobExecutionPager(resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobTargetExecutionsClientListByJobExecutionOptions) *runtime.Pager[JobTargetExecutionsClientListByJobExecutionResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobTargetExecutionsClientListByJobExecutionResponse]{
		More: func(page JobTargetExecutionsClientListByJobExecutionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobTargetExecutionsClientListByJobExecutionResponse) (JobTargetExecutionsClientListByJobExecutionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByJobExecutionCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobTargetExecutionsClientListByJobExecutionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return JobTargetExecutionsClientListByJobExecutionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobTargetExecutionsClientListByJobExecutionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByJobExecutionHandleResponse(resp)
		},
	})
}

// listByJobExecutionCreateRequest creates the ListByJobExecution request.
func (client *JobTargetExecutionsClient) listByJobExecutionCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobTargetExecutionsClientListByJobExecutionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}/targets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobExecutionId}", url.PathEscape(jobExecutionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CreateTimeMin != nil {
		reqQP.Set("createTimeMin", options.CreateTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.CreateTimeMax != nil {
		reqQP.Set("createTimeMax", options.CreateTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMin != nil {
		reqQP.Set("endTimeMin", options.EndTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMax != nil {
		reqQP.Set("endTimeMax", options.EndTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.IsActive != nil {
		reqQP.Set("isActive", strconv.FormatBool(*options.IsActive))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(*options.Skip, 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(*options.Top, 10))
	}
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByJobExecutionHandleResponse handles the ListByJobExecution response.
func (client *JobTargetExecutionsClient) listByJobExecutionHandleResponse(resp *http.Response) (JobTargetExecutionsClientListByJobExecutionResponse, error) {
	result := JobTargetExecutionsClientListByJobExecutionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecutionListResult); err != nil {
		return JobTargetExecutionsClientListByJobExecutionResponse{}, err
	}
	return result, nil
}

// NewListByStepPager - Lists the target executions of a job step execution.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - jobName - The name of the job to get.
//   - jobExecutionID - The id of the job execution
//   - stepName - The name of the step.
//   - options - JobTargetExecutionsClientListByStepOptions contains the optional parameters for the JobTargetExecutionsClient.NewListByStepPager
//     method.
func (client *JobTargetExecutionsClient) NewListByStepPager(resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, stepName string, options *JobTargetExecutionsClientListByStepOptions) *runtime.Pager[JobTargetExecutionsClientListByStepResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobTargetExecutionsClientListByStepResponse]{
		More: func(page JobTargetExecutionsClientListByStepResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobTargetExecutionsClientListByStepResponse) (JobTargetExecutionsClientListByStepResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByStepCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, stepName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobTargetExecutionsClientListByStepResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return JobTargetExecutionsClientListByStepResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobTargetExecutionsClientListByStepResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByStepHandleResponse(resp)
		},
	})
}

// listByStepCreateRequest creates the ListByStep request.
func (client *JobTargetExecutionsClient) listByStepCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, stepName string, options *JobTargetExecutionsClientListByStepOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}/steps/{stepName}/targets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobExecutionId}", url.PathEscape(jobExecutionID))
	if stepName == "" {
		return nil, errors.New("parameter stepName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stepName}", url.PathEscape(stepName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CreateTimeMin != nil {
		reqQP.Set("createTimeMin", options.CreateTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.CreateTimeMax != nil {
		reqQP.Set("createTimeMax", options.CreateTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMin != nil {
		reqQP.Set("endTimeMin", options.EndTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMax != nil {
		reqQP.Set("endTimeMax", options.EndTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.IsActive != nil {
		reqQP.Set("isActive", strconv.FormatBool(*options.IsActive))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(*options.Skip, 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(*options.Top, 10))
	}
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByStepHandleResponse handles the ListByStep response.
func (client *JobTargetExecutionsClient) listByStepHandleResponse(resp *http.Response) (JobTargetExecutionsClientListByStepResponse, error) {
	result := JobTargetExecutionsClientListByStepResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecutionListResult); err != nil {
		return JobTargetExecutionsClientListByStepResponse{}, err
	}
	return result, nil
}
