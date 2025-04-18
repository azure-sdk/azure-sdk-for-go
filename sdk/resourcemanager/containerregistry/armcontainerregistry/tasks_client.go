// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

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
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
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

// BeginCreate - Creates a task for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - taskName - The name of the container registry task.
//   - taskCreateParameters - The parameters for creating a task.
//   - options - TasksClientBeginCreateOptions contains the optional parameters for the TasksClient.BeginCreate method.
func (client *TasksClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, taskName string, taskCreateParameters Task, options *TasksClientBeginCreateOptions) (*runtime.Poller[TasksClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, taskName, taskCreateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TasksClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TasksClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates a task for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *TasksClient) create(ctx context.Context, resourceGroupName string, registryName string, taskName string, taskCreateParameters Task, options *TasksClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "TasksClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, taskName, taskCreateParameters, options)
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

// createCreateRequest creates the Create request.
func (client *TasksClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, taskName string, taskCreateParameters Task, _ *TasksClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, taskCreateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a specified task.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - taskName - The name of the container registry task.
//   - options - TasksClientBeginDeleteOptions contains the optional parameters for the TasksClient.BeginDelete method.
func (client *TasksClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, taskName string, options *TasksClientBeginDeleteOptions) (*runtime.Poller[TasksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, taskName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TasksClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TasksClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a specified task.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *TasksClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, taskName string, options *TasksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "TasksClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, taskName, options)
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
func (client *TasksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, taskName string, _ *TasksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the properties of a specified task.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - taskName - The name of the container registry task.
//   - options - TasksClientGetOptions contains the optional parameters for the TasksClient.Get method.
func (client *TasksClient) Get(ctx context.Context, resourceGroupName string, registryName string, taskName string, options *TasksClientGetOptions) (TasksClientGetResponse, error) {
	var err error
	const operationName = "TasksClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, taskName, options)
	if err != nil {
		return TasksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TasksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TasksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TasksClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, taskName string, _ *TasksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TasksClient) getHandleResponse(resp *http.Response) (TasksClientGetResponse, error) {
	result := TasksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Task); err != nil {
		return TasksClientGetResponse{}, err
	}
	return result, nil
}

// GetDetails - Returns a task with extended information that includes all secrets.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - taskName - The name of the container registry task.
//   - options - TasksClientGetDetailsOptions contains the optional parameters for the TasksClient.GetDetails method.
func (client *TasksClient) GetDetails(ctx context.Context, resourceGroupName string, registryName string, taskName string, options *TasksClientGetDetailsOptions) (TasksClientGetDetailsResponse, error) {
	var err error
	const operationName = "TasksClient.GetDetails"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDetailsCreateRequest(ctx, resourceGroupName, registryName, taskName, options)
	if err != nil {
		return TasksClientGetDetailsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TasksClientGetDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TasksClientGetDetailsResponse{}, err
	}
	resp, err := client.getDetailsHandleResponse(httpResp)
	return resp, err
}

// getDetailsCreateRequest creates the GetDetails request.
func (client *TasksClient) getDetailsCreateRequest(ctx context.Context, resourceGroupName string, registryName string, taskName string, _ *TasksClientGetDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}/listDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDetailsHandleResponse handles the GetDetails response.
func (client *TasksClient) getDetailsHandleResponse(resp *http.Response) (TasksClientGetDetailsResponse, error) {
	result := TasksClientGetDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Task); err != nil {
		return TasksClientGetDetailsResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the tasks for a specified container registry.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - options - TasksClientListOptions contains the optional parameters for the TasksClient.NewListPager method.
func (client *TasksClient) NewListPager(resourceGroupName string, registryName string, options *TasksClientListOptions) *runtime.Pager[TasksClientListResponse] {
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
				return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
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
func (client *TasksClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, _ *TasksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TasksClient) listHandleResponse(resp *http.Response) (TasksClientListResponse, error) {
	result := TasksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskListResult); err != nil {
		return TasksClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a task with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - taskName - The name of the container registry task.
//   - taskUpdateParameters - The parameters for updating a task.
//   - options - TasksClientBeginUpdateOptions contains the optional parameters for the TasksClient.BeginUpdate method.
func (client *TasksClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, taskName string, taskUpdateParameters TaskUpdateParameters, options *TasksClientBeginUpdateOptions) (*runtime.Poller[TasksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, registryName, taskName, taskUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TasksClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TasksClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates a task with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *TasksClient) update(ctx context.Context, resourceGroupName string, registryName string, taskName string, taskUpdateParameters TaskUpdateParameters, options *TasksClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "TasksClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, taskName, taskUpdateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *TasksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, taskName string, taskUpdateParameters TaskUpdateParameters, _ *TasksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, taskUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
