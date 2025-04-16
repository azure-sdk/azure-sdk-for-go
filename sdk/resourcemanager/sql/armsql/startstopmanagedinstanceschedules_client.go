// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// StartStopManagedInstanceSchedulesClient contains the methods for the StartStopManagedInstanceSchedules group.
// Don't use this type directly, use NewStartStopManagedInstanceSchedulesClient() instead.
type StartStopManagedInstanceSchedulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStartStopManagedInstanceSchedulesClient creates a new instance of StartStopManagedInstanceSchedulesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStartStopManagedInstanceSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StartStopManagedInstanceSchedulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StartStopManagedInstanceSchedulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the managed instance's Start/Stop schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - startStopScheduleName - Name of the managed instance Start/Stop schedule.
//   - parameters - The requested managed instance Start/Stop schedule.
//   - options - StartStopManagedInstanceSchedulesClientCreateOrUpdateOptions contains the optional parameters for the StartStopManagedInstanceSchedulesClient.CreateOrUpdate
//     method.
func (client *StartStopManagedInstanceSchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, startStopScheduleName StartStopScheduleName, parameters StartStopManagedInstanceSchedule, options *StartStopManagedInstanceSchedulesClientCreateOrUpdateOptions) (StartStopManagedInstanceSchedulesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "StartStopManagedInstanceSchedulesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, startStopScheduleName, parameters, options)
	if err != nil {
		return StartStopManagedInstanceSchedulesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StartStopManagedInstanceSchedulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return StartStopManagedInstanceSchedulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StartStopManagedInstanceSchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, startStopScheduleName StartStopScheduleName, parameters StartStopManagedInstanceSchedule, _ *StartStopManagedInstanceSchedulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/startStopSchedules/{startStopScheduleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if startStopScheduleName == "" {
		return nil, errors.New("parameter startStopScheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{startStopScheduleName}", url.PathEscape(string(startStopScheduleName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *StartStopManagedInstanceSchedulesClient) createOrUpdateHandleResponse(resp *http.Response) (StartStopManagedInstanceSchedulesClientCreateOrUpdateResponse, error) {
	result := StartStopManagedInstanceSchedulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StartStopManagedInstanceSchedule); err != nil {
		return StartStopManagedInstanceSchedulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the managed instance's Start/Stop schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - startStopScheduleName - Name of the managed instance Start/Stop schedule.
//   - options - StartStopManagedInstanceSchedulesClientDeleteOptions contains the optional parameters for the StartStopManagedInstanceSchedulesClient.Delete
//     method.
func (client *StartStopManagedInstanceSchedulesClient) Delete(ctx context.Context, resourceGroupName string, managedInstanceName string, startStopScheduleName StartStopScheduleName, options *StartStopManagedInstanceSchedulesClientDeleteOptions) (StartStopManagedInstanceSchedulesClientDeleteResponse, error) {
	var err error
	const operationName = "StartStopManagedInstanceSchedulesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, startStopScheduleName, options)
	if err != nil {
		return StartStopManagedInstanceSchedulesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StartStopManagedInstanceSchedulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return StartStopManagedInstanceSchedulesClientDeleteResponse{}, err
	}
	return StartStopManagedInstanceSchedulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StartStopManagedInstanceSchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, startStopScheduleName StartStopScheduleName, _ *StartStopManagedInstanceSchedulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/startStopSchedules/{startStopScheduleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if startStopScheduleName == "" {
		return nil, errors.New("parameter startStopScheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{startStopScheduleName}", url.PathEscape(string(startStopScheduleName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the managed instance's Start/Stop schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - startStopScheduleName - Name of the managed instance Start/Stop schedule.
//   - options - StartStopManagedInstanceSchedulesClientGetOptions contains the optional parameters for the StartStopManagedInstanceSchedulesClient.Get
//     method.
func (client *StartStopManagedInstanceSchedulesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, startStopScheduleName StartStopScheduleName, options *StartStopManagedInstanceSchedulesClientGetOptions) (StartStopManagedInstanceSchedulesClientGetResponse, error) {
	var err error
	const operationName = "StartStopManagedInstanceSchedulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, startStopScheduleName, options)
	if err != nil {
		return StartStopManagedInstanceSchedulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StartStopManagedInstanceSchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StartStopManagedInstanceSchedulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *StartStopManagedInstanceSchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, startStopScheduleName StartStopScheduleName, _ *StartStopManagedInstanceSchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/startStopSchedules/{startStopScheduleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if startStopScheduleName == "" {
		return nil, errors.New("parameter startStopScheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{startStopScheduleName}", url.PathEscape(string(startStopScheduleName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StartStopManagedInstanceSchedulesClient) getHandleResponse(resp *http.Response) (StartStopManagedInstanceSchedulesClientGetResponse, error) {
	result := StartStopManagedInstanceSchedulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StartStopManagedInstanceSchedule); err != nil {
		return StartStopManagedInstanceSchedulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInstancePager - Lists the managed instance's Start/Stop schedules.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - StartStopManagedInstanceSchedulesClientListByInstanceOptions contains the optional parameters for the StartStopManagedInstanceSchedulesClient.NewListByInstancePager
//     method.
func (client *StartStopManagedInstanceSchedulesClient) NewListByInstancePager(resourceGroupName string, managedInstanceName string, options *StartStopManagedInstanceSchedulesClientListByInstanceOptions) *runtime.Pager[StartStopManagedInstanceSchedulesClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[StartStopManagedInstanceSchedulesClientListByInstanceResponse]{
		More: func(page StartStopManagedInstanceSchedulesClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StartStopManagedInstanceSchedulesClientListByInstanceResponse) (StartStopManagedInstanceSchedulesClientListByInstanceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StartStopManagedInstanceSchedulesClient.NewListByInstancePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			}, nil)
			if err != nil {
				return StartStopManagedInstanceSchedulesClientListByInstanceResponse{}, err
			}
			return client.listByInstanceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *StartStopManagedInstanceSchedulesClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, _ *StartStopManagedInstanceSchedulesClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/startStopSchedules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *StartStopManagedInstanceSchedulesClient) listByInstanceHandleResponse(resp *http.Response) (StartStopManagedInstanceSchedulesClientListByInstanceResponse, error) {
	result := StartStopManagedInstanceSchedulesClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StartStopManagedInstanceScheduleListResult); err != nil {
		return StartStopManagedInstanceSchedulesClientListByInstanceResponse{}, err
	}
	return result, nil
}
