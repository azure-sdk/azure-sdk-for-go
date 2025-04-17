// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

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

// ApplicationTypesClient contains the methods for the ApplicationTypes group.
// Don't use this type directly, use NewApplicationTypesClient() instead.
type ApplicationTypesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplicationTypesClient creates a new instance of ApplicationTypesClient with the specified values.
//   - subscriptionID - The customer subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationTypesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationTypesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a Service Fabric application type name resource with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - applicationTypeName - The name of the application type name resource.
//   - parameters - The application type name resource.
//   - options - ApplicationTypesClientCreateOrUpdateOptions contains the optional parameters for the ApplicationTypesClient.CreateOrUpdate
//     method.
func (client *ApplicationTypesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, parameters ApplicationTypeResource, options *ApplicationTypesClientCreateOrUpdateOptions) (ApplicationTypesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ApplicationTypesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, parameters, options)
	if err != nil {
		return ApplicationTypesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationTypesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationTypesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationTypesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, parameters ApplicationTypeResource, _ *ApplicationTypesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}"
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
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ApplicationTypesClient) createOrUpdateHandleResponse(resp *http.Response) (ApplicationTypesClientCreateOrUpdateResponse, error) {
	result := ApplicationTypesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationTypeResource); err != nil {
		return ApplicationTypesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete a Service Fabric application type name resource with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - applicationTypeName - The name of the application type name resource.
//   - options - ApplicationTypesClientBeginDeleteOptions contains the optional parameters for the ApplicationTypesClient.BeginDelete
//     method.
func (client *ApplicationTypesClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypesClientBeginDeleteOptions) (*runtime.Poller[ApplicationTypesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, applicationTypeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ApplicationTypesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ApplicationTypesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Service Fabric application type name resource with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
func (client *ApplicationTypesClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ApplicationTypesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationTypesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, _ *ApplicationTypesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}"
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
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Service Fabric application type name resource created or in the process of being created in the Service Fabric
// cluster resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - applicationTypeName - The name of the application type name resource.
//   - options - ApplicationTypesClientGetOptions contains the optional parameters for the ApplicationTypesClient.Get method.
func (client *ApplicationTypesClient) Get(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypesClientGetOptions) (ApplicationTypesClientGetResponse, error) {
	var err error
	const operationName = "ApplicationTypesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, options)
	if err != nil {
		return ApplicationTypesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationTypesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ApplicationTypesClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, _ *ApplicationTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}"
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
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationTypesClient) getHandleResponse(resp *http.Response) (ApplicationTypesClientGetResponse, error) {
	result := ApplicationTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationTypeResource); err != nil {
		return ApplicationTypesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all application type name resources created or in the process of being created in the Service Fabric
// cluster resource.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - options - ApplicationTypesClientListOptions contains the optional parameters for the ApplicationTypesClient.NewListPager
//     method.
func (client *ApplicationTypesClient) NewListPager(resourceGroupName string, clusterName string, options *ApplicationTypesClientListOptions) *runtime.Pager[ApplicationTypesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationTypesClientListResponse]{
		More: func(page ApplicationTypesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationTypesClientListResponse) (ApplicationTypesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ApplicationTypesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, clusterName, options)
			}, nil)
			if err != nil {
				return ApplicationTypesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ApplicationTypesClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, _ *ApplicationTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes"
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
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationTypesClient) listHandleResponse(resp *http.Response) (ApplicationTypesClientListResponse, error) {
	result := ApplicationTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationTypeResourceList); err != nil {
		return ApplicationTypesClientListResponse{}, err
	}
	return result, nil
}
