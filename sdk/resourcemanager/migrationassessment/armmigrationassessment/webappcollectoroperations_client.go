//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationassessment

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

// WebAppCollectorOperationsClient contains the methods for the WebAppCollectorOperations group.
// Don't use this type directly, use NewWebAppCollectorOperationsClient() instead.
type WebAppCollectorOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWebAppCollectorOperationsClient creates a new instance of WebAppCollectorOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWebAppCollectorOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebAppCollectorOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WebAppCollectorOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a WebAppCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - collectorName - Web app collector ARM name.
//   - resource - Resource create parameters.
//   - options - WebAppCollectorOperationsClientBeginCreateOptions contains the optional parameters for the WebAppCollectorOperationsClient.BeginCreate
//     method.
func (client *WebAppCollectorOperationsClient) BeginCreate(ctx context.Context, resourceGroupName string, projectName string, collectorName string, resource WebAppCollector, options *WebAppCollectorOperationsClientBeginCreateOptions) (*runtime.Poller[WebAppCollectorOperationsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, projectName, collectorName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WebAppCollectorOperationsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WebAppCollectorOperationsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a WebAppCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *WebAppCollectorOperationsClient) create(ctx context.Context, resourceGroupName string, projectName string, collectorName string, resource WebAppCollector, options *WebAppCollectorOperationsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "WebAppCollectorOperationsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, projectName, collectorName, resource, options)
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

// createCreateRequest creates the Create request.
func (client *WebAppCollectorOperationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, projectName string, collectorName string, resource WebAppCollector, options *WebAppCollectorOperationsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/webAppCollectors/{collectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if collectorName == "" {
		return nil, errors.New("parameter collectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectorName}", url.PathEscape(collectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete a WebAppCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - collectorName - Web app collector ARM name.
//   - options - WebAppCollectorOperationsClientDeleteOptions contains the optional parameters for the WebAppCollectorOperationsClient.Delete
//     method.
func (client *WebAppCollectorOperationsClient) Delete(ctx context.Context, resourceGroupName string, projectName string, collectorName string, options *WebAppCollectorOperationsClientDeleteOptions) (WebAppCollectorOperationsClientDeleteResponse, error) {
	var err error
	const operationName = "WebAppCollectorOperationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, collectorName, options)
	if err != nil {
		return WebAppCollectorOperationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebAppCollectorOperationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WebAppCollectorOperationsClientDeleteResponse{}, err
	}
	return WebAppCollectorOperationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WebAppCollectorOperationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, collectorName string, options *WebAppCollectorOperationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/webAppCollectors/{collectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if collectorName == "" {
		return nil, errors.New("parameter collectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectorName}", url.PathEscape(collectorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a WebAppCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - collectorName - Web app collector ARM name.
//   - options - WebAppCollectorOperationsClientGetOptions contains the optional parameters for the WebAppCollectorOperationsClient.Get
//     method.
func (client *WebAppCollectorOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, collectorName string, options *WebAppCollectorOperationsClientGetOptions) (WebAppCollectorOperationsClientGetResponse, error) {
	var err error
	const operationName = "WebAppCollectorOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, collectorName, options)
	if err != nil {
		return WebAppCollectorOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebAppCollectorOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebAppCollectorOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WebAppCollectorOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, collectorName string, options *WebAppCollectorOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/webAppCollectors/{collectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if collectorName == "" {
		return nil, errors.New("parameter collectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectorName}", url.PathEscape(collectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebAppCollectorOperationsClient) getHandleResponse(resp *http.Response) (WebAppCollectorOperationsClientGetResponse, error) {
	result := WebAppCollectorOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppCollector); err != nil {
		return WebAppCollectorOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAssessmentProjectPager - List WebAppCollector resources by AssessmentProject
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - options - WebAppCollectorOperationsClientListByAssessmentProjectOptions contains the optional parameters for the WebAppCollectorOperationsClient.NewListByAssessmentProjectPager
//     method.
func (client *WebAppCollectorOperationsClient) NewListByAssessmentProjectPager(resourceGroupName string, projectName string, options *WebAppCollectorOperationsClientListByAssessmentProjectOptions) *runtime.Pager[WebAppCollectorOperationsClientListByAssessmentProjectResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebAppCollectorOperationsClientListByAssessmentProjectResponse]{
		More: func(page WebAppCollectorOperationsClientListByAssessmentProjectResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebAppCollectorOperationsClientListByAssessmentProjectResponse) (WebAppCollectorOperationsClientListByAssessmentProjectResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WebAppCollectorOperationsClient.NewListByAssessmentProjectPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAssessmentProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			}, nil)
			if err != nil {
				return WebAppCollectorOperationsClientListByAssessmentProjectResponse{}, err
			}
			return client.listByAssessmentProjectHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAssessmentProjectCreateRequest creates the ListByAssessmentProject request.
func (client *WebAppCollectorOperationsClient) listByAssessmentProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *WebAppCollectorOperationsClientListByAssessmentProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/webAppCollectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAssessmentProjectHandleResponse handles the ListByAssessmentProject response.
func (client *WebAppCollectorOperationsClient) listByAssessmentProjectHandleResponse(resp *http.Response) (WebAppCollectorOperationsClientListByAssessmentProjectResponse, error) {
	result := WebAppCollectorOperationsClientListByAssessmentProjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppCollectorListResult); err != nil {
		return WebAppCollectorOperationsClientListByAssessmentProjectResponse{}, err
	}
	return result, nil
}
