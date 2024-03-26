//go:build go1.18
// +build go1.18

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

// InferencePoolsClient contains the methods for the InferencePools group.
// Don't use this type directly, use NewInferencePoolsClient() instead.
type InferencePoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInferencePoolsClient creates a new instance of InferencePoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInferencePoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InferencePoolsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InferencePoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update InferencePool (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - inferencePoolName - Name of InferencePool
//   - body - InferencePool entity to apply during operation.
//   - options - InferencePoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the InferencePoolsClient.BeginCreateOrUpdate
//     method.
func (client *InferencePoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, body InferencePool, options *InferencePoolsClientBeginCreateOrUpdateOptions) (*runtime.Poller[InferencePoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, inferencePoolName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InferencePoolsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InferencePoolsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update InferencePool (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
func (client *InferencePoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, body InferencePool, options *InferencePoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "InferencePoolsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, inferencePoolName, body, options)
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
func (client *InferencePoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, body InferencePool, options *InferencePoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{inferencePoolName}"
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
	if inferencePoolName == "" {
		return nil, errors.New("parameter inferencePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inferencePoolName}", url.PathEscape(inferencePoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete InferencePool (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - inferencePoolName - Name of InferencePool
//   - options - InferencePoolsClientBeginDeleteOptions contains the optional parameters for the InferencePoolsClient.BeginDelete
//     method.
func (client *InferencePoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, options *InferencePoolsClientBeginDeleteOptions) (*runtime.Poller[InferencePoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, inferencePoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InferencePoolsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InferencePoolsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete InferencePool (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
func (client *InferencePoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, options *InferencePoolsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "InferencePoolsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, inferencePoolName, options)
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
func (client *InferencePoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, options *InferencePoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{inferencePoolName}"
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
	if inferencePoolName == "" {
		return nil, errors.New("parameter inferencePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inferencePoolName}", url.PathEscape(inferencePoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get InferencePool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - inferencePoolName - Name of InferencePool
//   - options - InferencePoolsClientGetOptions contains the optional parameters for the InferencePoolsClient.Get method.
func (client *InferencePoolsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, options *InferencePoolsClientGetOptions) (InferencePoolsClientGetResponse, error) {
	var err error
	const operationName = "InferencePoolsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, inferencePoolName, options)
	if err != nil {
		return InferencePoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InferencePoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InferencePoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *InferencePoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, options *InferencePoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{inferencePoolName}"
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
	if inferencePoolName == "" {
		return nil, errors.New("parameter inferencePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inferencePoolName}", url.PathEscape(inferencePoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InferencePoolsClient) getHandleResponse(resp *http.Response) (InferencePoolsClientGetResponse, error) {
	result := InferencePoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InferencePool); err != nil {
		return InferencePoolsClientGetResponse{}, err
	}
	return result, nil
}

// GetStatus - Retrieve inference pool status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - inferencePoolName - Name of InferencePool
//   - options - InferencePoolsClientGetStatusOptions contains the optional parameters for the InferencePoolsClient.GetStatus
//     method.
func (client *InferencePoolsClient) GetStatus(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, options *InferencePoolsClientGetStatusOptions) (InferencePoolsClientGetStatusResponse, error) {
	var err error
	const operationName = "InferencePoolsClient.GetStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, workspaceName, inferencePoolName, options)
	if err != nil {
		return InferencePoolsClientGetStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InferencePoolsClientGetStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InferencePoolsClientGetStatusResponse{}, err
	}
	resp, err := client.getStatusHandleResponse(httpResp)
	return resp, err
}

// getStatusCreateRequest creates the GetStatus request.
func (client *InferencePoolsClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, options *InferencePoolsClientGetStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{inferencePoolName}/getStatus"
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
	if inferencePoolName == "" {
		return nil, errors.New("parameter inferencePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inferencePoolName}", url.PathEscape(inferencePoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getStatusHandleResponse handles the GetStatus response.
func (client *InferencePoolsClient) getStatusHandleResponse(resp *http.Response) (InferencePoolsClientGetStatusResponse, error) {
	result := InferencePoolsClientGetStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PoolStatus); err != nil {
		return InferencePoolsClientGetStatusResponse{}, err
	}
	return result, nil
}

// NewListPager - List InferencePools.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - options - InferencePoolsClientListOptions contains the optional parameters for the InferencePoolsClient.NewListPager method.
func (client *InferencePoolsClient) NewListPager(resourceGroupName string, workspaceName string, options *InferencePoolsClientListOptions) *runtime.Pager[InferencePoolsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[InferencePoolsClientListResponse]{
		More: func(page InferencePoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InferencePoolsClientListResponse) (InferencePoolsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "InferencePoolsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return InferencePoolsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *InferencePoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *InferencePoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	reqQP.Set("api-version", "2024-04-01-preview")
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
func (client *InferencePoolsClient) listHandleResponse(resp *http.Response) (InferencePoolsClientListResponse, error) {
	result := InferencePoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InferencePoolTrackedResourceArmPaginatedResult); err != nil {
		return InferencePoolsClientListResponse{}, err
	}
	return result, nil
}

// NewListSKUsPager - List Inference Pool Skus.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - inferencePoolName - Inference Group name.
//   - options - InferencePoolsClientListSKUsOptions contains the optional parameters for the InferencePoolsClient.NewListSKUsPager
//     method.
func (client *InferencePoolsClient) NewListSKUsPager(resourceGroupName string, workspaceName string, inferencePoolName string, options *InferencePoolsClientListSKUsOptions) *runtime.Pager[InferencePoolsClientListSKUsResponse] {
	return runtime.NewPager(runtime.PagingHandler[InferencePoolsClientListSKUsResponse]{
		More: func(page InferencePoolsClientListSKUsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InferencePoolsClientListSKUsResponse) (InferencePoolsClientListSKUsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "InferencePoolsClient.NewListSKUsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSKUsCreateRequest(ctx, resourceGroupName, workspaceName, inferencePoolName, options)
			}, nil)
			if err != nil {
				return InferencePoolsClientListSKUsResponse{}, err
			}
			return client.listSKUsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *InferencePoolsClient) listSKUsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, options *InferencePoolsClientListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{inferencePoolName}/skus"
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
	if inferencePoolName == "" {
		return nil, errors.New("parameter inferencePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inferencePoolName}", url.PathEscape(inferencePoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	reqQP.Set("api-version", "2024-04-01-preview")
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatInt(int64(*options.Count), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *InferencePoolsClient) listSKUsHandleResponse(resp *http.Response) (InferencePoolsClientListSKUsResponse, error) {
	result := InferencePoolsClientListSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUResourceArmPaginatedResult); err != nil {
		return InferencePoolsClientListSKUsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update InferencePool (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - inferencePoolName - Name of InferencePool
//   - body - Inference Pool entity to apply during operation.
//   - options - InferencePoolsClientBeginUpdateOptions contains the optional parameters for the InferencePoolsClient.BeginUpdate
//     method.
func (client *InferencePoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, body PartialMinimalTrackedResourceWithSKUAndIdentity, options *InferencePoolsClientBeginUpdateOptions) (*runtime.Poller[InferencePoolsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, inferencePoolName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InferencePoolsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InferencePoolsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update InferencePool (asynchronous).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
func (client *InferencePoolsClient) update(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, body PartialMinimalTrackedResourceWithSKUAndIdentity, options *InferencePoolsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "InferencePoolsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, inferencePoolName, body, options)
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
func (client *InferencePoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, inferencePoolName string, body PartialMinimalTrackedResourceWithSKUAndIdentity, options *InferencePoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/inferencePools/{inferencePoolName}"
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
	if inferencePoolName == "" {
		return nil, errors.New("parameter inferencePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inferencePoolName}", url.PathEscape(inferencePoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
