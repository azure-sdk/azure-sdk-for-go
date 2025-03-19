// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

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

// PipelineRunsClient contains the methods for the PipelineRuns group.
// Don't use this type directly, use NewPipelineRunsClient() instead.
type PipelineRunsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPipelineRunsClient creates a new instance of PipelineRunsClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPipelineRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PipelineRunsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PipelineRunsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Cancel - Cancel a pipeline run by its run ID.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - runID - The pipeline run identifier.
//   - options - PipelineRunsClientCancelOptions contains the optional parameters for the PipelineRunsClient.Cancel method.
func (client *PipelineRunsClient) Cancel(ctx context.Context, resourceGroupName string, factoryName string, runID string, options *PipelineRunsClientCancelOptions) (PipelineRunsClientCancelResponse, error) {
	var err error
	const operationName = "PipelineRunsClient.Cancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, factoryName, runID, options)
	if err != nil {
		return PipelineRunsClientCancelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineRunsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelineRunsClientCancelResponse{}, err
	}
	return PipelineRunsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *PipelineRunsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, runID string, options *PipelineRunsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelineruns/{runId}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	if options != nil && options.IsRecursive != nil {
		reqQP.Set("isRecursive", strconv.FormatBool(*options.IsRecursive))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a pipeline run by its run ID.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - runID - The pipeline run identifier.
//   - options - PipelineRunsClientGetOptions contains the optional parameters for the PipelineRunsClient.Get method.
func (client *PipelineRunsClient) Get(ctx context.Context, resourceGroupName string, factoryName string, runID string, options *PipelineRunsClientGetOptions) (PipelineRunsClientGetResponse, error) {
	var err error
	const operationName = "PipelineRunsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, runID, options)
	if err != nil {
		return PipelineRunsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineRunsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelineRunsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PipelineRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, runID string, _ *PipelineRunsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelineruns/{runId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PipelineRunsClient) getHandleResponse(resp *http.Response) (PipelineRunsClientGetResponse, error) {
	result := PipelineRunsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRun); err != nil {
		return PipelineRunsClientGetResponse{}, err
	}
	return result, nil
}

// QueryByFactory - Query pipeline runs in the factory based on input filter conditions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - filterParameters - Parameters to filter the pipeline run.
//   - options - PipelineRunsClientQueryByFactoryOptions contains the optional parameters for the PipelineRunsClient.QueryByFactory
//     method.
func (client *PipelineRunsClient) QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string, filterParameters RunFilterParameters, options *PipelineRunsClientQueryByFactoryOptions) (PipelineRunsClientQueryByFactoryResponse, error) {
	var err error
	const operationName = "PipelineRunsClient.QueryByFactory"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.queryByFactoryCreateRequest(ctx, resourceGroupName, factoryName, filterParameters, options)
	if err != nil {
		return PipelineRunsClientQueryByFactoryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineRunsClientQueryByFactoryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelineRunsClientQueryByFactoryResponse{}, err
	}
	resp, err := client.queryByFactoryHandleResponse(httpResp)
	return resp, err
}

// queryByFactoryCreateRequest creates the QueryByFactory request.
func (client *PipelineRunsClient) queryByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, filterParameters RunFilterParameters, _ *PipelineRunsClientQueryByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/queryPipelineRuns"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, filterParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// queryByFactoryHandleResponse handles the QueryByFactory response.
func (client *PipelineRunsClient) queryByFactoryHandleResponse(resp *http.Response) (PipelineRunsClientQueryByFactoryResponse, error) {
	result := PipelineRunsClientQueryByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRunsQueryResponse); err != nil {
		return PipelineRunsClientQueryByFactoryResponse{}, err
	}
	return result, nil
}
