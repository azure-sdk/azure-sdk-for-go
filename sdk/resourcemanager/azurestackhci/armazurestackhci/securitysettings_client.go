//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

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

// SecuritySettingsClient contains the methods for the SecuritySettings group.
// Don't use this type directly, use NewSecuritySettingsClient() instead.
type SecuritySettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecuritySettingsClient creates a new instance of SecuritySettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecuritySettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecuritySettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecuritySettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a security setting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - securitySettingsName - Name of security setting
//   - resource - Resource create parameters.
//   - options - SecuritySettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the SecuritySettingsClient.BeginCreateOrUpdate
//     method.
func (client *SecuritySettingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, securitySettingsName string, resource SecuritySetting, options *SecuritySettingsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SecuritySettingsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, securitySettingsName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecuritySettingsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SecuritySettingsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a security setting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *SecuritySettingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, securitySettingsName string, resource SecuritySetting, options *SecuritySettingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SecuritySettingsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, securitySettingsName, resource, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecuritySettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, securitySettingsName string, resource SecuritySetting, options *SecuritySettingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/securitySettings/{securitySettingsName}"
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
	if securitySettingsName == "" {
		return nil, errors.New("parameter securitySettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securitySettingsName}", url.PathEscape(securitySettingsName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a SecuritySetting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - securitySettingsName - Name of security setting
//   - options - SecuritySettingsClientBeginDeleteOptions contains the optional parameters for the SecuritySettingsClient.BeginDelete
//     method.
func (client *SecuritySettingsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, securitySettingsName string, options *SecuritySettingsClientBeginDeleteOptions) (*runtime.Poller[SecuritySettingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, securitySettingsName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecuritySettingsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SecuritySettingsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a SecuritySetting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *SecuritySettingsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, securitySettingsName string, options *SecuritySettingsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SecuritySettingsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, securitySettingsName, options)
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
func (client *SecuritySettingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, securitySettingsName string, options *SecuritySettingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/securitySettings/{securitySettingsName}"
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
	if securitySettingsName == "" {
		return nil, errors.New("parameter securitySettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securitySettingsName}", url.PathEscape(securitySettingsName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a SecuritySetting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - securitySettingsName - Name of security setting
//   - options - SecuritySettingsClientGetOptions contains the optional parameters for the SecuritySettingsClient.Get method.
func (client *SecuritySettingsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, securitySettingsName string, options *SecuritySettingsClientGetOptions) (SecuritySettingsClientGetResponse, error) {
	var err error
	const operationName = "SecuritySettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, securitySettingsName, options)
	if err != nil {
		return SecuritySettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecuritySettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecuritySettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SecuritySettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, securitySettingsName string, options *SecuritySettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/securitySettings/{securitySettingsName}"
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
	if securitySettingsName == "" {
		return nil, errors.New("parameter securitySettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securitySettingsName}", url.PathEscape(securitySettingsName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecuritySettingsClient) getHandleResponse(resp *http.Response) (SecuritySettingsClientGetResponse, error) {
	result := SecuritySettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecuritySetting); err != nil {
		return SecuritySettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByClustersPager - List SecuritySetting resources by Clusters
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - SecuritySettingsClientListByClustersOptions contains the optional parameters for the SecuritySettingsClient.NewListByClustersPager
//     method.
func (client *SecuritySettingsClient) NewListByClustersPager(resourceGroupName string, clusterName string, options *SecuritySettingsClientListByClustersOptions) *runtime.Pager[SecuritySettingsClientListByClustersResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecuritySettingsClientListByClustersResponse]{
		More: func(page SecuritySettingsClientListByClustersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecuritySettingsClientListByClustersResponse) (SecuritySettingsClientListByClustersResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecuritySettingsClient.NewListByClustersPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByClustersCreateRequest(ctx, resourceGroupName, clusterName, options)
			}, nil)
			if err != nil {
				return SecuritySettingsClientListByClustersResponse{}, err
			}
			return client.listByClustersHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByClustersCreateRequest creates the ListByClusters request.
func (client *SecuritySettingsClient) listByClustersCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *SecuritySettingsClientListByClustersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/securitySettings"
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
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClustersHandleResponse handles the ListByClusters response.
func (client *SecuritySettingsClient) listByClustersHandleResponse(resp *http.Response) (SecuritySettingsClientListByClustersResponse, error) {
	result := SecuritySettingsClientListByClustersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecuritySettingListResult); err != nil {
		return SecuritySettingsClientListByClustersResponse{}, err
	}
	return result, nil
}
