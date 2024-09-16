//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// AuthorizationServerClient contains the methods for the AuthorizationServer group.
// Don't use this type directly, use NewAuthorizationServerClient() instead.
type AuthorizationServerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAuthorizationServerClient creates a new instance of AuthorizationServerClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAuthorizationServerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AuthorizationServerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AuthorizationServerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates new authorization server or updates an existing authorization server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - authsid - Identifier of the authorization server.
//   - parameters - Create or update parameters.
//   - options - AuthorizationServerClientCreateOrUpdateOptions contains the optional parameters for the AuthorizationServerClient.CreateOrUpdate
//     method.
func (client *AuthorizationServerClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters AuthorizationServerContract, options *AuthorizationServerClientCreateOrUpdateOptions) (AuthorizationServerClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "AuthorizationServerClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, authsid, parameters, options)
	if err != nil {
		return AuthorizationServerClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AuthorizationServerClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return AuthorizationServerClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AuthorizationServerClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters AuthorizationServerContract, options *AuthorizationServerClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if authsid == "" {
		return nil, errors.New("parameter authsid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authsid}", url.PathEscape(authsid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AuthorizationServerClient) createOrUpdateHandleResponse(resp *http.Response) (AuthorizationServerClientCreateOrUpdateResponse, error) {
	result := AuthorizationServerClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationServerContract); err != nil {
		return AuthorizationServerClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specific authorization server instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - authsid - Identifier of the authorization server.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - AuthorizationServerClientDeleteOptions contains the optional parameters for the AuthorizationServerClient.Delete
//     method.
func (client *AuthorizationServerClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, authsid string, ifMatch string, options *AuthorizationServerClientDeleteOptions) (AuthorizationServerClientDeleteResponse, error) {
	var err error
	const operationName = "AuthorizationServerClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, authsid, ifMatch, options)
	if err != nil {
		return AuthorizationServerClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AuthorizationServerClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AuthorizationServerClientDeleteResponse{}, err
	}
	return AuthorizationServerClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AuthorizationServerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, authsid string, ifMatch string, options *AuthorizationServerClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if authsid == "" {
		return nil, errors.New("parameter authsid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authsid}", url.PathEscape(authsid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	return req, nil
}

// Get - Gets the details of the authorization server specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - authsid - Identifier of the authorization server.
//   - options - AuthorizationServerClientGetOptions contains the optional parameters for the AuthorizationServerClient.Get method.
func (client *AuthorizationServerClient) Get(ctx context.Context, resourceGroupName string, serviceName string, authsid string, options *AuthorizationServerClientGetOptions) (AuthorizationServerClientGetResponse, error) {
	var err error
	const operationName = "AuthorizationServerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, authsid, options)
	if err != nil {
		return AuthorizationServerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AuthorizationServerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AuthorizationServerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AuthorizationServerClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, authsid string, options *AuthorizationServerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if authsid == "" {
		return nil, errors.New("parameter authsid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authsid}", url.PathEscape(authsid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AuthorizationServerClient) getHandleResponse(resp *http.Response) (AuthorizationServerClientGetResponse, error) {
	result := AuthorizationServerClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationServerContract); err != nil {
		return AuthorizationServerClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the authorizationServer specified by its identifier.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - authsid - Identifier of the authorization server.
//   - options - AuthorizationServerClientGetEntityTagOptions contains the optional parameters for the AuthorizationServerClient.GetEntityTag
//     method.
func (client *AuthorizationServerClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, authsid string, options *AuthorizationServerClientGetEntityTagOptions) (AuthorizationServerClientGetEntityTagResponse, error) {
	var err error
	const operationName = "AuthorizationServerClient.GetEntityTag"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, authsid, options)
	if err != nil {
		return AuthorizationServerClientGetEntityTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AuthorizationServerClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AuthorizationServerClientGetEntityTagResponse{}, err
	}
	resp, err := client.getEntityTagHandleResponse(httpResp)
	return resp, err
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *AuthorizationServerClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, authsid string, options *AuthorizationServerClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if authsid == "" {
		return nil, errors.New("parameter authsid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authsid}", url.PathEscape(authsid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *AuthorizationServerClient) getEntityTagHandleResponse(resp *http.Response) (AuthorizationServerClientGetEntityTagResponse, error) {
	result := AuthorizationServerClientGetEntityTagResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}

// NewListByServicePager - Lists a collection of authorization servers defined within a service instance.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - AuthorizationServerClientListByServiceOptions contains the optional parameters for the AuthorizationServerClient.NewListByServicePager
//     method.
func (client *AuthorizationServerClient) NewListByServicePager(resourceGroupName string, serviceName string, options *AuthorizationServerClientListByServiceOptions) *runtime.Pager[AuthorizationServerClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[AuthorizationServerClientListByServiceResponse]{
		More: func(page AuthorizationServerClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AuthorizationServerClientListByServiceResponse) (AuthorizationServerClientListByServiceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AuthorizationServerClient.NewListByServicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			}, nil)
			if err != nil {
				return AuthorizationServerClientListByServiceResponse{}, err
			}
			return client.listByServiceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *AuthorizationServerClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *AuthorizationServerClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *AuthorizationServerClient) listByServiceHandleResponse(resp *http.Response) (AuthorizationServerClientListByServiceResponse, error) {
	result := AuthorizationServerClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationServerCollection); err != nil {
		return AuthorizationServerClientListByServiceResponse{}, err
	}
	return result, nil
}

// ListSecrets - Gets the client secret details of the authorization server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - authsid - Identifier of the authorization server.
//   - options - AuthorizationServerClientListSecretsOptions contains the optional parameters for the AuthorizationServerClient.ListSecrets
//     method.
func (client *AuthorizationServerClient) ListSecrets(ctx context.Context, resourceGroupName string, serviceName string, authsid string, options *AuthorizationServerClientListSecretsOptions) (AuthorizationServerClientListSecretsResponse, error) {
	var err error
	const operationName = "AuthorizationServerClient.ListSecrets"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, serviceName, authsid, options)
	if err != nil {
		return AuthorizationServerClientListSecretsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AuthorizationServerClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AuthorizationServerClientListSecretsResponse{}, err
	}
	resp, err := client.listSecretsHandleResponse(httpResp)
	return resp, err
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *AuthorizationServerClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, authsid string, options *AuthorizationServerClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}/listSecrets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if authsid == "" {
		return nil, errors.New("parameter authsid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authsid}", url.PathEscape(authsid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *AuthorizationServerClient) listSecretsHandleResponse(resp *http.Response) (AuthorizationServerClientListSecretsResponse, error) {
	result := AuthorizationServerClientListSecretsResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationServerSecretsContract); err != nil {
		return AuthorizationServerClientListSecretsResponse{}, err
	}
	return result, nil
}

// Update - Updates the details of the authorization server specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - authsid - Identifier of the authorization server.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - parameters - OAuth2 Server settings Update parameters.
//   - options - AuthorizationServerClientUpdateOptions contains the optional parameters for the AuthorizationServerClient.Update
//     method.
func (client *AuthorizationServerClient) Update(ctx context.Context, resourceGroupName string, serviceName string, authsid string, ifMatch string, parameters AuthorizationServerUpdateContract, options *AuthorizationServerClientUpdateOptions) (AuthorizationServerClientUpdateResponse, error) {
	var err error
	const operationName = "AuthorizationServerClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, authsid, ifMatch, parameters, options)
	if err != nil {
		return AuthorizationServerClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AuthorizationServerClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AuthorizationServerClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AuthorizationServerClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, authsid string, ifMatch string, parameters AuthorizationServerUpdateContract, options *AuthorizationServerClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if authsid == "" {
		return nil, errors.New("parameter authsid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authsid}", url.PathEscape(authsid))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AuthorizationServerClient) updateHandleResponse(resp *http.Response) (AuthorizationServerClientUpdateResponse, error) {
	result := AuthorizationServerClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationServerContract); err != nil {
		return AuthorizationServerClientUpdateResponse{}, err
	}
	return result, nil
}
