//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

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

// TaskAssignmentsClient contains the methods for the StorageTaskAssignments group.
// Don't use this type directly, use NewTaskAssignmentsClient() instead.
type TaskAssignmentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTaskAssignmentsClient creates a new instance of TaskAssignmentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTaskAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TaskAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TaskAssignmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Asynchronously creates a new storage task assignment sub-resource with the specified parameters. If a storage
// task assignment is already created and a subsequent create request is issued with
// different properties, the storage task assignment properties will be updated. If a storage task assignment is already created
// and a subsequent create or update request is issued with the exact same
// set of properties, the request will succeed.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - storageTaskAssignmentName - The name of the storage task assignment within the specified resource group. Storage task assignment
//     names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
//   - parameters - The parameters to create a Storage Task Assignment.
//   - options - TaskAssignmentsClientBeginCreateOptions contains the optional parameters for the TaskAssignmentsClient.BeginCreate
//     method.
func (client *TaskAssignmentsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, parameters TaskAssignment, options *TaskAssignmentsClientBeginCreateOptions) (*runtime.Poller[TaskAssignmentsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, storageTaskAssignmentName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TaskAssignmentsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TaskAssignmentsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Asynchronously creates a new storage task assignment sub-resource with the specified parameters. If a storage
// task assignment is already created and a subsequent create request is issued with
// different properties, the storage task assignment properties will be updated. If a storage task assignment is already created
// and a subsequent create or update request is issued with the exact same
// set of properties, the request will succeed.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
func (client *TaskAssignmentsClient) create(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, parameters TaskAssignment, options *TaskAssignmentsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "TaskAssignmentsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, storageTaskAssignmentName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *TaskAssignmentsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, parameters TaskAssignment, options *TaskAssignmentsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/storageTaskAssignments/{storageTaskAssignmentName}"
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
	if storageTaskAssignmentName == "" {
		return nil, errors.New("parameter storageTaskAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTaskAssignmentName}", url.PathEscape(storageTaskAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the storage task assignment sub-resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - storageTaskAssignmentName - The name of the storage task assignment within the specified resource group. Storage task assignment
//     names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
//   - options - TaskAssignmentsClientBeginDeleteOptions contains the optional parameters for the TaskAssignmentsClient.BeginDelete
//     method.
func (client *TaskAssignmentsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, options *TaskAssignmentsClientBeginDeleteOptions) (*runtime.Poller[TaskAssignmentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, storageTaskAssignmentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TaskAssignmentsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TaskAssignmentsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the storage task assignment sub-resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
func (client *TaskAssignmentsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, options *TaskAssignmentsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "TaskAssignmentsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, storageTaskAssignmentName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TaskAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, options *TaskAssignmentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/storageTaskAssignments/{storageTaskAssignmentName}"
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
	if storageTaskAssignmentName == "" {
		return nil, errors.New("parameter storageTaskAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTaskAssignmentName}", url.PathEscape(storageTaskAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the storage task assignment properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - storageTaskAssignmentName - The name of the storage task assignment within the specified resource group. Storage task assignment
//     names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
//   - options - TaskAssignmentsClientGetOptions contains the optional parameters for the TaskAssignmentsClient.Get method.
func (client *TaskAssignmentsClient) Get(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, options *TaskAssignmentsClientGetOptions) (TaskAssignmentsClientGetResponse, error) {
	var err error
	const operationName = "TaskAssignmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, storageTaskAssignmentName, options)
	if err != nil {
		return TaskAssignmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TaskAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TaskAssignmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TaskAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, options *TaskAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/storageTaskAssignments/{storageTaskAssignmentName}"
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
	if storageTaskAssignmentName == "" {
		return nil, errors.New("parameter storageTaskAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTaskAssignmentName}", url.PathEscape(storageTaskAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TaskAssignmentsClient) getHandleResponse(resp *http.Response) (TaskAssignmentsClientGetResponse, error) {
	result := TaskAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskAssignment); err != nil {
		return TaskAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all the storage task assignments in an account
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - options - TaskAssignmentsClientListOptions contains the optional parameters for the TaskAssignmentsClient.NewListPager
//     method.
func (client *TaskAssignmentsClient) NewListPager(resourceGroupName string, accountName string, options *TaskAssignmentsClientListOptions) *runtime.Pager[TaskAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TaskAssignmentsClientListResponse]{
		More: func(page TaskAssignmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TaskAssignmentsClientListResponse) (TaskAssignmentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TaskAssignmentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return TaskAssignmentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TaskAssignmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *TaskAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/storageTaskAssignments"
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
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("$maxpagesize", strconv.FormatInt(int64(*options.Maxpagesize), 10))
	}
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TaskAssignmentsClient) listHandleResponse(resp *http.Response) (TaskAssignmentsClientListResponse, error) {
	result := TaskAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskAssignmentsList); err != nil {
		return TaskAssignmentsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update storage task assignment properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the storage account within the specified resource group. Storage account names must be between
//     3 and 24 characters in length and use numbers and lower-case letters only.
//   - storageTaskAssignmentName - The name of the storage task assignment within the specified resource group. Storage task assignment
//     names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
//   - parameters - The parameters to update a Storage Task Assignment.
//   - options - TaskAssignmentsClientBeginUpdateOptions contains the optional parameters for the TaskAssignmentsClient.BeginUpdate
//     method.
func (client *TaskAssignmentsClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, parameters TaskAssignmentUpdateParameters, options *TaskAssignmentsClientBeginUpdateOptions) (*runtime.Poller[TaskAssignmentsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, storageTaskAssignmentName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TaskAssignmentsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TaskAssignmentsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update storage task assignment properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
func (client *TaskAssignmentsClient) update(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, parameters TaskAssignmentUpdateParameters, options *TaskAssignmentsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "TaskAssignmentsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, storageTaskAssignmentName, parameters, options)
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
func (client *TaskAssignmentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageTaskAssignmentName string, parameters TaskAssignmentUpdateParameters, options *TaskAssignmentsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/storageTaskAssignments/{storageTaskAssignmentName}"
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
	if storageTaskAssignmentName == "" {
		return nil, errors.New("parameter storageTaskAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTaskAssignmentName}", url.PathEscape(storageTaskAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
