//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbotservice

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

// ChannelsClient contains the methods for the Channels group.
// Don't use this type directly, use NewChannelsClient() instead.
type ChannelsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewChannelsClient creates a new instance of ChannelsClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewChannelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ChannelsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ChannelsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates a Channel registration for a Bot Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - channelName - The name of the Channel resource.
//   - parameters - The parameters to provide for the created bot.
//   - options - ChannelsClientCreateOptions contains the optional parameters for the ChannelsClient.Create method.
func (client *ChannelsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, parameters BotChannel, options *ChannelsClientCreateOptions) (ChannelsClientCreateResponse, error) {
	var err error
	const operationName = "ChannelsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, channelName, parameters, options)
	if err != nil {
		return ChannelsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChannelsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ChannelsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ChannelsClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, parameters BotChannel, options *ChannelsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(string(channelName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ChannelsClient) createHandleResponse(resp *http.Response) (ChannelsClientCreateResponse, error) {
	result := ChannelsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BotChannel); err != nil {
		return ChannelsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Channel registration from a Bot Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - channelName - The name of the Channel resource.
//   - options - ChannelsClientDeleteOptions contains the optional parameters for the ChannelsClient.Delete method.
func (client *ChannelsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, options *ChannelsClientDeleteOptions) (ChannelsClientDeleteResponse, error) {
	var err error
	const operationName = "ChannelsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, channelName, options)
	if err != nil {
		return ChannelsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChannelsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ChannelsClientDeleteResponse{}, err
	}
	return ChannelsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ChannelsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, options *ChannelsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(string(channelName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns a BotService Channel registration specified by the parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - channelName - The name of the Channel resource.
//   - options - ChannelsClientGetOptions contains the optional parameters for the ChannelsClient.Get method.
func (client *ChannelsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, options *ChannelsClientGetOptions) (ChannelsClientGetResponse, error) {
	var err error
	const operationName = "ChannelsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, channelName, options)
	if err != nil {
		return ChannelsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChannelsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ChannelsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ChannelsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, options *ChannelsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(string(channelName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ChannelsClient) getHandleResponse(resp *http.Response) (ChannelsClientGetResponse, error) {
	result := ChannelsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BotChannel); err != nil {
		return ChannelsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Returns all the Channel registrations of a particular BotService resource
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - options - ChannelsClientListByResourceGroupOptions contains the optional parameters for the ChannelsClient.NewListByResourceGroupPager
//     method.
func (client *ChannelsClient) NewListByResourceGroupPager(resourceGroupName string, resourceName string, options *ChannelsClientListByResourceGroupOptions) *runtime.Pager[ChannelsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ChannelsClientListByResourceGroupResponse]{
		More: func(page ChannelsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ChannelsClientListByResourceGroupResponse) (ChannelsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ChannelsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return ChannelsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ChannelsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ChannelsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ChannelsClient) listByResourceGroupHandleResponse(resp *http.Response) (ChannelsClientListByResourceGroupResponse, error) {
	result := ChannelsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChannelResponseList); err != nil {
		return ChannelsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListWithKeys - Lists a Channel registration for a Bot Service including secrets
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - channelName - The name of the Channel resource.
//   - options - ChannelsClientListWithKeysOptions contains the optional parameters for the ChannelsClient.ListWithKeys method.
func (client *ChannelsClient) ListWithKeys(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, options *ChannelsClientListWithKeysOptions) (ChannelsClientListWithKeysResponse, error) {
	var err error
	const operationName = "ChannelsClient.ListWithKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listWithKeysCreateRequest(ctx, resourceGroupName, resourceName, channelName, options)
	if err != nil {
		return ChannelsClientListWithKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChannelsClientListWithKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ChannelsClientListWithKeysResponse{}, err
	}
	resp, err := client.listWithKeysHandleResponse(httpResp)
	return resp, err
}

// listWithKeysCreateRequest creates the ListWithKeys request.
func (client *ChannelsClient) listWithKeysCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, options *ChannelsClientListWithKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}/listChannelWithKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(string(channelName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWithKeysHandleResponse handles the ListWithKeys response.
func (client *ChannelsClient) listWithKeysHandleResponse(resp *http.Response) (ChannelsClientListWithKeysResponse, error) {
	result := ChannelsClientListWithKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListChannelWithKeysResponse); err != nil {
		return ChannelsClientListWithKeysResponse{}, err
	}
	return result, nil
}

// Update - Updates a Channel registration for a Bot Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - channelName - The name of the Channel resource.
//   - parameters - The parameters to provide for the created bot.
//   - options - ChannelsClientUpdateOptions contains the optional parameters for the ChannelsClient.Update method.
func (client *ChannelsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, parameters BotChannel, options *ChannelsClientUpdateOptions) (ChannelsClientUpdateResponse, error) {
	var err error
	const operationName = "ChannelsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, channelName, parameters, options)
	if err != nil {
		return ChannelsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChannelsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ChannelsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ChannelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, channelName ChannelName, parameters BotChannel, options *ChannelsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if channelName == "" {
		return nil, errors.New("parameter channelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{channelName}", url.PathEscape(string(channelName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ChannelsClient) updateHandleResponse(resp *http.Response) (ChannelsClientUpdateResponse, error) {
	result := ChannelsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BotChannel); err != nil {
		return ChannelsClientUpdateResponse{}, err
	}
	return result, nil
}
