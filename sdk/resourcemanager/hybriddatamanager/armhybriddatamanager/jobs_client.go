// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybriddatamanager

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

// JobsClient contains the methods for the Jobs group.
// Don't use this type directly, use NewJobsClient() instead.
type JobsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewJobsClient creates a new instance of JobsClient with the specified values.
//   - subscriptionID - The Subscription Id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &JobsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCancel - Cancels the given job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01
//   - dataServiceName - The name of the data service of the job definition.
//   - jobDefinitionName - The name of the job definition of the job.
//   - jobID - The job id of the job queried.
//   - resourceGroupName - The Resource Group Name
//   - dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
//     3 and 24 characters in length and use any alphanumeric and underscore only
//   - options - JobsClientBeginCancelOptions contains the optional parameters for the JobsClient.BeginCancel method.
func (client *JobsClient) BeginCancel(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsClientBeginCancelOptions) (*runtime.Poller[JobsClientCancelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancel(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[JobsClientCancelResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[JobsClientCancelResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Cancel - Cancels the given job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01
func (client *JobsClient) cancel(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsClientBeginCancelOptions) (*http.Response, error) {
	var err error
	const operationName = "JobsClient.BeginCancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *JobsClient) cancelCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, _ *JobsClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}/cancel"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if jobID == "" {
		return nil, errors.New("parameter jobID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobId}", url.PathEscape(jobID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - This method gets a data manager job given the jobId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01
//   - dataServiceName - The name of the data service of the job definition.
//   - jobDefinitionName - The name of the job definition of the job.
//   - jobID - The job id of the job queried.
//   - resourceGroupName - The Resource Group Name
//   - dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
//     3 and 24 characters in length and use any alphanumeric and underscore only
//   - options - JobsClientGetOptions contains the optional parameters for the JobsClient.Get method.
func (client *JobsClient) Get(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsClientGetOptions) (JobsClientGetResponse, error) {
	var err error
	const operationName = "JobsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, options)
	if err != nil {
		return JobsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JobsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *JobsClient) getCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if jobID == "" {
		return nil, errors.New("parameter jobID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobId}", url.PathEscape(jobID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobsClient) getHandleResponse(resp *http.Response) (JobsClientGetResponse, error) {
	result := JobsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return JobsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDataManagerPager - This method gets all the jobs at the data manager resource level.
//
// Generated from API version 2019-06-01
//   - resourceGroupName - The Resource Group Name
//   - dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
//     3 and 24 characters in length and use any alphanumeric and underscore only
//   - options - JobsClientListByDataManagerOptions contains the optional parameters for the JobsClient.NewListByDataManagerPager
//     method.
func (client *JobsClient) NewListByDataManagerPager(resourceGroupName string, dataManagerName string, options *JobsClientListByDataManagerOptions) *runtime.Pager[JobsClientListByDataManagerResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobsClientListByDataManagerResponse]{
		More: func(page JobsClientListByDataManagerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobsClientListByDataManagerResponse) (JobsClientListByDataManagerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "JobsClient.NewListByDataManagerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDataManagerCreateRequest(ctx, resourceGroupName, dataManagerName, options)
			}, nil)
			if err != nil {
				return JobsClientListByDataManagerResponse{}, err
			}
			return client.listByDataManagerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDataManagerCreateRequest creates the ListByDataManager request.
func (client *JobsClient) listByDataManagerCreateRequest(ctx context.Context, resourceGroupName string, dataManagerName string, options *JobsClientListByDataManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataManagerHandleResponse handles the ListByDataManager response.
func (client *JobsClient) listByDataManagerHandleResponse(resp *http.Response) (JobsClientListByDataManagerResponse, error) {
	result := JobsClientListByDataManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobList); err != nil {
		return JobsClientListByDataManagerResponse{}, err
	}
	return result, nil
}

// NewListByDataServicePager - This method gets all the jobs of a data service type in a given resource.
//
// Generated from API version 2019-06-01
//   - dataServiceName - The name of the data service of interest.
//   - resourceGroupName - The Resource Group Name
//   - dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
//     3 and 24 characters in length and use any alphanumeric and underscore only
//   - options - JobsClientListByDataServiceOptions contains the optional parameters for the JobsClient.NewListByDataServicePager
//     method.
func (client *JobsClient) NewListByDataServicePager(dataServiceName string, resourceGroupName string, dataManagerName string, options *JobsClientListByDataServiceOptions) *runtime.Pager[JobsClientListByDataServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobsClientListByDataServiceResponse]{
		More: func(page JobsClientListByDataServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobsClientListByDataServiceResponse) (JobsClientListByDataServiceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "JobsClient.NewListByDataServicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDataServiceCreateRequest(ctx, dataServiceName, resourceGroupName, dataManagerName, options)
			}, nil)
			if err != nil {
				return JobsClientListByDataServiceResponse{}, err
			}
			return client.listByDataServiceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDataServiceCreateRequest creates the ListByDataService request.
func (client *JobsClient) listByDataServiceCreateRequest(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, options *JobsClientListByDataServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobs"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataServiceHandleResponse handles the ListByDataService response.
func (client *JobsClient) listByDataServiceHandleResponse(resp *http.Response) (JobsClientListByDataServiceResponse, error) {
	result := JobsClientListByDataServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobList); err != nil {
		return JobsClientListByDataServiceResponse{}, err
	}
	return result, nil
}

// NewListByJobDefinitionPager - This method gets all the jobs of a given job definition.
//
// Generated from API version 2019-06-01
//   - dataServiceName - The name of the data service of the job definition.
//   - jobDefinitionName - The name of the job definition for which jobs are needed.
//   - resourceGroupName - The Resource Group Name
//   - dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
//     3 and 24 characters in length and use any alphanumeric and underscore only
//   - options - JobsClientListByJobDefinitionOptions contains the optional parameters for the JobsClient.NewListByJobDefinitionPager
//     method.
func (client *JobsClient) NewListByJobDefinitionPager(dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, options *JobsClientListByJobDefinitionOptions) *runtime.Pager[JobsClientListByJobDefinitionResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobsClientListByJobDefinitionResponse]{
		More: func(page JobsClientListByJobDefinitionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobsClientListByJobDefinitionResponse) (JobsClientListByJobDefinitionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "JobsClient.NewListByJobDefinitionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByJobDefinitionCreateRequest(ctx, dataServiceName, jobDefinitionName, resourceGroupName, dataManagerName, options)
			}, nil)
			if err != nil {
				return JobsClientListByJobDefinitionResponse{}, err
			}
			return client.listByJobDefinitionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByJobDefinitionCreateRequest creates the ListByJobDefinition request.
func (client *JobsClient) listByJobDefinitionCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, resourceGroupName string, dataManagerName string, options *JobsClientListByJobDefinitionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByJobDefinitionHandleResponse handles the ListByJobDefinition response.
func (client *JobsClient) listByJobDefinitionHandleResponse(resp *http.Response) (JobsClientListByJobDefinitionResponse, error) {
	result := JobsClientListByJobDefinitionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobList); err != nil {
		return JobsClientListByJobDefinitionResponse{}, err
	}
	return result, nil
}

// BeginResume - Resumes the given job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01
//   - dataServiceName - The name of the data service of the job definition.
//   - jobDefinitionName - The name of the job definition of the job.
//   - jobID - The job id of the job queried.
//   - resourceGroupName - The Resource Group Name
//   - dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
//     3 and 24 characters in length and use any alphanumeric and underscore only
//   - options - JobsClientBeginResumeOptions contains the optional parameters for the JobsClient.BeginResume method.
func (client *JobsClient) BeginResume(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsClientBeginResumeOptions) (*runtime.Poller[JobsClientResumeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resume(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[JobsClientResumeResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[JobsClientResumeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Resume - Resumes the given job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01
func (client *JobsClient) resume(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, options *JobsClientBeginResumeOptions) (*http.Response, error) {
	var err error
	const operationName = "JobsClient.BeginResume"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.resumeCreateRequest(ctx, dataServiceName, jobDefinitionName, jobID, resourceGroupName, dataManagerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// resumeCreateRequest creates the Resume request.
func (client *JobsClient) resumeCreateRequest(ctx context.Context, dataServiceName string, jobDefinitionName string, jobID string, resourceGroupName string, dataManagerName string, _ *JobsClientBeginResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}/resume"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if jobDefinitionName == "" {
		return nil, errors.New("parameter jobDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobDefinitionName}", url.PathEscape(jobDefinitionName))
	if jobID == "" {
		return nil, errors.New("parameter jobID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobId}", url.PathEscape(jobID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
