//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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

// ArtifactSourcesClient contains the methods for the ArtifactSources group.
// Don't use this type directly, use NewArtifactSourcesClient() instead.
type ArtifactSourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewArtifactSourcesClient creates a new instance of ArtifactSourcesClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewArtifactSourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ArtifactSourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ArtifactSourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or replace an existing artifact source.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - name - The name of the artifact source.
//   - artifactSource - Properties of an artifact source.
//   - options - ArtifactSourcesClientCreateOrUpdateOptions contains the optional parameters for the ArtifactSourcesClient.CreateOrUpdate
//     method.
func (client *ArtifactSourcesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, artifactSource ArtifactSource, options *ArtifactSourcesClientCreateOrUpdateOptions) (ArtifactSourcesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ArtifactSourcesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, name, artifactSource, options)
	if err != nil {
		return ArtifactSourcesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArtifactSourcesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ArtifactSourcesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ArtifactSourcesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, artifactSource ArtifactSource, options *ArtifactSourcesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, artifactSource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ArtifactSourcesClient) createOrUpdateHandleResponse(resp *http.Response) (ArtifactSourcesClientCreateOrUpdateResponse, error) {
	result := ArtifactSourcesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactSource); err != nil {
		return ArtifactSourcesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete artifact source.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - name - The name of the artifact source.
//   - options - ArtifactSourcesClientDeleteOptions contains the optional parameters for the ArtifactSourcesClient.Delete method.
func (client *ArtifactSourcesClient) Delete(ctx context.Context, resourceGroupName string, labName string, name string, options *ArtifactSourcesClientDeleteOptions) (ArtifactSourcesClientDeleteResponse, error) {
	var err error
	const operationName = "ArtifactSourcesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return ArtifactSourcesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArtifactSourcesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ArtifactSourcesClientDeleteResponse{}, err
	}
	return ArtifactSourcesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ArtifactSourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *ArtifactSourcesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get artifact source.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - name - The name of the artifact source.
//   - options - ArtifactSourcesClientGetOptions contains the optional parameters for the ArtifactSourcesClient.Get method.
func (client *ArtifactSourcesClient) Get(ctx context.Context, resourceGroupName string, labName string, name string, options *ArtifactSourcesClientGetOptions) (ArtifactSourcesClientGetResponse, error) {
	var err error
	const operationName = "ArtifactSourcesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return ArtifactSourcesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArtifactSourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArtifactSourcesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ArtifactSourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *ArtifactSourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArtifactSourcesClient) getHandleResponse(resp *http.Response) (ArtifactSourcesClientGetResponse, error) {
	result := ArtifactSourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactSource); err != nil {
		return ArtifactSourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List artifact sources in a given lab.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - options - ArtifactSourcesClientListOptions contains the optional parameters for the ArtifactSourcesClient.NewListPager
//     method.
func (client *ArtifactSourcesClient) NewListPager(resourceGroupName string, labName string, options *ArtifactSourcesClientListOptions) *runtime.Pager[ArtifactSourcesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ArtifactSourcesClientListResponse]{
		More: func(page ArtifactSourcesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ArtifactSourcesClientListResponse) (ArtifactSourcesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ArtifactSourcesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, labName, options)
			}, nil)
			if err != nil {
				return ArtifactSourcesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ArtifactSourcesClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *ArtifactSourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ArtifactSourcesClient) listHandleResponse(resp *http.Response) (ArtifactSourcesClientListResponse, error) {
	result := ArtifactSourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactSourceList); err != nil {
		return ArtifactSourcesClientListResponse{}, err
	}
	return result, nil
}

// Update - Allows modifying tags of artifact sources. All other properties will be ignored.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - name - The name of the artifact source.
//   - artifactSource - Properties of an artifact source.
//   - options - ArtifactSourcesClientUpdateOptions contains the optional parameters for the ArtifactSourcesClient.Update method.
func (client *ArtifactSourcesClient) Update(ctx context.Context, resourceGroupName string, labName string, name string, artifactSource ArtifactSourceFragment, options *ArtifactSourcesClientUpdateOptions) (ArtifactSourcesClientUpdateResponse, error) {
	var err error
	const operationName = "ArtifactSourcesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, name, artifactSource, options)
	if err != nil {
		return ArtifactSourcesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArtifactSourcesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArtifactSourcesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ArtifactSourcesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, artifactSource ArtifactSourceFragment, options *ArtifactSourcesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, artifactSource); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ArtifactSourcesClient) updateHandleResponse(resp *http.Response) (ArtifactSourcesClientUpdateResponse, error) {
	result := ArtifactSourcesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactSource); err != nil {
		return ArtifactSourcesClientUpdateResponse{}, err
	}
	return result, nil
}
