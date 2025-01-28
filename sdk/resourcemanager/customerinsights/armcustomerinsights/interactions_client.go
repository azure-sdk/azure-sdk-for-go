//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

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

// InteractionsClient contains the methods for the Interactions group.
// Don't use this type directly, use NewInteractionsClient() instead.
type InteractionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInteractionsClient creates a new instance of InteractionsClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInteractionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InteractionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InteractionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates an interaction or updates an existing interaction within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - interactionName - The name of the interaction.
//   - parameters - Parameters supplied to the CreateOrUpdate Interaction operation.
//   - options - InteractionsClientBeginCreateOrUpdateOptions contains the optional parameters for the InteractionsClient.BeginCreateOrUpdate
//     method.
func (client *InteractionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, interactionName string, parameters InteractionResourceFormat, options *InteractionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[InteractionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, hubName, interactionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InteractionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[InteractionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates an interaction or updates an existing interaction within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
func (client *InteractionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, hubName string, interactionName string, parameters InteractionResourceFormat, options *InteractionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "InteractionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, interactionName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InteractionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, interactionName string, parameters InteractionResourceFormat, options *InteractionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if interactionName == "" {
		return nil, errors.New("parameter interactionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{interactionName}", url.PathEscape(interactionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets information about the specified interaction.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - interactionName - The name of the interaction.
//   - options - InteractionsClientGetOptions contains the optional parameters for the InteractionsClient.Get method.
func (client *InteractionsClient) Get(ctx context.Context, resourceGroupName string, hubName string, interactionName string, options *InteractionsClientGetOptions) (InteractionsClientGetResponse, error) {
	var err error
	const operationName = "InteractionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, interactionName, options)
	if err != nil {
		return InteractionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InteractionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InteractionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *InteractionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, interactionName string, options *InteractionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if interactionName == "" {
		return nil, errors.New("parameter interactionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{interactionName}", url.PathEscape(interactionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	if options != nil && options.LocaleCode != nil {
		reqQP.Set("locale-code", *options.LocaleCode)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InteractionsClient) getHandleResponse(resp *http.Response) (InteractionsClientGetResponse, error) {
	result := InteractionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InteractionResourceFormat); err != nil {
		return InteractionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHubPager - Gets all interactions in the hub.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - options - InteractionsClientListByHubOptions contains the optional parameters for the InteractionsClient.NewListByHubPager
//     method.
func (client *InteractionsClient) NewListByHubPager(resourceGroupName string, hubName string, options *InteractionsClientListByHubOptions) *runtime.Pager[InteractionsClientListByHubResponse] {
	return runtime.NewPager(runtime.PagingHandler[InteractionsClientListByHubResponse]{
		More: func(page InteractionsClientListByHubResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InteractionsClientListByHubResponse) (InteractionsClientListByHubResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "InteractionsClient.NewListByHubPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByHubCreateRequest(ctx, resourceGroupName, hubName, options)
			}, nil)
			if err != nil {
				return InteractionsClientListByHubResponse{}, err
			}
			return client.listByHubHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByHubCreateRequest creates the ListByHub request.
func (client *InteractionsClient) listByHubCreateRequest(ctx context.Context, resourceGroupName string, hubName string, options *InteractionsClientListByHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	if options != nil && options.LocaleCode != nil {
		reqQP.Set("locale-code", *options.LocaleCode)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHubHandleResponse handles the ListByHub response.
func (client *InteractionsClient) listByHubHandleResponse(resp *http.Response) (InteractionsClientListByHubResponse, error) {
	result := InteractionsClientListByHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InteractionListResult); err != nil {
		return InteractionsClientListByHubResponse{}, err
	}
	return result, nil
}

// SuggestRelationshipLinks - Suggests relationships to create relationship links.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - interactionName - The name of the interaction.
//   - options - InteractionsClientSuggestRelationshipLinksOptions contains the optional parameters for the InteractionsClient.SuggestRelationshipLinks
//     method.
func (client *InteractionsClient) SuggestRelationshipLinks(ctx context.Context, resourceGroupName string, hubName string, interactionName string, options *InteractionsClientSuggestRelationshipLinksOptions) (InteractionsClientSuggestRelationshipLinksResponse, error) {
	var err error
	const operationName = "InteractionsClient.SuggestRelationshipLinks"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.suggestRelationshipLinksCreateRequest(ctx, resourceGroupName, hubName, interactionName, options)
	if err != nil {
		return InteractionsClientSuggestRelationshipLinksResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InteractionsClientSuggestRelationshipLinksResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InteractionsClientSuggestRelationshipLinksResponse{}, err
	}
	resp, err := client.suggestRelationshipLinksHandleResponse(httpResp)
	return resp, err
}

// suggestRelationshipLinksCreateRequest creates the SuggestRelationshipLinks request.
func (client *InteractionsClient) suggestRelationshipLinksCreateRequest(ctx context.Context, resourceGroupName string, hubName string, interactionName string, options *InteractionsClientSuggestRelationshipLinksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}/suggestRelationshipLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if interactionName == "" {
		return nil, errors.New("parameter interactionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{interactionName}", url.PathEscape(interactionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// suggestRelationshipLinksHandleResponse handles the SuggestRelationshipLinks response.
func (client *InteractionsClient) suggestRelationshipLinksHandleResponse(resp *http.Response) (InteractionsClientSuggestRelationshipLinksResponse, error) {
	result := InteractionsClientSuggestRelationshipLinksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SuggestRelationshipLinksResponse); err != nil {
		return InteractionsClientSuggestRelationshipLinksResponse{}, err
	}
	return result, nil
}
