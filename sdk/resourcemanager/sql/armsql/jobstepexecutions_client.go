//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// JobStepExecutionsClient contains the methods for the JobStepExecutions group.
// Don't use this type directly, use NewJobStepExecutionsClient() instead.
type JobStepExecutionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewJobStepExecutionsClient creates a new instance of JobStepExecutionsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewJobStepExecutionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobStepExecutionsClient, error) {
	cl, err := arm.NewClient(moduleName+".JobStepExecutionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &JobStepExecutionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a step execution of a job execution.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - jobName - The name of the job to get.
//   - jobExecutionID - The unique id of the job execution
//   - stepName - The name of the step.
//   - options - JobStepExecutionsClientGetOptions contains the optional parameters for the JobStepExecutionsClient.Get method.
func (client *JobStepExecutionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, stepName string, options *JobStepExecutionsClientGetOptions) (JobStepExecutionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, stepName, options)
	if err != nil {
		return JobStepExecutionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobStepExecutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobStepExecutionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobStepExecutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, stepName string, options *JobStepExecutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}/steps/{stepName}"
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
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobStepExecutionsClient) getHandleResponse(resp *http.Response) (JobStepExecutionsClientGetResponse, error) {
	result := JobStepExecutionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecution); err != nil {
		return JobStepExecutionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByJobExecutionPager - Lists the step executions of a job execution.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - jobName - The name of the job to get.
//   - jobExecutionID - The id of the job execution
//   - options - JobStepExecutionsClientListByJobExecutionOptions contains the optional parameters for the JobStepExecutionsClient.NewListByJobExecutionPager
//     method.
func (client *JobStepExecutionsClient) NewListByJobExecutionPager(resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobStepExecutionsClientListByJobExecutionOptions) *runtime.Pager[JobStepExecutionsClientListByJobExecutionResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobStepExecutionsClientListByJobExecutionResponse]{
		More: func(page JobStepExecutionsClientListByJobExecutionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobStepExecutionsClientListByJobExecutionResponse) (JobStepExecutionsClientListByJobExecutionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByJobExecutionCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobStepExecutionsClientListByJobExecutionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return JobStepExecutionsClientListByJobExecutionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobStepExecutionsClientListByJobExecutionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByJobExecutionHandleResponse(resp)
		},
	})
}

// listByJobExecutionCreateRequest creates the ListByJobExecution request.
func (client *JobStepExecutionsClient) listByJobExecutionCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobStepExecutionsClientListByJobExecutionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}/steps"
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
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByJobExecutionHandleResponse handles the ListByJobExecution response.
func (client *JobStepExecutionsClient) listByJobExecutionHandleResponse(resp *http.Response) (JobStepExecutionsClientListByJobExecutionResponse, error) {
	result := JobStepExecutionsClientListByJobExecutionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecutionListResult); err != nil {
		return JobStepExecutionsClientListByJobExecutionResponse{}, err
	}
	return result, nil
}
