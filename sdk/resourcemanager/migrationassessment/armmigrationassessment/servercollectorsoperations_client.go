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

// ServerCollectorsOperationsClient contains the methods for the ServerCollectorsOperations group.
// Don't use this type directly, use NewServerCollectorsOperationsClient() instead.
type ServerCollectorsOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerCollectorsOperationsClient creates a new instance of ServerCollectorsOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerCollectorsOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerCollectorsOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerCollectorsOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a ServerCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - serverCollectorName - Physical server collector ARM name
//   - resource - Resource create parameters.
//   - options - ServerCollectorsOperationsClientBeginCreateOptions contains the optional parameters for the ServerCollectorsOperationsClient.BeginCreate
//     method.
func (client *ServerCollectorsOperationsClient) BeginCreate(ctx context.Context, resourceGroupName string, projectName string, serverCollectorName string, resource ServerCollector, options *ServerCollectorsOperationsClientBeginCreateOptions) (*runtime.Poller[ServerCollectorsOperationsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, projectName, serverCollectorName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerCollectorsOperationsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerCollectorsOperationsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a ServerCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
func (client *ServerCollectorsOperationsClient) create(ctx context.Context, resourceGroupName string, projectName string, serverCollectorName string, resource ServerCollector, options *ServerCollectorsOperationsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerCollectorsOperationsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, projectName, serverCollectorName, resource, options)
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
func (client *ServerCollectorsOperationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, projectName string, serverCollectorName string, resource ServerCollector, _ *ServerCollectorsOperationsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/servercollectors/{serverCollectorName}"
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
	if serverCollectorName == "" {
		return nil, errors.New("parameter serverCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverCollectorName}", url.PathEscape(serverCollectorName))
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

// Delete - Delete a ServerCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - serverCollectorName - Physical server collector ARM name
//   - options - ServerCollectorsOperationsClientDeleteOptions contains the optional parameters for the ServerCollectorsOperationsClient.Delete
//     method.
func (client *ServerCollectorsOperationsClient) Delete(ctx context.Context, resourceGroupName string, projectName string, serverCollectorName string, options *ServerCollectorsOperationsClientDeleteOptions) (ServerCollectorsOperationsClientDeleteResponse, error) {
	var err error
	const operationName = "ServerCollectorsOperationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, serverCollectorName, options)
	if err != nil {
		return ServerCollectorsOperationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerCollectorsOperationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServerCollectorsOperationsClientDeleteResponse{}, err
	}
	return ServerCollectorsOperationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServerCollectorsOperationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, serverCollectorName string, _ *ServerCollectorsOperationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/servercollectors/{serverCollectorName}"
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
	if serverCollectorName == "" {
		return nil, errors.New("parameter serverCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverCollectorName}", url.PathEscape(serverCollectorName))
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

// Get - Get a ServerCollector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - serverCollectorName - Physical server collector ARM name
//   - options - ServerCollectorsOperationsClientGetOptions contains the optional parameters for the ServerCollectorsOperationsClient.Get
//     method.
func (client *ServerCollectorsOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, serverCollectorName string, options *ServerCollectorsOperationsClientGetOptions) (ServerCollectorsOperationsClientGetResponse, error) {
	var err error
	const operationName = "ServerCollectorsOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, serverCollectorName, options)
	if err != nil {
		return ServerCollectorsOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerCollectorsOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerCollectorsOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerCollectorsOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, serverCollectorName string, _ *ServerCollectorsOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/servercollectors/{serverCollectorName}"
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
	if serverCollectorName == "" {
		return nil, errors.New("parameter serverCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverCollectorName}", url.PathEscape(serverCollectorName))
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
func (client *ServerCollectorsOperationsClient) getHandleResponse(resp *http.Response) (ServerCollectorsOperationsClientGetResponse, error) {
	result := ServerCollectorsOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerCollector); err != nil {
		return ServerCollectorsOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAssessmentProjectPager - List ServerCollector resources by AssessmentProject
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - options - ServerCollectorsOperationsClientListByAssessmentProjectOptions contains the optional parameters for the ServerCollectorsOperationsClient.NewListByAssessmentProjectPager
//     method.
func (client *ServerCollectorsOperationsClient) NewListByAssessmentProjectPager(resourceGroupName string, projectName string, options *ServerCollectorsOperationsClientListByAssessmentProjectOptions) *runtime.Pager[ServerCollectorsOperationsClientListByAssessmentProjectResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerCollectorsOperationsClientListByAssessmentProjectResponse]{
		More: func(page ServerCollectorsOperationsClientListByAssessmentProjectResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerCollectorsOperationsClientListByAssessmentProjectResponse) (ServerCollectorsOperationsClientListByAssessmentProjectResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerCollectorsOperationsClient.NewListByAssessmentProjectPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAssessmentProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			}, nil)
			if err != nil {
				return ServerCollectorsOperationsClientListByAssessmentProjectResponse{}, err
			}
			return client.listByAssessmentProjectHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAssessmentProjectCreateRequest creates the ListByAssessmentProject request.
func (client *ServerCollectorsOperationsClient) listByAssessmentProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, _ *ServerCollectorsOperationsClientListByAssessmentProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/servercollectors"
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
func (client *ServerCollectorsOperationsClient) listByAssessmentProjectHandleResponse(resp *http.Response) (ServerCollectorsOperationsClientListByAssessmentProjectResponse, error) {
	result := ServerCollectorsOperationsClientListByAssessmentProjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerCollectorListResult); err != nil {
		return ServerCollectorsOperationsClientListByAssessmentProjectResponse{}, err
	}
	return result, nil
}
