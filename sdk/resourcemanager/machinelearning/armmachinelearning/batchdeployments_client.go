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

// BatchDeploymentsClient contains the methods for the BatchDeployments group.
// Don't use this type directly, use NewBatchDeploymentsClient() instead.
type BatchDeploymentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBatchDeploymentsClient creates a new instance of BatchDeploymentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBatchDeploymentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BatchDeploymentsClient, error) {
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
	client := &BatchDeploymentsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates/updates a batch inference deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Inference endpoint name
//   - deploymentName - The identifier for the Batch inference deployment.
//   - body - Batch inference deployment definition object.
//   - options - BatchDeploymentsClientBeginCreateOrUpdateOptions contains the optional parameters for the BatchDeploymentsClient.BeginCreateOrUpdate
//     method.
func (client *BatchDeploymentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body BatchDeployment, options *BatchDeploymentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[BatchDeploymentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[BatchDeploymentsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[BatchDeploymentsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates/updates a batch inference deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
func (client *BatchDeploymentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body BatchDeployment, options *BatchDeploymentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
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
func (client *BatchDeploymentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body BatchDeployment, options *BatchDeploymentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/batchEndpoints/{endpointName}/deployments/{deploymentName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
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

// BeginDelete - Delete Batch Inference deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Endpoint name
//   - deploymentName - Inference deployment identifier.
//   - options - BatchDeploymentsClientBeginDeleteOptions contains the optional parameters for the BatchDeploymentsClient.BeginDelete
//     method.
func (client *BatchDeploymentsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *BatchDeploymentsClientBeginDeleteOptions) (*runtime.Poller[BatchDeploymentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[BatchDeploymentsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[BatchDeploymentsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete Batch Inference deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
func (client *BatchDeploymentsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *BatchDeploymentsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BatchDeploymentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *BatchDeploymentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/batchEndpoints/{endpointName}/deployments/{deploymentName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
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

// Get - Gets a batch inference deployment by id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Endpoint name
//   - deploymentName - The identifier for the Batch deployments.
//   - options - BatchDeploymentsClientGetOptions contains the optional parameters for the BatchDeploymentsClient.Get method.
func (client *BatchDeploymentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *BatchDeploymentsClientGetOptions) (BatchDeploymentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, options)
	if err != nil {
		return BatchDeploymentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BatchDeploymentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BatchDeploymentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BatchDeploymentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, options *BatchDeploymentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/batchEndpoints/{endpointName}/deployments/{deploymentName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BatchDeploymentsClient) getHandleResponse(resp *http.Response) (BatchDeploymentsClientGetResponse, error) {
	result := BatchDeploymentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BatchDeployment); err != nil {
		return BatchDeploymentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists Batch inference deployments in the workspace.
//
// Generated from API version 2022-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Endpoint name
//   - options - BatchDeploymentsClientListOptions contains the optional parameters for the BatchDeploymentsClient.NewListPager
//     method.
func (client *BatchDeploymentsClient) NewListPager(resourceGroupName string, workspaceName string, endpointName string, options *BatchDeploymentsClientListOptions) *runtime.Pager[BatchDeploymentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BatchDeploymentsClientListResponse]{
		More: func(page BatchDeploymentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BatchDeploymentsClientListResponse) (BatchDeploymentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BatchDeploymentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BatchDeploymentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BatchDeploymentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BatchDeploymentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *BatchDeploymentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/batchEndpoints/{endpointName}/deployments"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderBy", *options.OrderBy)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BatchDeploymentsClient) listHandleResponse(resp *http.Response) (BatchDeploymentsClientListResponse, error) {
	result := BatchDeploymentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BatchDeploymentTrackedResourceArmPaginatedResult); err != nil {
		return BatchDeploymentsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a batch inference deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - endpointName - Inference endpoint name
//   - deploymentName - The identifier for the Batch inference deployment.
//   - body - Batch inference deployment definition object.
//   - options - BatchDeploymentsClientBeginUpdateOptions contains the optional parameters for the BatchDeploymentsClient.BeginUpdate
//     method.
func (client *BatchDeploymentsClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body PartialBatchDeploymentPartialMinimalTrackedResourceWithProperties, options *BatchDeploymentsClientBeginUpdateOptions) (*runtime.Poller[BatchDeploymentsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[BatchDeploymentsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[BatchDeploymentsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update a batch inference deployment (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
func (client *BatchDeploymentsClient) update(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body PartialBatchDeploymentPartialMinimalTrackedResourceWithProperties, options *BatchDeploymentsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, endpointName, deploymentName, body, options)
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

// updateCreateRequest creates the Update request.
func (client *BatchDeploymentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, deploymentName string, body PartialBatchDeploymentPartialMinimalTrackedResourceWithProperties, options *BatchDeploymentsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/batchEndpoints/{endpointName}/deployments/{deploymentName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if deploymentName == "" {
		return nil, errors.New("parameter deploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deploymentName}", url.PathEscape(deploymentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
