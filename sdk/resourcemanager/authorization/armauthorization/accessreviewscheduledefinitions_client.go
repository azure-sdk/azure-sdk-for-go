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

// AccessReviewScheduleDefinitionsClient contains the methods for the AccessReviewScheduleDefinitions group.
// Don't use this type directly, use NewAccessReviewScheduleDefinitionsClient() instead.
type AccessReviewScheduleDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessReviewScheduleDefinitionsClient creates a new instance of AccessReviewScheduleDefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessReviewScheduleDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewScheduleDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessReviewScheduleDefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdateByID - Create or Update access review schedule definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - properties - Access review schedule definition properties.
//   - options - AccessReviewScheduleDefinitionsClientCreateOrUpdateByIDOptions contains the optional parameters for the AccessReviewScheduleDefinitionsClient.CreateOrUpdateByID
//     method.
func (client *AccessReviewScheduleDefinitionsClient) CreateOrUpdateByID(ctx context.Context, scheduleDefinitionID string, properties AccessReviewScheduleDefinitionProperties, options *AccessReviewScheduleDefinitionsClientCreateOrUpdateByIDOptions) (AccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse, error) {
	var err error
	const operationName = "AccessReviewScheduleDefinitionsClient.CreateOrUpdateByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateByIDCreateRequest(ctx, scheduleDefinitionID, properties, options)
	if err != nil {
		return AccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	resp, err := client.createOrUpdateByIDHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateByIDCreateRequest creates the CreateOrUpdateByID request.
func (client *AccessReviewScheduleDefinitionsClient) createOrUpdateByIDCreateRequest(ctx context.Context, scheduleDefinitionID string, properties AccessReviewScheduleDefinitionProperties, _ *AccessReviewScheduleDefinitionsClientCreateOrUpdateByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *AccessReviewScheduleDefinitionsClient) createOrUpdateByIDHandleResponse(resp *http.Response) (AccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse, error) {
	result := AccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewScheduleDefinition); err != nil {
		return AccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	return result, nil
}

// DeleteByID - Delete access review schedule definition
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - options - AccessReviewScheduleDefinitionsClientDeleteByIDOptions contains the optional parameters for the AccessReviewScheduleDefinitionsClient.DeleteByID
//     method.
func (client *AccessReviewScheduleDefinitionsClient) DeleteByID(ctx context.Context, scheduleDefinitionID string, options *AccessReviewScheduleDefinitionsClientDeleteByIDOptions) (AccessReviewScheduleDefinitionsClientDeleteByIDResponse, error) {
	var err error
	const operationName = "AccessReviewScheduleDefinitionsClient.DeleteByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteByIDCreateRequest(ctx, scheduleDefinitionID, options)
	if err != nil {
		return AccessReviewScheduleDefinitionsClientDeleteByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewScheduleDefinitionsClientDeleteByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewScheduleDefinitionsClientDeleteByIDResponse{}, err
	}
	return AccessReviewScheduleDefinitionsClientDeleteByIDResponse{}, nil
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *AccessReviewScheduleDefinitionsClient) deleteByIDCreateRequest(ctx context.Context, scheduleDefinitionID string, _ *AccessReviewScheduleDefinitionsClientDeleteByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - options - AccessReviewScheduleDefinitionsClientGetByIDOptions contains the optional parameters for the AccessReviewScheduleDefinitionsClient.GetByID
//     method.
func (client *AccessReviewScheduleDefinitionsClient) GetByID(ctx context.Context, scheduleDefinitionID string, options *AccessReviewScheduleDefinitionsClientGetByIDOptions) (AccessReviewScheduleDefinitionsClientGetByIDResponse, error) {
	var err error
	const operationName = "AccessReviewScheduleDefinitionsClient.GetByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByIDCreateRequest(ctx, scheduleDefinitionID, options)
	if err != nil {
		return AccessReviewScheduleDefinitionsClientGetByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewScheduleDefinitionsClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewScheduleDefinitionsClientGetByIDResponse{}, err
	}
	resp, err := client.getByIDHandleResponse(httpResp)
	return resp, err
}

// getByIDCreateRequest creates the GetByID request.
func (client *AccessReviewScheduleDefinitionsClient) getByIDCreateRequest(ctx context.Context, scheduleDefinitionID string, _ *AccessReviewScheduleDefinitionsClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *AccessReviewScheduleDefinitionsClient) getByIDHandleResponse(resp *http.Response) (AccessReviewScheduleDefinitionsClientGetByIDResponse, error) {
	result := AccessReviewScheduleDefinitionsClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewScheduleDefinition); err != nil {
		return AccessReviewScheduleDefinitionsClientGetByIDResponse{}, err
	}
	return result, nil
}

// NewListPager - Get access review schedule definitions
//
// Generated from API version 2021-12-01-preview
//   - options - AccessReviewScheduleDefinitionsClientListOptions contains the optional parameters for the AccessReviewScheduleDefinitionsClient.NewListPager
//     method.
func (client *AccessReviewScheduleDefinitionsClient) NewListPager(options *AccessReviewScheduleDefinitionsClientListOptions) *runtime.Pager[AccessReviewScheduleDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessReviewScheduleDefinitionsClientListResponse]{
		More: func(page AccessReviewScheduleDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessReviewScheduleDefinitionsClientListResponse) (AccessReviewScheduleDefinitionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccessReviewScheduleDefinitionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AccessReviewScheduleDefinitionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AccessReviewScheduleDefinitionsClient) listCreateRequest(ctx context.Context, options *AccessReviewScheduleDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *AccessReviewScheduleDefinitionsClient) listHandleResponse(resp *http.Response) (AccessReviewScheduleDefinitionsClientListResponse, error) {
	result := AccessReviewScheduleDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewScheduleDefinitionListResult); err != nil {
		return AccessReviewScheduleDefinitionsClientListResponse{}, err
	}
	return result, nil
}

// Stop - Stop access review definition
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - options - AccessReviewScheduleDefinitionsClientStopOptions contains the optional parameters for the AccessReviewScheduleDefinitionsClient.Stop
//     method.
func (client *AccessReviewScheduleDefinitionsClient) Stop(ctx context.Context, scheduleDefinitionID string, options *AccessReviewScheduleDefinitionsClientStopOptions) (AccessReviewScheduleDefinitionsClientStopResponse, error) {
	var err error
	const operationName = "AccessReviewScheduleDefinitionsClient.Stop"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopCreateRequest(ctx, scheduleDefinitionID, options)
	if err != nil {
		return AccessReviewScheduleDefinitionsClientStopResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessReviewScheduleDefinitionsClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AccessReviewScheduleDefinitionsClientStopResponse{}, err
	}
	return AccessReviewScheduleDefinitionsClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *AccessReviewScheduleDefinitionsClient) stopCreateRequest(ctx context.Context, scheduleDefinitionID string, _ *AccessReviewScheduleDefinitionsClientStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
