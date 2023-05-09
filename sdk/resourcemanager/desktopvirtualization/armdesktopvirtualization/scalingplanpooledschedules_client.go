//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdesktopvirtualization

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

// ScalingPlanPooledSchedulesClient contains the methods for the ScalingPlanPooledSchedules group.
// Don't use this type directly, use NewScalingPlanPooledSchedulesClient() instead.
type ScalingPlanPooledSchedulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewScalingPlanPooledSchedulesClient creates a new instance of ScalingPlanPooledSchedulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScalingPlanPooledSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ScalingPlanPooledSchedulesClient, error) {
	cl, err := arm.NewClient(moduleName+".ScalingPlanPooledSchedulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScalingPlanPooledSchedulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create or update a ScalingPlanPooledSchedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scalingPlanName - The name of the scaling plan.
//   - scalingPlanScheduleName - The name of the ScalingPlanSchedule
//   - scalingPlanSchedule - Object containing ScalingPlanPooledSchedule definitions.
//   - options - ScalingPlanPooledSchedulesClientCreateOptions contains the optional parameters for the ScalingPlanPooledSchedulesClient.Create
//     method.
func (client *ScalingPlanPooledSchedulesClient) Create(ctx context.Context, resourceGroupName string, scalingPlanName string, scalingPlanScheduleName string, scalingPlanSchedule ScalingPlanPooledSchedule, options *ScalingPlanPooledSchedulesClientCreateOptions) (ScalingPlanPooledSchedulesClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, scalingPlanName, scalingPlanScheduleName, scalingPlanSchedule, options)
	if err != nil {
		return ScalingPlanPooledSchedulesClientCreateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalingPlanPooledSchedulesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ScalingPlanPooledSchedulesClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ScalingPlanPooledSchedulesClient) createCreateRequest(ctx context.Context, resourceGroupName string, scalingPlanName string, scalingPlanScheduleName string, scalingPlanSchedule ScalingPlanPooledSchedule, options *ScalingPlanPooledSchedulesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/scalingPlans/{scalingPlanName}/pooledSchedules/{scalingPlanScheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scalingPlanName == "" {
		return nil, errors.New("parameter scalingPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scalingPlanName}", url.PathEscape(scalingPlanName))
	if scalingPlanScheduleName == "" {
		return nil, errors.New("parameter scalingPlanScheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scalingPlanScheduleName}", url.PathEscape(scalingPlanScheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, scalingPlanSchedule)
}

// createHandleResponse handles the Create response.
func (client *ScalingPlanPooledSchedulesClient) createHandleResponse(resp *http.Response) (ScalingPlanPooledSchedulesClientCreateResponse, error) {
	result := ScalingPlanPooledSchedulesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScalingPlanPooledSchedule); err != nil {
		return ScalingPlanPooledSchedulesClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Remove a ScalingPlanPooledSchedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scalingPlanName - The name of the scaling plan.
//   - scalingPlanScheduleName - The name of the ScalingPlanSchedule
//   - options - ScalingPlanPooledSchedulesClientDeleteOptions contains the optional parameters for the ScalingPlanPooledSchedulesClient.Delete
//     method.
func (client *ScalingPlanPooledSchedulesClient) Delete(ctx context.Context, resourceGroupName string, scalingPlanName string, scalingPlanScheduleName string, options *ScalingPlanPooledSchedulesClientDeleteOptions) (ScalingPlanPooledSchedulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, scalingPlanName, scalingPlanScheduleName, options)
	if err != nil {
		return ScalingPlanPooledSchedulesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalingPlanPooledSchedulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ScalingPlanPooledSchedulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ScalingPlanPooledSchedulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ScalingPlanPooledSchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, scalingPlanName string, scalingPlanScheduleName string, options *ScalingPlanPooledSchedulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/scalingPlans/{scalingPlanName}/pooledSchedules/{scalingPlanScheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scalingPlanName == "" {
		return nil, errors.New("parameter scalingPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scalingPlanName}", url.PathEscape(scalingPlanName))
	if scalingPlanScheduleName == "" {
		return nil, errors.New("parameter scalingPlanScheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scalingPlanScheduleName}", url.PathEscape(scalingPlanScheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a ScalingPlanPooledSchedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scalingPlanName - The name of the scaling plan.
//   - scalingPlanScheduleName - The name of the ScalingPlanSchedule
//   - options - ScalingPlanPooledSchedulesClientGetOptions contains the optional parameters for the ScalingPlanPooledSchedulesClient.Get
//     method.
func (client *ScalingPlanPooledSchedulesClient) Get(ctx context.Context, resourceGroupName string, scalingPlanName string, scalingPlanScheduleName string, options *ScalingPlanPooledSchedulesClientGetOptions) (ScalingPlanPooledSchedulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, scalingPlanName, scalingPlanScheduleName, options)
	if err != nil {
		return ScalingPlanPooledSchedulesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalingPlanPooledSchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScalingPlanPooledSchedulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ScalingPlanPooledSchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, scalingPlanName string, scalingPlanScheduleName string, options *ScalingPlanPooledSchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/scalingPlans/{scalingPlanName}/pooledSchedules/{scalingPlanScheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scalingPlanName == "" {
		return nil, errors.New("parameter scalingPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scalingPlanName}", url.PathEscape(scalingPlanName))
	if scalingPlanScheduleName == "" {
		return nil, errors.New("parameter scalingPlanScheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scalingPlanScheduleName}", url.PathEscape(scalingPlanScheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ScalingPlanPooledSchedulesClient) getHandleResponse(resp *http.Response) (ScalingPlanPooledSchedulesClientGetResponse, error) {
	result := ScalingPlanPooledSchedulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScalingPlanPooledSchedule); err != nil {
		return ScalingPlanPooledSchedulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List ScalingPlanPooledSchedules.
//
// Generated from API version 2023-01-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scalingPlanName - The name of the scaling plan.
//   - options - ScalingPlanPooledSchedulesClientListOptions contains the optional parameters for the ScalingPlanPooledSchedulesClient.NewListPager
//     method.
func (client *ScalingPlanPooledSchedulesClient) NewListPager(resourceGroupName string, scalingPlanName string, options *ScalingPlanPooledSchedulesClientListOptions) *runtime.Pager[ScalingPlanPooledSchedulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScalingPlanPooledSchedulesClientListResponse]{
		More: func(page ScalingPlanPooledSchedulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScalingPlanPooledSchedulesClientListResponse) (ScalingPlanPooledSchedulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, scalingPlanName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ScalingPlanPooledSchedulesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ScalingPlanPooledSchedulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ScalingPlanPooledSchedulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ScalingPlanPooledSchedulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, scalingPlanName string, options *ScalingPlanPooledSchedulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/scalingPlans/{scalingPlanName}/pooledSchedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scalingPlanName == "" {
		return nil, errors.New("parameter scalingPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scalingPlanName}", url.PathEscape(scalingPlanName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-30-preview")
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.IsDescending != nil {
		reqQP.Set("isDescending", strconv.FormatBool(*options.IsDescending))
	}
	if options != nil && options.InitialSkip != nil {
		reqQP.Set("initialSkip", strconv.FormatInt(int64(*options.InitialSkip), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ScalingPlanPooledSchedulesClient) listHandleResponse(resp *http.Response) (ScalingPlanPooledSchedulesClientListResponse, error) {
	result := ScalingPlanPooledSchedulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScalingPlanPooledScheduleList); err != nil {
		return ScalingPlanPooledSchedulesClientListResponse{}, err
	}
	return result, nil
}

// Update - Update a ScalingPlanPooledSchedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scalingPlanName - The name of the scaling plan.
//   - scalingPlanScheduleName - The name of the ScalingPlanSchedule
//   - options - ScalingPlanPooledSchedulesClientUpdateOptions contains the optional parameters for the ScalingPlanPooledSchedulesClient.Update
//     method.
func (client *ScalingPlanPooledSchedulesClient) Update(ctx context.Context, resourceGroupName string, scalingPlanName string, scalingPlanScheduleName string, options *ScalingPlanPooledSchedulesClientUpdateOptions) (ScalingPlanPooledSchedulesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, scalingPlanName, scalingPlanScheduleName, options)
	if err != nil {
		return ScalingPlanPooledSchedulesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalingPlanPooledSchedulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScalingPlanPooledSchedulesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ScalingPlanPooledSchedulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, scalingPlanName string, scalingPlanScheduleName string, options *ScalingPlanPooledSchedulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/scalingPlans/{scalingPlanName}/pooledSchedules/{scalingPlanScheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scalingPlanName == "" {
		return nil, errors.New("parameter scalingPlanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scalingPlanName}", url.PathEscape(scalingPlanName))
	if scalingPlanScheduleName == "" {
		return nil, errors.New("parameter scalingPlanScheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scalingPlanScheduleName}", url.PathEscape(scalingPlanScheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ScalingPlanSchedule != nil {
		return req, runtime.MarshalAsJSON(req, *options.ScalingPlanSchedule)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ScalingPlanPooledSchedulesClient) updateHandleResponse(resp *http.Response) (ScalingPlanPooledSchedulesClientUpdateResponse, error) {
	result := ScalingPlanPooledSchedulesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScalingPlanPooledSchedule); err != nil {
		return ScalingPlanPooledSchedulesClientUpdateResponse{}, err
	}
	return result, nil
}
