//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ScalingPlansClient contains the methods for the ScalingPlans group.
// Don't use this type directly, use NewScalingPlansClient() instead.
type ScalingPlansClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewScalingPlansClient creates a new instance of ScalingPlansClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScalingPlansClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ScalingPlansClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScalingPlansClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create or update a scaling plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scalingPlanName - The name of the scaling plan.
//   - scalingPlan - Object containing scaling plan definitions.
//   - options - ScalingPlansClientCreateOptions contains the optional parameters for the ScalingPlansClient.Create method.
func (client *ScalingPlansClient) Create(ctx context.Context, resourceGroupName string, scalingPlanName string, scalingPlan ScalingPlan, options *ScalingPlansClientCreateOptions) (ScalingPlansClientCreateResponse, error) {
	var err error
	const operationName = "ScalingPlansClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, scalingPlanName, scalingPlan, options)
	if err != nil {
		return ScalingPlansClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalingPlansClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ScalingPlansClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ScalingPlansClient) createCreateRequest(ctx context.Context, resourceGroupName string, scalingPlanName string, scalingPlan ScalingPlan, options *ScalingPlansClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/scalingPlans/{scalingPlanName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, scalingPlan); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ScalingPlansClient) createHandleResponse(resp *http.Response) (ScalingPlansClientCreateResponse, error) {
	result := ScalingPlansClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScalingPlan); err != nil {
		return ScalingPlansClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Remove a scaling plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scalingPlanName - The name of the scaling plan.
//   - options - ScalingPlansClientDeleteOptions contains the optional parameters for the ScalingPlansClient.Delete method.
func (client *ScalingPlansClient) Delete(ctx context.Context, resourceGroupName string, scalingPlanName string, options *ScalingPlansClientDeleteOptions) (ScalingPlansClientDeleteResponse, error) {
	var err error
	const operationName = "ScalingPlansClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, scalingPlanName, options)
	if err != nil {
		return ScalingPlansClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalingPlansClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScalingPlansClientDeleteResponse{}, err
	}
	return ScalingPlansClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ScalingPlansClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, scalingPlanName string, options *ScalingPlansClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/scalingPlans/{scalingPlanName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a scaling plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scalingPlanName - The name of the scaling plan.
//   - options - ScalingPlansClientGetOptions contains the optional parameters for the ScalingPlansClient.Get method.
func (client *ScalingPlansClient) Get(ctx context.Context, resourceGroupName string, scalingPlanName string, options *ScalingPlansClientGetOptions) (ScalingPlansClientGetResponse, error) {
	var err error
	const operationName = "ScalingPlansClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, scalingPlanName, options)
	if err != nil {
		return ScalingPlansClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalingPlansClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScalingPlansClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ScalingPlansClient) getCreateRequest(ctx context.Context, resourceGroupName string, scalingPlanName string, options *ScalingPlansClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/scalingPlans/{scalingPlanName}"
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
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ScalingPlansClient) getHandleResponse(resp *http.Response) (ScalingPlansClientGetResponse, error) {
	result := ScalingPlansClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScalingPlan); err != nil {
		return ScalingPlansClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHostPoolPager - List scaling plan associated with hostpool.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - ScalingPlansClientListByHostPoolOptions contains the optional parameters for the ScalingPlansClient.NewListByHostPoolPager
//     method.
func (client *ScalingPlansClient) NewListByHostPoolPager(resourceGroupName string, hostPoolName string, options *ScalingPlansClientListByHostPoolOptions) *runtime.Pager[ScalingPlansClientListByHostPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScalingPlansClientListByHostPoolResponse]{
		More: func(page ScalingPlansClientListByHostPoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScalingPlansClientListByHostPoolResponse) (ScalingPlansClientListByHostPoolResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScalingPlansClient.NewListByHostPoolPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByHostPoolCreateRequest(ctx, resourceGroupName, hostPoolName, options)
			}, nil)
			if err != nil {
				return ScalingPlansClientListByHostPoolResponse{}, err
			}
			return client.listByHostPoolHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByHostPoolCreateRequest creates the ListByHostPool request.
func (client *ScalingPlansClient) listByHostPoolCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *ScalingPlansClientListByHostPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/scalingPlans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
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

// listByHostPoolHandleResponse handles the ListByHostPool response.
func (client *ScalingPlansClient) listByHostPoolHandleResponse(resp *http.Response) (ScalingPlansClientListByHostPoolResponse, error) {
	result := ScalingPlansClientListByHostPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScalingPlanList); err != nil {
		return ScalingPlansClientListByHostPoolResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List scaling plans.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ScalingPlansClientListByResourceGroupOptions contains the optional parameters for the ScalingPlansClient.NewListByResourceGroupPager
//     method.
func (client *ScalingPlansClient) NewListByResourceGroupPager(resourceGroupName string, options *ScalingPlansClientListByResourceGroupOptions) *runtime.Pager[ScalingPlansClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScalingPlansClientListByResourceGroupResponse]{
		More: func(page ScalingPlansClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScalingPlansClientListByResourceGroupResponse) (ScalingPlansClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScalingPlansClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ScalingPlansClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ScalingPlansClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ScalingPlansClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/scalingPlans"
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
	reqQP.Set("api-version", "2023-11-01-preview")
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ScalingPlansClient) listByResourceGroupHandleResponse(resp *http.Response) (ScalingPlansClientListByResourceGroupResponse, error) {
	result := ScalingPlansClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScalingPlanList); err != nil {
		return ScalingPlansClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List scaling plans in subscription.
//
// Generated from API version 2023-11-01-preview
//   - options - ScalingPlansClientListBySubscriptionOptions contains the optional parameters for the ScalingPlansClient.NewListBySubscriptionPager
//     method.
func (client *ScalingPlansClient) NewListBySubscriptionPager(options *ScalingPlansClientListBySubscriptionOptions) *runtime.Pager[ScalingPlansClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScalingPlansClientListBySubscriptionResponse]{
		More: func(page ScalingPlansClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScalingPlansClientListBySubscriptionResponse) (ScalingPlansClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScalingPlansClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ScalingPlansClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ScalingPlansClient) listBySubscriptionCreateRequest(ctx context.Context, options *ScalingPlansClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/scalingPlans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ScalingPlansClient) listBySubscriptionHandleResponse(resp *http.Response) (ScalingPlansClientListBySubscriptionResponse, error) {
	result := ScalingPlansClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScalingPlanList); err != nil {
		return ScalingPlansClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a scaling plan.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scalingPlanName - The name of the scaling plan.
//   - options - ScalingPlansClientUpdateOptions contains the optional parameters for the ScalingPlansClient.Update method.
func (client *ScalingPlansClient) Update(ctx context.Context, resourceGroupName string, scalingPlanName string, options *ScalingPlansClientUpdateOptions) (ScalingPlansClientUpdateResponse, error) {
	var err error
	const operationName = "ScalingPlansClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, scalingPlanName, options)
	if err != nil {
		return ScalingPlansClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScalingPlansClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScalingPlansClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ScalingPlansClient) updateCreateRequest(ctx context.Context, resourceGroupName string, scalingPlanName string, options *ScalingPlansClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/scalingPlans/{scalingPlanName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ScalingPlan != nil {
		if err := runtime.MarshalAsJSON(req, *options.ScalingPlan); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ScalingPlansClient) updateHandleResponse(resp *http.Response) (ScalingPlansClientUpdateResponse, error) {
	result := ScalingPlansClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScalingPlan); err != nil {
		return ScalingPlansClientUpdateResponse{}, err
	}
	return result, nil
}
