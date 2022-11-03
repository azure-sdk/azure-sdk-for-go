//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearning

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

// LabelingJobsClient contains the methods for the LabelingJobs group.
// Don't use this type directly, use NewLabelingJobsClient() instead.
type LabelingJobsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLabelingJobsClient creates a new instance of LabelingJobsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLabelingJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LabelingJobsClient, error) {
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
	client := &LabelingJobsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a labeling job (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// id - The name and identifier for the LabelingJob.
// body - LabelingJob definition object.
// options - LabelingJobsClientBeginCreateOrUpdateOptions contains the optional parameters for the LabelingJobsClient.BeginCreateOrUpdate
// method.
func (client *LabelingJobsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, id string, body LabelingJob, options *LabelingJobsClientBeginCreateOrUpdateOptions) (*runtime.Poller[LabelingJobsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, id, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LabelingJobsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LabelingJobsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a labeling job (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01-preview
func (client *LabelingJobsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, id string, body LabelingJob, options *LabelingJobsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, id, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LabelingJobsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, id string, body LabelingJob, options *LabelingJobsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/labelingJobs/{id}"
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
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// Delete - Delete a labeling job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// id - The name and identifier for the LabelingJob.
// options - LabelingJobsClientDeleteOptions contains the optional parameters for the LabelingJobsClient.Delete method.
func (client *LabelingJobsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *LabelingJobsClientDeleteOptions) (LabelingJobsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, id, options)
	if err != nil {
		return LabelingJobsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LabelingJobsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return LabelingJobsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return LabelingJobsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LabelingJobsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *LabelingJobsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/labelingJobs/{id}"
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
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExportLabels - Export labels from a labeling job (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// id - The name and identifier for the LabelingJob.
// body - The export summary.
// options - LabelingJobsClientBeginExportLabelsOptions contains the optional parameters for the LabelingJobsClient.BeginExportLabels
// method.
func (client *LabelingJobsClient) BeginExportLabels(ctx context.Context, resourceGroupName string, workspaceName string, id string, body ExportSummaryClassification, options *LabelingJobsClientBeginExportLabelsOptions) (*runtime.Poller[LabelingJobsClientExportLabelsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportLabels(ctx, resourceGroupName, workspaceName, id, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LabelingJobsClientExportLabelsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LabelingJobsClientExportLabelsResponse](options.ResumeToken, client.pl, nil)
	}
}

// ExportLabels - Export labels from a labeling job (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01-preview
func (client *LabelingJobsClient) exportLabels(ctx context.Context, resourceGroupName string, workspaceName string, id string, body ExportSummaryClassification, options *LabelingJobsClientBeginExportLabelsOptions) (*http.Response, error) {
	req, err := client.exportLabelsCreateRequest(ctx, resourceGroupName, workspaceName, id, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// exportLabelsCreateRequest creates the ExportLabels request.
func (client *LabelingJobsClient) exportLabelsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, id string, body ExportSummaryClassification, options *LabelingJobsClientBeginExportLabelsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/labelingJobs/{id}/exportLabels"
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
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// Get - Gets a labeling job by name/id.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// id - The name and identifier for the LabelingJob.
// options - LabelingJobsClientGetOptions contains the optional parameters for the LabelingJobsClient.Get method.
func (client *LabelingJobsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *LabelingJobsClientGetOptions) (LabelingJobsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, id, options)
	if err != nil {
		return LabelingJobsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LabelingJobsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LabelingJobsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LabelingJobsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *LabelingJobsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/labelingJobs/{id}"
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
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	if options != nil && options.IncludeJobInstructions != nil {
		reqQP.Set("includeJobInstructions", strconv.FormatBool(*options.IncludeJobInstructions))
	}
	if options != nil && options.IncludeLabelCategories != nil {
		reqQP.Set("includeLabelCategories", strconv.FormatBool(*options.IncludeLabelCategories))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LabelingJobsClient) getHandleResponse(resp *http.Response) (LabelingJobsClientGetResponse, error) {
	result := LabelingJobsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LabelingJob); err != nil {
		return LabelingJobsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists labeling jobs in the workspace.
// Generated from API version 2022-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// options - LabelingJobsClientListOptions contains the optional parameters for the LabelingJobsClient.List method.
func (client *LabelingJobsClient) NewListPager(resourceGroupName string, workspaceName string, options *LabelingJobsClientListOptions) *runtime.Pager[LabelingJobsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LabelingJobsClientListResponse]{
		More: func(page LabelingJobsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LabelingJobsClientListResponse) (LabelingJobsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LabelingJobsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LabelingJobsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LabelingJobsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LabelingJobsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *LabelingJobsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/labelingJobs"
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
	reqQP.Set("api-version", "2022-12-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatInt(int64(*options.Count), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LabelingJobsClient) listHandleResponse(resp *http.Response) (LabelingJobsClientListResponse, error) {
	result := LabelingJobsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LabelingJobResourceArmPaginatedResult); err != nil {
		return LabelingJobsClientListResponse{}, err
	}
	return result, nil
}

// Pause - Pause a labeling job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// id - The name and identifier for the LabelingJob.
// options - LabelingJobsClientPauseOptions contains the optional parameters for the LabelingJobsClient.Pause method.
func (client *LabelingJobsClient) Pause(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *LabelingJobsClientPauseOptions) (LabelingJobsClientPauseResponse, error) {
	req, err := client.pauseCreateRequest(ctx, resourceGroupName, workspaceName, id, options)
	if err != nil {
		return LabelingJobsClientPauseResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LabelingJobsClientPauseResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LabelingJobsClientPauseResponse{}, runtime.NewResponseError(resp)
	}
	return LabelingJobsClientPauseResponse{}, nil
}

// pauseCreateRequest creates the Pause request.
func (client *LabelingJobsClient) pauseCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *LabelingJobsClientPauseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/labelingJobs/{id}/pause"
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
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginResume - Resume a labeling job (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// id - The name and identifier for the LabelingJob.
// options - LabelingJobsClientBeginResumeOptions contains the optional parameters for the LabelingJobsClient.BeginResume
// method.
func (client *LabelingJobsClient) BeginResume(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *LabelingJobsClientBeginResumeOptions) (*runtime.Poller[LabelingJobsClientResumeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resume(ctx, resourceGroupName, workspaceName, id, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LabelingJobsClientResumeResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LabelingJobsClientResumeResponse](options.ResumeToken, client.pl, nil)
	}
}

// Resume - Resume a labeling job (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01-preview
func (client *LabelingJobsClient) resume(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *LabelingJobsClientBeginResumeOptions) (*http.Response, error) {
	req, err := client.resumeCreateRequest(ctx, resourceGroupName, workspaceName, id, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// resumeCreateRequest creates the Resume request.
func (client *LabelingJobsClient) resumeCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *LabelingJobsClientBeginResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/labelingJobs/{id}/resume"
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
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
