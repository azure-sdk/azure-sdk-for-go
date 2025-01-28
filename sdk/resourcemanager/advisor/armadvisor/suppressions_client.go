//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

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

// SuppressionsClient contains the methods for the Suppressions group.
// Don't use this type directly, use NewSuppressionsClient() instead.
type SuppressionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSuppressionsClient creates a new instance of SuppressionsClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSuppressionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SuppressionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SuppressionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Enables the snoozed or dismissed attribute of a recommendation. The snoozed or dismissed attribute is referred
// to as a suppression. Use this API to create or update the snoozed or dismissed status of
// a recommendation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceURI - The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
//   - recommendationID - The recommendation ID.
//   - name - The name of the suppression.
//   - suppressionContract - The snoozed or dismissed attribute; for example, the snooze duration.
//   - options - SuppressionsClientCreateOptions contains the optional parameters for the SuppressionsClient.Create method.
func (client *SuppressionsClient) Create(ctx context.Context, resourceURI string, recommendationID string, name string, suppressionContract SuppressionContract, options *SuppressionsClientCreateOptions) (SuppressionsClientCreateResponse, error) {
	var err error
	const operationName = "SuppressionsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceURI, recommendationID, name, suppressionContract, options)
	if err != nil {
		return SuppressionsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SuppressionsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SuppressionsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SuppressionsClient) createCreateRequest(ctx context.Context, resourceURI string, recommendationID string, name string, suppressionContract SuppressionContract, options *SuppressionsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", url.PathEscape(resourceURI))
	if recommendationID == "" {
		return nil, errors.New("parameter recommendationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendationId}", url.PathEscape(recommendationID))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, suppressionContract); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SuppressionsClient) createHandleResponse(resp *http.Response) (SuppressionsClientCreateResponse, error) {
	result := SuppressionsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SuppressionContract); err != nil {
		return SuppressionsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Enables the activation of a snoozed or dismissed recommendation. The snoozed or dismissed attribute of a recommendation
// is referred to as a suppression.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceURI - The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
//   - recommendationID - The recommendation ID.
//   - name - The name of the suppression.
//   - options - SuppressionsClientDeleteOptions contains the optional parameters for the SuppressionsClient.Delete method.
func (client *SuppressionsClient) Delete(ctx context.Context, resourceURI string, recommendationID string, name string, options *SuppressionsClientDeleteOptions) (SuppressionsClientDeleteResponse, error) {
	var err error
	const operationName = "SuppressionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceURI, recommendationID, name, options)
	if err != nil {
		return SuppressionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SuppressionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SuppressionsClientDeleteResponse{}, err
	}
	return SuppressionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SuppressionsClient) deleteCreateRequest(ctx context.Context, resourceURI string, recommendationID string, name string, options *SuppressionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", url.PathEscape(resourceURI))
	if recommendationID == "" {
		return nil, errors.New("parameter recommendationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendationId}", url.PathEscape(recommendationID))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Obtains the details of a suppression.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceURI - The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
//   - recommendationID - The recommendation ID.
//   - name - The name of the suppression.
//   - options - SuppressionsClientGetOptions contains the optional parameters for the SuppressionsClient.Get method.
func (client *SuppressionsClient) Get(ctx context.Context, resourceURI string, recommendationID string, name string, options *SuppressionsClientGetOptions) (SuppressionsClientGetResponse, error) {
	var err error
	const operationName = "SuppressionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceURI, recommendationID, name, options)
	if err != nil {
		return SuppressionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SuppressionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SuppressionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SuppressionsClient) getCreateRequest(ctx context.Context, resourceURI string, recommendationID string, name string, options *SuppressionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", url.PathEscape(resourceURI))
	if recommendationID == "" {
		return nil, errors.New("parameter recommendationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendationId}", url.PathEscape(recommendationID))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SuppressionsClient) getHandleResponse(resp *http.Response) (SuppressionsClientGetResponse, error) {
	result := SuppressionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SuppressionContract); err != nil {
		return SuppressionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves the list of snoozed or dismissed suppressions for a subscription. The snoozed or dismissed attribute
// of a recommendation is referred to as a suppression.
//
// Generated from API version 2023-01-01
//   - options - SuppressionsClientListOptions contains the optional parameters for the SuppressionsClient.NewListPager method.
func (client *SuppressionsClient) NewListPager(options *SuppressionsClientListOptions) *runtime.Pager[SuppressionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SuppressionsClientListResponse]{
		More: func(page SuppressionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SuppressionsClientListResponse) (SuppressionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SuppressionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SuppressionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SuppressionsClient) listCreateRequest(ctx context.Context, options *SuppressionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/suppressions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SuppressionsClient) listHandleResponse(resp *http.Response) (SuppressionsClientListResponse, error) {
	result := SuppressionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SuppressionContractListResult); err != nil {
		return SuppressionsClientListResponse{}, err
	}
	return result, nil
}
