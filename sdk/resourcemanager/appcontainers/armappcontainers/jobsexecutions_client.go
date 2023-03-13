//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

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

// JobsExecutionsClient contains the methods for the JobsExecutions group.
// Don't use this type directly, use NewJobsExecutionsClient() instead.
type JobsExecutionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobsExecutionsClient creates a new instance of JobsExecutionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewJobsExecutionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobsExecutionsClient, error) {
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
	client := &JobsExecutionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginList - Get a Container Apps Job's executions
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - jobName - Name of the Container Apps Job.
//   - options - JobsExecutionsClientBeginListOptions contains the optional parameters for the JobsExecutionsClient.BeginList
//     method.
func (client *JobsExecutionsClient) BeginList(ctx context.Context, resourceGroupName string, jobName string, options *JobsExecutionsClientBeginListOptions) (*runtime.Poller[*runtime.Pager[JobsExecutionsClientListResponse]], error) {
	pager := runtime.NewPager(runtime.PagingHandler[JobsExecutionsClientListResponse]{
		More: func(page JobsExecutionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobsExecutionsClientListResponse) (JobsExecutionsClientListResponse, error) {
			req, err := runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			if err != nil {
				return JobsExecutionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobsExecutionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobsExecutionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listOperation(ctx, resourceGroupName, jobName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[*runtime.Pager[JobsExecutionsClientListResponse]]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Response:      &pager,
		})
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.pl, &runtime.NewPollerFromResumeTokenOptions[*runtime.Pager[JobsExecutionsClientListResponse]]{
			Response: &pager,
		})
	}
}

// List - Get a Container Apps Job's executions
//
// Generated from API version 2022-11-01-preview
func (client *JobsExecutionsClient) listOperation(ctx context.Context, resourceGroupName string, jobName string, options *JobsExecutionsClientBeginListOptions) (*http.Response, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, jobName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// listCreateRequest creates the List request.
func (client *JobsExecutionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *JobsExecutionsClientBeginListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/jobs/{jobName}/executions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *JobsExecutionsClient) listHandleResponse(resp *http.Response) (JobsExecutionsClientListResponse, error) {
	result := JobsExecutionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAppJobExecutions); err != nil {
		return JobsExecutionsClientListResponse{}, err
	}
	return result, nil
}
