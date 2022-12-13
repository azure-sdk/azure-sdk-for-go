//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazurestackhci

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// UpdatesClient contains the methods for the Updates group.
// Don't use this type directly, use NewUpdatesClient() instead.
type UpdatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewUpdatesClient creates a new instance of UpdatesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUpdatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UpdatesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &UpdatesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginDelete - Delete specified Update
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-30-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the cluster.
// updateName - The name of the Update
// options - UpdatesClientBeginDeleteOptions contains the optional parameters for the UpdatesClient.BeginDelete method.
func (client *UpdatesClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, updateName string, options *UpdatesClientBeginDeleteOptions) (*runtime.Poller[UpdatesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, updateName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[UpdatesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[UpdatesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete specified Update
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-30-preview
func (client *UpdatesClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, updateName string, options *UpdatesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, updateName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *UpdatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, updateName string, options *UpdatesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/updates/{updateName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get specified Update
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-30-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the cluster.
// updateName - The name of the Update
// options - UpdatesClientGetOptions contains the optional parameters for the UpdatesClient.Get method.
func (client *UpdatesClient) Get(ctx context.Context, resourceGroupName string, clusterName string, updateName string, options *UpdatesClientGetOptions) (UpdatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, updateName, options)
	if err != nil {
		return UpdatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UpdatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UpdatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *UpdatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, updateName string, options *UpdatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/updates/{updateName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UpdatesClient) getHandleResponse(resp *http.Response) (UpdatesClientGetResponse, error) {
	result := UpdatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Update); err != nil {
		return UpdatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all Updates
// Generated from API version 2022-12-30-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the cluster.
// options - UpdatesClientListOptions contains the optional parameters for the UpdatesClient.List method.
func (client *UpdatesClient) NewListPager(resourceGroupName string, clusterName string, options *UpdatesClientListOptions) *runtime.Pager[UpdatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[UpdatesClientListResponse]{
		More: func(page UpdatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *UpdatesClientListResponse) (UpdatesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return UpdatesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return UpdatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UpdatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *UpdatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *UpdatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/updates"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UpdatesClient) listHandleResponse(resp *http.Response) (UpdatesClientListResponse, error) {
	result := UpdatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateList); err != nil {
		return UpdatesClientListResponse{}, err
	}
	return result, nil
}

// BeginPost - Apply Update
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-30-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the cluster.
// updateName - The name of the Update
// options - UpdatesClientBeginPostOptions contains the optional parameters for the UpdatesClient.BeginPost method.
func (client *UpdatesClient) BeginPost(ctx context.Context, resourceGroupName string, clusterName string, updateName string, options *UpdatesClientBeginPostOptions) (*runtime.Poller[UpdatesClientPostResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.post(ctx, resourceGroupName, clusterName, updateName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[UpdatesClientPostResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[UpdatesClientPostResponse](options.ResumeToken, client.pl, nil)
	}
}

// Post - Apply Update
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-30-preview
func (client *UpdatesClient) post(ctx context.Context, resourceGroupName string, clusterName string, updateName string, options *UpdatesClientBeginPostOptions) (*http.Response, error) {
	req, err := client.postCreateRequest(ctx, resourceGroupName, clusterName, updateName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// postCreateRequest creates the Post request.
func (client *UpdatesClient) postCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, updateName string, options *UpdatesClientBeginPostOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/updates/{updateName}/apply"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Put - Put specified Update
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-30-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the cluster.
// updateName - The name of the Update
// updateProperties - Properties of the Updates object
// options - UpdatesClientPutOptions contains the optional parameters for the UpdatesClient.Put method.
func (client *UpdatesClient) Put(ctx context.Context, resourceGroupName string, clusterName string, updateName string, updateProperties Update, options *UpdatesClientPutOptions) (UpdatesClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, resourceGroupName, clusterName, updateName, updateProperties, options)
	if err != nil {
		return UpdatesClientPutResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UpdatesClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UpdatesClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *UpdatesClient) putCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, updateName string, updateProperties Update, options *UpdatesClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/updates/{updateName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, updateProperties)
}

// putHandleResponse handles the Put response.
func (client *UpdatesClient) putHandleResponse(resp *http.Response) (UpdatesClientPutResponse, error) {
	result := UpdatesClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Update); err != nil {
		return UpdatesClientPutResponse{}, err
	}
	return result, nil
}
