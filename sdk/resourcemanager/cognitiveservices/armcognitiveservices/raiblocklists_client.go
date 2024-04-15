//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

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

// RaiBlocklistsClient contains the methods for the RaiBlocklists group.
// Don't use this type directly, use NewRaiBlocklistsClient() instead.
type RaiBlocklistsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRaiBlocklistsClient creates a new instance of RaiBlocklistsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRaiBlocklistsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RaiBlocklistsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RaiBlocklistsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Update the state of specified blocklist associated with the Azure OpenAI account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - raiBlocklistName - The name of the RaiBlocklist associated with the Cognitive Services Account
//   - raiBlocklist - Properties describing the custom blocklist.
//   - options - RaiBlocklistsClientCreateOrUpdateOptions contains the optional parameters for the RaiBlocklistsClient.CreateOrUpdate
//     method.
func (client *RaiBlocklistsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, raiBlocklist RaiBlocklist, options *RaiBlocklistsClientCreateOrUpdateOptions) (RaiBlocklistsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "RaiBlocklistsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, raiBlocklistName, raiBlocklist, options)
	if err != nil {
		return RaiBlocklistsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RaiBlocklistsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return RaiBlocklistsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RaiBlocklistsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, raiBlocklist RaiBlocklist, options *RaiBlocklistsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiBlocklists/{raiBlocklistName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if raiBlocklistName == "" {
		return nil, errors.New("parameter raiBlocklistName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistName}", url.PathEscape(raiBlocklistName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, raiBlocklist); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RaiBlocklistsClient) createOrUpdateHandleResponse(resp *http.Response) (RaiBlocklistsClientCreateOrUpdateResponse, error) {
	result := RaiBlocklistsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RaiBlocklist); err != nil {
		return RaiBlocklistsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the specified custom blocklist associated with the Azure OpenAI account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - raiBlocklistName - The name of the RaiBlocklist associated with the Cognitive Services Account
//   - options - RaiBlocklistsClientBeginDeleteOptions contains the optional parameters for the RaiBlocklistsClient.BeginDelete
//     method.
func (client *RaiBlocklistsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, options *RaiBlocklistsClientBeginDeleteOptions) (*runtime.Poller[RaiBlocklistsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, raiBlocklistName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RaiBlocklistsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RaiBlocklistsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified custom blocklist associated with the Azure OpenAI account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *RaiBlocklistsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, options *RaiBlocklistsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RaiBlocklistsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, raiBlocklistName, options)
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
func (client *RaiBlocklistsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, options *RaiBlocklistsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiBlocklists/{raiBlocklistName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if raiBlocklistName == "" {
		return nil, errors.New("parameter raiBlocklistName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistName}", url.PathEscape(raiBlocklistName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified custom blocklist associated with the Azure OpenAI account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - raiBlocklistName - The name of the RaiBlocklist associated with the Cognitive Services Account
//   - options - RaiBlocklistsClientGetOptions contains the optional parameters for the RaiBlocklistsClient.Get method.
func (client *RaiBlocklistsClient) Get(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, options *RaiBlocklistsClientGetOptions) (RaiBlocklistsClientGetResponse, error) {
	var err error
	const operationName = "RaiBlocklistsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, raiBlocklistName, options)
	if err != nil {
		return RaiBlocklistsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RaiBlocklistsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RaiBlocklistsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RaiBlocklistsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, options *RaiBlocklistsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiBlocklists/{raiBlocklistName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if raiBlocklistName == "" {
		return nil, errors.New("parameter raiBlocklistName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistName}", url.PathEscape(raiBlocklistName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RaiBlocklistsClient) getHandleResponse(resp *http.Response) (RaiBlocklistsClientGetResponse, error) {
	result := RaiBlocklistsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RaiBlocklist); err != nil {
		return RaiBlocklistsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the custom blocklists associated with the Azure OpenAI account.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - options - RaiBlocklistsClientListOptions contains the optional parameters for the RaiBlocklistsClient.NewListPager method.
func (client *RaiBlocklistsClient) NewListPager(resourceGroupName string, accountName string, options *RaiBlocklistsClientListOptions) *runtime.Pager[RaiBlocklistsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RaiBlocklistsClientListResponse]{
		More: func(page RaiBlocklistsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RaiBlocklistsClientListResponse) (RaiBlocklistsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RaiBlocklistsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return RaiBlocklistsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RaiBlocklistsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *RaiBlocklistsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiBlocklists"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RaiBlocklistsClient) listHandleResponse(resp *http.Response) (RaiBlocklistsClientListResponse, error) {
	result := RaiBlocklistsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RaiBlockListResult); err != nil {
		return RaiBlocklistsClientListResponse{}, err
	}
	return result, nil
}
