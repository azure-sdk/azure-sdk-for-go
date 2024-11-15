//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer

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
)

// PipelineTopologiesClient contains the methods for the PipelineTopologies group.
// Don't use this type directly, use NewPipelineTopologiesClient() instead.
type PipelineTopologiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPipelineTopologiesClient creates a new instance of PipelineTopologiesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPipelineTopologiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PipelineTopologiesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PipelineTopologiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new pipeline topology or updates an existing one, with the given name. A pipeline topology describes
// the processing steps to be applied when processing content for a particular outcome. The
// topology should be defined according to the scenario to be achieved and can be reused across many pipeline instances which
// share the same processing characteristics.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The Azure Video Analyzer account name.
//   - pipelineTopologyName - Pipeline topology unique identifier.
//   - parameters - The request parameters
//   - options - PipelineTopologiesClientCreateOrUpdateOptions contains the optional parameters for the PipelineTopologiesClient.CreateOrUpdate
//     method.
func (client *PipelineTopologiesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, pipelineTopologyName string, parameters PipelineTopology, options *PipelineTopologiesClientCreateOrUpdateOptions) (PipelineTopologiesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "PipelineTopologiesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, pipelineTopologyName, parameters, options)
	if err != nil {
		return PipelineTopologiesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineTopologiesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return PipelineTopologiesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PipelineTopologiesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, pipelineTopologyName string, parameters PipelineTopology, options *PipelineTopologiesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/pipelineTopologies/{pipelineTopologyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if pipelineTopologyName == "" {
		return nil, errors.New("parameter pipelineTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineTopologyName}", url.PathEscape(pipelineTopologyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PipelineTopologiesClient) createOrUpdateHandleResponse(resp *http.Response) (PipelineTopologiesClientCreateOrUpdateResponse, error) {
	result := PipelineTopologiesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineTopology); err != nil {
		return PipelineTopologiesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a pipeline topology with the given name. This method should be called after all instances of the topology
// have been stopped and deleted.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The Azure Video Analyzer account name.
//   - pipelineTopologyName - Pipeline topology unique identifier.
//   - options - PipelineTopologiesClientDeleteOptions contains the optional parameters for the PipelineTopologiesClient.Delete
//     method.
func (client *PipelineTopologiesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, pipelineTopologyName string, options *PipelineTopologiesClientDeleteOptions) (PipelineTopologiesClientDeleteResponse, error) {
	var err error
	const operationName = "PipelineTopologiesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, pipelineTopologyName, options)
	if err != nil {
		return PipelineTopologiesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineTopologiesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PipelineTopologiesClientDeleteResponse{}, err
	}
	return PipelineTopologiesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PipelineTopologiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, pipelineTopologyName string, options *PipelineTopologiesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/pipelineTopologies/{pipelineTopologyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if pipelineTopologyName == "" {
		return nil, errors.New("parameter pipelineTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineTopologyName}", url.PathEscape(pipelineTopologyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves a specific pipeline topology by name. If a topology with that name has been previously created, the call
// will return the JSON representation of that topology.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The Azure Video Analyzer account name.
//   - pipelineTopologyName - Pipeline topology unique identifier.
//   - options - PipelineTopologiesClientGetOptions contains the optional parameters for the PipelineTopologiesClient.Get method.
func (client *PipelineTopologiesClient) Get(ctx context.Context, resourceGroupName string, accountName string, pipelineTopologyName string, options *PipelineTopologiesClientGetOptions) (PipelineTopologiesClientGetResponse, error) {
	var err error
	const operationName = "PipelineTopologiesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, pipelineTopologyName, options)
	if err != nil {
		return PipelineTopologiesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineTopologiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelineTopologiesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PipelineTopologiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, pipelineTopologyName string, options *PipelineTopologiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/pipelineTopologies/{pipelineTopologyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if pipelineTopologyName == "" {
		return nil, errors.New("parameter pipelineTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineTopologyName}", url.PathEscape(pipelineTopologyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PipelineTopologiesClient) getHandleResponse(resp *http.Response) (PipelineTopologiesClientGetResponse, error) {
	result := PipelineTopologiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineTopology); err != nil {
		return PipelineTopologiesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves a list of pipeline topologies that have been added to the account, if any, along with their JSON
// representation.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The Azure Video Analyzer account name.
//   - options - PipelineTopologiesClientListOptions contains the optional parameters for the PipelineTopologiesClient.NewListPager
//     method.
func (client *PipelineTopologiesClient) NewListPager(resourceGroupName string, accountName string, options *PipelineTopologiesClientListOptions) *runtime.Pager[PipelineTopologiesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PipelineTopologiesClientListResponse]{
		More: func(page PipelineTopologiesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PipelineTopologiesClientListResponse) (PipelineTopologiesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PipelineTopologiesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return PipelineTopologiesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PipelineTopologiesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *PipelineTopologiesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/pipelineTopologies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PipelineTopologiesClient) listHandleResponse(resp *http.Response) (PipelineTopologiesClientListResponse, error) {
	result := PipelineTopologiesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineTopologyCollection); err != nil {
		return PipelineTopologiesClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing pipeline topology with the given name. If the associated live pipelines or pipeline jobs are
// in active or processing state, respectively, then only the description can be updated.
// Else, the properties that can be updated include: description, parameter declarations, sources, processors, and sinks.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The Azure Video Analyzer account name.
//   - pipelineTopologyName - Pipeline topology unique identifier.
//   - parameters - The request parameters
//   - options - PipelineTopologiesClientUpdateOptions contains the optional parameters for the PipelineTopologiesClient.Update
//     method.
func (client *PipelineTopologiesClient) Update(ctx context.Context, resourceGroupName string, accountName string, pipelineTopologyName string, parameters PipelineTopologyUpdate, options *PipelineTopologiesClientUpdateOptions) (PipelineTopologiesClientUpdateResponse, error) {
	var err error
	const operationName = "PipelineTopologiesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, pipelineTopologyName, parameters, options)
	if err != nil {
		return PipelineTopologiesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineTopologiesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelineTopologiesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *PipelineTopologiesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, pipelineTopologyName string, parameters PipelineTopologyUpdate, options *PipelineTopologiesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/pipelineTopologies/{pipelineTopologyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if pipelineTopologyName == "" {
		return nil, errors.New("parameter pipelineTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineTopologyName}", url.PathEscape(pipelineTopologyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *PipelineTopologiesClient) updateHandleResponse(resp *http.Response) (PipelineTopologiesClientUpdateResponse, error) {
	result := PipelineTopologiesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineTopology); err != nil {
		return PipelineTopologiesClientUpdateResponse{}, err
	}
	return result, nil
}
