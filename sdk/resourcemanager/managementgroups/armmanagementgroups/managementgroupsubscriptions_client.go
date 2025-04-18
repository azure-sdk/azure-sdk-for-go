// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

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

// ManagementGroupSubscriptionsClient contains the methods for the ManagementGroupSubscriptions group.
// Don't use this type directly, use NewManagementGroupSubscriptionsClient() instead.
type ManagementGroupSubscriptionsClient struct {
	internal *arm.Client
}

// NewManagementGroupSubscriptionsClient creates a new instance of ManagementGroupSubscriptionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementGroupSubscriptionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementGroupSubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagementGroupSubscriptionsClient{
		internal: cl,
	}
	return client, nil
}

// Create - Associates existing subscription with the management group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - groupID - Management Group ID.
//   - subscriptionID - Subscription ID.
//   - options - ManagementGroupSubscriptionsClientCreateOptions contains the optional parameters for the ManagementGroupSubscriptionsClient.Create
//     method.
func (client *ManagementGroupSubscriptionsClient) Create(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientCreateOptions) (ManagementGroupSubscriptionsClientCreateResponse, error) {
	var err error
	const operationName = "ManagementGroupSubscriptionsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, groupID, subscriptionID, options)
	if err != nil {
		return ManagementGroupSubscriptionsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementGroupSubscriptionsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagementGroupSubscriptionsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ManagementGroupSubscriptionsClient) createCreateRequest(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.CacheControl != nil {
		req.Raw().Header["Cache-Control"] = []string{*options.CacheControl}
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ManagementGroupSubscriptionsClient) createHandleResponse(resp *http.Response) (ManagementGroupSubscriptionsClientCreateResponse, error) {
	result := ManagementGroupSubscriptionsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionUnderManagementGroup); err != nil {
		return ManagementGroupSubscriptionsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - De-associates subscription from the management group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - groupID - Management Group ID.
//   - subscriptionID - Subscription ID.
//   - options - ManagementGroupSubscriptionsClientDeleteOptions contains the optional parameters for the ManagementGroupSubscriptionsClient.Delete
//     method.
func (client *ManagementGroupSubscriptionsClient) Delete(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientDeleteOptions) (ManagementGroupSubscriptionsClientDeleteResponse, error) {
	var err error
	const operationName = "ManagementGroupSubscriptionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, groupID, subscriptionID, options)
	if err != nil {
		return ManagementGroupSubscriptionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementGroupSubscriptionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ManagementGroupSubscriptionsClientDeleteResponse{}, err
	}
	return ManagementGroupSubscriptionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagementGroupSubscriptionsClient) deleteCreateRequest(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.CacheControl != nil {
		req.Raw().Header["Cache-Control"] = []string{*options.CacheControl}
	}
	return req, nil
}

// GetSubscription - Retrieves details about given subscription which is associated with the management group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - groupID - Management Group ID.
//   - subscriptionID - Subscription ID.
//   - options - ManagementGroupSubscriptionsClientGetSubscriptionOptions contains the optional parameters for the ManagementGroupSubscriptionsClient.GetSubscription
//     method.
func (client *ManagementGroupSubscriptionsClient) GetSubscription(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientGetSubscriptionOptions) (ManagementGroupSubscriptionsClientGetSubscriptionResponse, error) {
	var err error
	const operationName = "ManagementGroupSubscriptionsClient.GetSubscription"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getSubscriptionCreateRequest(ctx, groupID, subscriptionID, options)
	if err != nil {
		return ManagementGroupSubscriptionsClientGetSubscriptionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementGroupSubscriptionsClientGetSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagementGroupSubscriptionsClientGetSubscriptionResponse{}, err
	}
	resp, err := client.getSubscriptionHandleResponse(httpResp)
	return resp, err
}

// getSubscriptionCreateRequest creates the GetSubscription request.
func (client *ManagementGroupSubscriptionsClient) getSubscriptionCreateRequest(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientGetSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.CacheControl != nil {
		req.Raw().Header["Cache-Control"] = []string{*options.CacheControl}
	}
	return req, nil
}

// getSubscriptionHandleResponse handles the GetSubscription response.
func (client *ManagementGroupSubscriptionsClient) getSubscriptionHandleResponse(resp *http.Response) (ManagementGroupSubscriptionsClientGetSubscriptionResponse, error) {
	result := ManagementGroupSubscriptionsClientGetSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionUnderManagementGroup); err != nil {
		return ManagementGroupSubscriptionsClientGetSubscriptionResponse{}, err
	}
	return result, nil
}

// NewGetSubscriptionsUnderManagementGroupPager - Retrieves details about all subscriptions which are associated with the
// management group.
//
// Generated from API version 2023-04-01
//   - groupID - Management Group ID.
//   - options - ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupOptions contains the optional parameters
//     for the ManagementGroupSubscriptionsClient.NewGetSubscriptionsUnderManagementGroupPager method.
func (client *ManagementGroupSubscriptionsClient) NewGetSubscriptionsUnderManagementGroupPager(groupID string, options *ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupOptions) *runtime.Pager[ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse]{
		More: func(page ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse) (ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagementGroupSubscriptionsClient.NewGetSubscriptionsUnderManagementGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.getSubscriptionsUnderManagementGroupCreateRequest(ctx, groupID, options)
			}, nil)
			if err != nil {
				return ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse{}, err
			}
			return client.getSubscriptionsUnderManagementGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// getSubscriptionsUnderManagementGroupCreateRequest creates the GetSubscriptionsUnderManagementGroup request.
func (client *ManagementGroupSubscriptionsClient) getSubscriptionsUnderManagementGroupCreateRequest(ctx context.Context, groupID string, options *ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSubscriptionsUnderManagementGroupHandleResponse handles the GetSubscriptionsUnderManagementGroup response.
func (client *ManagementGroupSubscriptionsClient) getSubscriptionsUnderManagementGroupHandleResponse(resp *http.Response) (ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse, error) {
	result := ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListSubscriptionUnderManagementGroup); err != nil {
		return ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse{}, err
	}
	return result, nil
}
