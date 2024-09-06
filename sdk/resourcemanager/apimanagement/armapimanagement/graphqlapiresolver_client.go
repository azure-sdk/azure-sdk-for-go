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

// GraphQLAPIResolverClient contains the methods for the GraphQLAPIResolver group.
// Don't use this type directly, use NewGraphQLAPIResolverClient() instead.
type GraphQLAPIResolverClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGraphQLAPIResolverClient creates a new instance of GraphQLAPIResolverClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGraphQLAPIResolverClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GraphQLAPIResolverClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GraphQLAPIResolverClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new resolver in the GraphQL API or updates an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - resolverID - Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
//   - parameters - Create parameters.
//   - options - GraphQLAPIResolverClientCreateOrUpdateOptions contains the optional parameters for the GraphQLAPIResolverClient.CreateOrUpdate
//     method.
func (client *GraphQLAPIResolverClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, parameters ResolverContract, options *GraphQLAPIResolverClientCreateOrUpdateOptions) (GraphQLAPIResolverClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "GraphQLAPIResolverClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiID, resolverID, parameters, options)
	if err != nil {
		return GraphQLAPIResolverClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GraphQLAPIResolverClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return GraphQLAPIResolverClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GraphQLAPIResolverClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, parameters ResolverContract, options *GraphQLAPIResolverClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers/{resolverId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if resolverID == "" {
		return nil, errors.New("parameter resolverID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resolverId}", url.PathEscape(resolverID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
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
func (client *GraphQLAPIResolverClient) createOrUpdateHandleResponse(resp *http.Response) (GraphQLAPIResolverClientCreateOrUpdateResponse, error) {
	result := GraphQLAPIResolverClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResolverContract); err != nil {
		return GraphQLAPIResolverClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified resolver in the GraphQL API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - resolverID - Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - GraphQLAPIResolverClientDeleteOptions contains the optional parameters for the GraphQLAPIResolverClient.Delete
//     method.
func (client *GraphQLAPIResolverClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, ifMatch string, options *GraphQLAPIResolverClientDeleteOptions) (GraphQLAPIResolverClientDeleteResponse, error) {
	var err error
	const operationName = "GraphQLAPIResolverClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiID, resolverID, ifMatch, options)
	if err != nil {
		return GraphQLAPIResolverClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GraphQLAPIResolverClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return GraphQLAPIResolverClientDeleteResponse{}, err
	}
	return GraphQLAPIResolverClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GraphQLAPIResolverClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, ifMatch string, options *GraphQLAPIResolverClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers/{resolverId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if resolverID == "" {
		return nil, errors.New("parameter resolverID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resolverId}", url.PathEscape(resolverID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	return req, nil
}

// Get - Gets the details of the GraphQL API Resolver specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - resolverID - Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
//   - options - GraphQLAPIResolverClientGetOptions contains the optional parameters for the GraphQLAPIResolverClient.Get method.
func (client *GraphQLAPIResolverClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, options *GraphQLAPIResolverClientGetOptions) (GraphQLAPIResolverClientGetResponse, error) {
	var err error
	const operationName = "GraphQLAPIResolverClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiID, resolverID, options)
	if err != nil {
		return GraphQLAPIResolverClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GraphQLAPIResolverClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GraphQLAPIResolverClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GraphQLAPIResolverClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, options *GraphQLAPIResolverClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers/{resolverId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if resolverID == "" {
		return nil, errors.New("parameter resolverID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resolverId}", url.PathEscape(resolverID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GraphQLAPIResolverClient) getHandleResponse(resp *http.Response) (GraphQLAPIResolverClientGetResponse, error) {
	result := GraphQLAPIResolverClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResolverContract); err != nil {
		return GraphQLAPIResolverClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the GraphQL API resolver specified by its identifier.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - resolverID - Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
//   - options - GraphQLAPIResolverClientGetEntityTagOptions contains the optional parameters for the GraphQLAPIResolverClient.GetEntityTag
//     method.
func (client *GraphQLAPIResolverClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, options *GraphQLAPIResolverClientGetEntityTagOptions) (GraphQLAPIResolverClientGetEntityTagResponse, error) {
	var err error
	const operationName = "GraphQLAPIResolverClient.GetEntityTag"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, apiID, resolverID, options)
	if err != nil {
		return GraphQLAPIResolverClientGetEntityTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GraphQLAPIResolverClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GraphQLAPIResolverClientGetEntityTagResponse{}, err
	}
	resp, err := client.getEntityTagHandleResponse(httpResp)
	return resp, err
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *GraphQLAPIResolverClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, options *GraphQLAPIResolverClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers/{resolverId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if resolverID == "" {
		return nil, errors.New("parameter resolverID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resolverId}", url.PathEscape(resolverID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *GraphQLAPIResolverClient) getEntityTagHandleResponse(resp *http.Response) (GraphQLAPIResolverClientGetEntityTagResponse, error) {
	result := GraphQLAPIResolverClientGetEntityTagResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}

// NewListByAPIPager - Lists a collection of the resolvers for the specified GraphQL API.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - options - GraphQLAPIResolverClientListByAPIOptions contains the optional parameters for the GraphQLAPIResolverClient.NewListByAPIPager
//     method.
func (client *GraphQLAPIResolverClient) NewListByAPIPager(resourceGroupName string, serviceName string, apiID string, options *GraphQLAPIResolverClientListByAPIOptions) *runtime.Pager[GraphQLAPIResolverClientListByAPIResponse] {
	return runtime.NewPager(runtime.PagingHandler[GraphQLAPIResolverClientListByAPIResponse]{
		More: func(page GraphQLAPIResolverClientListByAPIResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GraphQLAPIResolverClientListByAPIResponse) (GraphQLAPIResolverClientListByAPIResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GraphQLAPIResolverClient.NewListByAPIPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAPICreateRequest(ctx, resourceGroupName, serviceName, apiID, options)
			}, nil)
			if err != nil {
				return GraphQLAPIResolverClientListByAPIResponse{}, err
			}
			return client.listByAPIHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAPICreateRequest creates the ListByAPI request.
func (client *GraphQLAPIResolverClient) listByAPICreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, options *GraphQLAPIResolverClientListByAPIOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
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
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAPIHandleResponse handles the ListByAPI response.
func (client *GraphQLAPIResolverClient) listByAPIHandleResponse(resp *http.Response) (GraphQLAPIResolverClientListByAPIResponse, error) {
	result := GraphQLAPIResolverClientListByAPIResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResolverCollection); err != nil {
		return GraphQLAPIResolverClientListByAPIResponse{}, err
	}
	return result, nil
}

// Update - Updates the details of the resolver in the GraphQL API specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - resolverID - Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - parameters - GraphQL API Resolver Update parameters.
//   - options - GraphQLAPIResolverClientUpdateOptions contains the optional parameters for the GraphQLAPIResolverClient.Update
//     method.
func (client *GraphQLAPIResolverClient) Update(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, ifMatch string, parameters ResolverUpdateContract, options *GraphQLAPIResolverClientUpdateOptions) (GraphQLAPIResolverClientUpdateResponse, error) {
	var err error
	const operationName = "GraphQLAPIResolverClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, apiID, resolverID, ifMatch, parameters, options)
	if err != nil {
		return GraphQLAPIResolverClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GraphQLAPIResolverClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GraphQLAPIResolverClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *GraphQLAPIResolverClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, ifMatch string, parameters ResolverUpdateContract, options *GraphQLAPIResolverClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers/{resolverId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if resolverID == "" {
		return nil, errors.New("parameter resolverID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resolverId}", url.PathEscape(resolverID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *GraphQLAPIResolverClient) updateHandleResponse(resp *http.Response) (GraphQLAPIResolverClientUpdateResponse, error) {
	result := GraphQLAPIResolverClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResolverContract); err != nil {
		return GraphQLAPIResolverClientUpdateResponse{}, err
	}
	return result, nil
}
