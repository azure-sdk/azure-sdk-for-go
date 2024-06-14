//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armverifiedid

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

// AuthoritiesClient contains the methods for the Authorities group.
// Don't use this type directly, use NewAuthoritiesClient() instead.
type AuthoritiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAuthoritiesClient creates a new instance of AuthoritiesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAuthoritiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AuthoritiesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AuthoritiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a Authority
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-26-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - authorityName - The ID of the authority
//   - resource - Resource create parameters.
//   - options - AuthoritiesClientBeginCreateOrUpdateOptions contains the optional parameters for the AuthoritiesClient.BeginCreateOrUpdate
//     method.
func (client *AuthoritiesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, authorityName string, resource Authority, options *AuthoritiesClientBeginCreateOrUpdateOptions) (*runtime.Poller[AuthoritiesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, authorityName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AuthoritiesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AuthoritiesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a Authority
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-26-preview
func (client *AuthoritiesClient) createOrUpdate(ctx context.Context, resourceGroupName string, authorityName string, resource Authority, options *AuthoritiesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AuthoritiesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, authorityName, resource, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AuthoritiesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, authorityName string, resource Authority, options *AuthoritiesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VerifiedId/authorities/{authorityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if authorityName == "" {
		return nil, errors.New("parameter authorityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorityName}", url.PathEscape(authorityName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-26-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete a Authority
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-26-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - authorityName - The ID of the authority
//   - options - AuthoritiesClientDeleteOptions contains the optional parameters for the AuthoritiesClient.Delete method.
func (client *AuthoritiesClient) Delete(ctx context.Context, resourceGroupName string, authorityName string, options *AuthoritiesClientDeleteOptions) (AuthoritiesClientDeleteResponse, error) {
	var err error
	const operationName = "AuthoritiesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, authorityName, options)
	if err != nil {
		return AuthoritiesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AuthoritiesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AuthoritiesClientDeleteResponse{}, err
	}
	return AuthoritiesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AuthoritiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, authorityName string, options *AuthoritiesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VerifiedId/authorities/{authorityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if authorityName == "" {
		return nil, errors.New("parameter authorityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorityName}", url.PathEscape(authorityName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-26-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Authority
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-26-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - authorityName - The ID of the authority
//   - options - AuthoritiesClientGetOptions contains the optional parameters for the AuthoritiesClient.Get method.
func (client *AuthoritiesClient) Get(ctx context.Context, resourceGroupName string, authorityName string, options *AuthoritiesClientGetOptions) (AuthoritiesClientGetResponse, error) {
	var err error
	const operationName = "AuthoritiesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, authorityName, options)
	if err != nil {
		return AuthoritiesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AuthoritiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AuthoritiesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AuthoritiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, authorityName string, options *AuthoritiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VerifiedId/authorities/{authorityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if authorityName == "" {
		return nil, errors.New("parameter authorityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorityName}", url.PathEscape(authorityName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-26-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AuthoritiesClient) getHandleResponse(resp *http.Response) (AuthoritiesClientGetResponse, error) {
	result := AuthoritiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Authority); err != nil {
		return AuthoritiesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List Authority resources by resource group
//
// Generated from API version 2024-01-26-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - AuthoritiesClientListByResourceGroupOptions contains the optional parameters for the AuthoritiesClient.NewListByResourceGroupPager
//     method.
func (client *AuthoritiesClient) NewListByResourceGroupPager(resourceGroupName string, options *AuthoritiesClientListByResourceGroupOptions) *runtime.Pager[AuthoritiesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AuthoritiesClientListByResourceGroupResponse]{
		More: func(page AuthoritiesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AuthoritiesClientListByResourceGroupResponse) (AuthoritiesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AuthoritiesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AuthoritiesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AuthoritiesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AuthoritiesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VerifiedId/authorities"
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
	reqQP.Set("api-version", "2024-01-26-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AuthoritiesClient) listByResourceGroupHandleResponse(resp *http.Response) (AuthoritiesClientListByResourceGroupResponse, error) {
	result := AuthoritiesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorityListResult); err != nil {
		return AuthoritiesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List Authority resources by subscription ID
//
// Generated from API version 2024-01-26-preview
//   - options - AuthoritiesClientListBySubscriptionOptions contains the optional parameters for the AuthoritiesClient.NewListBySubscriptionPager
//     method.
func (client *AuthoritiesClient) NewListBySubscriptionPager(options *AuthoritiesClientListBySubscriptionOptions) *runtime.Pager[AuthoritiesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AuthoritiesClientListBySubscriptionResponse]{
		More: func(page AuthoritiesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AuthoritiesClientListBySubscriptionResponse) (AuthoritiesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AuthoritiesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AuthoritiesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AuthoritiesClient) listBySubscriptionCreateRequest(ctx context.Context, options *AuthoritiesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VerifiedId/authorities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-26-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AuthoritiesClient) listBySubscriptionHandleResponse(resp *http.Response) (AuthoritiesClientListBySubscriptionResponse, error) {
	result := AuthoritiesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorityListResult); err != nil {
		return AuthoritiesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a Authority
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-26-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - authorityName - The ID of the authority
//   - properties - The resource properties to be updated.
//   - options - AuthoritiesClientUpdateOptions contains the optional parameters for the AuthoritiesClient.Update method.
func (client *AuthoritiesClient) Update(ctx context.Context, resourceGroupName string, authorityName string, properties AuthorityUpdate, options *AuthoritiesClientUpdateOptions) (AuthoritiesClientUpdateResponse, error) {
	var err error
	const operationName = "AuthoritiesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, authorityName, properties, options)
	if err != nil {
		return AuthoritiesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AuthoritiesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AuthoritiesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AuthoritiesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, authorityName string, properties AuthorityUpdate, options *AuthoritiesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VerifiedId/authorities/{authorityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if authorityName == "" {
		return nil, errors.New("parameter authorityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorityName}", url.PathEscape(authorityName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-26-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AuthoritiesClient) updateHandleResponse(resp *http.Response) (AuthoritiesClientUpdateResponse, error) {
	result := AuthoritiesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Authority); err != nil {
		return AuthoritiesClientUpdateResponse{}, err
	}
	return result, nil
}