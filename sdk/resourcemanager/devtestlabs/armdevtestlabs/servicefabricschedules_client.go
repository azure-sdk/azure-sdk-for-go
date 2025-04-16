// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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

// ServiceFabricSchedulesClient contains the methods for the ServiceFabricSchedules group.
// Don't use this type directly, use NewServiceFabricSchedulesClient() instead.
type ServiceFabricSchedulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServiceFabricSchedulesClient creates a new instance of ServiceFabricSchedulesClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServiceFabricSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceFabricSchedulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceFabricSchedulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or replace an existing schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - serviceFabricName - The name of the service fabric.
//   - name - The name of the schedule.
//   - schedule - A schedule.
//   - options - ServiceFabricSchedulesClientCreateOrUpdateOptions contains the optional parameters for the ServiceFabricSchedulesClient.CreateOrUpdate
//     method.
func (client *ServiceFabricSchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, schedule Schedule, options *ServiceFabricSchedulesClientCreateOrUpdateOptions) (ServiceFabricSchedulesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ServiceFabricSchedulesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, name, schedule, options)
	if err != nil {
		return ServiceFabricSchedulesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceFabricSchedulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ServiceFabricSchedulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceFabricSchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, schedule Schedule, _ *ServiceFabricSchedulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, schedule); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServiceFabricSchedulesClient) createOrUpdateHandleResponse(resp *http.Response) (ServiceFabricSchedulesClientCreateOrUpdateResponse, error) {
	result := ServiceFabricSchedulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return ServiceFabricSchedulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - serviceFabricName - The name of the service fabric.
//   - name - The name of the schedule.
//   - options - ServiceFabricSchedulesClientDeleteOptions contains the optional parameters for the ServiceFabricSchedulesClient.Delete
//     method.
func (client *ServiceFabricSchedulesClient) Delete(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesClientDeleteOptions) (ServiceFabricSchedulesClientDeleteResponse, error) {
	var err error
	const operationName = "ServiceFabricSchedulesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, name, options)
	if err != nil {
		return ServiceFabricSchedulesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceFabricSchedulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceFabricSchedulesClientDeleteResponse{}, err
	}
	return ServiceFabricSchedulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceFabricSchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, _ *ServiceFabricSchedulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExecute - Execute a schedule. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - serviceFabricName - The name of the service fabric.
//   - name - The name of the schedule.
//   - options - ServiceFabricSchedulesClientBeginExecuteOptions contains the optional parameters for the ServiceFabricSchedulesClient.BeginExecute
//     method.
func (client *ServiceFabricSchedulesClient) BeginExecute(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesClientBeginExecuteOptions) (*runtime.Poller[ServiceFabricSchedulesClientExecuteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.execute(ctx, resourceGroupName, labName, userName, serviceFabricName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServiceFabricSchedulesClientExecuteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServiceFabricSchedulesClientExecuteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Execute - Execute a schedule. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *ServiceFabricSchedulesClient) execute(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesClientBeginExecuteOptions) (*http.Response, error) {
	var err error
	const operationName = "ServiceFabricSchedulesClient.BeginExecute"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.executeCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, name, options)
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

// executeCreateRequest creates the Execute request.
func (client *ServiceFabricSchedulesClient) executeCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, _ *ServiceFabricSchedulesClientBeginExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}/execute"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - serviceFabricName - The name of the service fabric.
//   - name - The name of the schedule.
//   - options - ServiceFabricSchedulesClientGetOptions contains the optional parameters for the ServiceFabricSchedulesClient.Get
//     method.
func (client *ServiceFabricSchedulesClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesClientGetOptions) (ServiceFabricSchedulesClientGetResponse, error) {
	var err error
	const operationName = "ServiceFabricSchedulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, name, options)
	if err != nil {
		return ServiceFabricSchedulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceFabricSchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceFabricSchedulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServiceFabricSchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, options *ServiceFabricSchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceFabricSchedulesClient) getHandleResponse(resp *http.Response) (ServiceFabricSchedulesClientGetResponse, error) {
	result := ServiceFabricSchedulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return ServiceFabricSchedulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List schedules in a given service fabric.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - serviceFabricName - The name of the service fabric.
//   - options - ServiceFabricSchedulesClientListOptions contains the optional parameters for the ServiceFabricSchedulesClient.NewListPager
//     method.
func (client *ServiceFabricSchedulesClient) NewListPager(resourceGroupName string, labName string, userName string, serviceFabricName string, options *ServiceFabricSchedulesClientListOptions) *runtime.Pager[ServiceFabricSchedulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceFabricSchedulesClientListResponse]{
		More: func(page ServiceFabricSchedulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceFabricSchedulesClientListResponse) (ServiceFabricSchedulesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServiceFabricSchedulesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, options)
			}, nil)
			if err != nil {
				return ServiceFabricSchedulesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ServiceFabricSchedulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, options *ServiceFabricSchedulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceFabricSchedulesClient) listHandleResponse(resp *http.Response) (ServiceFabricSchedulesClientListResponse, error) {
	result := ServiceFabricSchedulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduleList); err != nil {
		return ServiceFabricSchedulesClientListResponse{}, err
	}
	return result, nil
}

// Update - Allows modifying tags of schedules. All other properties will be ignored.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - serviceFabricName - The name of the service fabric.
//   - name - The name of the schedule.
//   - schedule - A schedule.
//   - options - ServiceFabricSchedulesClientUpdateOptions contains the optional parameters for the ServiceFabricSchedulesClient.Update
//     method.
func (client *ServiceFabricSchedulesClient) Update(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, schedule ScheduleFragment, options *ServiceFabricSchedulesClientUpdateOptions) (ServiceFabricSchedulesClientUpdateResponse, error) {
	var err error
	const operationName = "ServiceFabricSchedulesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, userName, serviceFabricName, name, schedule, options)
	if err != nil {
		return ServiceFabricSchedulesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceFabricSchedulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceFabricSchedulesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ServiceFabricSchedulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, schedule ScheduleFragment, _ *ServiceFabricSchedulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if serviceFabricName == "" {
		return nil, errors.New("parameter serviceFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceFabricName}", url.PathEscape(serviceFabricName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, schedule); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ServiceFabricSchedulesClient) updateHandleResponse(resp *http.Response) (ServiceFabricSchedulesClientUpdateResponse, error) {
	result := ServiceFabricSchedulesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return ServiceFabricSchedulesClientUpdateResponse{}, err
	}
	return result, nil
}
