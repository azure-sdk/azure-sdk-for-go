//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevhub

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

// WorkflowClient contains the methods for the Workflow group.
// Don't use this type directly, use NewWorkflowClient() instead.
type WorkflowClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkflowClient creates a new instance of WorkflowClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkflowClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkflowClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkflowClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a workflow
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workflowName - The name of the workflow resource.
//   - options - WorkflowClientCreateOrUpdateOptions contains the optional parameters for the WorkflowClient.CreateOrUpdate method.
func (client *WorkflowClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workflowName string, parameters Workflow, options *WorkflowClientCreateOrUpdateOptions) (WorkflowClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "WorkflowClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workflowName, parameters, options)
	if err != nil {
		return WorkflowClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkflowClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, parameters Workflow, options *WorkflowClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevHub/workflows/{workflowName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkflowClient) createOrUpdateHandleResponse(resp *http.Response) (WorkflowClientCreateOrUpdateResponse, error) {
	result := WorkflowClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workflow); err != nil {
		return WorkflowClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a workflow
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workflowName - The name of the workflow resource.
//   - options - WorkflowClientDeleteOptions contains the optional parameters for the WorkflowClient.Delete method.
func (client *WorkflowClient) Delete(ctx context.Context, resourceGroupName string, workflowName string, options *WorkflowClientDeleteOptions) (WorkflowClientDeleteResponse, error) {
	var err error
	const operationName = "WorkflowClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workflowName, options)
	if err != nil {
		return WorkflowClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowClientDeleteResponse{}, err
	}
	resp, err := client.deleteHandleResponse(httpResp)
	return resp, err
}

// deleteCreateRequest creates the Delete request.
func (client *WorkflowClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, options *WorkflowClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevHub/workflows/{workflowName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *WorkflowClient) deleteHandleResponse(resp *http.Response) (WorkflowClientDeleteResponse, error) {
	result := WorkflowClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeleteWorkflowResponse); err != nil {
		return WorkflowClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Gets a workflow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workflowName - The name of the workflow resource.
//   - options - WorkflowClientGetOptions contains the optional parameters for the WorkflowClient.Get method.
func (client *WorkflowClient) Get(ctx context.Context, resourceGroupName string, workflowName string, options *WorkflowClientGetOptions) (WorkflowClientGetResponse, error) {
	var err error
	const operationName = "WorkflowClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, options)
	if err != nil {
		return WorkflowClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkflowClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, options *WorkflowClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevHub/workflows/{workflowName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowClient) getHandleResponse(resp *http.Response) (WorkflowClientGetResponse, error) {
	result := WorkflowClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workflow); err != nil {
		return WorkflowClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of workflows associated with the specified subscription.
//
// Generated from API version 2023-08-01
//   - options - WorkflowClientListOptions contains the optional parameters for the WorkflowClient.NewListPager method.
func (client *WorkflowClient) NewListPager(options *WorkflowClientListOptions) *runtime.Pager[WorkflowClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkflowClientListResponse]{
		More: func(page WorkflowClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkflowClientListResponse) (WorkflowClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkflowClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return WorkflowClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *WorkflowClient) listCreateRequest(ctx context.Context, options *WorkflowClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/workflows"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowClient) listHandleResponse(resp *http.Response) (WorkflowClientListResponse, error) {
	result := WorkflowClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowListResult); err != nil {
		return WorkflowClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of workflows within a resource group.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - WorkflowClientListByResourceGroupOptions contains the optional parameters for the WorkflowClient.NewListByResourceGroupPager
//     method.
func (client *WorkflowClient) NewListByResourceGroupPager(resourceGroupName string, options *WorkflowClientListByResourceGroupOptions) *runtime.Pager[WorkflowClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkflowClientListByResourceGroupResponse]{
		More: func(page WorkflowClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkflowClientListByResourceGroupResponse) (WorkflowClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkflowClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return WorkflowClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *WorkflowClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *WorkflowClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevHub/workflows"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	if options != nil && options.ManagedClusterResource != nil {
		reqQP.Set("managedClusterResource", *options.ManagedClusterResource)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *WorkflowClient) listByResourceGroupHandleResponse(resp *http.Response) (WorkflowClientListByResourceGroupResponse, error) {
	result := WorkflowClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowListResult); err != nil {
		return WorkflowClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates tags on a workflow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workflowName - The name of the workflow resource.
//   - parameters - Parameters supplied to the Update Workflow Tags operation.
//   - options - WorkflowClientUpdateTagsOptions contains the optional parameters for the WorkflowClient.UpdateTags method.
func (client *WorkflowClient) UpdateTags(ctx context.Context, resourceGroupName string, workflowName string, parameters TagsObject, options *WorkflowClientUpdateTagsOptions) (WorkflowClientUpdateTagsResponse, error) {
	var err error
	const operationName = "WorkflowClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, workflowName, parameters, options)
	if err != nil {
		return WorkflowClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *WorkflowClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, parameters TagsObject, options *WorkflowClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevHub/workflows/{workflowName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *WorkflowClient) updateTagsHandleResponse(resp *http.Response) (WorkflowClientUpdateTagsResponse, error) {
	result := WorkflowClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workflow); err != nil {
		return WorkflowClientUpdateTagsResponse{}, err
	}
	return result, nil
}
