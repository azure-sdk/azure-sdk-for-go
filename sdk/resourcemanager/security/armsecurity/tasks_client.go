// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// TasksClient contains the methods for the Tasks group.
// Don't use this type directly, use NewTasksClient() instead.
type TasksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTasksClient creates a new instance of TasksClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTasksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TasksClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TasksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetResourceGroupLevelTask - Recommended tasks that will help improve the security of the subscription proactively
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - taskName - Name of the task object, will be a GUID
//   - options - TasksClientGetResourceGroupLevelTaskOptions contains the optional parameters for the TasksClient.GetResourceGroupLevelTask
//     method.
func (client *TasksClient) GetResourceGroupLevelTask(ctx context.Context, resourceGroupName string, ascLocation string, taskName string, options *TasksClientGetResourceGroupLevelTaskOptions) (TasksClientGetResourceGroupLevelTaskResponse, error) {
	var err error
	const operationName = "TasksClient.GetResourceGroupLevelTask"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getResourceGroupLevelTaskCreateRequest(ctx, resourceGroupName, ascLocation, taskName, options)
	if err != nil {
		return TasksClientGetResourceGroupLevelTaskResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TasksClientGetResourceGroupLevelTaskResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TasksClientGetResourceGroupLevelTaskResponse{}, err
	}
	resp, err := client.getResourceGroupLevelTaskHandleResponse(httpResp)
	return resp, err
}

// getResourceGroupLevelTaskCreateRequest creates the GetResourceGroupLevelTask request.
func (client *TasksClient) getResourceGroupLevelTaskCreateRequest(ctx context.Context, resourceGroupName string, ascLocation string, taskName string, _ *TasksClientGetResourceGroupLevelTaskOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/tasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getResourceGroupLevelTaskHandleResponse handles the GetResourceGroupLevelTask response.
func (client *TasksClient) getResourceGroupLevelTaskHandleResponse(resp *http.Response) (TasksClientGetResourceGroupLevelTaskResponse, error) {
	result := TasksClientGetResourceGroupLevelTaskResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Task); err != nil {
		return TasksClientGetResourceGroupLevelTaskResponse{}, err
	}
	return result, nil
}

// GetSubscriptionLevelTask - Recommended tasks that will help improve the security of the subscription proactively
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-06-01-preview
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - taskName - Name of the task object, will be a GUID
//   - options - TasksClientGetSubscriptionLevelTaskOptions contains the optional parameters for the TasksClient.GetSubscriptionLevelTask
//     method.
func (client *TasksClient) GetSubscriptionLevelTask(ctx context.Context, ascLocation string, taskName string, options *TasksClientGetSubscriptionLevelTaskOptions) (TasksClientGetSubscriptionLevelTaskResponse, error) {
	var err error
	const operationName = "TasksClient.GetSubscriptionLevelTask"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getSubscriptionLevelTaskCreateRequest(ctx, ascLocation, taskName, options)
	if err != nil {
		return TasksClientGetSubscriptionLevelTaskResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TasksClientGetSubscriptionLevelTaskResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TasksClientGetSubscriptionLevelTaskResponse{}, err
	}
	resp, err := client.getSubscriptionLevelTaskHandleResponse(httpResp)
	return resp, err
}

// getSubscriptionLevelTaskCreateRequest creates the GetSubscriptionLevelTask request.
func (client *TasksClient) getSubscriptionLevelTaskCreateRequest(ctx context.Context, ascLocation string, taskName string, _ *TasksClientGetSubscriptionLevelTaskOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/tasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSubscriptionLevelTaskHandleResponse handles the GetSubscriptionLevelTask response.
func (client *TasksClient) getSubscriptionLevelTaskHandleResponse(resp *http.Response) (TasksClientGetSubscriptionLevelTaskResponse, error) {
	result := TasksClientGetSubscriptionLevelTaskResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Task); err != nil {
		return TasksClientGetSubscriptionLevelTaskResponse{}, err
	}
	return result, nil
}

// NewListPager - Recommended tasks that will help improve the security of the subscription proactively
//
// Generated from API version 2015-06-01-preview
//   - options - TasksClientListOptions contains the optional parameters for the TasksClient.NewListPager method.
func (client *TasksClient) NewListPager(options *TasksClientListOptions) *runtime.Pager[TasksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TasksClientListResponse]{
		More: func(page TasksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TasksClientListResponse) (TasksClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TasksClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TasksClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TasksClient) listCreateRequest(ctx context.Context, options *TasksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/tasks"
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
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TasksClient) listHandleResponse(resp *http.Response) (TasksClientListResponse, error) {
	result := TasksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskList); err != nil {
		return TasksClientListResponse{}, err
	}
	return result, nil
}

// NewListByHomeRegionPager - Recommended tasks that will help improve the security of the subscription proactively
//
// Generated from API version 2015-06-01-preview
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - options - TasksClientListByHomeRegionOptions contains the optional parameters for the TasksClient.NewListByHomeRegionPager
//     method.
func (client *TasksClient) NewListByHomeRegionPager(ascLocation string, options *TasksClientListByHomeRegionOptions) *runtime.Pager[TasksClientListByHomeRegionResponse] {
	return runtime.NewPager(runtime.PagingHandler[TasksClientListByHomeRegionResponse]{
		More: func(page TasksClientListByHomeRegionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TasksClientListByHomeRegionResponse) (TasksClientListByHomeRegionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TasksClient.NewListByHomeRegionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByHomeRegionCreateRequest(ctx, ascLocation, options)
			}, nil)
			if err != nil {
				return TasksClientListByHomeRegionResponse{}, err
			}
			return client.listByHomeRegionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByHomeRegionCreateRequest creates the ListByHomeRegion request.
func (client *TasksClient) listByHomeRegionCreateRequest(ctx context.Context, ascLocation string, options *TasksClientListByHomeRegionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/tasks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHomeRegionHandleResponse handles the ListByHomeRegion response.
func (client *TasksClient) listByHomeRegionHandleResponse(resp *http.Response) (TasksClientListByHomeRegionResponse, error) {
	result := TasksClientListByHomeRegionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskList); err != nil {
		return TasksClientListByHomeRegionResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Recommended tasks that will help improve the security of the subscription proactively
//
// Generated from API version 2015-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - options - TasksClientListByResourceGroupOptions contains the optional parameters for the TasksClient.NewListByResourceGroupPager
//     method.
func (client *TasksClient) NewListByResourceGroupPager(resourceGroupName string, ascLocation string, options *TasksClientListByResourceGroupOptions) *runtime.Pager[TasksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[TasksClientListByResourceGroupResponse]{
		More: func(page TasksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TasksClientListByResourceGroupResponse) (TasksClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TasksClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, ascLocation, options)
			}, nil)
			if err != nil {
				return TasksClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *TasksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, ascLocation string, options *TasksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/tasks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *TasksClient) listByResourceGroupHandleResponse(resp *http.Response) (TasksClientListByResourceGroupResponse, error) {
	result := TasksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskList); err != nil {
		return TasksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateResourceGroupLevelTaskState - Recommended tasks that will help improve the security of the subscription proactively
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - taskName - Name of the task object, will be a GUID
//   - taskUpdateActionType - Type of the action to do on the task
//   - options - TasksClientUpdateResourceGroupLevelTaskStateOptions contains the optional parameters for the TasksClient.UpdateResourceGroupLevelTaskState
//     method.
func (client *TasksClient) UpdateResourceGroupLevelTaskState(ctx context.Context, resourceGroupName string, ascLocation string, taskName string, taskUpdateActionType TaskUpdateActionType, options *TasksClientUpdateResourceGroupLevelTaskStateOptions) (TasksClientUpdateResourceGroupLevelTaskStateResponse, error) {
	var err error
	const operationName = "TasksClient.UpdateResourceGroupLevelTaskState"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateResourceGroupLevelTaskStateCreateRequest(ctx, resourceGroupName, ascLocation, taskName, taskUpdateActionType, options)
	if err != nil {
		return TasksClientUpdateResourceGroupLevelTaskStateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TasksClientUpdateResourceGroupLevelTaskStateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TasksClientUpdateResourceGroupLevelTaskStateResponse{}, err
	}
	return TasksClientUpdateResourceGroupLevelTaskStateResponse{}, nil
}

// updateResourceGroupLevelTaskStateCreateRequest creates the UpdateResourceGroupLevelTaskState request.
func (client *TasksClient) updateResourceGroupLevelTaskStateCreateRequest(ctx context.Context, resourceGroupName string, ascLocation string, taskName string, taskUpdateActionType TaskUpdateActionType, _ *TasksClientUpdateResourceGroupLevelTaskStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/tasks/{taskName}/{taskUpdateActionType}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	if taskUpdateActionType == "" {
		return nil, errors.New("parameter taskUpdateActionType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskUpdateActionType}", url.PathEscape(string(taskUpdateActionType)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// UpdateSubscriptionLevelTaskState - Recommended tasks that will help improve the security of the subscription proactively
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-06-01-preview
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - taskName - Name of the task object, will be a GUID
//   - taskUpdateActionType - Type of the action to do on the task
//   - options - TasksClientUpdateSubscriptionLevelTaskStateOptions contains the optional parameters for the TasksClient.UpdateSubscriptionLevelTaskState
//     method.
func (client *TasksClient) UpdateSubscriptionLevelTaskState(ctx context.Context, ascLocation string, taskName string, taskUpdateActionType TaskUpdateActionType, options *TasksClientUpdateSubscriptionLevelTaskStateOptions) (TasksClientUpdateSubscriptionLevelTaskStateResponse, error) {
	var err error
	const operationName = "TasksClient.UpdateSubscriptionLevelTaskState"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateSubscriptionLevelTaskStateCreateRequest(ctx, ascLocation, taskName, taskUpdateActionType, options)
	if err != nil {
		return TasksClientUpdateSubscriptionLevelTaskStateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TasksClientUpdateSubscriptionLevelTaskStateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TasksClientUpdateSubscriptionLevelTaskStateResponse{}, err
	}
	return TasksClientUpdateSubscriptionLevelTaskStateResponse{}, nil
}

// updateSubscriptionLevelTaskStateCreateRequest creates the UpdateSubscriptionLevelTaskState request.
func (client *TasksClient) updateSubscriptionLevelTaskStateCreateRequest(ctx context.Context, ascLocation string, taskName string, taskUpdateActionType TaskUpdateActionType, _ *TasksClientUpdateSubscriptionLevelTaskStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/tasks/{taskName}/{taskUpdateActionType}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	if taskUpdateActionType == "" {
		return nil, errors.New("parameter taskUpdateActionType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskUpdateActionType}", url.PathEscape(string(taskUpdateActionType)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
