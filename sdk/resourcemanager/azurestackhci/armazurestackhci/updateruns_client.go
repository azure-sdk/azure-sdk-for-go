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

// UpdateRunsClient contains the methods for the UpdateRuns group.
// Don't use this type directly, use NewUpdateRunsClient() instead.
type UpdateRunsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUpdateRunsClient creates a new instance of UpdateRunsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUpdateRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UpdateRunsClient, error) {
	cl, err := arm.NewClient(moduleName+".UpdateRunsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UpdateRunsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDelete - Delete specified Update Run
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - updateName - The name of the Update
//   - updateRunName - The name of the Update Run
//   - options - UpdateRunsClientBeginDeleteOptions contains the optional parameters for the UpdateRunsClient.BeginDelete method.
func (client *UpdateRunsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, updateName string, updateRunName string, options *UpdateRunsClientBeginDeleteOptions) (*runtime.Poller[UpdateRunsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, updateName, updateRunName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[UpdateRunsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[UpdateRunsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete specified Update Run
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *UpdateRunsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, updateName string, updateRunName string, options *UpdateRunsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, updateName, updateRunName, options)
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
func (client *UpdateRunsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, updateName string, updateRunName string, options *UpdateRunsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/updates/{updateName}/updateRuns/{updateRunName}"
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
	if updateName == "" {
		return nil, errors.New("parameter updateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateName}", url.PathEscape(updateName))
	if updateRunName == "" {
		return nil, errors.New("parameter updateRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateRunName}", url.PathEscape(updateRunName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Update run for a specified update
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - updateName - The name of the Update
//   - updateRunName - The name of the Update Run
//   - options - UpdateRunsClientGetOptions contains the optional parameters for the UpdateRunsClient.Get method.
func (client *UpdateRunsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, updateName string, updateRunName string, options *UpdateRunsClientGetOptions) (UpdateRunsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, updateName, updateRunName, options)
	if err != nil {
		return UpdateRunsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UpdateRunsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UpdateRunsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *UpdateRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, updateName string, updateRunName string, options *UpdateRunsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/updates/{updateName}/updateRuns/{updateRunName}"
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
	if updateName == "" {
		return nil, errors.New("parameter updateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateName}", url.PathEscape(updateName))
	if updateRunName == "" {
		return nil, errors.New("parameter updateRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateRunName}", url.PathEscape(updateRunName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UpdateRunsClient) getHandleResponse(resp *http.Response) (UpdateRunsClientGetResponse, error) {
	result := UpdateRunsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateRun); err != nil {
		return UpdateRunsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all Update runs for a specified update
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - updateName - The name of the Update
//   - options - UpdateRunsClientListOptions contains the optional parameters for the UpdateRunsClient.NewListPager method.
func (client *UpdateRunsClient) NewListPager(resourceGroupName string, clusterName string, updateName string, options *UpdateRunsClientListOptions) *runtime.Pager[UpdateRunsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[UpdateRunsClientListResponse]{
		More: func(page UpdateRunsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *UpdateRunsClientListResponse) (UpdateRunsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, clusterName, updateName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return UpdateRunsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return UpdateRunsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UpdateRunsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *UpdateRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, updateName string, options *UpdateRunsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/updates/{updateName}/updateRuns"
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
	if updateName == "" {
		return nil, errors.New("parameter updateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateName}", url.PathEscape(updateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UpdateRunsClient) listHandleResponse(resp *http.Response) (UpdateRunsClientListResponse, error) {
	result := UpdateRunsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateRunList); err != nil {
		return UpdateRunsClientListResponse{}, err
	}
	return result, nil
}

// Put - Put Update runs for a specified update
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - updateName - The name of the Update
//   - updateRunName - The name of the Update Run
//   - updateRunsProperties - Properties of the updateRuns object
//   - options - UpdateRunsClientPutOptions contains the optional parameters for the UpdateRunsClient.Put method.
func (client *UpdateRunsClient) Put(ctx context.Context, resourceGroupName string, clusterName string, updateName string, updateRunName string, updateRunsProperties UpdateRun, options *UpdateRunsClientPutOptions) (UpdateRunsClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, resourceGroupName, clusterName, updateName, updateRunName, updateRunsProperties, options)
	if err != nil {
		return UpdateRunsClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UpdateRunsClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UpdateRunsClientPutResponse{}, err
	}
	resp, err := client.putHandleResponse(httpResp)
	return resp, err
}

// putCreateRequest creates the Put request.
func (client *UpdateRunsClient) putCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, updateName string, updateRunName string, updateRunsProperties UpdateRun, options *UpdateRunsClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/updates/{updateName}/updateRuns/{updateRunName}"
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
	if updateName == "" {
		return nil, errors.New("parameter updateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateName}", url.PathEscape(updateName))
	if updateRunName == "" {
		return nil, errors.New("parameter updateRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateRunName}", url.PathEscape(updateRunName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateRunsProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *UpdateRunsClient) putHandleResponse(resp *http.Response) (UpdateRunsClientPutResponse, error) {
	result := UpdateRunsClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateRun); err != nil {
		return UpdateRunsClientPutResponse{}, err
	}
	return result, nil
}
