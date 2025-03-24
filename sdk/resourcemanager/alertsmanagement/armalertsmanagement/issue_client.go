// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

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

// IssueClient contains the methods for the Issue group.
// Don't use this type directly, use NewIssueClient() instead.
type IssueClient struct {
	internal  *arm.Client
	issueName string
}

// NewIssueClient creates a new instance of IssueClient with the specified values.
//   - issueName - The name of the IssueResource
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIssueClient(issueName string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IssueClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IssueClient{
		issueName: issueName,
		internal:  cl,
	}
	return client, nil
}

// AddOrUpdateAlerts - Add or update alerts in the issue
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - body - The content of the action request
//   - options - IssueClientAddOrUpdateAlertsOptions contains the optional parameters for the IssueClient.AddOrUpdateAlerts method.
func (client *IssueClient) AddOrUpdateAlerts(ctx context.Context, resourceURI string, body RelatedAlerts, options *IssueClientAddOrUpdateAlertsOptions) (IssueClientAddOrUpdateAlertsResponse, error) {
	var err error
	const operationName = "IssueClient.AddOrUpdateAlerts"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.addOrUpdateAlertsCreateRequest(ctx, resourceURI, body, options)
	if err != nil {
		return IssueClientAddOrUpdateAlertsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IssueClientAddOrUpdateAlertsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IssueClientAddOrUpdateAlertsResponse{}, err
	}
	return IssueClientAddOrUpdateAlertsResponse{}, nil
}

// addOrUpdateAlertsCreateRequest creates the AddOrUpdateAlerts request.
func (client *IssueClient) addOrUpdateAlertsCreateRequest(ctx context.Context, resourceURI string, body RelatedAlerts, _ *IssueClientAddOrUpdateAlertsOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues/{issueName}/addOrUpdateAlerts"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if client.issueName == "" {
		return nil, errors.New("parameter client.issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(client.issueName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// AddOrUpdateResources - Add or update resources in the issue
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - body - The content of the action request
//   - options - IssueClientAddOrUpdateResourcesOptions contains the optional parameters for the IssueClient.AddOrUpdateResources
//     method.
func (client *IssueClient) AddOrUpdateResources(ctx context.Context, resourceURI string, body RelatedResources, options *IssueClientAddOrUpdateResourcesOptions) (IssueClientAddOrUpdateResourcesResponse, error) {
	var err error
	const operationName = "IssueClient.AddOrUpdateResources"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.addOrUpdateResourcesCreateRequest(ctx, resourceURI, body, options)
	if err != nil {
		return IssueClientAddOrUpdateResourcesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IssueClientAddOrUpdateResourcesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IssueClientAddOrUpdateResourcesResponse{}, err
	}
	return IssueClientAddOrUpdateResourcesResponse{}, nil
}

// addOrUpdateResourcesCreateRequest creates the AddOrUpdateResources request.
func (client *IssueClient) addOrUpdateResourcesCreateRequest(ctx context.Context, resourceURI string, body RelatedResources, _ *IssueClientAddOrUpdateResourcesOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues/{issueName}/addOrUpdateResources"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if client.issueName == "" {
		return nil, errors.New("parameter client.issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(client.issueName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// Create - Create a new issue or updates an existing one
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - resource - Resource create parameters.
//   - options - IssueClientCreateOptions contains the optional parameters for the IssueClient.Create method.
func (client *IssueClient) Create(ctx context.Context, resourceURI string, resource IssueResource, options *IssueClientCreateOptions) (IssueClientCreateResponse, error) {
	var err error
	const operationName = "IssueClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceURI, resource, options)
	if err != nil {
		return IssueClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IssueClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return IssueClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *IssueClient) createCreateRequest(ctx context.Context, resourceURI string, resource IssueResource, _ *IssueClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues/{issueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if client.issueName == "" {
		return nil, errors.New("parameter client.issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(client.issueName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *IssueClient) createHandleResponse(resp *http.Response) (IssueClientCreateResponse, error) {
	result := IssueClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IssueResource); err != nil {
		return IssueClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an issue
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - IssueClientDeleteOptions contains the optional parameters for the IssueClient.Delete method.
func (client *IssueClient) Delete(ctx context.Context, resourceURI string, options *IssueClientDeleteOptions) (IssueClientDeleteResponse, error) {
	var err error
	const operationName = "IssueClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return IssueClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IssueClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IssueClientDeleteResponse{}, err
	}
	return IssueClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IssueClient) deleteCreateRequest(ctx context.Context, resourceURI string, _ *IssueClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues/{issueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if client.issueName == "" {
		return nil, errors.New("parameter client.issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(client.issueName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// FetchInvestigationResult - Fetch investigation results
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - body - The content of the action request
//   - options - IssueClientFetchInvestigationResultOptions contains the optional parameters for the IssueClient.FetchInvestigationResult
//     method.
func (client *IssueClient) FetchInvestigationResult(ctx context.Context, resourceURI string, body FetchInvestigationResultParameters, options *IssueClientFetchInvestigationResultOptions) (IssueClientFetchInvestigationResultResponse, error) {
	var err error
	const operationName = "IssueClient.FetchInvestigationResult"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fetchInvestigationResultCreateRequest(ctx, resourceURI, body, options)
	if err != nil {
		return IssueClientFetchInvestigationResultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IssueClientFetchInvestigationResultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IssueClientFetchInvestigationResultResponse{}, err
	}
	resp, err := client.fetchInvestigationResultHandleResponse(httpResp)
	return resp, err
}

// fetchInvestigationResultCreateRequest creates the FetchInvestigationResult request.
func (client *IssueClient) fetchInvestigationResultCreateRequest(ctx context.Context, resourceURI string, body FetchInvestigationResultParameters, _ *IssueClientFetchInvestigationResultOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues/{issueName}/fetchInvestigationResult"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if client.issueName == "" {
		return nil, errors.New("parameter client.issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(client.issueName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// fetchInvestigationResultHandleResponse handles the FetchInvestigationResult response.
func (client *IssueClient) fetchInvestigationResultHandleResponse(resp *http.Response) (IssueClientFetchInvestigationResultResponse, error) {
	result := IssueClientFetchInvestigationResultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvestigationResult); err != nil {
		return IssueClientFetchInvestigationResultResponse{}, err
	}
	return result, nil
}

// Get - Get issue properties
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - IssueClientGetOptions contains the optional parameters for the IssueClient.Get method.
func (client *IssueClient) Get(ctx context.Context, resourceURI string, options *IssueClientGetOptions) (IssueClientGetResponse, error) {
	var err error
	const operationName = "IssueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return IssueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IssueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IssueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IssueClient) getCreateRequest(ctx context.Context, resourceURI string, _ *IssueClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues/{issueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if client.issueName == "" {
		return nil, errors.New("parameter client.issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(client.issueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IssueClient) getHandleResponse(resp *http.Response) (IssueClientGetResponse, error) {
	result := IssueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IssueResource); err != nil {
		return IssueClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all issues under the parent
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - options - IssueClientListOptions contains the optional parameters for the IssueClient.NewListPager method.
func (client *IssueClient) NewListPager(resourceURI string, options *IssueClientListOptions) *runtime.Pager[IssueClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IssueClientListResponse]{
		More: func(page IssueClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IssueClientListResponse) (IssueClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IssueClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceURI, options)
			}, nil)
			if err != nil {
				return IssueClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *IssueClient) listCreateRequest(ctx context.Context, resourceURI string, _ *IssueClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IssueClient) listHandleResponse(resp *http.Response) (IssueClientListResponse, error) {
	result := IssueClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IssueResourceListResult); err != nil {
		return IssueClientListResponse{}, err
	}
	return result, nil
}

// NewListAlertsPager - List all alerts in the issue - this method uses pagination to return all alerts
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - body - The content of the action request
//   - options - IssueClientListAlertsOptions contains the optional parameters for the IssueClient.NewListAlertsPager method.
func (client *IssueClient) NewListAlertsPager(resourceURI string, body ListParameter, options *IssueClientListAlertsOptions) *runtime.Pager[IssueClientListAlertsResponse] {
	return runtime.NewPager(runtime.PagingHandler[IssueClientListAlertsResponse]{
		More: func(page IssueClientListAlertsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IssueClientListAlertsResponse) (IssueClientListAlertsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IssueClient.NewListAlertsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAlertsCreateRequest(ctx, resourceURI, body, options)
			}, nil)
			if err != nil {
				return IssueClientListAlertsResponse{}, err
			}
			return client.listAlertsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAlertsCreateRequest creates the ListAlerts request.
func (client *IssueClient) listAlertsCreateRequest(ctx context.Context, resourceURI string, body ListParameter, _ *IssueClientListAlertsOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues/{issueName}/listAlerts"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if client.issueName == "" {
		return nil, errors.New("parameter client.issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(client.issueName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// listAlertsHandleResponse handles the ListAlerts response.
func (client *IssueClient) listAlertsHandleResponse(resp *http.Response) (IssueClientListAlertsResponse, error) {
	result := IssueClientListAlertsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedRelatedAlert); err != nil {
		return IssueClientListAlertsResponse{}, err
	}
	return result, nil
}

// NewListResourcesPager - List all resources in the issue - this method uses pagination to return all resources
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - body - The content of the action request
//   - options - IssueClientListResourcesOptions contains the optional parameters for the IssueClient.NewListResourcesPager method.
func (client *IssueClient) NewListResourcesPager(resourceURI string, body ListParameter, options *IssueClientListResourcesOptions) *runtime.Pager[IssueClientListResourcesResponse] {
	return runtime.NewPager(runtime.PagingHandler[IssueClientListResourcesResponse]{
		More: func(page IssueClientListResourcesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IssueClientListResourcesResponse) (IssueClientListResourcesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IssueClient.NewListResourcesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listResourcesCreateRequest(ctx, resourceURI, body, options)
			}, nil)
			if err != nil {
				return IssueClientListResourcesResponse{}, err
			}
			return client.listResourcesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listResourcesCreateRequest creates the ListResources request.
func (client *IssueClient) listResourcesCreateRequest(ctx context.Context, resourceURI string, body ListParameter, _ *IssueClientListResourcesOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues/{issueName}/listResources"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if client.issueName == "" {
		return nil, errors.New("parameter client.issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(client.issueName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// listResourcesHandleResponse handles the ListResources response.
func (client *IssueClient) listResourcesHandleResponse(resp *http.Response) (IssueClientListResourcesResponse, error) {
	result := IssueClientListResourcesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedRelatedResource); err != nil {
		return IssueClientListResourcesResponse{}, err
	}
	return result, nil
}

// StartInvestigation - Start a new investigation
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - body - The content of the action request
//   - options - IssueClientStartInvestigationOptions contains the optional parameters for the IssueClient.StartInvestigation
//     method.
func (client *IssueClient) StartInvestigation(ctx context.Context, resourceURI string, body StartInvestigationParameters, options *IssueClientStartInvestigationOptions) (IssueClientStartInvestigationResponse, error) {
	var err error
	const operationName = "IssueClient.StartInvestigation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startInvestigationCreateRequest(ctx, resourceURI, body, options)
	if err != nil {
		return IssueClientStartInvestigationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IssueClientStartInvestigationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IssueClientStartInvestigationResponse{}, err
	}
	resp, err := client.startInvestigationHandleResponse(httpResp)
	return resp, err
}

// startInvestigationCreateRequest creates the StartInvestigation request.
func (client *IssueClient) startInvestigationCreateRequest(ctx context.Context, resourceURI string, body StartInvestigationParameters, _ *IssueClientStartInvestigationOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues/{issueName}/startInvestigation"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if client.issueName == "" {
		return nil, errors.New("parameter client.issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(client.issueName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// startInvestigationHandleResponse handles the StartInvestigation response.
func (client *IssueClient) startInvestigationHandleResponse(resp *http.Response) (IssueClientStartInvestigationResponse, error) {
	result := IssueClientStartInvestigationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InvestigationMetadata); err != nil {
		return IssueClientStartInvestigationResponse{}, err
	}
	return result, nil
}

// Update - Update an issue
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource.
//   - properties - The resource properties to be updated.
//   - options - IssueClientUpdateOptions contains the optional parameters for the IssueClient.Update method.
func (client *IssueClient) Update(ctx context.Context, resourceURI string, properties IssueResourceUpdate, options *IssueClientUpdateOptions) (IssueClientUpdateResponse, error) {
	var err error
	const operationName = "IssueClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceURI, properties, options)
	if err != nil {
		return IssueClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IssueClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IssueClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *IssueClient) updateCreateRequest(ctx context.Context, resourceURI string, properties IssueResourceUpdate, _ *IssueClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/issues/{issueName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if client.issueName == "" {
		return nil, errors.New("parameter client.issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(client.issueName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *IssueClient) updateHandleResponse(resp *http.Response) (IssueClientUpdateResponse, error) {
	result := IssueClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IssueResource); err != nil {
		return IssueClientUpdateResponse{}, err
	}
	return result, nil
}
