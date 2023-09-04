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

// APIIssueCommentClient contains the methods for the APIIssueComment group.
// Don't use this type directly, use NewAPIIssueCommentClient() instead.
type APIIssueCommentClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAPIIssueCommentClient creates a new instance of APIIssueCommentClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAPIIssueCommentClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIIssueCommentClient, error) {
	cl, err := arm.NewClient(moduleName+".APIIssueCommentClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &APIIssueCommentClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new Comment for the Issue in an API or updates an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - issueID - Issue identifier. Must be unique in the current API Management service instance.
//   - commentID - Comment identifier within an Issue. Must be unique in the current Issue.
//   - parameters - Create parameters.
//   - options - APIIssueCommentClientCreateOrUpdateOptions contains the optional parameters for the APIIssueCommentClient.CreateOrUpdate
//     method.
func (client *APIIssueCommentClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, commentID string, parameters IssueCommentContract, options *APIIssueCommentClientCreateOrUpdateOptions) (APIIssueCommentClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiID, issueID, commentID, parameters, options)
	if err != nil {
		return APIIssueCommentClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIIssueCommentClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return APIIssueCommentClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIIssueCommentClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, commentID string, parameters IssueCommentContract, options *APIIssueCommentClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}"
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
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
	if commentID == "" {
		return nil, errors.New("parameter commentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{commentId}", url.PathEscape(commentID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *APIIssueCommentClient) createOrUpdateHandleResponse(resp *http.Response) (APIIssueCommentClientCreateOrUpdateResponse, error) {
	result := APIIssueCommentClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.IssueCommentContract); err != nil {
		return APIIssueCommentClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified comment from an Issue.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - issueID - Issue identifier. Must be unique in the current API Management service instance.
//   - commentID - Comment identifier within an Issue. Must be unique in the current Issue.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - APIIssueCommentClientDeleteOptions contains the optional parameters for the APIIssueCommentClient.Delete method.
func (client *APIIssueCommentClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, commentID string, ifMatch string, options *APIIssueCommentClientDeleteOptions) (APIIssueCommentClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiID, issueID, commentID, ifMatch, options)
	if err != nil {
		return APIIssueCommentClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIIssueCommentClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return APIIssueCommentClientDeleteResponse{}, err
	}
	return APIIssueCommentClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIIssueCommentClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, commentID string, ifMatch string, options *APIIssueCommentClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}"
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
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
	if commentID == "" {
		return nil, errors.New("parameter commentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{commentId}", url.PathEscape(commentID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the issue Comment for an API specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - issueID - Issue identifier. Must be unique in the current API Management service instance.
//   - commentID - Comment identifier within an Issue. Must be unique in the current Issue.
//   - options - APIIssueCommentClientGetOptions contains the optional parameters for the APIIssueCommentClient.Get method.
func (client *APIIssueCommentClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, commentID string, options *APIIssueCommentClientGetOptions) (APIIssueCommentClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiID, issueID, commentID, options)
	if err != nil {
		return APIIssueCommentClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIIssueCommentClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return APIIssueCommentClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *APIIssueCommentClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, commentID string, options *APIIssueCommentClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}"
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
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
	if commentID == "" {
		return nil, errors.New("parameter commentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{commentId}", url.PathEscape(commentID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIIssueCommentClient) getHandleResponse(resp *http.Response) (APIIssueCommentClientGetResponse, error) {
	result := APIIssueCommentClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.IssueCommentContract); err != nil {
		return APIIssueCommentClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the issue Comment for an API specified by its identifier.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - issueID - Issue identifier. Must be unique in the current API Management service instance.
//   - commentID - Comment identifier within an Issue. Must be unique in the current Issue.
//   - options - APIIssueCommentClientGetEntityTagOptions contains the optional parameters for the APIIssueCommentClient.GetEntityTag
//     method.
func (client *APIIssueCommentClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, commentID string, options *APIIssueCommentClientGetEntityTagOptions) (APIIssueCommentClientGetEntityTagResponse, error) {
	var err error
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, apiID, issueID, commentID, options)
	if err != nil {
		return APIIssueCommentClientGetEntityTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIIssueCommentClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return APIIssueCommentClientGetEntityTagResponse{}, err
	}
	resp, err := client.getEntityTagHandleResponse(httpResp)
	return resp, err
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *APIIssueCommentClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, commentID string, options *APIIssueCommentClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}"
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
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
	if commentID == "" {
		return nil, errors.New("parameter commentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{commentId}", url.PathEscape(commentID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *APIIssueCommentClient) getEntityTagHandleResponse(resp *http.Response) (APIIssueCommentClientGetEntityTagResponse, error) {
	result := APIIssueCommentClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Lists all comments for the Issue associated with the specified API.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - issueID - Issue identifier. Must be unique in the current API Management service instance.
//   - options - APIIssueCommentClientListByServiceOptions contains the optional parameters for the APIIssueCommentClient.NewListByServicePager
//     method.
func (client *APIIssueCommentClient) NewListByServicePager(resourceGroupName string, serviceName string, apiID string, issueID string, options *APIIssueCommentClientListByServiceOptions) *runtime.Pager[APIIssueCommentClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[APIIssueCommentClientListByServiceResponse]{
		More: func(page APIIssueCommentClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APIIssueCommentClientListByServiceResponse) (APIIssueCommentClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, apiID, issueID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APIIssueCommentClientListByServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return APIIssueCommentClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APIIssueCommentClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *APIIssueCommentClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, options *APIIssueCommentClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments"
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
	if issueID == "" {
		return nil, errors.New("parameter issueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueId}", url.PathEscape(issueID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *APIIssueCommentClient) listByServiceHandleResponse(resp *http.Response) (APIIssueCommentClientListByServiceResponse, error) {
	result := APIIssueCommentClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IssueCommentCollection); err != nil {
		return APIIssueCommentClientListByServiceResponse{}, err
	}
	return result, nil
}
