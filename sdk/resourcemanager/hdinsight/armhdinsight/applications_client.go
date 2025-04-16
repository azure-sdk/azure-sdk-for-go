// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight

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

// ApplicationsClient contains the methods for the Applications group.
// Don't use this type directly, use NewApplicationsClient() instead.
type ApplicationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplicationsClient creates a new instance of ApplicationsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates applications for the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-15-preview
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster.
//   - applicationName - The constant value for the application name.
//   - parameters - The application create request.
//   - options - ApplicationsClientBeginCreateOptions contains the optional parameters for the ApplicationsClient.BeginCreate
//     method.
func (client *ApplicationsClient) BeginCreate(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters Application, options *ApplicationsClientBeginCreateOptions) (*runtime.Poller[ApplicationsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, clusterName, applicationName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ApplicationsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ApplicationsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates applications for the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-15-preview
func (client *ApplicationsClient) create(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters Application, options *ApplicationsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ApplicationsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterName, applicationName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *ApplicationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters Application, _ *ApplicationsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified application on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-15-preview
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster.
//   - applicationName - The constant value for the application name.
//   - options - ApplicationsClientBeginDeleteOptions contains the optional parameters for the ApplicationsClient.BeginDelete
//     method.
func (client *ApplicationsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, options *ApplicationsClientBeginDeleteOptions) (*runtime.Poller[ApplicationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, applicationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ApplicationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ApplicationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified application on the HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-15-preview
func (client *ApplicationsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, options *ApplicationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ApplicationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, applicationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, _ *ApplicationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets properties of the specified application.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-15-preview
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster.
//   - applicationName - The constant value for the application name.
//   - options - ApplicationsClientGetOptions contains the optional parameters for the ApplicationsClient.Get method.
func (client *ApplicationsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, options *ApplicationsClientGetOptions) (ApplicationsClientGetResponse, error) {
	var err error
	const operationName = "ApplicationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, applicationName, options)
	if err != nil {
		return ApplicationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ApplicationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, _ *ApplicationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationsClient) getHandleResponse(resp *http.Response) (ApplicationsClientGetResponse, error) {
	result := ApplicationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Application); err != nil {
		return ApplicationsClientGetResponse{}, err
	}
	return result, nil
}

// GetAzureAsyncOperationStatus - Gets the async operation status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-15-preview
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster.
//   - applicationName - The constant value for the application name.
//   - operationID - The long running operation id.
//   - options - ApplicationsClientGetAzureAsyncOperationStatusOptions contains the optional parameters for the ApplicationsClient.GetAzureAsyncOperationStatus
//     method.
func (client *ApplicationsClient) GetAzureAsyncOperationStatus(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, operationID string, options *ApplicationsClientGetAzureAsyncOperationStatusOptions) (ApplicationsClientGetAzureAsyncOperationStatusResponse, error) {
	var err error
	const operationName = "ApplicationsClient.GetAzureAsyncOperationStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAzureAsyncOperationStatusCreateRequest(ctx, resourceGroupName, clusterName, applicationName, operationID, options)
	if err != nil {
		return ApplicationsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	resp, err := client.getAzureAsyncOperationStatusHandleResponse(httpResp)
	return resp, err
}

// getAzureAsyncOperationStatusCreateRequest creates the GetAzureAsyncOperationStatus request.
func (client *ApplicationsClient) getAzureAsyncOperationStatusCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, operationID string, _ *ApplicationsClientGetAzureAsyncOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}/azureasyncoperations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAzureAsyncOperationStatusHandleResponse handles the GetAzureAsyncOperationStatus response.
func (client *ApplicationsClient) getAzureAsyncOperationStatusHandleResponse(resp *http.Response) (ApplicationsClientGetAzureAsyncOperationStatusResponse, error) {
	result := ApplicationsClientGetAzureAsyncOperationStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AsyncOperationResult); err != nil {
		return ApplicationsClientGetAzureAsyncOperationStatusResponse{}, err
	}
	return result, nil
}

// NewListByClusterPager - Lists all of the applications for the HDInsight cluster.
//
// Generated from API version 2025-01-15-preview
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster.
//   - options - ApplicationsClientListByClusterOptions contains the optional parameters for the ApplicationsClient.NewListByClusterPager
//     method.
func (client *ApplicationsClient) NewListByClusterPager(resourceGroupName string, clusterName string, options *ApplicationsClientListByClusterOptions) *runtime.Pager[ApplicationsClientListByClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationsClientListByClusterResponse]{
		More: func(page ApplicationsClientListByClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationsClientListByClusterResponse) (ApplicationsClientListByClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ApplicationsClient.NewListByClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
			}, nil)
			if err != nil {
				return ApplicationsClientListByClusterResponse{}, err
			}
			return client.listByClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *ApplicationsClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, _ *ApplicationsClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *ApplicationsClient) listByClusterHandleResponse(resp *http.Response) (ApplicationsClientListByClusterResponse, error) {
	result := ApplicationsClientListByClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationListResult); err != nil {
		return ApplicationsClientListByClusterResponse{}, err
	}
	return result, nil
}
