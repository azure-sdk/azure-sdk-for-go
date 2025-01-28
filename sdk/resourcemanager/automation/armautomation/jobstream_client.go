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
	"net/http"
	"net/url"
	"strings"
)

// JobStreamClient contains the methods for the JobStream group.
// Don't use this type directly, use NewJobStreamClient() instead.
type JobStreamClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewJobStreamClient creates a new instance of JobStreamClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewJobStreamClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobStreamClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &JobStreamClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Retrieve the job stream identified by job stream id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - jobName - The job name.
//   - jobStreamID - The job stream id.
//   - options - JobStreamClientGetOptions contains the optional parameters for the JobStreamClient.Get method.
func (client *JobStreamClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, jobStreamID string, options *JobStreamClientGetOptions) (JobStreamClientGetResponse, error) {
	var err error
	const operationName = "JobStreamClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, jobStreamID, options)
	if err != nil {
		return JobStreamClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobStreamClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JobStreamClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *JobStreamClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, jobStreamID string, options *JobStreamClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/streams/{jobStreamId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if jobStreamID == "" {
		return nil, errors.New("parameter jobStreamID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobStreamId}", url.PathEscape(jobStreamID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["clientRequestId"] = []string{*options.ClientRequestID}
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobStreamClient) getHandleResponse(resp *http.Response) (JobStreamClientGetResponse, error) {
	result := JobStreamClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStream); err != nil {
		return JobStreamClientGetResponse{}, err
	}
	return result, nil
}

// NewListByJobPager - Retrieve a list of jobs streams identified by job name.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - jobName - The job name.
//   - options - JobStreamClientListByJobOptions contains the optional parameters for the JobStreamClient.NewListByJobPager method.
func (client *JobStreamClient) NewListByJobPager(resourceGroupName string, automationAccountName string, jobName string, options *JobStreamClientListByJobOptions) *runtime.Pager[JobStreamClientListByJobResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobStreamClientListByJobResponse]{
		More: func(page JobStreamClientListByJobResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobStreamClientListByJobResponse) (JobStreamClientListByJobResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "JobStreamClient.NewListByJobPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByJobCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, options)
			}, nil)
			if err != nil {
				return JobStreamClientListByJobResponse{}, err
			}
			return client.listByJobHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByJobCreateRequest creates the ListByJob request.
func (client *JobStreamClient) listByJobCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobStreamClientListByJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/streams"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
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
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["clientRequestId"] = []string{*options.ClientRequestID}
	}
	return req, nil
}

// listByJobHandleResponse handles the ListByJob response.
func (client *JobStreamClient) listByJobHandleResponse(resp *http.Response) (JobStreamClientListByJobResponse, error) {
	result := JobStreamClientListByJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStreamListResult); err != nil {
		return JobStreamClientListByJobResponse{}, err
	}
	return result, nil
}
