// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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

// ReplicationJobsClient contains the methods for the ReplicationJobs group.
// Don't use this type directly, use NewReplicationJobsClient() instead.
type ReplicationJobsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationJobsClient creates a new instance of ReplicationJobsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationJobsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationJobsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCancel - The operation to cancel an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - jobName - Job identifier.
//   - options - ReplicationJobsClientBeginCancelOptions contains the optional parameters for the ReplicationJobsClient.BeginCancel
//     method.
func (client *ReplicationJobsClient) BeginCancel(ctx context.Context, resourceGroupName string, resourceName string, jobName string, options *ReplicationJobsClientBeginCancelOptions) (*runtime.Poller[ReplicationJobsClientCancelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancel(ctx, resourceGroupName, resourceName, jobName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationJobsClientCancelResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationJobsClientCancelResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Cancel - The operation to cancel an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *ReplicationJobsClient) cancel(ctx context.Context, resourceGroupName string, resourceName string, jobName string, options *ReplicationJobsClientBeginCancelOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationJobsClient.BeginCancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, resourceName, jobName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *ReplicationJobsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, jobName string, _ *ReplicationJobsClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExport - The operation to export the details of the Azure Site Recovery jobs of the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - jobQueryParameter - job query filter.
//   - options - ReplicationJobsClientBeginExportOptions contains the optional parameters for the ReplicationJobsClient.BeginExport
//     method.
func (client *ReplicationJobsClient) BeginExport(ctx context.Context, resourceGroupName string, resourceName string, jobQueryParameter JobQueryParameter, options *ReplicationJobsClientBeginExportOptions) (*runtime.Poller[ReplicationJobsClientExportResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.export(ctx, resourceGroupName, resourceName, jobQueryParameter, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationJobsClientExportResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationJobsClientExportResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Export - The operation to export the details of the Azure Site Recovery jobs of the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *ReplicationJobsClient) export(ctx context.Context, resourceGroupName string, resourceName string, jobQueryParameter JobQueryParameter, options *ReplicationJobsClientBeginExportOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationJobsClient.BeginExport"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportCreateRequest(ctx, resourceGroupName, resourceName, jobQueryParameter, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// exportCreateRequest creates the Export request.
func (client *ReplicationJobsClient) exportCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, jobQueryParameter JobQueryParameter, _ *ReplicationJobsClientBeginExportOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/export"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, jobQueryParameter); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get the details of an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - jobName - Job identifier.
//   - options - ReplicationJobsClientGetOptions contains the optional parameters for the ReplicationJobsClient.Get method.
func (client *ReplicationJobsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, jobName string, options *ReplicationJobsClientGetOptions) (ReplicationJobsClientGetResponse, error) {
	var err error
	const operationName = "ReplicationJobsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, jobName, options)
	if err != nil {
		return ReplicationJobsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationJobsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReplicationJobsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReplicationJobsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, jobName string, _ *ReplicationJobsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationJobsClient) getHandleResponse(resp *http.Response) (ReplicationJobsClientGetResponse, error) {
	result := ReplicationJobsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return ReplicationJobsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of Azure Site Recovery Jobs for the vault.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - options - ReplicationJobsClientListOptions contains the optional parameters for the ReplicationJobsClient.NewListPager
//     method.
func (client *ReplicationJobsClient) NewListPager(resourceGroupName string, resourceName string, options *ReplicationJobsClientListOptions) *runtime.Pager[ReplicationJobsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationJobsClientListResponse]{
		More: func(page ReplicationJobsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationJobsClientListResponse) (ReplicationJobsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReplicationJobsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return ReplicationJobsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationJobsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ReplicationJobsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationJobsClient) listHandleResponse(resp *http.Response) (ReplicationJobsClientListResponse, error) {
	result := ReplicationJobsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollection); err != nil {
		return ReplicationJobsClientListResponse{}, err
	}
	return result, nil
}

// BeginRestart - The operation to restart an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - jobName - Job identifier.
//   - options - ReplicationJobsClientBeginRestartOptions contains the optional parameters for the ReplicationJobsClient.BeginRestart
//     method.
func (client *ReplicationJobsClient) BeginRestart(ctx context.Context, resourceGroupName string, resourceName string, jobName string, options *ReplicationJobsClientBeginRestartOptions) (*runtime.Poller[ReplicationJobsClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceGroupName, resourceName, jobName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationJobsClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationJobsClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Restart - The operation to restart an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *ReplicationJobsClient) restart(ctx context.Context, resourceGroupName string, resourceName string, jobName string, options *ReplicationJobsClientBeginRestartOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationJobsClient.BeginRestart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartCreateRequest(ctx, resourceGroupName, resourceName, jobName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// restartCreateRequest creates the Restart request.
func (client *ReplicationJobsClient) restartCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, jobName string, _ *ReplicationJobsClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}/restart"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginResume - The operation to resume an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - jobName - Job identifier.
//   - resumeJobParams - Resume rob comments.
//   - options - ReplicationJobsClientBeginResumeOptions contains the optional parameters for the ReplicationJobsClient.BeginResume
//     method.
func (client *ReplicationJobsClient) BeginResume(ctx context.Context, resourceGroupName string, resourceName string, jobName string, resumeJobParams ResumeJobParams, options *ReplicationJobsClientBeginResumeOptions) (*runtime.Poller[ReplicationJobsClientResumeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resume(ctx, resourceGroupName, resourceName, jobName, resumeJobParams, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationJobsClientResumeResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationJobsClientResumeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Resume - The operation to resume an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *ReplicationJobsClient) resume(ctx context.Context, resourceGroupName string, resourceName string, jobName string, resumeJobParams ResumeJobParams, options *ReplicationJobsClientBeginResumeOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationJobsClient.BeginResume"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.resumeCreateRequest(ctx, resourceGroupName, resourceName, jobName, resumeJobParams, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// resumeCreateRequest creates the Resume request.
func (client *ReplicationJobsClient) resumeCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, jobName string, resumeJobParams ResumeJobParams, _ *ReplicationJobsClientBeginResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}/resume"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resumeJobParams); err != nil {
		return nil, err
	}
	return req, nil
}
