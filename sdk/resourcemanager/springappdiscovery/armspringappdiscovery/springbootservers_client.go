//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armspringappdiscovery

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

// SpringbootserversClient contains the methods for the Springbootservers group.
// Don't use this type directly, use NewSpringbootserversClient() instead.
type SpringbootserversClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSpringbootserversClient creates a new instance of SpringbootserversClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSpringbootserversClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SpringbootserversClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SpringbootserversClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create springbootservers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - The springbootsites name.
//   - springbootserversName - The springbootservers name.
//   - springbootservers - Create a springbootservers payload.
//   - options - SpringbootserversClientCreateOrUpdateOptions contains the optional parameters for the SpringbootserversClient.CreateOrUpdate
//     method.
func (client *SpringbootserversClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, siteName string, springbootserversName string, springbootservers SpringbootserversModel, options *SpringbootserversClientCreateOrUpdateOptions) (SpringbootserversClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "SpringbootserversClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, siteName, springbootserversName, springbootservers, options)
	if err != nil {
		return SpringbootserversClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpringbootserversClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SpringbootserversClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SpringbootserversClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, siteName string, springbootserversName string, springbootservers SpringbootserversModel, options *SpringbootserversClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers/{springbootserversName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if springbootserversName == "" {
		return nil, errors.New("parameter springbootserversName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{springbootserversName}", url.PathEscape(springbootserversName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, springbootservers); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SpringbootserversClient) createOrUpdateHandleResponse(resp *http.Response) (SpringbootserversClientCreateOrUpdateResponse, error) {
	result := SpringbootserversClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversModel); err != nil {
		return SpringbootserversClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete springbootservers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - The springbootsites name.
//   - springbootserversName - The springbootservers name.
//   - options - SpringbootserversClientBeginDeleteOptions contains the optional parameters for the SpringbootserversClient.BeginDelete
//     method.
func (client *SpringbootserversClient) BeginDelete(ctx context.Context, resourceGroupName string, siteName string, springbootserversName string, options *SpringbootserversClientBeginDeleteOptions) (*runtime.Poller[SpringbootserversClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, siteName, springbootserversName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SpringbootserversClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SpringbootserversClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete springbootservers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *SpringbootserversClient) deleteOperation(ctx context.Context, resourceGroupName string, siteName string, springbootserversName string, options *SpringbootserversClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SpringbootserversClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, siteName, springbootserversName, options)
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
func (client *SpringbootserversClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, springbootserversName string, options *SpringbootserversClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers/{springbootserversName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if springbootserversName == "" {
		return nil, errors.New("parameter springbootserversName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{springbootserversName}", url.PathEscape(springbootserversName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - List springbootservers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - The springbootsites name.
//   - springbootserversName - The springbootservers name.
//   - options - SpringbootserversClientGetOptions contains the optional parameters for the SpringbootserversClient.Get method.
func (client *SpringbootserversClient) Get(ctx context.Context, resourceGroupName string, siteName string, springbootserversName string, options *SpringbootserversClientGetOptions) (SpringbootserversClientGetResponse, error) {
	var err error
	const operationName = "SpringbootserversClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, springbootserversName, options)
	if err != nil {
		return SpringbootserversClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpringbootserversClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SpringbootserversClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SpringbootserversClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, springbootserversName string, options *SpringbootserversClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers/{springbootserversName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if springbootserversName == "" {
		return nil, errors.New("parameter springbootserversName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{springbootserversName}", url.PathEscape(springbootserversName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SpringbootserversClient) getHandleResponse(resp *http.Response) (SpringbootserversClientGetResponse, error) {
	result := SpringbootserversClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversModel); err != nil {
		return SpringbootserversClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List springbootservers resource by resourceGroup
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - The springbootsites name.
//   - options - SpringbootserversClientListByResourceGroupOptions contains the optional parameters for the SpringbootserversClient.NewListByResourceGroupPager
//     method.
func (client *SpringbootserversClient) NewListByResourceGroupPager(resourceGroupName string, siteName string, options *SpringbootserversClientListByResourceGroupOptions) *runtime.Pager[SpringbootserversClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SpringbootserversClientListByResourceGroupResponse]{
		More: func(page SpringbootserversClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SpringbootserversClientListByResourceGroupResponse) (SpringbootserversClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SpringbootserversClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return SpringbootserversClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SpringbootserversClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *SpringbootserversClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SpringbootserversClient) listByResourceGroupHandleResponse(resp *http.Response) (SpringbootserversClientListByResourceGroupResponse, error) {
	result := SpringbootserversClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversListResult); err != nil {
		return SpringbootserversClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List springbootservers resource by subscription
//
// Generated from API version 2023-01-01-preview
//   - siteName - The springbootsites name.
//   - options - SpringbootserversClientListBySubscriptionOptions contains the optional parameters for the SpringbootserversClient.NewListBySubscriptionPager
//     method.
func (client *SpringbootserversClient) NewListBySubscriptionPager(siteName string, options *SpringbootserversClientListBySubscriptionOptions) *runtime.Pager[SpringbootserversClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SpringbootserversClientListBySubscriptionResponse]{
		More: func(page SpringbootserversClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SpringbootserversClientListBySubscriptionResponse) (SpringbootserversClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SpringbootserversClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, siteName, options)
			}, nil)
			if err != nil {
				return SpringbootserversClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SpringbootserversClient) listBySubscriptionCreateRequest(ctx context.Context, siteName string, options *SpringbootserversClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SpringbootserversClient) listBySubscriptionHandleResponse(resp *http.Response) (SpringbootserversClientListBySubscriptionResponse, error) {
	result := SpringbootserversClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversListResult); err != nil {
		return SpringbootserversClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update springbootservers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - The springbootsites name.
//   - springbootserversName - The springbootservers name.
//   - springbootservers - Update a springbootservers payload.
//   - options - SpringbootserversClientUpdateOptions contains the optional parameters for the SpringbootserversClient.Update
//     method.
func (client *SpringbootserversClient) Update(ctx context.Context, resourceGroupName string, siteName string, springbootserversName string, springbootservers SpringbootserversPatch, options *SpringbootserversClientUpdateOptions) (SpringbootserversClientUpdateResponse, error) {
	var err error
	const operationName = "SpringbootserversClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, siteName, springbootserversName, springbootservers, options)
	if err != nil {
		return SpringbootserversClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpringbootserversClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SpringbootserversClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SpringbootserversClient) updateCreateRequest(ctx context.Context, resourceGroupName string, siteName string, springbootserversName string, springbootservers SpringbootserversPatch, options *SpringbootserversClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers/{springbootserversName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if springbootserversName == "" {
		return nil, errors.New("parameter springbootserversName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{springbootserversName}", url.PathEscape(springbootserversName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, springbootservers); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SpringbootserversClient) updateHandleResponse(resp *http.Response) (SpringbootserversClientUpdateResponse, error) {
	result := SpringbootserversClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversModel); err != nil {
		return SpringbootserversClientUpdateResponse{}, err
	}
	return result, nil
}
