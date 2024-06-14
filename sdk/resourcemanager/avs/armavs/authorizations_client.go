//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// AuthorizationsClient contains the methods for the Authorizations group.
// Don't use this type directly, use NewAuthorizationsClient() instead.
type AuthorizationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAuthorizationsClient creates a new instance of AuthorizationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAuthorizationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AuthorizationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AuthorizationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a ExpressRouteAuthorization
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - authorizationName - Name of the ExpressRoute Circuit Authorization
//   - authorization - Resource create parameters.
//   - options - AuthorizationsClientBeginCreateOrUpdateOptions contains the optional parameters for the AuthorizationsClient.BeginCreateOrUpdate
//     method.
func (client *AuthorizationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, authorization ExpressRouteAuthorization, options *AuthorizationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AuthorizationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, authorizationName, authorization, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AuthorizationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AuthorizationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a ExpressRouteAuthorization
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *AuthorizationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, authorization ExpressRouteAuthorization, options *AuthorizationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AuthorizationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, authorizationName, authorization, options)
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
func (client *AuthorizationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, authorization ExpressRouteAuthorization, options *AuthorizationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/authorizations/{authorizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, authorization); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a ExpressRouteAuthorization
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - authorizationName - Name of the ExpressRoute Circuit Authorization
//   - options - AuthorizationsClientBeginDeleteOptions contains the optional parameters for the AuthorizationsClient.BeginDelete
//     method.
func (client *AuthorizationsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, options *AuthorizationsClientBeginDeleteOptions) (*runtime.Poller[AuthorizationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, authorizationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AuthorizationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AuthorizationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a ExpressRouteAuthorization
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *AuthorizationsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, options *AuthorizationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AuthorizationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, authorizationName, options)
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
func (client *AuthorizationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, options *AuthorizationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/authorizations/{authorizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a ExpressRouteAuthorization
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - authorizationName - Name of the ExpressRoute Circuit Authorization
//   - options - AuthorizationsClientGetOptions contains the optional parameters for the AuthorizationsClient.Get method.
func (client *AuthorizationsClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, options *AuthorizationsClientGetOptions) (AuthorizationsClientGetResponse, error) {
	var err error
	const operationName = "AuthorizationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, authorizationName, options)
	if err != nil {
		return AuthorizationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AuthorizationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AuthorizationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AuthorizationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, options *AuthorizationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/authorizations/{authorizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AuthorizationsClient) getHandleResponse(resp *http.Response) (AuthorizationsClientGetResponse, error) {
	result := AuthorizationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteAuthorization); err != nil {
		return AuthorizationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List ExpressRouteAuthorization resources by PrivateCloud
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - options - AuthorizationsClientListOptions contains the optional parameters for the AuthorizationsClient.NewListPager method.
func (client *AuthorizationsClient) NewListPager(resourceGroupName string, privateCloudName string, options *AuthorizationsClientListOptions) *runtime.Pager[AuthorizationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AuthorizationsClientListResponse]{
		More: func(page AuthorizationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AuthorizationsClientListResponse) (AuthorizationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AuthorizationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, privateCloudName, options)
			}, nil)
			if err != nil {
				return AuthorizationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AuthorizationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *AuthorizationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/authorizations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AuthorizationsClient) listHandleResponse(resp *http.Response) (AuthorizationsClientListResponse, error) {
	result := AuthorizationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteAuthorizationList); err != nil {
		return AuthorizationsClientListResponse{}, err
	}
	return result, nil
}
