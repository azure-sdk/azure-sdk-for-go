// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ReachabilityAnalysisIntentsClient contains the methods for the ReachabilityAnalysisIntents group.
// Don't use this type directly, use NewReachabilityAnalysisIntentsClient() instead.
type ReachabilityAnalysisIntentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReachabilityAnalysisIntentsClient creates a new instance of ReachabilityAnalysisIntentsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReachabilityAnalysisIntentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReachabilityAnalysisIntentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReachabilityAnalysisIntentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates Reachability Analysis Intent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group.
//   - networkManagerName - The name of the network manager.
//   - workspaceName - Workspace name.
//   - reachabilityAnalysisIntentName - Reachability Analysis Intent name.
//   - body - Reachability Analysis Intent object to create/update.
//   - options - ReachabilityAnalysisIntentsClientCreateOptions contains the optional parameters for the ReachabilityAnalysisIntentsClient.Create
//     method.
func (client *ReachabilityAnalysisIntentsClient) Create(ctx context.Context, resourceGroupName string, networkManagerName string, workspaceName string, reachabilityAnalysisIntentName string, body ReachabilityAnalysisIntent, options *ReachabilityAnalysisIntentsClientCreateOptions) (ReachabilityAnalysisIntentsClientCreateResponse, error) {
	var err error
	const operationName = "ReachabilityAnalysisIntentsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkManagerName, workspaceName, reachabilityAnalysisIntentName, body, options)
	if err != nil {
		return ReachabilityAnalysisIntentsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReachabilityAnalysisIntentsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ReachabilityAnalysisIntentsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ReachabilityAnalysisIntentsClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, workspaceName string, reachabilityAnalysisIntentName string, body ReachabilityAnalysisIntent, _ *ReachabilityAnalysisIntentsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/verifierWorkspaces/{workspaceName}/reachabilityAnalysisIntents/{reachabilityAnalysisIntentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if reachabilityAnalysisIntentName == "" {
		return nil, errors.New("parameter reachabilityAnalysisIntentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reachabilityAnalysisIntentName}", url.PathEscape(reachabilityAnalysisIntentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ReachabilityAnalysisIntentsClient) createHandleResponse(resp *http.Response) (ReachabilityAnalysisIntentsClientCreateResponse, error) {
	result := ReachabilityAnalysisIntentsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReachabilityAnalysisIntent); err != nil {
		return ReachabilityAnalysisIntentsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes Reachability Analysis Intent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group.
//   - networkManagerName - The name of the network manager.
//   - workspaceName - Workspace name.
//   - reachabilityAnalysisIntentName - Reachability Analysis Intent name.
//   - options - ReachabilityAnalysisIntentsClientDeleteOptions contains the optional parameters for the ReachabilityAnalysisIntentsClient.Delete
//     method.
func (client *ReachabilityAnalysisIntentsClient) Delete(ctx context.Context, resourceGroupName string, networkManagerName string, workspaceName string, reachabilityAnalysisIntentName string, options *ReachabilityAnalysisIntentsClientDeleteOptions) (ReachabilityAnalysisIntentsClientDeleteResponse, error) {
	var err error
	const operationName = "ReachabilityAnalysisIntentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkManagerName, workspaceName, reachabilityAnalysisIntentName, options)
	if err != nil {
		return ReachabilityAnalysisIntentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReachabilityAnalysisIntentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ReachabilityAnalysisIntentsClientDeleteResponse{}, err
	}
	return ReachabilityAnalysisIntentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReachabilityAnalysisIntentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, workspaceName string, reachabilityAnalysisIntentName string, _ *ReachabilityAnalysisIntentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/verifierWorkspaces/{workspaceName}/reachabilityAnalysisIntents/{reachabilityAnalysisIntentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if reachabilityAnalysisIntentName == "" {
		return nil, errors.New("parameter reachabilityAnalysisIntentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reachabilityAnalysisIntentName}", url.PathEscape(reachabilityAnalysisIntentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Reachability Analysis Intent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group.
//   - networkManagerName - The name of the network manager.
//   - workspaceName - Workspace name.
//   - reachabilityAnalysisIntentName - Reachability Analysis Intent name.
//   - options - ReachabilityAnalysisIntentsClientGetOptions contains the optional parameters for the ReachabilityAnalysisIntentsClient.Get
//     method.
func (client *ReachabilityAnalysisIntentsClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string, workspaceName string, reachabilityAnalysisIntentName string, options *ReachabilityAnalysisIntentsClientGetOptions) (ReachabilityAnalysisIntentsClientGetResponse, error) {
	var err error
	const operationName = "ReachabilityAnalysisIntentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkManagerName, workspaceName, reachabilityAnalysisIntentName, options)
	if err != nil {
		return ReachabilityAnalysisIntentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReachabilityAnalysisIntentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReachabilityAnalysisIntentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReachabilityAnalysisIntentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, workspaceName string, reachabilityAnalysisIntentName string, _ *ReachabilityAnalysisIntentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/verifierWorkspaces/{workspaceName}/reachabilityAnalysisIntents/{reachabilityAnalysisIntentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if reachabilityAnalysisIntentName == "" {
		return nil, errors.New("parameter reachabilityAnalysisIntentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reachabilityAnalysisIntentName}", url.PathEscape(reachabilityAnalysisIntentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReachabilityAnalysisIntentsClient) getHandleResponse(resp *http.Response) (ReachabilityAnalysisIntentsClientGetResponse, error) {
	result := ReachabilityAnalysisIntentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReachabilityAnalysisIntent); err != nil {
		return ReachabilityAnalysisIntentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets list of Reachability Analysis Intents .
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group.
//   - networkManagerName - The name of the network manager.
//   - workspaceName - Workspace name.
//   - options - ReachabilityAnalysisIntentsClientListOptions contains the optional parameters for the ReachabilityAnalysisIntentsClient.NewListPager
//     method.
func (client *ReachabilityAnalysisIntentsClient) NewListPager(resourceGroupName string, networkManagerName string, workspaceName string, options *ReachabilityAnalysisIntentsClientListOptions) *runtime.Pager[ReachabilityAnalysisIntentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReachabilityAnalysisIntentsClientListResponse]{
		More: func(page ReachabilityAnalysisIntentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReachabilityAnalysisIntentsClientListResponse) (ReachabilityAnalysisIntentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReachabilityAnalysisIntentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, networkManagerName, workspaceName, options)
			}, nil)
			if err != nil {
				return ReachabilityAnalysisIntentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ReachabilityAnalysisIntentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, workspaceName string, options *ReachabilityAnalysisIntentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/verifierWorkspaces/{workspaceName}/reachabilityAnalysisIntents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("skipToken", *options.SkipToken)
	}
	if options != nil && options.SortKey != nil {
		reqQP.Set("sortKey", *options.SortKey)
	}
	if options != nil && options.SortValue != nil {
		reqQP.Set("sortValue", *options.SortValue)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReachabilityAnalysisIntentsClient) listHandleResponse(resp *http.Response) (ReachabilityAnalysisIntentsClientListResponse, error) {
	result := ReachabilityAnalysisIntentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReachabilityAnalysisIntentListResult); err != nil {
		return ReachabilityAnalysisIntentsClientListResponse{}, err
	}
	return result, nil
}
