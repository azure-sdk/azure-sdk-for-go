// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

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

// InferenceEndpointsClient contains the methods for the InferenceEndpoints group.
// Don't use this type directly, use NewInferenceEndpointsClient() instead.
type InferenceEndpointsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInferenceEndpointsClient creates a new instance of InferenceEndpointsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInferenceEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InferenceEndpointsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InferenceEndpointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update InferenceEndpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - InferencePool name.
//   - endpointName - InferenceEndpoint name.
//   - body - InferenceEndpoint entity to apply during operation.
//   - options - InferenceEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the InferenceEndpointsClient.BeginCreateOrUpdate
//     method.
func (client *InferenceEndpointsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, body InferenceEndpoint, options *InferenceEndpointsClientBeginCreateOrUpdateOptions) (*runtime.Poller[InferenceEndpointsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, poolName, endpointName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InferenceEndpointsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InferenceEndpointsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update InferenceEndpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
func (client *InferenceEndpointsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, body InferenceEndpoint, options *InferenceEndpointsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "InferenceEndpointsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, poolName, endpointName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InferenceEndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, body InferenceEndpoint, _ *InferenceEndpointsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/endpoints/{endpointName}"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete InferenceEndpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - InferencePool name.
//   - endpointName - InferenceEndpoint name.
//   - options - InferenceEndpointsClientBeginDeleteOptions contains the optional parameters for the InferenceEndpointsClient.BeginDelete
//     method.
func (client *InferenceEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, options *InferenceEndpointsClientBeginDeleteOptions) (*runtime.Poller[InferenceEndpointsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, poolName, endpointName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InferenceEndpointsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InferenceEndpointsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete InferenceEndpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
func (client *InferenceEndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, options *InferenceEndpointsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "InferenceEndpointsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, poolName, endpointName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *InferenceEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, _ *InferenceEndpointsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/endpoints/{endpointName}"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get InferenceEndpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - InferencePool name.
//   - endpointName - InferenceEndpoint name.
//   - options - InferenceEndpointsClientGetOptions contains the optional parameters for the InferenceEndpointsClient.Get method.
func (client *InferenceEndpointsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, options *InferenceEndpointsClientGetOptions) (InferenceEndpointsClientGetResponse, error) {
	var err error
	const operationName = "InferenceEndpointsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, poolName, endpointName, options)
	if err != nil {
		return InferenceEndpointsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InferenceEndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InferenceEndpointsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *InferenceEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, _ *InferenceEndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/endpoints/{endpointName}"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InferenceEndpointsClient) getHandleResponse(resp *http.Response) (InferenceEndpointsClientGetResponse, error) {
	result := InferenceEndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InferenceEndpoint); err != nil {
		return InferenceEndpointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List Inference Endpoints.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - Name of the InferencePool.
//   - options - InferenceEndpointsClientListOptions contains the optional parameters for the InferenceEndpointsClient.NewListPager
//     method.
func (client *InferenceEndpointsClient) NewListPager(resourceGroupName string, workspaceName string, poolName string, options *InferenceEndpointsClientListOptions) *runtime.Pager[InferenceEndpointsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[InferenceEndpointsClientListResponse]{
		More: func(page InferenceEndpointsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InferenceEndpointsClientListResponse) (InferenceEndpointsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "InferenceEndpointsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, poolName, options)
			}, nil)
			if err != nil {
				return InferenceEndpointsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *InferenceEndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, options *InferenceEndpointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/endpoints"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	reqQP.Set("api-version", "2025-01-01-preview")
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatInt(int64(*options.Count), 10))
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderBy", string(*options.OrderBy))
	}
	if options != nil && options.Properties != nil {
		reqQP.Set("properties", *options.Properties)
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", *options.Tags)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *InferenceEndpointsClient) listHandleResponse(resp *http.Response) (InferenceEndpointsClientListResponse, error) {
	result := InferenceEndpointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InferenceEndpointTrackedResourceArmPaginatedResult); err != nil {
		return InferenceEndpointsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update InferenceEndpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - poolName - InferencePool name.
//   - endpointName - InferenceEndpoint name.
//   - body - Online Endpoint entity to apply during operation.
//   - options - InferenceEndpointsClientBeginUpdateOptions contains the optional parameters for the InferenceEndpointsClient.BeginUpdate
//     method.
func (client *InferenceEndpointsClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, body any, options *InferenceEndpointsClientBeginUpdateOptions) (*runtime.Poller[InferenceEndpointsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, poolName, endpointName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InferenceEndpointsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InferenceEndpointsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update InferenceEndpoint (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
func (client *InferenceEndpointsClient) update(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, body any, options *InferenceEndpointsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "InferenceEndpointsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, poolName, endpointName, body, options)
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

// updateCreateRequest creates the Update request.
func (client *InferenceEndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, body any, _ *InferenceEndpointsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{poolName}/endpoints/{endpointName}"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
