//go:build go1.18
// +build go1.18

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

// CustomEntityStoreAssignmentsClient contains the methods for the CustomEntityStoreAssignments group.
// Don't use this type directly, use NewCustomEntityStoreAssignmentsClient() instead.
type CustomEntityStoreAssignmentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCustomEntityStoreAssignmentsClient creates a new instance of CustomEntityStoreAssignmentsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCustomEntityStoreAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CustomEntityStoreAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CustomEntityStoreAssignmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - [DEPRECATED NOTICE: This operation will soon be removed] Creates a custom entity store assignment for the provided
// subscription, if not already exists.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - customEntityStoreAssignmentName - Name of the custom entity store assignment. Generated name is GUID.
//   - customEntityStoreAssignmentRequestBody - Custom entity store assignment body
//   - options - CustomEntityStoreAssignmentsClientCreateOptions contains the optional parameters for the CustomEntityStoreAssignmentsClient.Create
//     method.
func (client *CustomEntityStoreAssignmentsClient) Create(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, customEntityStoreAssignmentRequestBody CustomEntityStoreAssignmentRequest, options *CustomEntityStoreAssignmentsClientCreateOptions) (CustomEntityStoreAssignmentsClientCreateResponse, error) {
	var err error
	const operationName = "CustomEntityStoreAssignmentsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, customEntityStoreAssignmentName, customEntityStoreAssignmentRequestBody, options)
	if err != nil {
		return CustomEntityStoreAssignmentsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomEntityStoreAssignmentsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return CustomEntityStoreAssignmentsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *CustomEntityStoreAssignmentsClient) createCreateRequest(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, customEntityStoreAssignmentRequestBody CustomEntityStoreAssignmentRequest, options *CustomEntityStoreAssignmentsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customEntityStoreAssignments/{customEntityStoreAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customEntityStoreAssignmentName == "" {
		return nil, errors.New("parameter customEntityStoreAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customEntityStoreAssignmentName}", url.PathEscape(customEntityStoreAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, customEntityStoreAssignmentRequestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *CustomEntityStoreAssignmentsClient) createHandleResponse(resp *http.Response) (CustomEntityStoreAssignmentsClientCreateResponse, error) {
	result := CustomEntityStoreAssignmentsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomEntityStoreAssignment); err != nil {
		return CustomEntityStoreAssignmentsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - [DEPRECATED NOTICE: This operation will soon be removed] Delete a custom entity store assignment by name for a
// provided subscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - customEntityStoreAssignmentName - Name of the custom entity store assignment. Generated name is GUID.
//   - options - CustomEntityStoreAssignmentsClientDeleteOptions contains the optional parameters for the CustomEntityStoreAssignmentsClient.Delete
//     method.
func (client *CustomEntityStoreAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, options *CustomEntityStoreAssignmentsClientDeleteOptions) (CustomEntityStoreAssignmentsClientDeleteResponse, error) {
	var err error
	const operationName = "CustomEntityStoreAssignmentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, customEntityStoreAssignmentName, options)
	if err != nil {
		return CustomEntityStoreAssignmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomEntityStoreAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CustomEntityStoreAssignmentsClientDeleteResponse{}, err
	}
	return CustomEntityStoreAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomEntityStoreAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, options *CustomEntityStoreAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customEntityStoreAssignments/{customEntityStoreAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customEntityStoreAssignmentName == "" {
		return nil, errors.New("parameter customEntityStoreAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customEntityStoreAssignmentName}", url.PathEscape(customEntityStoreAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - [DEPRECATED NOTICE: This operation will soon be removed] Gets a single custom entity store assignment by name for
// the provided subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - customEntityStoreAssignmentName - Name of the custom entity store assignment. Generated name is GUID.
//   - options - CustomEntityStoreAssignmentsClientGetOptions contains the optional parameters for the CustomEntityStoreAssignmentsClient.Get
//     method.
func (client *CustomEntityStoreAssignmentsClient) Get(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, options *CustomEntityStoreAssignmentsClientGetOptions) (CustomEntityStoreAssignmentsClientGetResponse, error) {
	var err error
	const operationName = "CustomEntityStoreAssignmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, customEntityStoreAssignmentName, options)
	if err != nil {
		return CustomEntityStoreAssignmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomEntityStoreAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CustomEntityStoreAssignmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CustomEntityStoreAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, options *CustomEntityStoreAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customEntityStoreAssignments/{customEntityStoreAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customEntityStoreAssignmentName == "" {
		return nil, errors.New("parameter customEntityStoreAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customEntityStoreAssignmentName}", url.PathEscape(customEntityStoreAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomEntityStoreAssignmentsClient) getHandleResponse(resp *http.Response) (CustomEntityStoreAssignmentsClientGetResponse, error) {
	result := CustomEntityStoreAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomEntityStoreAssignment); err != nil {
		return CustomEntityStoreAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - [DEPRECATED NOTICE: This operation will soon be removed] List custom entity store assignments
// by a provided subscription and resource group
//
// Generated from API version 2021-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - options - CustomEntityStoreAssignmentsClientListByResourceGroupOptions contains the optional parameters for the CustomEntityStoreAssignmentsClient.NewListByResourceGroupPager
//     method.
func (client *CustomEntityStoreAssignmentsClient) NewListByResourceGroupPager(resourceGroupName string, options *CustomEntityStoreAssignmentsClientListByResourceGroupOptions) *runtime.Pager[CustomEntityStoreAssignmentsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomEntityStoreAssignmentsClientListByResourceGroupResponse]{
		More: func(page CustomEntityStoreAssignmentsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomEntityStoreAssignmentsClientListByResourceGroupResponse) (CustomEntityStoreAssignmentsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CustomEntityStoreAssignmentsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return CustomEntityStoreAssignmentsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CustomEntityStoreAssignmentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CustomEntityStoreAssignmentsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customEntityStoreAssignments"
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
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CustomEntityStoreAssignmentsClient) listByResourceGroupHandleResponse(resp *http.Response) (CustomEntityStoreAssignmentsClientListByResourceGroupResponse, error) {
	result := CustomEntityStoreAssignmentsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomEntityStoreAssignmentsListResult); err != nil {
		return CustomEntityStoreAssignmentsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - [DEPRECATED NOTICE: This operation will soon be removed] List custom entity store assignments
// by provided subscription
//
// Generated from API version 2021-07-01-preview
//   - options - CustomEntityStoreAssignmentsClientListBySubscriptionOptions contains the optional parameters for the CustomEntityStoreAssignmentsClient.NewListBySubscriptionPager
//     method.
func (client *CustomEntityStoreAssignmentsClient) NewListBySubscriptionPager(options *CustomEntityStoreAssignmentsClientListBySubscriptionOptions) *runtime.Pager[CustomEntityStoreAssignmentsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomEntityStoreAssignmentsClientListBySubscriptionResponse]{
		More: func(page CustomEntityStoreAssignmentsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomEntityStoreAssignmentsClientListBySubscriptionResponse) (CustomEntityStoreAssignmentsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CustomEntityStoreAssignmentsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CustomEntityStoreAssignmentsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CustomEntityStoreAssignmentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CustomEntityStoreAssignmentsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/customEntityStoreAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CustomEntityStoreAssignmentsClient) listBySubscriptionHandleResponse(resp *http.Response) (CustomEntityStoreAssignmentsClientListBySubscriptionResponse, error) {
	result := CustomEntityStoreAssignmentsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomEntityStoreAssignmentsListResult); err != nil {
		return CustomEntityStoreAssignmentsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
