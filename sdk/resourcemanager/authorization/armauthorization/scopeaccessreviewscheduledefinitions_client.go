// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// ScopeAccessReviewScheduleDefinitionsClient contains the methods for the ScopeAccessReviewScheduleDefinitions group.
// Don't use this type directly, use NewScopeAccessReviewScheduleDefinitionsClient() instead.
type ScopeAccessReviewScheduleDefinitionsClient struct {
	internal *arm.Client
}

// NewScopeAccessReviewScheduleDefinitionsClient creates a new instance of ScopeAccessReviewScheduleDefinitionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScopeAccessReviewScheduleDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ScopeAccessReviewScheduleDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScopeAccessReviewScheduleDefinitionsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdateByID - Create or Update access review schedule definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - properties - Access review schedule definition properties.
//   - options - ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDOptions contains the optional parameters for the
//     ScopeAccessReviewScheduleDefinitionsClient.CreateOrUpdateByID method.
func (client *ScopeAccessReviewScheduleDefinitionsClient) CreateOrUpdateByID(ctx context.Context, scope string, scheduleDefinitionID string, properties AccessReviewScheduleDefinitionProperties, options *ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDOptions) (ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewScheduleDefinitionsClient.CreateOrUpdateByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateByIDCreateRequest(ctx, scope, scheduleDefinitionID, properties, options)
	if err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	resp, err := client.createOrUpdateByIDHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateByIDCreateRequest creates the CreateOrUpdateByID request.
func (client *ScopeAccessReviewScheduleDefinitionsClient) createOrUpdateByIDCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, properties AccessReviewScheduleDefinitionProperties, _ *ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateByIDHandleResponse handles the CreateOrUpdateByID response.
func (client *ScopeAccessReviewScheduleDefinitionsClient) createOrUpdateByIDHandleResponse(resp *http.Response) (ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse, error) {
	result := ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewScheduleDefinition); err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	return result, nil
}

// DeleteByID - Delete access review schedule definition
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - options - ScopeAccessReviewScheduleDefinitionsClientDeleteByIDOptions contains the optional parameters for the ScopeAccessReviewScheduleDefinitionsClient.DeleteByID
//     method.
func (client *ScopeAccessReviewScheduleDefinitionsClient) DeleteByID(ctx context.Context, scope string, scheduleDefinitionID string, options *ScopeAccessReviewScheduleDefinitionsClientDeleteByIDOptions) (ScopeAccessReviewScheduleDefinitionsClientDeleteByIDResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewScheduleDefinitionsClient.DeleteByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteByIDCreateRequest(ctx, scope, scheduleDefinitionID, options)
	if err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientDeleteByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientDeleteByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewScheduleDefinitionsClientDeleteByIDResponse{}, err
	}
	return ScopeAccessReviewScheduleDefinitionsClientDeleteByIDResponse{}, nil
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *ScopeAccessReviewScheduleDefinitionsClient) deleteByIDCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, _ *ScopeAccessReviewScheduleDefinitionsClientDeleteByIDOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetByID - Get single access review definition
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - options - ScopeAccessReviewScheduleDefinitionsClientGetByIDOptions contains the optional parameters for the ScopeAccessReviewScheduleDefinitionsClient.GetByID
//     method.
func (client *ScopeAccessReviewScheduleDefinitionsClient) GetByID(ctx context.Context, scope string, scheduleDefinitionID string, options *ScopeAccessReviewScheduleDefinitionsClientGetByIDOptions) (ScopeAccessReviewScheduleDefinitionsClientGetByIDResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewScheduleDefinitionsClient.GetByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByIDCreateRequest(ctx, scope, scheduleDefinitionID, options)
	if err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientGetByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewScheduleDefinitionsClientGetByIDResponse{}, err
	}
	resp, err := client.getByIDHandleResponse(httpResp)
	return resp, err
}

// getByIDCreateRequest creates the GetByID request.
func (client *ScopeAccessReviewScheduleDefinitionsClient) getByIDCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, _ *ScopeAccessReviewScheduleDefinitionsClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *ScopeAccessReviewScheduleDefinitionsClient) getByIDHandleResponse(resp *http.Response) (ScopeAccessReviewScheduleDefinitionsClientGetByIDResponse, error) {
	result := ScopeAccessReviewScheduleDefinitionsClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewScheduleDefinition); err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientGetByIDResponse{}, err
	}
	return result, nil
}

// NewListPager - Get access review schedule definitions
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - options - ScopeAccessReviewScheduleDefinitionsClientListOptions contains the optional parameters for the ScopeAccessReviewScheduleDefinitionsClient.NewListPager
//     method.
func (client *ScopeAccessReviewScheduleDefinitionsClient) NewListPager(scope string, options *ScopeAccessReviewScheduleDefinitionsClientListOptions) *runtime.Pager[ScopeAccessReviewScheduleDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScopeAccessReviewScheduleDefinitionsClientListResponse]{
		More: func(page ScopeAccessReviewScheduleDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScopeAccessReviewScheduleDefinitionsClientListResponse) (ScopeAccessReviewScheduleDefinitionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScopeAccessReviewScheduleDefinitionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return ScopeAccessReviewScheduleDefinitionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ScopeAccessReviewScheduleDefinitionsClient) listCreateRequest(ctx context.Context, scope string, options *ScopeAccessReviewScheduleDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ScopeAccessReviewScheduleDefinitionsClient) listHandleResponse(resp *http.Response) (ScopeAccessReviewScheduleDefinitionsClientListResponse, error) {
	result := ScopeAccessReviewScheduleDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewScheduleDefinitionListResult); err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientListResponse{}, err
	}
	return result, nil
}

// Stop - Stop access review definition
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - options - ScopeAccessReviewScheduleDefinitionsClientStopOptions contains the optional parameters for the ScopeAccessReviewScheduleDefinitionsClient.Stop
//     method.
func (client *ScopeAccessReviewScheduleDefinitionsClient) Stop(ctx context.Context, scope string, scheduleDefinitionID string, options *ScopeAccessReviewScheduleDefinitionsClientStopOptions) (ScopeAccessReviewScheduleDefinitionsClientStopResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewScheduleDefinitionsClient.Stop"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopCreateRequest(ctx, scope, scheduleDefinitionID, options)
	if err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientStopResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewScheduleDefinitionsClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewScheduleDefinitionsClientStopResponse{}, err
	}
	return ScopeAccessReviewScheduleDefinitionsClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *ScopeAccessReviewScheduleDefinitionsClient) stopCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, _ *ScopeAccessReviewScheduleDefinitionsClientStopOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/stop"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
