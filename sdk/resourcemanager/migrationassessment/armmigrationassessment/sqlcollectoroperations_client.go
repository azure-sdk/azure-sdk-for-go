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

// SQLCollectorOperationsClient contains the methods for the SQLCollectorOperations group.
// Don't use this type directly, use NewSQLCollectorOperationsClient() instead.
type SQLCollectorOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLCollectorOperationsClient creates a new instance of SQLCollectorOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLCollectorOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLCollectorOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLCollectorOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a SqlCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - collectorName - Sql collector ARM name.
//   - resource - Resource create parameters.
//   - options - SQLCollectorOperationsClientBeginCreateOptions contains the optional parameters for the SQLCollectorOperationsClient.BeginCreate
//     method.
func (client *SQLCollectorOperationsClient) BeginCreate(ctx context.Context, resourceGroupName string, projectName string, collectorName string, resource SQLCollector, options *SQLCollectorOperationsClientBeginCreateOptions) (*runtime.Poller[SQLCollectorOperationsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, projectName, collectorName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SQLCollectorOperationsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SQLCollectorOperationsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a SqlCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *SQLCollectorOperationsClient) create(ctx context.Context, resourceGroupName string, projectName string, collectorName string, resource SQLCollector, options *SQLCollectorOperationsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "SQLCollectorOperationsClient.BeginCreate"
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
func (client *SQLCollectorOperationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, projectName string, collectorName string, resource SQLCollector, options *SQLCollectorOperationsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/sqlcollectors/{collectorName}"
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

// Delete - Delete a SqlCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - collectorName - Sql collector ARM name.
//   - options - SQLCollectorOperationsClientDeleteOptions contains the optional parameters for the SQLCollectorOperationsClient.Delete
//     method.
func (client *SQLCollectorOperationsClient) Delete(ctx context.Context, resourceGroupName string, projectName string, collectorName string, options *SQLCollectorOperationsClientDeleteOptions) (SQLCollectorOperationsClientDeleteResponse, error) {
	var err error
	const operationName = "SQLCollectorOperationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, collectorName, options)
	if err != nil {
		return SQLCollectorOperationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLCollectorOperationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SQLCollectorOperationsClientDeleteResponse{}, err
	}
	return SQLCollectorOperationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLCollectorOperationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, collectorName string, options *SQLCollectorOperationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/sqlcollectors/{collectorName}"
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

// Get - Get a SqlCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - collectorName - Sql collector ARM name.
//   - options - SQLCollectorOperationsClientGetOptions contains the optional parameters for the SQLCollectorOperationsClient.Get
//     method.
func (client *SQLCollectorOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, collectorName string, options *SQLCollectorOperationsClientGetOptions) (SQLCollectorOperationsClientGetResponse, error) {
	var err error
	const operationName = "SQLCollectorOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, collectorName, options)
	if err != nil {
		return SQLCollectorOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLCollectorOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLCollectorOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SQLCollectorOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, collectorName string, options *SQLCollectorOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/sqlcollectors/{collectorName}"
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
func (client *SQLCollectorOperationsClient) getHandleResponse(resp *http.Response) (SQLCollectorOperationsClientGetResponse, error) {
	result := SQLCollectorOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLCollector); err != nil {
		return SQLCollectorOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAssessmentProjectPager - List SqlCollector resources by AssessmentProject
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - options - SQLCollectorOperationsClientListByAssessmentProjectOptions contains the optional parameters for the SQLCollectorOperationsClient.NewListByAssessmentProjectPager
//     method.
func (client *SQLCollectorOperationsClient) NewListByAssessmentProjectPager(resourceGroupName string, projectName string, options *SQLCollectorOperationsClientListByAssessmentProjectOptions) *runtime.Pager[SQLCollectorOperationsClientListByAssessmentProjectResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLCollectorOperationsClientListByAssessmentProjectResponse]{
		More: func(page SQLCollectorOperationsClientListByAssessmentProjectResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLCollectorOperationsClientListByAssessmentProjectResponse) (SQLCollectorOperationsClientListByAssessmentProjectResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SQLCollectorOperationsClient.NewListByAssessmentProjectPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAssessmentProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			}, nil)
			if err != nil {
				return SQLCollectorOperationsClientListByAssessmentProjectResponse{}, err
			}
			return client.listByAssessmentProjectHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAssessmentProjectCreateRequest creates the ListByAssessmentProject request.
func (client *SQLCollectorOperationsClient) listByAssessmentProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *SQLCollectorOperationsClientListByAssessmentProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/sqlcollectors"
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
func (client *SQLCollectorOperationsClient) listByAssessmentProjectHandleResponse(resp *http.Response) (SQLCollectorOperationsClientListByAssessmentProjectResponse, error) {
	result := SQLCollectorOperationsClientListByAssessmentProjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLCollectorListResult); err != nil {
		return SQLCollectorOperationsClientListByAssessmentProjectResponse{}, err
	}
	return result, nil
}
