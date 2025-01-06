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
	"strings"
)

// ComputeClient contains the methods for the Compute group.
// Don't use this type directly, use NewComputeClient() instead.
type ComputeClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewComputeClient creates a new instance of ComputeClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewComputeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ComputeClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ComputeClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable
// operation. If your intent is to create a new compute, do a GET first to verify that it does not
// exist yet.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - computeName - Name of the Azure Machine Learning compute.
//   - parameters - Payload with Machine Learning compute definition.
//   - options - ComputeClientBeginCreateOrUpdateOptions contains the optional parameters for the ComputeClient.BeginCreateOrUpdate
//     method.
func (client *ComputeClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource, options *ComputeClientBeginCreateOrUpdateOptions) (*runtime.Poller[ComputeClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ComputeClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ComputeClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates compute. This call will overwrite a compute if it exists. This is a nonrecoverable
// operation. If your intent is to create a new compute, do a GET first to verify that it does not
// exist yet.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ComputeClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource, options *ComputeClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ComputeClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
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
func (client *ComputeClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ComputeResource, options *ComputeClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
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
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes specified Machine Learning compute.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - computeName - Name of the Azure Machine Learning compute.
//   - underlyingResourceAction - Delete the underlying compute if 'Delete', or detach the underlying compute from workspace if
//     'Detach'.
//   - options - ComputeClientBeginDeleteOptions contains the optional parameters for the ComputeClient.BeginDelete method.
func (client *ComputeClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction, options *ComputeClientBeginDeleteOptions) (*runtime.Poller[ComputeClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, computeName, underlyingResourceAction, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ComputeClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ComputeClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes specified Machine Learning compute.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ComputeClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction, options *ComputeClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ComputeClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, computeName, underlyingResourceAction, options)
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
func (client *ComputeClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction UnderlyingResourceAction, options *ComputeClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
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
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	reqQP.Set("underlyingResourceAction", string(underlyingResourceAction))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets compute definition by its name. Any secrets (storage keys, service credentials, etc) are not returned - use
// 'keys' nested resource to get them.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - computeName - Name of the Azure Machine Learning compute.
//   - options - ComputeClientGetOptions contains the optional parameters for the ComputeClient.Get method.
func (client *ComputeClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientGetOptions) (ComputeClientGetResponse, error) {
	var err error
	const operationName = "ComputeClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComputeClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComputeClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ComputeClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
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
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ComputeClient) getHandleResponse(resp *http.Response) (ComputeClientGetResponse, error) {
	result := ComputeClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComputeResource); err != nil {
		return ComputeClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets computes in specified workspace.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - options - ComputeClientListOptions contains the optional parameters for the ComputeClient.NewListPager method.
func (client *ComputeClient) NewListPager(resourceGroupName string, workspaceName string, options *ComputeClientListOptions) *runtime.Pager[ComputeClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ComputeClientListResponse]{
		More: func(page ComputeClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ComputeClientListResponse) (ComputeClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ComputeClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return ComputeClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ComputeClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ComputeClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes"
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
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ComputeClient) listHandleResponse(resp *http.Response) (ComputeClientListResponse, error) {
	result := ComputeClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaginatedComputeResourcesList); err != nil {
		return ComputeClientListResponse{}, err
	}
	return result, nil
}

// ListKeys - Gets secrets related to Machine Learning compute (storage keys, service credentials, etc).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - computeName - Name of the Azure Machine Learning compute.
//   - options - ComputeClientListKeysOptions contains the optional parameters for the ComputeClient.ListKeys method.
func (client *ComputeClient) ListKeys(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientListKeysOptions) (ComputeClientListKeysResponse, error) {
	var err error
	const operationName = "ComputeClient.ListKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return ComputeClientListKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComputeClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComputeClientListKeysResponse{}, err
	}
	resp, err := client.listKeysHandleResponse(httpResp)
	return resp, err
}

// listKeysCreateRequest creates the ListKeys request.
func (client *ComputeClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/listKeys"
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
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *ComputeClient) listKeysHandleResponse(resp *http.Response) (ComputeClientListKeysResponse, error) {
	result := ComputeClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ComputeClientListKeysResponse{}, err
	}
	return result, nil
}

// NewListNodesPager - Get the details (e.g IP address, port etc) of all the compute nodes in the compute.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - computeName - Name of the Azure Machine Learning compute.
//   - options - ComputeClientListNodesOptions contains the optional parameters for the ComputeClient.NewListNodesPager method.
func (client *ComputeClient) NewListNodesPager(resourceGroupName string, workspaceName string, computeName string, options *ComputeClientListNodesOptions) *runtime.Pager[ComputeClientListNodesResponse] {
	return runtime.NewPager(runtime.PagingHandler[ComputeClientListNodesResponse]{
		More: func(page ComputeClientListNodesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ComputeClientListNodesResponse) (ComputeClientListNodesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ComputeClient.NewListNodesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listNodesCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
			}, nil)
			if err != nil {
				return ComputeClientListNodesResponse{}, err
			}
			return client.listNodesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listNodesCreateRequest creates the ListNodes request.
func (client *ComputeClient) listNodesCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientListNodesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/listNodes"
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
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listNodesHandleResponse handles the ListNodes response.
func (client *ComputeClient) listNodesHandleResponse(resp *http.Response) (ComputeClientListNodesResponse, error) {
	result := ComputeClientListNodesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AmlComputeNodesInformation); err != nil {
		return ComputeClientListNodesResponse{}, err
	}
	return result, nil
}

// BeginRestart - Posts a restart action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - computeName - Name of the Azure Machine Learning compute.
//   - options - ComputeClientBeginRestartOptions contains the optional parameters for the ComputeClient.BeginRestart method.
func (client *ComputeClient) BeginRestart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginRestartOptions) (*runtime.Poller[ComputeClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceGroupName, workspaceName, computeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ComputeClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ComputeClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Restart - Posts a restart action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ComputeClient) restart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginRestartOptions) (*http.Response, error) {
	var err error
	const operationName = "ComputeClient.BeginRestart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// restartCreateRequest creates the Restart request.
func (client *ComputeClient) restartCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/restart"
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
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStart - Posts a start action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - computeName - Name of the Azure Machine Learning compute.
//   - options - ComputeClientBeginStartOptions contains the optional parameters for the ComputeClient.BeginStart method.
func (client *ComputeClient) BeginStart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStartOptions) (*runtime.Poller[ComputeClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, workspaceName, computeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ComputeClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ComputeClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - Posts a start action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ComputeClient) start(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "ComputeClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// startCreateRequest creates the Start request.
func (client *ComputeClient) startCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/start"
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
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStop - Posts a stop action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - computeName - Name of the Azure Machine Learning compute.
//   - options - ComputeClientBeginStopOptions contains the optional parameters for the ComputeClient.BeginStop method.
func (client *ComputeClient) BeginStop(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStopOptions) (*runtime.Poller[ComputeClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, workspaceName, computeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ComputeClientStopResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ComputeClientStopResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Stop - Posts a stop action to a compute instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ComputeClient) stop(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStopOptions) (*http.Response, error) {
	var err error
	const operationName = "ComputeClient.BeginStop"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopCreateRequest(ctx, resourceGroupName, workspaceName, computeName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// stopCreateRequest creates the Stop request.
func (client *ComputeClient) stopCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, options *ComputeClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}/stop"
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
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Updates properties of a compute. This call will overwrite a compute if it exists. This is a nonrecoverable
// operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - computeName - Name of the Azure Machine Learning compute.
//   - parameters - Additional parameters for cluster update.
//   - options - ComputeClientBeginUpdateOptions contains the optional parameters for the ComputeClient.BeginUpdate method.
func (client *ComputeClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters, options *ComputeClientBeginUpdateOptions) (*runtime.Poller[ComputeClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ComputeClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ComputeClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates properties of a compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ComputeClient) update(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters, options *ComputeClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ComputeClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, computeName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *ComputeClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters ClusterUpdateParameters, options *ComputeClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}"
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
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
